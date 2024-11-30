package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"text/template"
	"time"

	"aoc/common"
	"github.com/spf13/pflag"
)

func main() {
	day := pflag.IntP("day", "d", time.Now().Day(), "Advent of Code Day")
	pflag.Parse()

	dirname := fmt.Sprintf("day%02d", *day)
	if err := os.MkdirAll(dirname, 0755); err != nil {
		log.Fatalf("Error creating destination directory %s: %s\n", dirname, err)
	}

	templates := template.Must(template.ParseFS(os.DirFS("templates"), "*.tmpl"))

	partOneFileName := "part_one.go"
	f, err := os.Create(path.Join(dirname, partOneFileName))
	common.Check(err, "unable to create file %s", partOneFileName)
	defer common.CloseFile(f, partOneFileName)

	partOneTemplateName := "part.tmpl"
	err = templates.ExecuteTemplate(f, partOneTemplateName, *day)
	common.Check(err, "unable to execute template %s", partOneTemplateName)

	testsFileName := "part_one_test.go"
	ft, err := os.Create(path.Join(dirname, testsFileName))
	common.Check(err, "unable to create file %s", testsFileName)
	defer common.CloseFile(ft, testsFileName)

	partTestTemplateName := "part_test.tmpl"
	err = templates.ExecuteTemplate(ft, partTestTemplateName, *day)
	common.Check(err, "unable to execute template %s", partTestTemplateName)
}
