package domain

func divFloor(dividend int64, divisor int64) (int64, int64) {
	rest := dividend % divisor
	quotient := dividend / divisor

	if (rest != 0) && (dividend < 0) != (divisor < 0) {
		quotient--
		rest += divisor
	}
	return quotient, rest
}

func splitEvenly(total Gold, n int) ([]Gold, error) {
	if n <= 0 {
		return nil, ErrInvalidPartySize
	}
	quotient, rest := divFloor(int64(total), int64(n))
	slices := make([]Gold, n)

	for i := 0; i < n; i++ {
		slices[i] = Gold(quotient)
	}
	for i := 0; i < int(rest); i++ {
		slices[i] = slices[i] + Gold(1)
	}

	return slices, nil
}
