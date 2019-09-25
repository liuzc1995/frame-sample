package controller

import (
	"html/template"
	"io/ioutil"
	"os"
)

//模板渲染
func PopulateTemplates() map[string]*template.Template {
	const basePath = "templates"
	result := make(map[string]*template.Template)
	//取根模板
	layout := template.Must(template.ParseFiles(basePath + "/_base.html"))
	//遍历模板根目录下的文件夹
	dirs, err := os.Open(basePath)
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fiss, err := dirs.Readdir(-1)
	if err != nil {
		panic("Failed to read dirs: " + err.Error())
	} else {
		//过滤在模板根目录下的html文件
		for _, fl := range fiss {
			lenPath := len(fl.Name()) //文件夹名长度
			if lenPath <= 5 || fl.Name()[lenPath-5:lenPath] != ".html" {
				dir, err := os.Open(basePath + "/" + fl.Name())
				if err != nil {
					panic("Failed to open " + fl.Name() + " blocks directory: " + err.Error())
				}
				fis, err := dir.Readdir(-1)
				if err != nil {
					panic("Failed to read " + fl.Name() + ": " + err.Error())
				}
				//渲染每个文件夹下的模板
				for _, fi := range fis {
					f, err := os.Open(basePath + "/" + fl.Name() + "/" + fi.Name())
					if err != nil {
						panic("Failed to open template '" + fi.Name() + "'")
					}
					//取模板页内容
					content, err := ioutil.ReadAll(f)
					if err != nil {
						panic("Failed to read content from file '" + fi.Name() + "'")
					}
					f.Close()
					//创建模板
					tmpl := template.Must(layout.Clone())
					_, err = tmpl.Parse(string(content))
					if err != nil {
						panic("Failed to parse contents of '" + fi.Name() + "' as template")
					}
					result[fl.Name()+"/"+fi.Name()] = tmpl
				}
			}

		}
	}
	return result
}
