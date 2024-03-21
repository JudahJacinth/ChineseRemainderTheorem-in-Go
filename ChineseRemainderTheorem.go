package main

import (
	"fmt"
	"os"
)

func GCD(A, B int32) int32 {
	if B == 0 {
		return A
	}
	return int32(GCD(B, A%B))
}

func Coprime(M, m int32) bool {
	var A int32 = M
	var B int32 = m
	var gcd int32 = GCD(A, B)

	if gcd == 1 {
		return true
	}
	return false
}

func ExtendedEuclideanAlgorithm(A, B, T1, T2 int32) int32 {
	if B == 0 {
		// fmt.Printf("Quotient: %d, Remainder: %d, T: %d\n", A/B, A%B, T1)
		if T1 < 0 {
			return A - T1
		}
		return T1
	}
	Q := A / B
	R := A % B
	T := T1 - T2*Q
	// fmt.Printf("Quotient: %d, A: %d, B: %d,  Remainder: %d, T1: %d, T2: %d, T: %d\n", Q, A, B, R, T1, T2, T)
	return ExtendedEuclideanAlgorithm(B, R, T2, T)
}

var M1Inverse, M2Inverse, M3Inverse int32
var MInverse_values []int32

func main() {
	var dividendOne, dividendTwo, dividendThree int32
	var divisorOne, divisorTwo, divisorThree int32

	dividendOne = 3
	dividendTwo = 5
	dividendThree = 2
	divisorOne = 4
	divisorTwo = 6
	divisorThree = 5

	// var n int32
	// fmt.Println("Enter the number of equations: ")
	// fmt.Scan(&n)

	// var a_values = [n]int32{}
	// var m_values = [n]int32{}
	// for i := 0; i < int(n); i++ {
	// 	fmt.Printf("a%d: ", i+1)
	// 	fmt.Scan(&a_values[i])
	// 	fmt.Printf("m%d: ", i+1)
	// 	fmt.Scan(&m_values[i])
	// }

	var a1 int32 = dividendOne
	var a2 int32 = dividendTwo
	var a3 int32 = dividendThree
	var a_values = [3]int32{a1, a2, a3}

	var m1 int32 = divisorOne
	var m2 int32 = divisorTwo
	var m3 int32 = divisorThree
	var m_values = [3]int32{m1, m2, m3}

	var M int32 = m1 * m2 * m3

	// fmt.Println("M is ", M)

	var M1 int32 = M / m1
	var M2 int32 = M / m2
	var M3 int32 = M / m3
	var M_values = [3]int32{M1, M2, M3}

	// fmt.Println("M1, M2, M3 = ", M_values)

	// fmt.Println(Coprime(35, 3))

	count := 0
	for i := 0; i < len(M_values); i++ {
		if Coprime(M_values[i], m_values[i]) {
			// fmt.Println(ExtendedEuclideanAlgorithm(M_values[i], m_values[i], 1, 0))
			MInverse_values = append(MInverse_values, ExtendedEuclideanAlgorithm(M_values[i], m_values[i], 1, 0))
			count++
			if count == 3 {
				break
			}
		} else {
			fmt.Println("Because ", M_values[i], " and ", m_values[i], " are not coprimes, We cannot calculate the remainder using the Chinese Remainder Theorem.")
			os.Exit(0)
		}
	}

	var X int32 = 0
	for i := 0; i < 3; i++ {
		// fmt.Println("Set ", i, "a: ", a_values[i], "M: ", M_values[i], "MInverse: ", MInverse_values[i])
		X += a_values[i] * M_values[i] * MInverse_values[i]
	}
	// fmt.Println("X is ", X)
	fmt.Println("The remainder of the system of equations is: ", X%M)

}
