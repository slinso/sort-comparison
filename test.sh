#!/bin/bash

# Check if benchstat is installed
if ! command -v benchstat &>/dev/null; then
    echo "Installing benchstat..."
    go install golang.org/x/perf/cmd/benchstat@latest
fi

echo "Running benchmarks..."
# Create temp file for benchmark results
TEMP_FILE=$(mktemp)

# Run benchmarks and save to temp file
go test -bench "///[10,100]$" -benchtime 1x -cpu 1 -count=10 >"$TEMP_FILE"
benchstat -col /dist -format csv out.txt | rg "Sort/" | head -50 >test.out
mlr --from test.out --icsv --opprint --hi --omd --ofmt %.9f cut -f 1,2,4,8,12,16,20,24,28,32,36,40,44,48,52,56,60,64,68,72,76 then sort -f 2 >random.md

# Check if benchmark succeeded
if [ $? -eq 0 ]; then
    echo "Processing results with benchstat..."
    benchstat -row /dataset,/size -col /algo "$TEMP_FILE"
else
    echo "Benchmark failed!"
    cat "$TEMP_FILE"
fi

# Cleanup
rm "$TEMP_FILE"
