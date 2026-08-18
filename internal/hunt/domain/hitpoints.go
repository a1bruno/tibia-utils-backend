package domain

import "errors"

type HitPoints int64

func (h HitPoints) Add(other HitPoints) HitPoints {
	return h + other
}

func (h HitPoints) IsZero() bool {
	return h == 0
}

func (h HitPoints) Int64() int64 {
	return int64(h)
}

func NewHitPoints(hitpoints int64) (HitPoints, error) {
	if hitpoints < 0 {
		return HitPoints(0), errors.New("hitpoints cannot be negative")
	}
	return HitPoints(hitpoints), nil
}
