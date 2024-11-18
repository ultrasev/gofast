package math

// Min returns the smaller of x or y.
func Min[T ~int | ~float64](x, y T) T {
    if x < y {
        return x
    }
    return y
}

// Max returns the larger of x or y.
func Max[T ~int | ~float64](x, y T) T {
	if x > y {
        return x
    }
	return y
}
