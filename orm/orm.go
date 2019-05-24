// 包orm负责将cli包解析出的数据映射为model中的对象
package orm

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

// Scan 对外暴露扫描更新数据接口
func Scan(v interface{}) error {
	return scan(reflect.ValueOf(v))
}

// scan　实际查找对应数据，并更新实例的接口
func scan(v reflect.Value) error {
	switch v.Kind() {
	case reflect.Ptr:
		ve := v.Elem()
		if v.IsNil() {
			return fmt.Errorf("cann't scan nil Ptr")
		}
		if ve.Kind() == reflect.Struct {
			return unmarshal(v)
		}
		if ve.Kind() == reflect.Slice || ve.Kind() == reflect.Array {
			prepare := func(v reflect.Value) (names []string, err error) {
				_filepath := "_data/metas"
				f, err := os.Open(_filepath)
				if err != nil {
					return
				}
				list, err := f.Readdir(-1)
				f.Close()
				if err != nil {
					return
				}
				names = make([]string, 0, len(list))
				for i := range list {
					name := strings.Split(list[i].Name(), ".")[0]
					names = append(names, name)
				}
				return
			}
			if ve.Len() < 1 {
				return nil
			}
			names, err := prepare(ve.Index(0))
			if err != nil {
				return err
			}
			t := ve.Index(0).Elem().Type()
			for i := 0; i < len(names); i++ {
				item := reflect.New(t)
				item.Elem().FieldByName("Name").SetString(names[i])
				scan(item)
				ve.Set(reflect.Append(ve, item))
			}
		}
	default:
		return fmt.Errorf("must be a Ptr")
	}
	return nil
}

func unmarshal(v reflect.Value) error {
	prepare := func(v reflect.Value) (filename string, _unmarshal func([]byte, interface{}) error) {
		if strings.HasSuffix(v.Type().String(), "Meta") {
			filename = "_data/metas/" + fmt.Sprint(v.FieldByName("Name")) + ".yml"
			_unmarshal = yaml.Unmarshal
		}
		if strings.HasSuffix(v.Type().String(), "Post") {
			filename = "_data/posts/" + fmt.Sprint(v.FieldByName("Name")) + ".html"
			_unmarshal = htmlunmarshal
		}
		return
	}

	ve := v.Elem()
	filename, _unmarshal := prepare(ve)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error ioutil.ReadFile(", filename, ") ", err)
	}

	err = _unmarshal(data, v.Interface())
	if err != nil {
		fmt.Println("error _unmarshal ", err)
	}

	return nil
}

func htmlunmarshal(data []byte, v interface{}) error {
	ve := reflect.ValueOf(v).Elem()
	content := string(data)
	ve.FieldByName("Content").SetString(content)
	return nil
}
