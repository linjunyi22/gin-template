package template

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"runtime"
)

type Template struct {
	projectName string
	projectPath string // working directory + project name
}

func getwd() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("error in get working directory.")
		errWrapper(err)
	}
	return dir
}

func errWrapper(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func NewTemplate(projectName string) *Template {
	return &Template{
		projectName: projectName,
		projectPath: path.Join(getwd(), projectName),
	}
}

func (t *Template) projectIsExist() bool {
	_, err := os.Stat(t.projectPath)
	return err == nil || os.IsExist(err)
}

func (t *Template) createDir(dirName string) error {
	p := path.Join(t.projectPath, dirName)
	return os.MkdirAll(p, 0755)
}

func (t *Template) createFile(dirPath, fileName, fileContent string) error {
	err := os.Chdir(path.Join(t.projectPath, dirPath))
	if err != nil {
		return err
	}
	f, err := os.Create(fileName)
	defer f.Close()

	if err != nil {
		return err
	}
	_, err = f.Write([]byte(fileContent))
	return err
}

func (t *Template) modInit() string {
	v := runtime.Version()
	r, _ := regexp.Compile("go(\\d\\.\\d+)(.*)")
	res := r.FindStringSubmatch(v)
	version := "1.13"
	if len(res) >= 1 {
		version = res[1]
	}
	return fmt.Sprintf(goModuleStr, t.projectName, version)
}

func (t *Template) Run() {
	isExist := t.projectIsExist()

	if isExist {
		errWrapper(errors.New(fmt.Sprintf("project %s is existed", t.projectName)))
	}

	dirs := []string{"conf", "models", "routers", "controllers", "utils"}
	for _, dir := range dirs {
		errWrapper(t.createDir(dir))
	}

	errWrapper(t.createFile("", "README.md", ""))
	errWrapper(t.createFile("", "main.go", mainStr))
	errWrapper(t.createFile("", "go.mod", t.modInit()))
	errWrapper(t.createFile("controllers", "controller.go", controllerStr))
	errWrapper(t.createFile("routers", "router.go", routerStr))
	errWrapper(t.createFile("models", "model.go", modelStr))
	errWrapper(t.createFile("conf", "app.conf", ""))

}
