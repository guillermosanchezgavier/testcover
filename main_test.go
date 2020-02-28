package main

import (
	"testing"
)

// func TestMain(m *testing.M) {
// 	// call flag.Parse() here if TestMain uses flags
// 	rc := m.Run()

// 	// rc 0 means we've passed,
// 	// and CoverMode will be non empty if run with -cover
// 	if rc == 0 && testing.CoverMode() != "" {
// 		c := testing.Coverage()
// 		if c < 0.8 {
// 			fmt.Println("Tests passed but coverage failed at", c)
// 			rc = -1
// 		}
// 	}
// 	os.Exit(rc)
// }
func Test_add(t *testing.T) {
	if got := add(1, 1); got != 2 {
		t.Errorf("add() = %v, want %v", got, 2)
	}
}
