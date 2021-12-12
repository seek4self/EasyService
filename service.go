package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/see4self/server/utils"
)

type Service struct {
	data  TemplateData
	files []string
}

type TemplateData struct {
	Name      string
	GoVersion string
}

func newService() *Service {
	return &Service{
		data: TemplateData{
			Name:      filepath.Base(svcName),
			GoVersion: goVersion(),
		},
	}
}

func (s *Service) createDirs() {
	dirs, files := utils.ReadDir()
	for i, v := range dirs {
		dirs[i] = strings.Replace(v, "templates", svcName, 1)
		// fmt.Println(dirs[i])
		utils.Mkdir(dirs[i])
	}
	s.files = files
}

func (s *Service) createGoFiles() {
	newfiles := make([]string, len(s.files))
	for i, file := range s.files {
		newfiles[i] = strings.TrimSuffix(strings.Replace(file, "templates", svcName, 1), ".tmpl")
		// fmt.Println(newfiles[i])
	}
	tpl, names := newTemplate(s.files)
	for i, name := range names {
		f, err := os.Create(newfiles[i])
		if err != nil {
			fmt.Printf("create %s err %v\n", newfiles[i], err)
			os.Exit(1)
		}
		if err = tpl.ExecuteTemplate(f, name, s.data); err != nil {
			fmt.Println("create template err", err)
			os.Exit(1)
		}
		f.Close()
	}
}

func (s *Service) installMod() {
	runCMD([]string{"go", "mod", "tidy"})
	runCMD([]string{"go", "get", "-u", "github.com/swaggo/swag/cmd/swag"})
	runCMD([]string{os.Getenv("GOPATH") + "/bin/swag", "init"})
}

func newTemplate(tplFiles []string) (*template.Template, []string) {
	var err error
	tpl := template.New("service").Funcs(template.FuncMap{
		"upper": strings.ToUpper,
	})
	tpl, err = tpl.ParseFiles(tplFiles...)
	if err != nil {
		fmt.Println("template parse files err", err)
		os.Exit(1)
	}
	names := make([]string, len(tplFiles))
	for i, file := range tplFiles {
		names[i] = filepath.Base(file)
	}
	return tpl, names
}

func runCMD(params []string) {
	cmd := exec.Command(params[0], params[1:]...)
	cmd.Dir = svcName
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Printf("run %s err: %v\n", cmd.String(), err)
		os.Exit(1)
	}
}
