package main

import (
	"fmt"
)

func main() {
	vec := Vector{5, 2}
	fmt.Printf("%v \n", vec)
	vec1 := Vector{3, 1}

	out := vec.Add(vec1)
	fmt.Printf("vec[%v] + vec1[%v] = [%v] \n", vec, vec1, out)

	out = vec.Sub(vec1)
	fmt.Printf("vec[%v] - vec1[%v] = [%v] \n", vec, vec1, out)

	out = vec.Mul(2)
	fmt.Printf("vec[%v] * [%v] = [%v] \n", vec, 2, out)

	zero := Zero(2)
	fmt.Printf("zero verctor : [%v] \n", zero)

	fmt.Printf("vec[%v] + zero[%v] = [%v] \n", vec, zero, vec.Add(zero))

	fmt.Printf("vec[%v] norm = [%v] normalize = [%v] \n", vec, vec.Norm(), vec.Normalize())

	fmt.Printf("vec[%v] dot [%v] = [%v] \n", vec, vec1, vec.Dot(vec1))
}
