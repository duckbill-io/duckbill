package cli

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"encoding/json"
	"gopkg.in/russross/blackfriday.v2"
	"gopkg.in/yaml.v2"
)

// Parser 解析器 用于从FromDir解析所有的markdown文件到ToDir
type Parser struct {
	inputdir   string        // 解析器获取数据的文件夹
	inputinfos []os.FileInfo // inputdir中有效的文件信息列表
	outputdir  string        // 解析器保存最终数据的文件夹
}

// NewParser　创建一个解析其
func NewParser(inputdir, outputdir string) *Parser {
	return &Parser{
		inputdir:  inputdir,
		outputdir: outputdir,
	}
}

// Run 开始解析
// 目前是全量解析，而非增量(要么全部成功, 要么全部失败)
func (p *Parser) Run() {
	p.prepare()
	// 遍历p.inputdir中的markdown文件, 完成解析
	p.parse()
	return
}

// prepare 准备工作(清空并重建p.outputdir/查找可用的input文件)
func (p *Parser) prepare() {
	// 清空并重建p.outputdir
	p.rebuild()
	// 查找文件夹中的所有内容
	f, err := os.Open(p.inputdir)
	checkerror(err)
	list, err := f.Readdir(-1)
	f.Close()
	checkerror(err)
	// 筛选出有效的输入文件
	for i, j := 0, 0; i < len(list); i++ {
		if ispostmd(list[i]) {
			list[j] = list[i]
			j++
		}
		if i == len(list) {
			list = list[:j]
		}
	}
	sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
	p.inputinfos = list
}

func (p *Parser) rebuild() {
	var err error
	// 清空文件夹
	err = os.RemoveAll(p.outputdir)
	checkerror(err)
	// 创建文件夹
	err = os.Mkdir(p.outputdir, os.ModePerm)
	checkerror(err)
	for _, dir := range []string{"metas", "posts", "tags"} {
		err = os.Mkdir(filepath.Join(p.outputdir, dir), os.ModePerm)
		checkerror(err)
	}
}

func (p *Parser) parse() {
	tagsmap := map[string][]string{}
	var err error

	for i := range p.inputinfos {
		mdfilename := filepath.Join(p.inputdir, p.inputinfos[i].Name())
		// 分割文章元信息与内容
		post, err := ioutil.ReadFile(mdfilename)
		checkerror(err)
		post = bytes.TrimSpace(post)
		slicepost := bytes.SplitN(post, []byte("---\n"), 3)
		ymlmeta, mdcontent := slicepost[1], slicepost[2]
		// 文章内容转换为html
		htmlcontent := blackfriday.Run(mdcontent)
		// 反序列化文章元信息
		metastruct := struct {
			Name, CreatedAt string
			Tags            []string
		}{}
		err = yaml.Unmarshal(ymlmeta, &metastruct)
		checkerror(err)
		// 序列化文章元信息为json格式
		jsonmeta, err := json.MarshalIndent(metastruct, "", "")
		// 保存jsonmeta与htmlcontent
		filename := metastruct.Name
		metafilepath := filepath.Join(p.outputdir, "metas", filename+".json")
		contentfilepath := filepath.Join(p.outputdir, "posts", filename+".html")
		metafile, err := os.Create(metafilepath)
		checkerror(err)
		_, err = metafile.Write(jsonmeta)
		checkerror(err)
		metafile.Close()
		contentfile, err := os.Create(contentfilepath)
		checkerror(err)
		_, err = contentfile.Write(htmlcontent)
		checkerror(err)
		contentfile.Close()
		// 通过文章元信息筛选出tagsmap映射
		for i := range metastruct.Tags {
			if _, exist := tagsmap[metastruct.Name]; !exist {
				tagsmap[metastruct.Tags[i]] = []string{metastruct.Name}
			} else {
				tagsmap[metastruct.Tags[i]] = append(tagsmap[metastruct.Tags[i]], metastruct.Name)
			}
		}
	}
	// 反序列化tagsmap映射为单篇json格式文件
	tagsfilepath := filepath.Join(p.outputdir, "tags", "tags.json")
	tagsfile, err := os.Create(tagsfilepath)
	checkerror(err)
	var buf bytes.Buffer
	buf.WriteString("[")
	for tagname, articles := range tagsmap {
		buf.WriteString(fmt.Sprintf("{\"name\": \"%s\",", tagname))
		buf.WriteString(fmt.Sprintf("\"count\": \"%d\",", len(articles)))
		buf.WriteString("\"articles\": [")
		for i := range articles {
			buf.WriteString(fmt.Sprintf("\"%s\"", articles[i]))
			if i < len(articles)-1 {
				buf.WriteString(",")
			}
		}
		buf.WriteString("]},\n")
	}
	buf.Truncate(buf.Len() - len(",\n"))
	buf.WriteString("]")
	_, err = tagsfile.Write([]byte(buf.String()))
	checkerror(err)
	tagsfile.Close()
}

// ispostmd 判断是否是文章的markdown文件
func ispostmd(fi os.FileInfo) bool {
	if fi.IsDir() {
		return false
	}
	if !strings.HasSuffix(fi.Name(), ".md") {
		return false
	}
	if fi.Name() == "README.md" {
		return false
	}
	return true
}

func checkerror(err error) {
	if err != nil {
		panic(err)
	}
}
