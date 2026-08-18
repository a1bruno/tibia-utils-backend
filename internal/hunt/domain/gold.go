package domain

type Gold int64

func (g Gold) Add(other Gold) Gold {
	return g + other
}

func (g Gold) Sub(other Gold) Gold {
	return g - other
}

func (g Gold) Neg() Gold {
	return -g
}

func (g Gold) Abs() Gold {
	if g < 0 {
		return -g
	}
	return g
}

func (g Gold) IsNegative() bool {
	return g < 0
}

func (g Gold) IsZero() bool {
	return g == 0
}

func (g Gold) Int64() int64 {
	return int64(g)
}

func NewGold(value int64) Gold {
	return Gold(value)
}
