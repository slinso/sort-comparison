package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/slinso/sortcomparison"
)

func generateSVG(data []int, distributionName string, filename string) error {
	var (
		height     = 800
		margin     = 80
		barWidth   = 2
		barGap     = 1
		plotHeight = height - 2*margin
		width      = len(data)*3 + 2*margin
	)

	// Calculate statistics
	min, max := sortcomparison.MinMaxValue(data)
	sum := 0
	for _, v := range data {
		sum += v
	}
	avg := sum / len(data)

	// Begin SVG with proper definitions
	svg := strings.Builder{}
	svg.WriteString(fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<svg width="%d" height="%d" xmlns="http://www.w3.org/2000/svg">
    <defs>
        <filter id="shadow" x="-20%%" y="-20%%" width="140%%" height="140%%">
            <feDropShadow dx="2" dy="2" stdDeviation="2" flood-opacity="0.1"/>
        </filter>
        <linearGradient id="barGradient" x1="0" x2="0" y1="0" y2="1">
            <stop offset="0%%" stop-color="#1976D2"/>
            <stop offset="100%%" stop-color="#2196F3"/>
        </linearGradient>
    </defs>
    <rect width="%d" height="%d" fill="#fafafa"/>`, width, height, width, height))

	// Draw grid (unchanged)
	for i := 0; i <= 10; i++ {
		y := margin + (plotHeight * i / 10)
		svg.WriteString(fmt.Sprintf(`
    <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="#e0e0e0" stroke-width="1"/>`,
			margin, y, width-margin, y))
	}

	// Draw bars with proper attributes
	for i, val := range data {
		x := margin + i*(barWidth+barGap)
		h := int(float64(val) * float64(plotHeight) / float64(max))
		y := height - margin - h
		svg.WriteString(fmt.Sprintf(`
    <rect x="%d" y="%d" width="%d" height="%d" fill="url(#barGradient)" filter="url(#shadow)"/>`,
			x, y, barWidth, h))
	}

	// Add labels and statistics
	svg.WriteString(fmt.Sprintf(`
        <text x="%d" y="40" font-family="Arial" font-size="24" text-anchor="middle" fill="#212121">
            %s Distribution
        </text>
        <text x="%d" y="%d" font-family="Arial" font-size="14" text-anchor="middle" fill="#757575">
            Index
        </text>
        <text x="30" y="%d" font-family="Arial" font-size="14" text-anchor="middle" 
            transform="rotate(-90,30,%d)" fill="#757575">
            Value
        </text>`,
		width/2, distributionName, width/2, height-10, height/2, height/2))

	// Add legend
	yStart := 10
	xStart := width - 300
	svg.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" width="150" height="90" fill="white" rx="5" filter="url(#shadow)"/>`, xStart, yStart))
	svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" font-family="Arial" font-size="12" fill="#757575">Minimum: %d</text>`, xStart+20, yStart+25, min))
	svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" font-family="Arial" font-size="12" fill="#757575">Maximum: %d</text>`, xStart+20, yStart+45, max))
	svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" font-family="Arial" font-size="12" fill="#757575">Average: %d</text>`, xStart+20, yStart+65, avg))

	svg.WriteString("</svg>")

	return os.WriteFile(filename, []byte(svg.String()), 0o644)
}

type DataGenerator struct {
	Name string
	Data func(int) []int
}

func main() {
	dataGenerators := []DataGenerator{
		{"Random", sortcomparison.GenerateRandomInts},
		{"RandomMaxN", sortcomparison.GenerateRandomIntsMaxN},
		{"AllZero", sortcomparison.GenerateAllZero},
		{"Sorted", sortcomparison.GenerateSortedInts},
		{"Rotated", sortcomparison.GenerateRotated},
		{"Reversed", sortcomparison.GenerateReversedInts},
		{"Mountain", sortcomparison.GenerateMountain},
		{"Valley", sortcomparison.GenerateValley},
		{"Plateau", sortcomparison.GeneratePlateau},
		{"SmallHills", sortcomparison.GenerateSmallHills},
		{"RandomMod8", sortcomparison.GenerateRandomMod8},
		{"RepeatedMod8", sortcomparison.GenerateRepeatedMod8},
		{"RandomMod16", sortcomparison.GenerateRandomMod16},
		{"RepeatedMod16", sortcomparison.GenerateRepeatedMod16},
		{"BackToFront", sortcomparison.GenerateBackToFront},
		{"FrontToBack", sortcomparison.GenerateFrontToBack},
		{"MiddleToBack", sortcomparison.GenerateMiddleToBack},
		{"PushMiddle", sortcomparison.GeneratePushMiddle},
		{"NearlySorted", sortcomparison.GenerateNearlySorted},
		{"NearlyReversed", sortcomparison.GenerateNearlyReversed},
	}

	for _, gen := range dataGenerators {
		data := gen.Data(300)
		filename := fmt.Sprintf("./images/%s.svg", strings.ReplaceAll(gen.Name, " ", ""))

		err := generateSVG(data, gen.Name, filename)
		if err != nil {
			panic(err)
		}
	}
}
