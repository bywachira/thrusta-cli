package logger

import "fmt"

type FmtOperator func(output string) (n int, err error)

type Console struct {
	Success FmtOperator
	Error   FmtOperator
	Info    FmtOperator
}

var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorYellow = "\033[33m"
var colorBlue = "\033[34m"
var colorPurple = "\033[35m"
var colorCyan = "\033[36m"
var colorWhite = "\033[37m"

var console Console = Console{
	Success: func(output string) (n int, err error) { return fmt.Println(string(colorGreen), output) },
	Error:   func(output string) (n int, err error) { return fmt.Println(string(colorRed), output) },
	Info:    func(output string) (n int, err error) { return fmt.Println(string(colorBlue), output) },
}
