#!/bin/bash

# Check if benchstat is installed
if ! command -v benchstat &> /dev/null; then
    echo "Installing benchstat..."
    go install golang.org/x/perf/cmd/benchstat@latest
fi

echo "Running benchmarks..."
# Create temp file for benchmark results
TEMP_FILE=$(mktemp)

# Run benchmarks and save to temp file
go test -bench "///[10,100]$" -benchtime 1x -cpu 1 -count=10 > "$TEMP_FILE"

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