package one4

// num can represent the mantissa as well as the sign bit. Note that there need
// only be as many didits as `prec` (the precision) however. Due to prec, to
// get a full range of reperesentaiton of digits, we need pow, so that
// `base`^`pow` gives the appropriate number.
type Float struct {
    num Int
    prec uint
    pow int
}

// Return Float a+b. The `prec` of result will be max(`a.prec`, `b.prec`)
func SumFloat(a Float, b Float) Float {
    return Float{}
}

// Return Float a*b. The `prec` of result will be max(`a.prec`, `b.prec`)
// This means that the result may not fit in the result
func MultFloat(a Float, b Float) Float {
    return Float{}
}

// Return Float a*b. The `prec` of the result will be the `prec` arg. If
// the arguement `prec` is 0, the precision will large enough to fit the
// representation of Float a*b without any rounding error
func MultFloatPrec(a Float, b Float, prec uint) Float {
    return Float{}
}

// TODO Consolodate MultFloat & MultFloatPrec into 1 function?

// Return Sqrt(a). The `prec` of the result will be the `prec` arg.
func SqRootFloat(a Float, prec uint) Float {
    return Float{}
}

// Return Float `n/d`. The `prec` of the result will be
// max(`n.prec`, `d.prec`).
func DivFloat(n Float, d Float) Float {
    return Float{}
}

// Return Float `n/d`. The `prec` of the resilt will be the `prec` arg. If the
// `prec` arg is 0, the `prec` of the result will be large enough to fit the
// representation of Float `n/d` without any round error, that is, unless
// there is some infinitly repeating sequence after some point, which will
// then be truncated and rounded at an as of yet undetermined point.
func DivFloatPrec(n Float, d Float, prec uint) Float {
    return Float{}
}

// TOOD Consolodate DivFloat and DivFloatPrec into 1 function?

// Return as Float the quotient of n/d
func QuoFloat(n Float, d Float) Float {
    return Float{}
}

// Return as Float the remainder of n/d
func RemFloat(n Float, d Float) Float{
    return Float{}
}