GOFILES := $(filter-out run.go, $(wildcard **/*.go))
TEMPLATES := $(wildcard templates/*.tmpl)
GO ?= go
.PHONY: run runall test clean

.DEFAULT_GOAL := build

run.go: $(GOFILES) $(TEMPLATES)
	$(GO) generate

# run the most recently edited day
run: main.go run.go
	@$(GO) run .

runday%p1: main.go run.go
	@$(GO) run . -d $(shell echo $* | sed 's/^0*//') -p 1

runday%p2: main.go run.go
	@$(GO) run . -d $(shell echo $* | sed 's/^0*//') -p 2

runall: main.go run.go
	@$(GO) run . -a

test:
	$(GO) test -cover ./day*

clean:
	$(RM) run.go

day%p1:
	$(GO) run ./start -d $(shell echo $* | sed 's/^0*//')

day%p2:
	sed 's/PartOne/PartTwo/' day$(*)/part_one.go > day$(*)/part_two.go
	sed 's/PartOne/PartTwo/' day$(*)/part_one_test.go > day$(*)/part_two_test.go
