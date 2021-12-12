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

func goVersion() string {
	v := runtime.Version()
	end := len(v)
	if strings.Count(v, ".") > 1 {
		end = strings.LastIndex(v, ".")
	}
	return v[2:end]
}

func main() {
	svc := newService()
	svc.createDirs()
	svc.createGoFiles()
	svc.installMod()
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
