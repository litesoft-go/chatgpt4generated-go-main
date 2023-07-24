package main

import (
	"github.com/litesoft-go/chatgpt4generated-go-calculator"
	"github.com/litesoft-go/chatgpt4generated-go-formatter"
	"github.com/litesoft-go/chatgpt4generated-go-printer"
)

func main() {
	sum := calculator.Add(5, 10)
	formattedSum := formatter.FormatFloat(float64(sum))
	printer.PrintString(formattedSum)
}
