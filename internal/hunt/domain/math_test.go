package domain

import (
	"errors"
	"slices"
	"testing"
)

func TestDivFloor(t *testing.T) {
	tests := []struct {
		name             string
		dividend         int64
		divisor          int64
		expectedQuotient int64
		expectedRest     int64
	}{
		{
			name:             "positive with rest",
			dividend:         7,
			divisor:          4,
			expectedQuotient: 1,
			expectedRest:     3,
		},
		{
			name:             "negative with rest",
			dividend:         -7,
			divisor:          4,
			expectedQuotient: -2,
			expectedRest:     1,
		},
		{
			name:             "exactly negative division",
			dividend:         -8,
			divisor:          4,
			expectedQuotient: -2,
			expectedRest:     0,
		},
		{
			name:             "zero",
			dividend:         0,
			divisor:          1,
			expectedQuotient: 0,
			expectedRest:     0,
		},
		{
			name:             "real balance of session",
			dividend:         24000878,
			divisor:          4,
			expectedQuotient: 6000219,
			expectedRest:     2,
		},
		{
			name:             "solo exactly negative",
			dividend:         -459,
			divisor:          1,
			expectedQuotient: -459,
			expectedRest:     0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultQuotient, resultRest := divFloor(tt.dividend, tt.divisor)
			if resultQuotient != tt.expectedQuotient {
				t.Errorf("quotient: found %d, expected %d", resultQuotient, tt.expectedQuotient)
			}
			if resultRest != tt.expectedRest {
				t.Errorf("rest: found %d, expected %d", resultRest, tt.expectedRest)
			}
		})
	}
}

func TestSplitEvenly(t *testing.T) {
	tests := []struct {
		name           string
		total          Gold
		n              int
		expectedShares []Gold
	}{
		{
			name:           "4 players real session",
			total:          24000878,
			n:              4,
			expectedShares: []Gold{6000220, 6000220, 6000219, 6000219},
		},
		{
			name:           "solo player negative session",
			total:          -459,
			n:              1,
			expectedShares: []Gold{-459},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shares, err := splitEvenly(tt.total, tt.n)
			if err != nil {
				t.Fatalf("unexpected error %v", err)
			}
			if !slices.Equal(shares, tt.expectedShares) {
				t.Errorf("shares: expected %v, found %v", tt.expectedShares, shares)
			}
			var total Gold
			for _, share := range shares {
				total += share
			}
			if total != tt.total {
				t.Errorf("total: expected %d, found %d", tt.total, total)
			}
		})
	}
}

func TestSplitEvenlyInvalidPartySize(t *testing.T) {
	shares, err := splitEvenly(4000, 0)
	if !errors.Is(err, ErrInvalidPartySize) {
		t.Errorf("error: expected %v, found %v", ErrInvalidPartySize, err)
	}
	if shares != nil {
		t.Errorf("expected nil, found %v", shares)
	}
}
