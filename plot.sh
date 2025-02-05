#!/bin/bash

echo "Plotting results..."

# create on big file with all the results
cat data/bench*.txt > tmp/out.txt

# distribution results
go run ./cmd/benchplot/main.go -categories a -series s -table d -width 2400 -filter "10$" -dir ./images/perf/distribution/ tmp/out.txt
go run ./cmd/benchplot/main.go -categories a -series s -table d -width 2400 -filter "100$" -dir ./images/perf/distribution/ tmp/out.txt
go run ./cmd/benchplot/main.go -categories a -series s -table d -width 2400 -filter "1000$" -dir ./images/perf/distribution/ tmp/out.txt
go run ./cmd/benchplot/main.go -categories a -series s -table d -width 2400 -filter "10000$" -dir ./images/perf/distribution/ tmp/out.txt

go run ./cmd/benchplot/main.go -categories a -series s -table d -width 1920 -filter "100000$" -dir ./images/perf/distribution/ tmp/out.txt
go run ./cmd/benchplot/main.go -categories a -series s -table d -width 1920 -filter "1000000$" -dir ./images/perf/distribution/ tmp/out.txt

go run ./cmd/benchplot/main.go -categories a -series s -table d -width 1200 -filter "10000000$" -dir ./images/perf/distribution/ tmp/out.txt
go run ./cmd/benchplot/main.go -categories a -series s -table d -width 1200 -filter "100000000$" -dir ./images/perf/distribution/ tmp/out.txt
go run ./cmd/benchplot/main.go -categories a -series s -table d -width 1200 -filter "1000000000$" -dir ./images/perf/distribution/ tmp/out.txt
