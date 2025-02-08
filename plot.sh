#!/bin/bash

echo "Plotting results..."

# create on big file with all the results
algos=($(go test -bench "///10$" -cpu 1 -count 5 -benchtime 1x | rg BenchmarkSort | sed -n 's/.*algo=\([^/]*\).*/\1/p' | sort | uniq))

cat data/bench-slices.Sort.txt > tmp/out.txt
for algo in "${algos[@]}"
do
    cat data/bench-${algo}.txt >> tmp/out.txt
done

# benchstat compare everything
benchstat -row /dataset,/size -col /algo tmp/out.txt > data/benchstat-result.txt

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

# algo results
go run ./cmd/benchplot/main.go -categories d -series s -table a -width 1000 -filter "10$" -dir ./images/perf/algo/ tmp/out.txt
go run ./cmd/benchplot/main.go -categories d -series s -table a -width 1000 -filter "100$" -dir ./images/perf/algo/ tmp/out.txt
go run ./cmd/benchplot/main.go -categories d -series s -table a -width 1000 -filter "1000$" -dir ./images/perf/algo/ tmp/out.txt
go run ./cmd/benchplot/main.go -categories d -series s -table a -width 1000 -filter "10000$" -dir ./images/perf/algo/ tmp/out.txt

go run ./cmd/benchplot/main.go -categories d -series s -table a -width 1000 -filter "100000$" -dir ./images/perf/algo/ tmp/out.txt
go run ./cmd/benchplot/main.go -categories d -series s -table a -width 1000 -filter "1000000$" -dir ./images/perf/algo/ tmp/out.txt

go run ./cmd/benchplot/main.go -categories d -series s -table a -width 1000 -filter "10000000$" -dir ./images/perf/algo/ tmp/out.txt
go run ./cmd/benchplot/main.go -categories d -series s -table a -width 1000 -filter "100000000$" -dir ./images/perf/algo/ tmp/out.txt
go run ./cmd/benchplot/main.go -categories d -series s -table a -width 1000 -filter "1000000000$" -dir ./images/perf/algo/ tmp/out.txt

# algo avg results
go run ./cmd/benchplot/main.go -categories a -series d -width 1920 -filter "10$" -dir ./images/perf/algo/ -avg tmp/out.txt
go run ./cmd/benchplot/main.go -categories a -series d -width 1920 -filter "100$" -dir ./images/perf/algo/ -avg tmp/out.txt
go run ./cmd/benchplot/main.go -categories a -series d -width 1920 -filter "1000$" -dir ./images/perf/algo/ -avg tmp/out.txt
go run ./cmd/benchplot/main.go -categories a -series d -width 1920 -filter "10000$" -dir ./images/perf/algo/ -avg tmp/out.txt

go run ./cmd/benchplot/main.go -categories a -series d -width 1920 -filter "100000$" -dir ./images/perf/algo/ -avg tmp/out.txt
go run ./cmd/benchplot/main.go -categories a -series d -width 1920 -filter "1000000$" -dir ./images/perf/algo/ -avg tmp/out.txt

go run ./cmd/benchplot/main.go -categories a -series d -width 1920 -filter "10000000$" -dir ./images/perf/algo/ -avg tmp/out.txt
go run ./cmd/benchplot/main.go -categories a -series d -width 1920 -filter "100000000$" -dir ./images/perf/algo/ -avg tmp/out.txt
go run ./cmd/benchplot/main.go -categories a -series d -width 1920 -filter "1000000000$" -dir ./images/perf/algo/ -avg tmp/out.txt
