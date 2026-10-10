package cli

import "math"

type mathStdLib struct{}

//
func (c *CliEngine) RegisterLibMath() {
	lib := &mathStdLib{}

	c.RegisterMethod(lib, "Math.Abs", "Abs returns the absolute value of x.", "x")
	c.RegisterMethod(lib, "Math.Acos", "Acos returns the arccosine, in radians, of x.", "x")
	c.RegisterMethod(lib, "Math.Acosh", "Acosh returns the inverse hyperbolic cosine of x.", "x")
	c.RegisterMethod(lib, "Math.Asin", "Asinh returns the inverse hyperbolic sine of x.", "x")
	c.RegisterMethod(lib, "Math.Asinh", "Asinh returns the inverse hyperbolic sine of x.", "x")
	c.RegisterMethod(lib, "Math.Atan", "Atan returns the arctangent, in radians, of x.", "x")
	c.RegisterMethod(lib, "Math.Atan2", "Atan2 returns the arc tangent of y/x, using the signs of the two to determine the quadrant of the return value.", "x", "y")
	c.RegisterMethod(lib, "Math.Atanh", "Atanh returns the inverse hyperbolic tangent of x.", "x")
	c.RegisterMethod(lib, "Math.Cbrt", "Cbrt returns the cube root of x.", "x")
	c.RegisterMethod(lib, "Math.Ceil", "Ceil returns the least integer value greater than or equal to x.", "x")
	c.RegisterMethod(lib, "Math.Cos", "Cos returns the cosine of the radian argument x.", "x")
	c.RegisterMethod(lib, "Math.Cosh", "Cosh returns the hyperbolic cosine of x.", "x")
	c.RegisterMethod(lib, "Math.Dim", "Dim returns the maximum of x-y or 0.", "x", "y")
	c.RegisterMethod(lib, "Math.Erf", "Erf returns the error function of x.", "x")
	c.RegisterMethod(lib, "Math.Erfc", "Erfc returns the complementary error function of x.", "x")
	c.RegisterMethod(lib, "Math.Erfcinv", "Erfcinv returns the inverse of Erfc(x).", "x")
	c.RegisterMethod(lib, "Math.Erfinv", "Erfinv returns the inverse error function of x.", "x")
	c.RegisterMethod(lib, "Math.Exp", "Exp returns e**x, the base-e exponential of x.", "x")
	c.RegisterMethod(lib, "Math.Exp2", "Exp2 returns 2**x, the base-2 exponential of x.", "x")
	c.RegisterMethod(lib, "Math.Expm1", "Expm1 returns e**x - 1, the base-e exponential of x minus 1. It is more accurate than Exp(x) - 1 when x is near zero.", "x")
	c.RegisterMethod(lib, "Math.FMA", "FMA returns x * y + z, computed with only one rounding. (That is, FMA returns the fused multiply-add of x, y, and z.)", "x", "y", "z")
	c.RegisterMethod(lib, "Math.Floor", "Floor returns the greatest integer value less than or equal to x.", "x")
	c.RegisterMethod(lib, "Math.Gamma", "Gamma returns the Gamma function of x.", "x")
	c.RegisterMethod(lib, "Math.Hypot", "Hypot returns Sqrt(p*p + q*q), taking care to avoid unnecessary overflow and underflow.", "x", "y")
	c.RegisterMethod(lib, "Math.Inf", "Inf returns positive infinity if sign >= 0, negative infinity if sign < 0.", "sign")
	c.RegisterMethod(lib, "Math.J0", "J0 returns the order-zero Bessel function of the first kind.", "x")
	c.RegisterMethod(lib, "Math.J1", "J1 returns the order-one Bessel function of the first kind.", "x")
	c.RegisterMethod(lib, "Math.Jn", "Jn returns the order-n Bessel function of the first kind.", "n", "x")
	c.RegisterMethod(lib, "Math.Ldexp", "Ldexp is the inverse of Frexp. It returns frac × 2**exp.", "frac", "exp")
	c.RegisterMethod(lib, "Math.Log", "Log returns the natural logarithm of x.", "x")
	c.RegisterMethod(lib, "Math.Log10", "Log10 returns the decimal logarithm of x. The special cases are the same as for Log.", "x")
	c.RegisterMethod(lib, "Math.Log1p", "Log1p returns the natural logarithm of 1 plus its argument x. It is more accurate than Log(1 + x) when x is near zero.", "x")
	c.RegisterMethod(lib, "Math.Log2", "Log2 returns the binary logarithm of x. The special cases are the same as for Log.", "x")
	c.RegisterMethod(lib, "Math.Logb", "Logb returns the binary exponent of x.", "x")
	c.RegisterMethod(lib, "Math.Max", "Max returns the larger of x or y.", "x", "y")
	c.RegisterMethod(lib, "Math.Min", "Min returns the smaller of x or y.", "x", "y")
	c.RegisterMethod(lib, "Math.Mod", "Mod returns the floating-point remainder of x/y. The magnitude of the result is less than y and its sign agrees with that of x.", "x", "y")
	c.RegisterMethod(lib, "Math.NaN", "NaN returns an “not-a-number” value.")
	c.RegisterMethod(lib, "Math.Nextafter", "Nextafter returns the next representable float64 value after x towards y.", "x", "y")
	c.RegisterMethod(lib, "Math.Pow", "Pow returns x**y, the base-x exponential of y.", "x", "y")
	c.RegisterMethod(lib, "Math.Pow10", "Pow10 returns 10**n, the base-10 exponential of n.", "n")
	c.RegisterMethod(lib, "Math.Remainder", "Remainder returns the floating-point remainder of x/y.", "x", "y")
	c.RegisterMethod(lib, "Math.Round", "Round returns the nearest integer, rounding half away from zero.", "x")
	c.RegisterMethod(lib, "Math.RoundToEven", "RoundToEven returns the nearest integer, rounding ties to even.", "x")
	c.RegisterMethod(lib, "Math.Sin", "Sin returns the sine of the radian argument x.", "x")
	c.RegisterMethod(lib, "Math.Sinh", "Sinh returns the hyperbolic sine of x.", "x")
	c.RegisterMethod(lib, "Math.Sqrt", "Sqrt returns the square root of x.", "x")
	c.RegisterMethod(lib, "Math.Tan", "Tan returns the tangent of the radian argument x.", "x")
	c.RegisterMethod(lib, "Math.Tanh", "Tanh returns the hyperbolic tangent of x.", "x")
	c.RegisterMethod(lib, "Math.Trunc", "Trunc returns the integer value of x.", "x")
	c.RegisterMethod(lib, "Math.Y0", "Y1 returns the order-zero Bessel function of the second kind.", "x")
	c.RegisterMethod(lib, "Math.Y1", "Y1 returns the order-one Bessel function of the second kind.", "x")
	c.RegisterMethod(lib, "Math.Yn", "Yn returns the order-n Bessel function of the second kind", "n", "x")
	c.RegisterMethod(lib, "Math.IsInf", "IsInf reports whether f is an infinity, according to sign. If sign > 0, IsInf reports whether f is positive infinity. If sign < 0, IsInf reports whether f is negative infinity. If sign == 0, IsInf reports whether f is either infinity.", "f", "sign")
	c.RegisterMethod(lib, "Math.IsNaN", "IsNaN reports whether f is an “not-a-number” value.", "f")
	c.RegisterMethod(lib, "Math.Pi", "Mathematical constant")
	c.RegisterMethod(lib, "Math.E", "Mathematical constant")
}

// Abs returns the absolute value of x.
func (c *mathStdLib) Abs(x float64) float64 {
	return math.Abs(x)
}

// Acos returns the arccosine, in radians, of x.
func (c *mathStdLib) Acos(x float64) float64 {
	return math.Acos(x)
}

// Acosh returns the inverse hyperbolic cosine of x.
func (c *mathStdLib) Acosh(x float64) float64 {
	return math.Acosh(x)
}

// Asinh returns the inverse hyperbolic sine of x.
func (c *mathStdLib) Asin(x float64) float64 {
	return math.Asin(x)
}

// Asinh returns the inverse hyperbolic sine of x.
func (c *mathStdLib) Asinh(x float64) float64 {
	return math.Asinh(x)
}

// Atan returns the arctangent, in radians, of x.
func (c *mathStdLib) Atan(x float64) float64 {
	return math.Atan(x)
}

// Atan2 returns the arc tangent of y/x, using the signs of the two to determine the quadrant of the return value.
func (c *mathStdLib) Atan2(x, y float64) float64 {
	return math.Atan2(x, y)
}

// Atanh returns the inverse hyperbolic tangent of x.
func (c *mathStdLib) Atanh(x float64) float64 {
	return math.Atanh(x)
}

// Cbrt returns the cube root of x.
func (c *mathStdLib) Cbrt(x float64) float64 {
	return math.Cbrt(x)
}

// Ceil returns the least integer value greater than or equal to x.
func (c *mathStdLib) Ceil(x float64) float64 {
	return math.Ceil(x)
}

// Cos returns the cosine of the radian argument x.
func (c *mathStdLib) Cos(x float64) float64 {
	return math.Cos(x)
}

// Cosh returns the hyperbolic cosine of x.
func (c *mathStdLib) Cosh(x float64) float64 {
	return math.Cosh(x)
}

// Dim returns the maximum of x-y or 0.
func (c *mathStdLib) Dim(x, y float64) float64 {
	return math.Dim(x, y)
}

// Erf returns the error function of x.
func (c *mathStdLib) Erf(x float64) float64 {
	return math.Erf(x)
}

// Erfc returns the complementary error function of x.
func (c *mathStdLib) Erfc(x float64) float64 {
	return math.Erfc(x)
}

// Erfcinv returns the inverse of Erfc(x).
func (c *mathStdLib) Erfcinv(x float64) float64 {
	return math.Erfcinv(x)
}

// Erfinv returns the inverse error function of x.
func (c *mathStdLib) Erfinv(x float64) float64 {
	return math.Erfinv(x)
}

// Exp returns e**x, the base-e exponential of x.
func (c *mathStdLib) Exp(x float64) float64 {
	return math.Exp(x)
}

// Exp2 returns 2**x, the base-2 exponential of x.
func (c *mathStdLib) Exp2(x float64) float64 {
	return math.Exp2(x)
}

// Expm1 returns e**x - 1, the base-e exponential of x minus 1. It is more accurate than Exp(x) - 1 when x is near zero.
func (c *mathStdLib) Expm1(x float64) float64 {
	return math.Expm1(x)
}

// FMA returns x * y + z, computed with only one rounding. (That is, FMA returns the fused multiply-add of x, y, and z.)
func (c *mathStdLib) FMA(x, y, z float64) float64 {
	return math.FMA(x, y, z)
}

// Floor returns the greatest integer value less than or equal to x.
func (c *mathStdLib) Floor(x float64) float64 {
	return math.Floor(x)
}

// Gamma returns the Gamma function of x.
func (c *mathStdLib) Gamma(x float64) float64 {
	return math.Gamma(x)
}

// Hypot returns Sqrt(p*p + q*q), taking care to avoid unnecessary overflow and underflow.
func (c *mathStdLib) Hypot(x, y float64) float64 {
	return math.Hypot(x, y)
}

// Inf returns positive infinity if sign >= 0, negative infinity if sign < 0.
func (c *mathStdLib) Inf(sign int) float64 {
	return math.Inf(sign)
}

// J0 returns the order-zero Bessel function of the first kind.
func (c *mathStdLib) J0(x float64) float64 {
	return math.J0(x)
}

// J1 returns the order-one Bessel function of the first kind.
func (c *mathStdLib) J1(x float64) float64 {
	return math.J1(x)
}

// Jn returns the order-n Bessel function of the first kind.
func (c *mathStdLib) Jn(n int, x float64) float64 {
	return math.Jn(n, x)
}

// Ldexp is the inverse of Frexp. It returns frac × 2**exp.
func (c *mathStdLib) Ldexp(frac float64, exp int) float64 {
	return math.Ldexp(frac, exp)
}

// Log returns the natural logarithm of x.
func (c *mathStdLib) Log(x float64) float64 {
	return math.Log(x)
}

// Log10 returns the decimal logarithm of x. The special cases are the same as for Log.
func (c *mathStdLib) Log10(x float64) float64 {
	return math.Log10(x)
}

// Log1p returns the natural logarithm of 1 plus its argument x. It is more accurate than Log(1 + x) when x is near zero.
func (c *mathStdLib) Log1p(x float64) float64 {
	return math.Log1p(x)
}

// Log2 returns the binary logarithm of x. The special cases are the same as for Log.
func (c *mathStdLib) Log2(x float64) float64 {
	return math.Log2(x)
}

// Logb returns the binary exponent of x.
func (c *mathStdLib) Logb(x float64) float64 {
	return math.Logb(x)
}

// Max returns the larger of x or y.
func (c *mathStdLib) Max(x, y float64) float64 {
	return math.Max(x, y)
}

// Min returns the smaller of x or y.
func (c *mathStdLib) Min(x, y float64) float64 {
	return math.Min(x, y)
}

// Mod returns the floating-point remainder of x/y. The magnitude of the result is less than y and its sign agrees with that of x.
func (c *mathStdLib) Mod(x, y float64) float64 {
	return math.Mod(x, y)
}

// NaN returns an “not-a-number” value.
func (c *mathStdLib) NaN() float64 {
	return math.NaN()
}

// Nextafter returns the next representable float64 value after x towards y.
func (c *mathStdLib) Nextafter(x, y float64) float64 {
	return math.Nextafter(x, y)
}

// Pow returns x**y, the base-x exponential of y.
func (c *mathStdLib) Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

// Pow10 returns 10**n, the base-10 exponential of n.
func (c *mathStdLib) Pow10(n int) float64 {
	return math.Pow10(n)
}

// Remainder returns the floating-point remainder of x/y.
func (c *mathStdLib) Remainder(x, y float64) float64 {
	return math.Remainder(x, y)
}

// Round returns the nearest integer, rounding half away from zero.
func (c *mathStdLib) Round(x float64) float64 {
	return math.Round(x)
}

// RoundToEven returns the nearest integer, rounding ties to even.
func (c *mathStdLib) RoundToEven(x float64) float64 {
	return math.RoundToEven(x)
}

// Sin returns the sine of the radian argument x.
func (c *mathStdLib) Sin(x float64) float64 {
	return math.Sin(x)
}

// Sinh returns the hyperbolic sine of x.
func (c *mathStdLib) Sinh(x float64) float64 {
	return math.Sinh(x)
}

// Sqrt returns the square root of x.
func (c *mathStdLib) Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

// Tan returns the tangent of the radian argument x.
func (c *mathStdLib) Tan(x float64) float64 {
	return math.Tan(x)
}

// Tanh returns the hyperbolic tangent of x.
func (c *mathStdLib) Tanh(x float64) float64 {
	return math.Tanh(x)
}

// Trunc returns the integer value of x.
func (c *mathStdLib) Trunc(x float64) float64 {
	return math.Trunc(x)
}

// Y1 returns the order-zero Bessel function of the second kind.
func (c *mathStdLib) Y0(x float64) float64 {
	return math.Y0(x)
}

// Y1 returns the order-one Bessel function of the second kind.
func (c *mathStdLib) Y1(x float64) float64 {
	return math.Y1(x)
}

// Yn returns the order-n Bessel function of the second kind
func (c *mathStdLib) Yn(n int, x float64) float64 {
	return math.Yn(n, x)
}

// IsInf reports whether f is an infinity, according to sign. If sign > 0, IsInf reports whether f is positive infinity. If sign < 0, IsInf reports whether f is negative infinity. If sign == 0, IsInf reports whether f is either infinity.
func (c *mathStdLib) IsInf(f float64, sign int) bool {
	return math.IsInf(f, sign)
}

// IsNaN reports whether f is an “not-a-number” value.
func (c *mathStdLib) IsNaN(f float64) bool {
	return math.IsNaN(f)
}

// Mathematical constant
func (c *mathStdLib) Pi() float64 {
	return math.Pi
}

// Mathematical constant
func (c *mathStdLib) E() float64 {
	return math.E
}
