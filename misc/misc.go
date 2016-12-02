package main

import "fmt"

//Given a list of bytes a, each representing one byte of a larger integer (ie. {0x12, 0x34, 0x56, 0x78} represents the integer 0x12345678), and an integer b, find a % b.
//http://www.byte-by-byte.com/bigintmod/

func mod(n []byte, a int) int {
	m := 0
	for i := 0; i < len(n); i++ {
		m = m << 8
		m = m + int(n[i]&0xFF)
		m = m % a
	}
	return m
}

func TestModulo() {
	b := []byte{0x03, 0xED}
	a := 10
	fmt.Println(b)
	fmt.Println("modulo:", mod(b, a))
}

func fizzbuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}

}

func TestFizzbuzz() {
	fizzbuzz(16)
}

//Given an input amount of change x, write a function to determine the minimum number of coins required to make that amount of change.
// assume american coins. [1,5,10,25]
// change > coin[i], change < coin[i], change = coin[i], change = 0
// 27 = 25 * 1 + 1 *2 = 3
// 25 = 25 *1
// 16 = 10 *1 + 5 *1 + 1 *1 = 3
// 0 = 0
// sort coins in descending order
// find smallest element m > change
// count = change / m
// change = change % m
// reapeat

// coints = {25, 5, 10, 1}
func smallestChange(n int) int {
	coins := []int{25, 10, 5, 1}
	var count int
	for i := 0; i < len(coins) && n > 0; i++ {
		if coins[i] > n {
			continue
		}
		count = count + (n / coins[i])
		n = n % coins[i]
	}
	return count
}

func TestSmallestChange() {
	cases := []struct {
		in, out int
	}{
		{0, 0},
		{27, 3},
		{25, 1},
		{16, 3},
	}
	for _, c := range cases {
		fmt.Println("In:", c.in, "Expected:", c.out, "result:", smallestChange(c.in))
	}
}

//Given an integer, write a function to compute the number of ones in the binary representation of the number.
func sumOnes(n int) int {
	var sum int
	for n > 0 {
		sum = sum + n&1
		n = n >> 1
	}
	return sum
}

func TestSumOnes() {
	fmt.Println("sumones(7)=", sumOnes(7))
}

//Given 2 sorted arrays, A and B, where A is long enough to hold the contents of A and B, write a function to copy the contents of B into A without using any buffer or additional memory.
//http://www.byte-by-byte.com/mergearrays/
func MergeTwoSortedArray(A, B []int) []int {
	if len(A) == 0 {
		return A
	}
	var aMax int
	for i := 0; i < len(A); i++ {
		if A[i] == 0 {
			aMax = i - 1
			break
		}
	}
	aLast, bMax := len(A)-1, len(B)-1
	fmt.Println(aMax, aLast, bMax)
	for ; aMax >= 0 && bMax >= 0; aLast-- {
		if A[aMax] > B[bMax] {
			A[aLast] = A[aMax]
			aMax--
		} else if B[bMax] > A[aMax] {
			A[aLast] = B[bMax]
			bMax--
		} else {
			A[aLast] = B[aMax]
			aLast--
			A[aLast] = A[aMax]
			aMax--
			bMax--
		}
	}
	if aMax == 0 {
		for bMax > 0 {
			A[aLast] = B[bMax]
			aLast--
			bMax--
		}
	}

	return A
}

func TestMergeTwoSortedArray() {
	A := []int{1, 3, 0, 0, 0}
	B := []int{1, 4, 6}
	fmt.Println("A=", A, "B=", B)
	C := MergeTwoSortedArray(A, B)
	fmt.Println("Result=", C)
}

// Write function to find permutation of string
// permutation: n! .
// if dupliactes then permutations are n!/(a! * b! *...) where a , b are number of duplicates
// {a} => {a}
// {ab} => {ab,ba}
//{abc} =>{ abc, acb, bac, bca, cab, cba}
// logic: taken one char and append permutation of rest of chars to it. and then iterate through it.
// abc
//  /
// a, bc
//   /   \
// ab, c  ac, b
//     /      \
//    abc, nil acb, nil
func stringPermutation(s string) []string {
	result := make([]string, 0)
	permutation("", s, &result)
	return result
}

func permutation(prefix, suffix string, results *[]string) {
	if len(suffix) == 0 {
		*results = append(*results, prefix)
	}
	for i := 0; i < len(suffix); i++ {
		permutation(prefix+string(suffix[i]), suffix[:i]+suffix[i+1:], results)
	}
}

func TestPermuatation() {
	fmt.Println("permuatation of abc", stringPermutation("abc"))
}

func main() {
	//TestModulo()
	//TestFizzbuzz()
	//TestSmallestChange()
	//TestSumOnes()
	//TestMergeTwoSortedArray()
	//TestNStack()
	TestPermuatation()
}
