package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"text/template"
	"time"

	"aoc/common"
)

type DayPart struct {
	Day  int
	Part int
}

type DayDirectory struct {
	Name  string
	Parts []DayPart
}

type TemplateData struct {
	Days        []DayDirectory
	CurrentPart DayPart
}

func main() {
	dirs, err := os.ReadDir(".")
	common.Check(err, "unable to read directory")

	td := TemplateData{}
	latest := time.Time{}

	for _, dir := range dirs {
		dirName := dir.Name()
		if dir.IsDir() && strings.HasPrefix(dirName, "day") {
			files, err := os.ReadDir(dirName)
			common.Check(err, "unable to read subdirectory %s", dirName)

			dayDir := DayDirectory{Name: dirName}
			day, _ := strconv.Atoi(dirName[3:])

			for _, file := range files {
				fileName := file.Name()
				var part int
				switch fileName {
				case "part_one.go":
					part = 1
				case "part_two.go":
					part = 2
				default:
					continue
				}

				dayPart := DayPart{Day: day, Part: part}
				dayDir.Parts = append(dayDir.Parts, dayPart)
				t, err := getModTime(file)
				common.Check(err, "unable to get mod time for %s/%s", dirName, fileName)
				if t.After(latest) {
					latest = t
					td.CurrentPart = dayPart
				}
			}

			td.Days = append(td.Days, dayDir)
		}
	}

	slices.SortFunc(td.Days, func(a, b DayDirectory) int {
		return cmp.Compare(a.Name, b.Name)
	})

	templateName := "run.tmpl"
	outputName := "run.go"

	templates := template.Must(template.ParseFiles(fmt.Sprintf("templates/%s", templateName)))
	outputFile, err := os.Create(outputName)
	common.Check(err, "unable to create %s", outputName)
	defer common.CloseFile(outputFile, outputName)

	err = templates.ExecuteTemplate(outputFile, templateName, td)
	common.Check(err, "unable to execute template %s", templateName)
}

func getModTime(de os.DirEntry) (time.Time, error) {
	fi, err := de.Info()
	if err != nil {
		return time.Time{}, fmt.Errorf("unable to get file information: %w", err)
	}

	return fi.ModTime(), nil
}
