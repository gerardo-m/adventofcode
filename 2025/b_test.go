package main

import (
	"fmt"
	"testing"
)

func BenchmarkAll(b *testing.B) {
	n := len(problems)
	for i := range n {
		name1 := fmt.Sprintf("P=p%dp%d", i, 1)
		name2 := fmt.Sprintf("P=p%dp%d", i, 2)
		b.Run(name1, func(b *testing.B) { runP(i, 1) })
		b.Run(name2, func(b *testing.B) { runP(i, 2) })
	}
}
