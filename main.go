package main

import (
	"github.com/litesoft-go/chatgpt4generated-go-calculator/v2/calculator"
	"github.com/litesoft-go/chatgpt4generated-go-formatter/v2/formatter"
	"github.com/litesoft-go/chatgpt4generated-go-printer/v2/printer"
)

func main() {
	sum := calculator.Add(5, 10)
	formattedSum := formatter.FormatFloat(float64(sum))
	printer.PrintString("5 + 10 = ")
	printer.PrintStringln(formattedSum)
}
