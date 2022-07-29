package generics

import "fmt"

//non-generic
func sumInt(m map[string]int64) int64 {
	var sum int64
	for _, i := range m {
		sum += i
	}
	return sum
}

func sumFloat(m map[string]float64) float64 {
	var sum float64
	for _, i := range m {
		sum += i
	}
	return sum
}

//generic

type Type interface {
	int64 | float64
}

func sum[C comparable, T Type](m map[C]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

func main() {

	ints := map[string]int64{
		"1": 34,
		"2": 12,
	}

	floats := map[string]float64{
		"1": 35.98,
		"2": 11.97,
	}

	fmt.Printf("%v , %v\n", sumInt(ints), sumFloat(floats))
	fmt.Printf("%v , %v\n", sum(ints), sum(floats))
}
