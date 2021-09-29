package problems

import (
	"fmt"
)

type Problem struct{}

func (p Problem) Default() {
    fmt.Println("Default called.")
}
