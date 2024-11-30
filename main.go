package main

//go:generate go run ./gen

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"slices"
	"strconv"
	"time"

	"aoc/common"
	"github.com/spf13/pflag"
)

var runAll bool
var selectedDay int
var selectedPart int
var profile string

func init() {
	pflag.BoolVarP(&runAll, "all", "a", false, "run all days")
	pflag.IntVarP(&selectedDay, "day", "d", 0, "run specific day")
	pflag.IntVarP(&selectedPart, "part", "p", 0, "run specific part")
	pflag.StringVar(&profile, "profile", "", "write cpu profile to `file`")
}

func main() {
	pflag.Parse()
	if profile != "" {
		f, err := os.Create(profile)
		common.Check(err, "could not create profile file %s", profile)
		_ = pprof.StartCPUProfile(f)
		defer common.CloseFile(f, profile)
	}

	if runAll {
		runAllDays()
	} else if selectedDay == 0 {
		runCurrent()
	} else {
		runDayPart(selectedDay, selectedPart)
	}
	pprof.StopCPUProfile()
}

type aocFunc func(io.Reader) any

type aocResult struct {
	result      string
	timeElapsed time.Duration
}

type aocRunnerInput struct {
	Func     aocFunc
	Filename string
	Day      int
	Part     int
}

func runAocPart(dayPart aocRunnerInput) aocResult {
	f, err := os.Open(dayPart.Filename)
	common.Check(err, "unable to open file %s", dayPart.Filename)
	defer common.CloseFile(f, dayPart.Filename)

	start := time.Now()
	r := dayPart.Func(f)
	duration := time.Since(start)

	var output string

	switch v := r.(type) {
	case int:
		output = strconv.Itoa(v)
	case string:
		output = v
	case fmt.Stringer:
		output = v.String()
	default:
		output = "unknown return value"
	}

	return aocResult{result: output, timeElapsed: duration}
}

func runAllDays() {
	var overallDuration time.Duration

	days := make([][]aocRunnerInput, 25)
	for _, dayPart := range dayParts {
		days[dayPart.Day-1] = append(days[dayPart.Day-1], dayPart)
	}

	for i, day := range days {
		if len(day) == 0 {
			continue
		}
		fmt.Printf("Day %d\n", i+1)
		for _, dayPart := range day {
			r := runAocPart(dayPart)
			result, elapsed := r.result, r.timeElapsed
			overallDuration += elapsed
			fmt.Printf("part %d: %s (duration: %s)\n", dayPart.Part, result, elapsed)
		}
		fmt.Println()
	}

	fmt.Printf("Overall time elapsed: %s\n", overallDuration)
}

func runDayPart(day int, part int) {
	partIndex := slices.IndexFunc(dayParts, func(daypart aocRunnerInput) bool {
		return daypart.Day == day && daypart.Part == part
	})

	if partIndex == -1 {
		log.Fatalf("Did not find a solution for day %d part %d\n", day, part)
	}

	dayPart := dayParts[partIndex]
	fmt.Printf("Day %d part %d\n", day, part)
	r := runAocPart(dayPart)
	fmt.Println(r.result)
	fmt.Printf("Time elapsed: %s\n", r.timeElapsed)
}
