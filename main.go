package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"

	"github.com/see4self/server/utils"
)

var svcName = "./tms"

func init() {
	flag.StringVar(&svcName, "n", svcName, "service name")
	flag.Parse()
}

type Service struct {
	Name      string
	GoVersion string
}

func goVersion() string {
	v := runtime.Version()
	end := len(v)
	if strings.Count(v, ".") > 1 {
		end = strings.LastIndex(v, ".")
	}
	return v[2:end]
}

func main() {
	svc := Service{
		Name:      filepath.Base(svcName),
		GoVersion: goVersion(),
	}
	dirs, files := utils.ReadDir()
	for i, v := range dirs {
		dirs[i] = strings.Replace(v, "templates", svcName, 1)
		// fmt.Println(dirs[i])
		utils.Mkdir(dirs[i])
	}
	newfiles := make([]string, len(files))
	for i, file := range files {
		newfiles[i] = strings.TrimSuffix(strings.Replace(file, "templates", svcName, 1), ".tmpl")
		// fmt.Println(newfiles[i])
	}
	tpl, names := newTemplate(files)
	for i, v := range names {
		f, err := os.Create(newfiles[i])
		if err != nil {
			fmt.Printf("create %s err %v\n", newfiles[i], err)
			os.Exit(1)
		}
		if err = tpl.ExecuteTemplate(f, v, svc); err != nil {
			fmt.Println("create template err", err)
			os.Exit(1)
		}
		f.Close()
	}

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

func testTemplate() {
	a := struct {
		Name string
	}{"sss"}
	tpl, _ := template.New("test").Parse("My name is {{.Name}}\n")
	tpl.Execute(os.Stdout, a)
	f, _ := os.Create("./service/main.go")
	defer f.Close()
	tpl1, err := template.ParseFiles("templates/main.go.tpl")
	fmt.Println(err)

	err = tpl1.Execute(f, a)
	fmt.Println(err)
}
