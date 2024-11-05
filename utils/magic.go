package utils

import (
	"math"
)

func MagicSum(n int) int {
	return n + n
}

func MagicPow(n int) float64 {
	return math.Pow(float64(n), float64(n))
}

func MagicOdd(n int) bool {
	return n%2 != 0
}

func MagicGrade(n int) string {
	switch n {
	case 0:
		return "zero"
	case 1:
		return "Bad"
	case 2:
		return "Ok"
	case 3:
		return "Nice"
	case 4:
		return "Awesome"
	case 5:
		return "Exceptional"
	default:
		return "Unknown"
	}
}

func MagicName(n int) []string {
	name := "dimas"
	names := make([]string, n)
	for i := 0; i < n; i++ {
		names[i] = name
	}
	return names
}

func MagicTria(n int) []int {

	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = i + 1
	}

	return result
}

type MagicNumber struct {
	num int
}

func (m *MagicNumber) Multiply(n int) {
	m.num *= n
}

func MagicChange(n *int) {
	*n *= 2
}
