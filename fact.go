package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/markbates/pkger"
)

type MashFact struct {
	Name     string
	FactText string
	FactFile string
}

var fileRes pkg.Resource

func (mf *MashFact) GetHashes() string {
	data := []byte(mf.Name)
	sha256b := sha256.Sum256(data)
	sha224b := sha256.Sum224(data)
	sha1b := sha1.Sum(data)

	sha256s := hex.EncodeToString(sha256b[:])
	sha224s := hex.EncodeToString(sha224b[:])
	sha1s := hex.EncodeToString(sha1b[:])

	return fmt.Sprintf("<br/>SHA256: %v<br/>SHA224: %v<br/>SHA1: %v<br/>", sha256s, sha224s, sha1s)
}

func init() {
	f, err := pkger.Open("/facts/rules.grl")
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fileRes = pkg.NewBytesResource(b)

	pkger.Walk("/facts", func(path string, info os.FileInfo, err error) error {
		return err
	})
	//Here we just walk the `/facts` directory so pkger packs all the files
}

func getFact(mashName string) template.HTML {
	factText := ""

	myFact := &MashFact{
		Name: mashName,
	}

	dataCtx := ast.NewDataContext()
	err := dataCtx.Add("Mash", myFact)
	if err != nil {
		panic(err)
	}

	workingMemory := ast.NewWorkingMemory()
	knowledgeBase := ast.NewKnowledgeBase("facts", GitCommit)

	ruleBuilder := builder.NewRuleBuilder(knowledgeBase, workingMemory)

	err = ruleBuilder.BuildRuleFromResource(fileRes)
	if err != nil {
		panic(err)
	}

	factEngine := engine.NewGruleEngine()

	err = factEngine.Execute(dataCtx, knowledgeBase, workingMemory)
	if err != nil {
		panic(err)
	}

	if myFact.FactFile != "" {

		f, err := pkger.Open("/facts/" + myFact.FactFile)
		if err != nil {
			panic(err)
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}

		factText = string(b)
	}

	factText = factText + myFact.FactText
	factText = strings.TrimSpace(factText)
	return template.HTML(factText)
}
