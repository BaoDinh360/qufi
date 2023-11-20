package template

import (
	"example/qufi/features/config"
	"example/qufi/helpers"
	"fmt"
	"os"
	"path/filepath"
	gotemplate "text/template"
	"time"
)

// type RecordTemplateData struct {
// 	Datetime   string
// 	Blockquote string
// }

type TemplateInfo struct {
	TemplateFile string
	TemplatePath string
}

// func GenerateTemplateData(blockquote string) RecordTemplateData {
// 	now := time.Now()
// 	templatedata := RecordTemplateData{
// 		Datetime:   now.Format("Monday, January 02, 2006"),
// 		Blockquote: blockquote,
// 	}
// 	return templatedata
// }

// check if there is template file
func GetTemplateFile(templatename string) (bool, TemplateInfo) {
	basecfg := config.GetBaseConfig()
	tmplfile := fmt.Sprintf("%s.tmpl", templatename)
	tmplpath := filepath.Join(basecfg.TemplatePath, tmplfile)

	isfileexists, _ := helpers.IsPathExists(tmplpath)
	if isfileexists {
		return true, TemplateInfo{
			TemplateFile: tmplfile,
			TemplatePath: tmplpath,
		}
	} else {
		return false, TemplateInfo{}
	}
}

func WriteTemplateIfExists(templatename string, destfile *os.File) error {

	//check if exists a template file with templatename in templates directory
	istmplfile, tmplinfo := GetTemplateFile(templatename)
	if !istmplfile {
		return nil
	}

	//define custom functions for template
	funcmap := gotemplate.FuncMap{
		"today": func(f string) string { return time.Now().Format(f) },
		"affirmation": func() string {
			result, err := config.GetRandomAffirmation()
			if err != nil {
				return "Error: Cannot get affirmation"
			}
			return result
		},
	}

	//only perform write template if template file exists
	tmpl, err := gotemplate.New(tmplinfo.TemplateFile).Funcs(funcmap).ParseFiles(tmplinfo.TemplatePath)
	if err != nil {
		return err
	}

	// tmpldata := struct {
	// 	Datetime   string
	// 	Dailyquote string
	// }{
	// 	Datetime:   time.Now().Format("Monday, January 02, 2006"),
	// 	Dailyquote: "You're an inspiration",
	// }

	err = tmpl.Execute(destfile, nil)
	if err != nil {
		return err
	}
	return nil
}
