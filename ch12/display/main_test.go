package display

import (
	"bytes"
	"os"
	"testing"
)

func TestDisplay(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		stdout := new(bytes.Buffer)
		a := "a"
		Display("a", a, stdout)
		expect := "a = \"a\"\n"
		if stdout.String() != expect {
			t.Errorf("actual: %sexpect: %s", stdout.String(), expect)
		}
	})
}

func Example_arrayMap() {
	mf := map[[2]int]int{
		{2, 1}: 1,
		{3, 1}: 1,
	}
	Display("magnetic field", mf, os.Stdout)
	// Output:
	// Unordered output:
	// magnetic field[[2]int{2, 1}] = 1
	// magnetic field[[2]int{3, 1}] = 1
}

type coord struct {
	x int
	y int
}

func Example_structMap() {
	mf := map[coord]int{
		{x: 2, y: 1}: 1,
		{x: 3, y: 1}: 1,
	}
	Display("magnetic field", mf, os.Stdout)
	// Output:
	// Unordered output:
	// magnetic field[display.coord{x: 2, y: 1}] = 1
	// magnetic field[display.coord{x: 3, y: 1}] = 1
}
