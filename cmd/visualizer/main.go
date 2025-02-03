package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/slinso/sortcomparison"
)

func generateSVG(data []int, distributionName, filename string) error {
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

	svg := strings.Builder{}
	// Start SVG with proper namespace and styling definitions, including a radial gradient for background.
	svg.WriteString(fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<svg width="%d" height="%d" viewBox="0 0 %d %d" xmlns="http://www.w3.org/2000/svg">
  <defs>
    <radialGradient id="bgGradient" cx="50%%" cy="50%%" r="75%%">
      <stop offset="0%%" stop-color="#ffffff"/>
      <stop offset="100%%" stop-color="#e0e0e0"/>
    </radialGradient>
    <filter id="shadow" x="-20%%" y="-20%%" width="140%%" height="140%%">
      <feDropShadow dx="3" dy="3" stdDeviation="3" flood-color="#999" flood-opacity="0.4"/>
    </filter>
    <linearGradient id="barGradient" x1="0" x2="0" y1="0" y2="1">
      <stop offset="0%%" stop-color="#1976D2"/>
      <stop offset="100%%" stop-color="#2196F3"/>
    </linearGradient>
    <style type="text/css">
      <![CDATA[
        .axis { fill: none; stroke: #424242; stroke-width: 2; }
        .grid { stroke: #cfcfcf; stroke-width: 1; }
        .label { font-family: "Segoe UI", Arial, sans-serif; fill: #424242; }
        .title { font-size: 28px; font-weight: bold; }
        .tick { font-size: 12px; }
      ]]>
    </style>
  </defs>
  <rect width="100%%" height="100%%" fill="url(#bgGradient)" />
`, width, height, width, height))

	// Draw grid lines and y-axis ticks
	for i := 0; i <= 10; i++ {
		y := margin + (plotHeight * i / 10)
		svg.WriteString(fmt.Sprintf(`
  <line x1="%d" y1="%d" x2="%d" y2="%d" class="grid" />`, margin, y, width-margin, y))
		// Tick label on y-axis
		value := max - (max * i / 10)
		svg.WriteString(fmt.Sprintf(`
  <text x="%d" y="%d" class="label tick" text-anchor="end" dy="4">%d</text>`, margin-10, y, value))
	}

	// Draw x-axis ticks (only a few to avoid clutter)
	numberOfXTicks := 10
	dataLen := len(data)
	for i := 0; i <= numberOfXTicks; i++ {
		idx := i * dataLen / numberOfXTicks
		x := margin + idx*(barWidth+barGap)
		svg.WriteString(fmt.Sprintf(`
  <line x1="%d" y1="%d" x2="%d" y2="%d" class="grid" />`, x, height-margin, x, height-margin+5))
		svg.WriteString(fmt.Sprintf(`
  <text x="%d" y="%d" class="label tick" text-anchor="middle" dy="15">%d</text>`, x, height-margin+5, idx))
	}

	// Draw axes
	svg.WriteString(fmt.Sprintf(`
  <line x1="%d" y1="%d" x2="%d" y2="%d" class="axis"/>`, margin, height-margin, width-margin, height-margin))
	svg.WriteString(fmt.Sprintf(`
  <line x1="%d" y1="%d" x2="%d" y2="%d" class="axis"/>`, margin, margin, margin, height-margin))

	// Draw bars with modern gradient and drop shadow
	for i, val := range data {
		x := margin + i*(barWidth+barGap)
		h := int(float64(val) * float64(plotHeight) / float64(max))
		y := height - margin - h
		svg.WriteString(fmt.Sprintf(`
  <rect x="%d" y="%d" width="%d" height="%d" fill="url(#barGradient)" filter="url(#shadow)" />`,
			x, y, barWidth, h))
	}

	// Add Title and axis labels
	svg.WriteString(fmt.Sprintf(`
  <text x="%d" y="40" class="label title" text-anchor="middle">%s Distribution</text>
  <text x="%d" y="%d" class="label" text-anchor="middle">Index</text>
  <text x="30" y="%d" class="label" transform="rotate(-90,30,%d)" text-anchor="middle">Value</text>`,
		width/2, distributionName, width/2, height-20, height/2, height/2))

	// Add a legend with statistics
	yStart := 10
	xStart := width - 300
	svg.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" rx="8" stroke="#b0b0b0" width="150" height="90" fill="white" filter="url(#shadow)"/>`, xStart, yStart))
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
