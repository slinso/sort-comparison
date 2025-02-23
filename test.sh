#!/bin/bash

# Check if benchstat is installed
if ! command -v benchstat &>/dev/null; then
    echo "Installing benchstat..."
    go install golang.org/x/perf/cmd/benchstat@latest
fi

echo "Running benchmarks..."

algos=($(go test -bench "///10$" -cpu 1 -count 5 -benchtime 1x | rg BenchmarkSort | sed -n 's/.*algo=\([^/]*\).*/\1/p' | sort | uniq))

# for each algo run the benchmark
for algo in "${algos[@]}"
do
    go test -bench "//${algo}/" -cpu 1 -count 5 -benchtime 100ms | tee data/bench-${algo}.txt
done

# testing a single size
#go test -bench "///10$" -cpu 1 -count 5 -benchtime 10ms | tee data/size-10.txt


# Run benchmarks and save to temp file
# go test -bench "///[10,100]$" -benchtime 1x -cpu 1 -count=10 >"$TEMP_FILE"
# benchstat -col /dist -format csv out.txt | rg "Sort/" | head -50 >test.out
# mlr --from test.out --icsv --opprint --hi --omd --ofmt %.9f cut -f 1,2,4,8,12,16,20,24,28,32,36,40,44,48,52,56,60,64,68,72,76 then sort -f 2 >random.md
# go run ./cmd/benchplot/main.go -categories d -series s -table a -filter "100$" -width 800 data/all-to-10k.txt

# distribution results
# go run ./cmd/benchplot/main.go -categories a -series s -table d -width 1920 -filter "10$" -dir ./images/perf/distribution/ data/all-to-10k.txt
# go run ./cmd/benchplot/main.go -categories a -series s -table d -width 1920 -filter "100$" -dir ./images/perf/distribution/ data/all-to-10k.txt
# go run ./cmd/benchplot/main.go -categories a -series s -table d -width 1920 -filter "1000$" -dir ./images/perf/distribution/ data/all-to-10k.txt
# go run ./cmd/benchplot/main.go -categories a -series s -table d -width 1920 -filter "10000$" -dir ./images/perf/distribution/ data/size-10000.txt

# Cleanup
rm "$TEMP_FILE"
