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

//func TestSumNullable1(t *testing.T) {
//	a1 := sql.NullInt64{ Valid: false, Int64: 1}
//	a2 := sql.NullInt64{ Valid: false, Int64: 2}
//	r1 := a1 + a2
//}

type keyIntStr struct {
	id 	 int
	name string
}

func TestMap1(t *testing.T) {
	m1 := make(map[keyIntStr]int)
	m1[keyIntStr{id: 1, name: "11"}] = 1
	m1[keyIntStr{id: 1, name: "11"}] = 2
	m1[keyIntStr{id: 1, name: "11"}] = 3
	fmt.Println("done")
}
