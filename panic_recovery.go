package BetterErrors

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

func RecoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Printf("╭─ [%s] Recovered from panic:\n", aurora.Yellow(aurora.Bold("recovery")))
		fmt.Printf("│   %v\n", r)
		fmt.Println("╰───────────────────────────────────")
	}
}
