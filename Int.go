package one4

import (
	"fmt"
	"math"
    "strconv"
)

var (
	base = getBase()
)

type Int struct {
	// Higher index -> more siginificant digit
	mant []uint
	neg  bool
}

// Takes in a string of form "dddddddd..." where each d is a digit.
// An arbitrary number of digits may appear. Additionally, a `+` or '-'
// may precede the form.
func MakeIntStr(num string) Int {
    var res Int
    baseLen, bitLen := 9, 64
    if base == 10000 {
        baseLen, bitLen = 4, 32
    }

    fstD := num[0]
    if fstD == '-' || fstD == '+' {
        res.neg = fstD == '-'
        num = num[1:]
    }

    mantLen := len(num)/baseLen
    if len(num) % baseLen > 0 {
        mantLen++
    }
    res.mant = make([]uint, mantLen)

    remD := len(num)
    i := 0
    var bin uint64
    for ; remD > baseLen; remD -= baseLen {
        bin, _ = strconv.ParseUint(num[remD-baseLen:remD], 10, bitLen)
        res.mant[i] = uint(bin)
        i++
    }
    if remD > 0 {
        bin, _ = strconv.ParseUint(num[:remD], 10, bitLen)
        res.mant[mantLen-1] = uint(bin)
    }

    return res
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

func MultInt(a Int, b Int) Int {
	var res Int
	if a.neg != b.neg {
		res.neg = true
	}
	res.mant = karatsuba(a.mant, b.mant)
	return res
}

func karatsuba(x []uint, y []uint) []uint {
	if len(x) <= 1 && len(y) <= 1 {
		// TODO Don't change to int?
		return MakeInt(int(x[0] * y[0])).mant
	}
	// Set y to be of smaller len
	if len(x) < len(y) {
		x, y = y, x
	}
	n := len(x)
	m := n / 2
	// Lo will be bits 0 to m-1. Hi will be m onwards
	var xl, xh, yl, yh []uint

	xl, xh = x[:m], x[m:]
	if m >= len(y) {
		yl = y
		yh = []uint{0}
	} else {
		yl, yh = y[:m], y[m:]
	}

	a := karatsuba(xh, yh)
	b := karatsuba(xl, yl)
	e := karatsuba(sumMant(xl, xh), sumMant(yl, yh))
	e = subMant(subMant(e, a), b)

	res := sumMant(sumMant(lShift(a, uint(2*m)), lShift(e, uint(m))), b)
	return res
}

// Increase pow of mant by base^pow
func lShift(mant []uint, pow uint) []uint {
	if len(mant) == 1 && mant[0] == 0 {
		return mant
	}
	newMant := make([]uint, uint(len(mant))+pow)
	c := uint(pow)
	for i := 0; i < len(mant); i++ {
		newMant[uint(i)+c] = mant[i]
	}
	return newMant
}

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
		// return 1000000000
		return 10000
	}
	// Base is 10^18
	// return 1000000000000000000
	return 1000000000
}

func (num Int) String() string {
	strFmt := ""
	// if base == 1000000000 {
	if base == 10000 {
		//strFmt = "%09d"
		strFmt = "%04d"
	} else {
		// strFmt = "%018d"
		strFmt = "%09d"
	}
	res := ""
	i := 0
	for ; i < len(num.mant)-1; i++ {
		res = fmt.Sprintf(strFmt, num.mant[i]) + res
	}
	res = fmt.Sprintf("%d", num.mant[i]) + res
	if num.neg {
		res = "-" + res
	}
	return res
}
