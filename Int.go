package one4

import (
	"fmt"
	"math"
)

var (
	base = getBase()
)

type Int struct {
	// Higher index -> more siginificant digit
	mant []uint
	neg  bool
}

func MakeInt(num int) Int {
	var newInt Int

	if num < 0 {
		newInt.neg = true
		num *= -1
	}
	carry := uint(num) / base
	if carry > 0 {
		newInt.mant = []uint{uint(num) % base, carry}
	} else {
		newInt.mant = []uint{uint(num)}
	}
	return newInt
}

func SumInt(a Int, b Int) Int {
	var sum Int
	// SAME SIGN
	if a.neg == b.neg {
		sum.neg = a.neg
		sum.mant = sumMant(a.mant, b.mant)
		return sum
	}

	// DIFFERENT SIGN
	if lteMant(b.mant, a.mant) {
		a, b = b, a
	}
	sum.neg = b.neg
	sum.mant = subMant(b.mant, a.mant)

	return sum
}

func sumMant(a []uint, b []uint) []uint {
	if len(a) > len(b) {
		a, b = b, a
	}
	sum := make([]uint, len(b)+1)
	carry, i := uint(0), 0

	for ; i < len(a); i++ {
		sum[i] = a[i] + b[i] + carry
		carry = sum[i] / base
		sum[i] = sum[i] % base
	}

	for ; i < len(b); i++ {
		sum[i] = b[i] + carry
		carry = sum[i] / base
		sum[i] = sum[i] % base
	}
	if carry == 0 {
		return sum[:len(b)]
	}
	sum[len(b)] = carry

	return sum
}

// minuend - subtrahend
func subMant(minu []uint, subt []uint) []uint {
	result := make([]uint, len(minu))
	i := 0
	// demand "carry" from the more significant digit
	dem := false
	for ; i < len(subt); i++ {
		minuD := minu[i]
		subtD := subt[i]
		if dem {
			if minuD > 0 {
				minuD--
				dem = false
			} else {
				minuD += base - 1
			}
		}
		if minuD < subtD {
			minuD += base
			dem = true
		}
		result[i] = minuD - subtD
	}
	for ; i < len(result); i++ {
		minuD := minu[i]
		if dem {
			if minuD > 0 {
				minuD--
				dem = false
			} else {
				minuD += base - 1
			}
		}
		result[i] = minuD
	}

	// STRIP LEADING ZEROS

	// i will now represent the highest digit place without a zero
	// After the line below, i = len(result)-1
	i--
	for i >= 0 && result[i] == 0 {
		i--
	}

	if i == -1 {
		return []uint{0}
	}
	return result[:i+1]
}

// Get the (Base) compelement of num
// func complement(digits []uint) []uint {
//     var comp = make([]uint, len(digits))
//     for i, d = range digits {
//         comp[i] = base - 1 - d
//     }
//     return comp
// }

// Return true if a < b and false otherwise
func lteMant(a []uint, b []uint) bool {
	if len(a) != len(b) {
		return len(a) < len(b)
	}

	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return true
}

// Return true of a <= b and false otherwise
func Lte(a Int, b Int) bool {
	if a.neg != b.neg {
		return a.neg
	}

	pos := !a.neg

	if len(a.mant) != len(b.mant) {
		aSmall := len(a.mant) < len(b.mant)
		// true if len(a.mant) is smaller and both nums are pos
		// or if len(a.mant) is larger and both nums are neg
		return (aSmall && pos) || !(aSmall || pos)
	}

	for i := len(a.mant) - 1; i >= 0; i-- {
		if a.mant[i] != b.mant[i] {
			aSmall := a.mant[i] < b.mant[i]
			return (aSmall && pos) || !(aSmall || pos)
		}
	}
	return true
}

func getBase() uint {
	maxUint := ^uint(0)
	if maxUint == math.MaxUint32 {
		// Base is 10^9
		return 1000000000
	}
	// Base is 10^18
	return 1000000000000000000
}

func (num Int) String() string {
	res := ""
	i := 0
	for ; i < len(num.mant)-1; i++ {
		res = fmt.Sprintf("%018d", num.mant[i]) + res
	}
	res = fmt.Sprintf("%d", num.mant[i]) + res
	if num.neg {
		res = "-" + res
	}
	return res
}
