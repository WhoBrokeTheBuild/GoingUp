package goingup

import (
	"io/ioutil"
	"path/filepath"
	
	"github.com/russross/blackfriday"
)

func parseContentGlob(pattern string) (map[string]string) {
	filenames, err := filepath.Glob(pattern)
	if err != nil {
		panic(err)
	}
	if len(filenames) == 0 {
		return nil
	}
	
	content := map[string]string{}
	for _, file := range filenames {
		ctnt, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		html := blackfriday.MarkdownCommon(ctnt)
		
		name := filepath.Base(file)
		ext := filepath.Ext(name)
		name = name[0:len(name) - len(ext)]
		content[name] = string(html)
	}
	
	return content
}
