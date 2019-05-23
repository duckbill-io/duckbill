// 包render用于渲染模板与数据
package render

import (
	"html/template"
	"io"
)

func toTemplateHTML(text string) template.HTML {
	return template.HTML(text)
}

// Execute 执行渲染过程
func Execute(w io.Writer, filename string, data interface{}) error {
	filenames := []string{"views/layouts/application.html", filename}

	funcMap := template.FuncMap{"toHTML": toTemplateHTML}
	t := template.Must(template.ParseFiles(filenames[0]))
	tc := t.New("content").Funcs(funcMap)
	tc.ParseFiles(filenames[1])

	err := t.Execute(w, data)
	return err
}

