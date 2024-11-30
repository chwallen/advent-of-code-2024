# Advent of Code 2024 solutions in Go

These are my solutions for [Advent of Code 2024](https://adventofcode.com/2024)
written in Go. Each day is in a separate folder with one file per part. Shared
code between parts of the same day is in the file _shared.go_ in that day's
folder. Things that are shared between days are found in the _common_ package.

## Instructions

The repository contains a _Makefile_ with a few helpful commands.
- `make rundayXXpY` will run a specific day
- `make runall` will run all days 
- `make run` will run the most recently edited day. This uses the filesystem
  _atime_. If you have disabled it via _noatime_, it may not work properly.
- `make test` will run all tests with code coverage

## Using as a template repository

The repository contains some additional code and Makefile commands to start new
days while the event is ongoing. To start a new day, for example day 1, you can
use the command `make day01p1`. This will create a new folder called _day01_ and
create the part one implementation and test file. Then, download your input from
the website and place it in that folder as _input.txt_. The `PartOne` function
will receive an `io.Reader` as argument (the input) and is expected to return
the solution of any type. When you are done, use `make run` or `make rundayXXp1`
to execute the code and generate the answer.

Once part 1 is successfully submitted on the website, you can start part 2 with
`make day01p2`. This will copy the code from part 1 into _part_two.go_.
