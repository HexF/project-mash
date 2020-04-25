package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/Masterminds/sprig"
	"github.com/markbates/pkger"
)

type Node struct {
	Uptime          time.Duration
	Hostname        string
	PreviousRequest string
	AppVersion      string
}

type Page struct {
	Mash        string
	Fact        template.HTML
	NextURL     string
	PreviousURL string
}

type Template struct {
	Node Node
	Page Page
}

var pageTemplate string
var myNode Node
var GitCommit string

func init() {
	f, err := pkger.Open("/template/page.html")
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	pageTemplate = string(b)
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	myNode = Node{
		Hostname:   hostname,
		AppVersion: GitCommit,
	}
}

func renderTemplate(node Node, mash Page, w io.Writer) {
	tmpl, err := template.New("test").Funcs(sprig.FuncMap()).Parse(pageTemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, Template{
		Node: node,
		Page: mash,
	})
}
