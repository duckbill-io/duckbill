package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// 负责将cli包解析出的数据映射为models中的对象
type Orm struct{}

func (orm *Orm) scan(vi interface{}) (err error) {
	switch v := vi.(type) {
	case *Tags:
		tagsfilepath := "_data/tags/tags.json"
		tagsfile, err := os.Open(tagsfilepath)
		defer tagsfile.Close()
		if err != nil {
			return err
		}
		err = json.NewDecoder(tagsfile).Decode(v)
		if err != nil {
			return err
		}
	case *Tag:
		tagsfilepath := "_data/tags/tags.json"
		tagsfile, err := os.Open(tagsfilepath)
		defer tagsfile.Close()
		if err != nil {
			return err
		}
		tags := []Tag{}
		err = json.NewDecoder(tagsfile).Decode(&tags)
		if err != nil {
			return err
		}
		for _, tag := range tags {
			if tag.Name == v.Name {
				*v = tag
			}
		}
	case *Metas:
		metasdir := "_data/metas"
		f, err := os.Open(metasdir)
		if err != nil {
			return err
		}
		list, err := f.Readdir(-1)
		f.Close()
		if err != nil {
			return err
		}
		// 筛选出为json格式的文件
		for i, j := 0, 0; i < len(list); i++ {
			if !list[i].IsDir() && strings.HasSuffix(list[i].Name(), ".json") {
				list[j] = list[i]
				j++
			}
			if i == len(list)-1 {
				list = list[:j]
			}
		}
		// 读取所有元信息数据
		buf := bytes.NewBuffer([]byte{})
		buf.WriteString("[")
		for i := range list {
			metafilepath := filepath.Join(metasdir, list[i].Name())
			metabytes, err := ioutil.ReadFile(metafilepath)
			if err != nil {
				return err
			}
			buf.WriteString(string(metabytes))
			buf.WriteString(",")
		}
		buf.Truncate(buf.Len() - len(","))
		buf.WriteString("]")
		// 反序列化元信息数据
		err = json.NewDecoder(buf).Decode(v)
		if err != nil {
			return err
		}
	case *Meta:
		metafilepath := filepath.Join("_data/metas", v.Name+".json")
		metafile, err := os.Open(metafilepath)
		defer metafile.Close()
		if err != nil {
			return err
		}
		err = json.NewDecoder(metafile).Decode(v)
		if err != nil {
			return err
		}
	case *Post:
		postfilepath := filepath.Join("_data/posts", v.Name+".html")
		htmlpost, err := ioutil.ReadFile(postfilepath)
		if err != nil {
			return err
		}
		v.Content = string(htmlpost)
	default:
		err = fmt.Errorf("ormx不支持次类型")
	}
	return
}
