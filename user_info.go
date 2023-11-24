package BetterErrors

import (
	"fmt"
	"runtime"

	"github.com/logrusorgru/aurora"
)

func PrintInfo() {
	fmt.Printf("╭─ %s Application Information %s\n", aurora.Magenta("─"), aurora.Magenta("─"))
	fmt.Printf("│  %s OS: %s\n", aurora.Cyan("→"), runtime.GOOS)
	fmt.Printf("│  %s Go Version: %s\n", aurora.Cyan("→"), runtime.Version())
	fmt.Printf("│  %s Number of CPUs: %d\n", aurora.Cyan("→"), runtime.NumCPU())
	fmt.Printf("│  %s GOMAXPROCS: %d\n", aurora.Cyan("→"), runtime.GOMAXPROCS(0))
	fmt.Println("╰───────────────────────────────────")
}
