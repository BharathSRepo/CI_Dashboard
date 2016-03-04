package populatetemplates

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

func PopulateTemplates() *template.Template {

	alltemplates := template.New("templates")

	basePath, err := filepath.Abs("./templates")

	if err != nil {
		fmt.Println("Could not find path for templates -", basePath)
	} else {
		basePath = filepath.ToSlash(basePath)
		templateFolder, err := os.Open(basePath)
		if err != nil {
			fmt.Println("Could not open templates - ", basePath)
		} else {
			defer templateFolder.Close()
			templatePathsRaw, _ := templateFolder.Readdir(-1)
			templatePaths := new([]string)
			for _, pathInfo := range templatePathsRaw {
				if !pathInfo.IsDir() {
					*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
				}
			}
			alltemplates.ParseFiles(*templatePaths...)
		}
	}
	return alltemplates
}
