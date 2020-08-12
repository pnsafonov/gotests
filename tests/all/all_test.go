package all

import (
	"fmt"
	"math"
	"testing"
)

func TestFloor(t *testing.T) {
	a1 := math.Floor(1.01)
	fmt.Printf("a1 = %f\n", a1)

	a2 := math.Floor(1.49)
	fmt.Printf("a2 = %f\n", a2)

	a3 := math.Floor(1.51)
	fmt.Printf("a3 = %f\n", a3)

	a4 := math.Floor(1.99)
	fmt.Printf("a4 = %f\n", a4)

	fmt.Println("done")
}

func TestCeil(t *testing.T) {
	a1 := math.Ceil(1.01)
	fmt.Printf("a1 = %f\n", a1)

	a2 := math.Ceil(1.49)
	fmt.Printf("a2 = %f\n", a2)

	a3 := math.Ceil(1.51)
	fmt.Printf("a3 = %f\n", a3)

	a4 := math.Ceil(1.99)
	fmt.Printf("a4 = %f\n", a4)

	fmt.Println("done")
}