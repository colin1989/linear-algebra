package main

import "math"

type Vector []float64

func Zero(n int) Vector {
	return make(Vector, n)
}

func (v Vector) Add(other Vector) Vector {
	var out Vector
	if len(v) != len(other) {
		return out
	}
	out = make(Vector, len(v))
	for i, v1 := range v {
		out[i] = v1 + other[i]
	}
	return out
}

func (v Vector) Sub(other Vector) Vector {
	var out Vector
	if len(v) != len(other) {
		return out
	}
	out = make(Vector, len(v))
	for i, v1 := range v {
		out[i] = v1 - other[i]
	}
	return out
}

func (v Vector) Mul(n float64) Vector {
	var out = make(Vector, len(v))
	for i, v1 := range v {
		out[i] = v1 * n
	}
	return out
}

func (v Vector) Norm() float64 {
	var n = float64(0)
	for _, v1 := range v {
		n += v1 * v1
	}
	return math.Sqrt(n)
}

func (v Vector) Normalize() Vector {
	n := v.Norm()
	var out = make(Vector, len(v))
	for i, v1 := range v {
		out[i] = v1 / n
	}
	return out
}

func (v Vector) Dot(other Vector) float64 {
	var n float64
	if len(v) != len(other) {
		return n
	}
	for i, _ := range v {
		n += v[i] * other[i]
	}
	return n
}
