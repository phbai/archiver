package util

import (
	"time"
	"fmt"
)

func Spinner(delay time.Duration, message string) {
	for {
			for _, r := range `-\|/` {
					fmt.Printf("\r%s  %c", message, r)
					time.Sleep(delay)
			}
	}
}