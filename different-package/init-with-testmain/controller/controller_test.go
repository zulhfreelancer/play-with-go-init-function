package controller

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	fmt.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
