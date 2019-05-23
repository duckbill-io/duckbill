package cli

import (
	"bytes"
	"gopkg.in/russross/blackfriday.v2"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Parser 解析器用于从FromDir解析所有的markdown文件到ToDir
type Parser struct {
	FromDir string // 解析器获取数据的文件夹
	ToDir   string // 解析器输出数据的文件夹
}

// NewParser 创建一Parser实例
func NewParser(from, to string) *Parser {
	return &Parser{
		FromDir: from,
		ToDir:   to,
	}
}

// Fire 开始解析
// 目前是全量解析，不是增量解析(要么全成功，要么全失败)
func (p *Parser) Fire() (err error) {
	filenames, err := getAllMarkDownFileNames(p.FromDir)
	if err != nil {
		return
	}
	// 清空p.ToDir
	err = p.rebuildToDir()
	if err != nil {
		return
	}
	// 遍历每一个markdown文件, 并解析
	for _, filename := range filenames {
		err = parse(filename, p.ToDir)
		if err != nil {
			return
		}
	}
	return
}

// rebuildToDir 清空p.ToDir并重建p.ToDir
func (p *Parser) rebuildToDir() (err error) {
	// 清空文件夹
	err = os.RemoveAll(p.ToDir)
	if err != nil {
		return err
	}
	// 创建文件夹
	err = os.Mkdir(p.ToDir, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.Mkdir(filepath.Join(p.ToDir, "posts"), os.ModePerm)
	if err != nil {
		return err
	}
	err = os.Mkdir(filepath.Join(p.ToDir, "metas"), os.ModePerm)
	if err != nil {
		return err
	}
	return
}

// getAllMarkDownFileNames 找出所有的markdown文件
func getAllMarkDownFileNames(dirname string) ([]string, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	// 筛选出list中为markdown文件的元素
	names := make([]string, 0, len(list))
	for i := range list {
		if ispostmd(list[i]) {
			names = append(names, filepath.Join(dirname, list[i].Name()))
		}
	}
	// 排序
	sort.Slice(names, func(i, j int) bool { return names[i] < names[j] })
	return names, nil
}

// ismd 判断是否是markdown文件
func ispostmd(fi os.FileInfo) bool {
	return strings.HasSuffix(fi.Name(), ".md") && !fi.IsDir() && fi.Name() != "README.md"
}

// parse 解析文件
func parse(filePath, toDir string) (err error) {
	// 分割md文件中的文章元信息与内容
	postinfo, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	postinfo = bytes.TrimSpace(postinfo)
	sep := []byte("---\n")
	re := bytes.SplitN(postinfo, sep, 3)
	meta, post := re[1], re[2]
	// 转换文件内容为html
	post = blackfriday.Run(post)
	_post := template.HTMLEscapeString(string(post))
	post = []byte(_post)
	post = bytes.ReplaceAll(post, []byte("\n"), []byte(""))
	// 转换文件内容为json格式
	post = tojson(post)
	// 分别单独存储文章的元信息与内容
	_, filename := filepath.Split(filePath)
	filename = strings.TrimSuffix(filename, ".md")
	metafilepath := filepath.Join(toDir, "metas", filename) + ".yml"
	postfilepath := filepath.Join(toDir, "posts", filename) + ".json"
	// 保存文件
	postfile, err := os.Create(postfilepath)
	if err != nil {
		return
	}
	defer postfile.Close()
	metafile, err := os.Create(metafilepath)
	if err != nil {
		return
	}
	defer metafile.Close()
	_, err = postfile.Write(post)
	if err != nil {
		return
	}
	_, err = metafile.Write(meta)
	if err != nil {
		return
	}
	return
}

// tojson 将内容装换为post的json形式
func tojson(post []byte) []byte {
	jsonpost := []byte("{\"content\": ")
	jsonpost = append(jsonpost, []byte("\"")...)
	jsonpost = append(jsonpost, post...)
	jsonpost = append(jsonpost, []byte("\"}")...)
	return jsonpost
}