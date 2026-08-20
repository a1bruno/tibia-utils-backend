package domain

import "strings"

type Participant struct {
	name     string
	isLeader bool
	loot     Gold
	supplies Gold
	balance  Gold
	damage   HitPoints
	healing  HitPoints
}

func NewParticipant(name string, isLeader bool, loot Gold, supplies Gold, balance Gold, damage HitPoints, healing HitPoints) (Participant, error) {
	clearedName := strings.TrimSpace(name)
	if clearedName == "" {
		return Participant{}, ErrEmptyPlayerName
	}
	if loot.IsNegative() {
		return Participant{}, ErrNegativeAmount
	}
	if supplies.IsNegative() {
		return Participant{}, ErrNegativeAmount
	}
	if damage < 0 {
		return Participant{}, ErrNegativeAmount
	}
	if healing < 0 {
		return Participant{}, ErrNegativeAmount
	}
	if balance != (loot.Sub(supplies)) {
		return Participant{}, ErrParticipantBalance
	}
	return Participant{name: clearedName, isLeader: isLeader, loot: loot, supplies: supplies, balance: balance, damage: damage, healing: healing}, nil
}

func (p Participant) Name() string {
	return p.name
}

func (p Participant) IsLeader() bool {
	return p.isLeader
}

func (p Participant) Loot() Gold {
	return p.loot
}

func (p Participant) Supplies() Gold {
	return p.supplies
}

func (p Participant) Balance() Gold {
	return p.balance
}

func (p Participant) Damage() HitPoints {
	return p.damage
}

func (p Participant) Healing() HitPoints {
	return p.healing
}
