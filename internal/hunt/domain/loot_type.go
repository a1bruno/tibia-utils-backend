package domain

import (
	"fmt"
	"strings"
)

type LootType string

const (
	LootTypeLeader LootType = "Leader"
	LootTypeMarket LootType = "Market"
)

func ParseLootType(lootType string) LootType {
	return LootType(strings.TrimSpace(lootType))
}

func (lt LootType) IsKnown() bool {
	switch lt {
	case LootTypeLeader, LootTypeMarket:
		return true
	default:
		return false
	}
}

func (lt LootType) String() string {
	return string(lt)
}

func NewLootType(loot string) (LootType, error) {
	lt := ParseLootType(loot)
	if !lt.IsKnown() {
		return LootType(""), fmt.Errorf("unknown loot type %s", loot)
	}
	return lt, nil
}
