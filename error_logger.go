package BetterErrors

import (
	"fmt"
	"log"
	"runtime"

	"github.com/logrusorgru/aurora"
)

type ErrorLogger struct {
	IncludeSensitiveInfo bool
	ExitMessage          string
}

func NewErrorLogger() *ErrorLogger {
	return &ErrorLogger{
		IncludeSensitiveInfo: true,
		ExitMessage:          "↑ Something went wrong! ↑",
	}
}

func (logger *ErrorLogger) SetIncludeSensitiveInfo(includeSensitiveInfo bool) {
	logger.IncludeSensitiveInfo = includeSensitiveInfo
}

func (logger *ErrorLogger) SetExitMessage(exitMessage string) {
	logger.ExitMessage = exitMessage
}

func (logger *ErrorLogger) LogError(err error) {
	if err != nil {
		pc, filename, line, _ := runtime.Caller(1)

		funcName := runtime.FuncForPC(pc).Name()
		sourceInfo := fmt.Sprintf("[%s:%d]", filename, line)

		os := runtime.GOOS
		goVersion := runtime.Version()

		if logger.IncludeSensitiveInfo {
			fmt.Printf("╭─ [%s] in %s %s %s\n", aurora.Red(aurora.Bold("error")), aurora.Blue(funcName), aurora.Green(sourceInfo), aurora.Cyan("↙"))
		} else {
			fmt.Printf("╭─ %s Error Information [%s] %s\n", aurora.Magenta("─"), aurora.Blue(funcName), aurora.Magenta("─"))
		}
		fmt.Printf("│   %v\n", err)
		fmt.Printf("│\n")
		fmt.Printf("│  %s Something went wrong! \n", aurora.Cyan("↑"))
		fmt.Printf("│  %s Please check the error message above for details.\n", aurora.Cyan("→"))
		if logger.IncludeSensitiveInfo {
			fmt.Printf("│  %s User OS: %s\n", aurora.Cyan("→"), os)
			fmt.Printf("│  %s Go Version: %s\n", aurora.Cyan("→"), goVersion)
		} else {
			fmt.Printf("│  %s User OS and Go Version information omitted for privacy.\n", aurora.Cyan("→"))
		}
		fmt.Println("╰───────────────────────────────────")

		//log.Fatalf(logger.ExitMessage)
	}
}
