package main

import "fmt"

// SumInts adds together the values of m and returns result (int64).
func SumInts(m []int64) int64 {
	var sum int64
	for _, v := range m {
		sum += v
	}
	return sum
}

// SumFloats adds together the values of m and returns result (float64).
func SumFloats(m []float64) float64 {
	var sum float64
	for _, v := range m {
		sum += v
	}
	return sum
}

// SumFloatsGenericSlice adds together the values of m and returns result
// (int64|float64).
// T type parameter a constraint that is a union of two types: int64 and float64.
// Using | specifies a union of the two types, meaning that this constraint
// allows either type. Either type will be permitted by the compiler as an
// argument in the calling code.
func SumFloatsGenericSlices[T int64 | float64](m []T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

// SumFloatsGenericMap adds together the values of m and returns result (int64|float64).
// K type parameter the type constraint comparable. Intended specifically for
// cases like these, the comparable constraint is predeclared in Go. It allows
// any type whose values may be used as an operand of the comparison operators
// == and !=. Go requires that map keys be comparable. So declaring K as comparable
// is necessary so you can use K as the key in the map variable.
func SumFloatsGenericMap[K comparable, T int64 | float64](m map[K]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

// Number interface type constraint.
type Number interface {
    int64 | float64
}

// SumFloatsGenericMapInterface adds together the values of m and returns result 
// (int64|float64). Type constraint can be declared as interface. The constraint 
// allows any type implementing the interface. In this case, the interface is a 
// union of two types: int64 and float64.
func SumFloatsGenericMapInterface[K comparable, T Number](m map[K]T) T {
    var sum T
    for _, v := range m {
        sum += v
    }
    return sum
}

func main() {
	m_int := []int64{1, 2, 3}
	m_float := []float64{1.1, 2.2, 3.3}

	// Print sums of non-generic functions.
    fmt.Println("Print sums of non-generic functions.")
	fmt.Println("SumInts(m_int):", SumInts(m_int))
    fmt.Println("SumFloats(m_float):", SumFloats(m_float))

	sum_int := SumFloatsGenericSlices[int64](m_int)       // type arguments can be omitted in function call without explicit type arguments
	sum_float := SumFloatsGenericSlices[float64](m_float) // 6.6
	// Print sums of generic functions.
    fmt.Println()
    fmt.Println("Print sums of generic functions (without interface).")
    fmt.Println("SumFloatsGenericSlices[int64](m_int):", sum_int)
    fmt.Println("SumFloatsGenericSlices[float64](m_float):", sum_float)

	// Declare a map of type map[int64]float64.
	m := map[int64]float64{
		1: 1.1,
		2: 2.2,
		3: 3.3,
	}
	// Print sum of generic function.
	fmt.Println("SumFloatsGenericMap(m):", SumFloatsGenericMap(m)) 

    // Print sum of generic function with interface type constraint.
    fmt.Println()
    fmt.Println("Print sums of generic functions (with interface type constraint).")
    fmt.Println("SumFloatsGenericMapInterface(m):", SumFloatsGenericMapInterface(m)) 
}
