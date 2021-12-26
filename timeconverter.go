package main

import (
	"fmt"
	"syscall/js"

	"github.com/airthee/timeconverter"
)

func main() {
	fmt.Println("Exporting hoursToSeconds")
	js.Global().Set("timeconverter", js.ValueOf(map[string]interface{}{
		"hoursToSeconds":   js.FuncOf(HoursToSecondsFunction),
		"factorizeSeconds": js.FuncOf(FactorizeFunction),
	}))
	select {}
}

func HoursToSecondsFunction(this js.Value, p []js.Value) interface{} {
	hours := p[0].Int()
	seconds := timeconverter.HoursToSeconds(hours)
	return js.ValueOf(seconds)
}

func FactorizeFunction(this js.Value, p []js.Value) interface{} {
	seconds := p[0].Int()
	factorized := timeconverter.Factorize(seconds)
	result := map[string]interface{}{
		"weeks":   factorized.Weeks,
		"days":    factorized.Days,
		"hours":   factorized.Hours,
		"minutes": factorized.Minutes,
		"seconds": factorized.Seconds,
	}
	return js.ValueOf(result)
}
