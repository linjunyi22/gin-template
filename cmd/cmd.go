package cmd

import (
	"flag"
	"github.com/linjunyi22/gin-template/template"
	"log"
)

type Cmd struct {
}

func NewCmd() *Cmd {
	return &Cmd{}
}

func (c *Cmd) Run() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 || (len(args) != 0 && args[0] != "new") {
		log.Fatal("use command gin-template new [project name], e.g 'gin-template new hello' ")
		return
	}

	t := template.NewTemplate(args[1])
	t.Run()
}
