package domain

import "errors"

var (
	ErrInvalidPartySize     = errors.New("invalid party size")
	ErrDuplicateParticipant = errors.New("duplicated participant")
	ErrEmptyPlayerName      = errors.New("empty player name")
	ErrInvalidPeriod        = errors.New("invalid period")
	ErrNegativeDuration     = errors.New("negative duration")
	ErrLootMismatch         = errors.New("loot mismatch")
	ErrSuppliesMismatch     = errors.New("supplies mismatch")
	ErrBalanceMismatch      = errors.New("balance mismatch")
	ErrParticipantBalance   = errors.New("participant balance mismatch")
	ErrNegativeAmount       = errors.New("negative amount")
	ErrNoParticipants       = errors.New("no participants")
)
