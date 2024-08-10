package quaver

type GameMode int

const (
	GameMode4K GameMode = iota + 1
	GameMode7K
)

func GameModeFromInt(i int) GameMode {
	switch i {
	case 1:
		return GameMode4K
	case 2:
		return GameMode7K
	default:
		return GameMode4K
	}
}

type Grade string

const (
	GradeX  Grade = "X"
	GradeSS Grade = "SS"
	GradeS  Grade = "S"
	GradeA  Grade = "A"
	GradeB  Grade = "B"
	GradeC  Grade = "C"
	GradeD  Grade = "D"
)

type Modifier int64

const (
	ModifierNone Modifier = 1 << iota
	ModifierNoSliderVelocity
	ModifierSpeed05X
	ModifierSpeed06X
	ModifierSpeed07X
	ModifierSpeed08X
	ModifierSpeed09X
	ModifierSpeed11X
	ModifierSpeed12X
	ModifierSpeed13X
	ModifierSpeed14X
	ModifierSpeed15X
	ModifierSpeed16X
	ModifierSpeed17X
	ModifierSpeed18X
	ModifierSpeed19X
	ModifierSpeed20X
	ModifierStrict
	ModifierChill
	ModifierNoPause
	ModifierAutoplay
	ModifierPaused
	ModifierNoFail
	ModifierNoLongNotes
	ModifierRandomize
	ModifierSpeed055X
	ModifierSpeed065X
	ModifierSpeed075X
	ModifierSpeed085X
	ModifierSpeed095X
	ModifierInverse
	ModifierFullLN
	ModifierMirror
	ModifierCoop
	ModifierSpeed105X
	ModifierSpeed115X
	ModifierSpeed125X
	ModifierSpeed135X
	ModifierSpeed145X
	ModifierSpeed155X
	ModifierSpeed165X
	ModifierSpeed175X
	ModifierSpeed185X
	ModifierSpeed195X
	ModifierHealthAdjust
	ModifierNoMiss
)

const (
	RateModifiers = ModifierSpeed05X | ModifierSpeed06X | ModifierSpeed07X |
		ModifierSpeed08X | ModifierSpeed09X | ModifierSpeed11X |
		ModifierSpeed12X | ModifierSpeed13X | ModifierSpeed14X |
		ModifierSpeed15X | ModifierSpeed16X | ModifierSpeed17X |
		ModifierSpeed18X | ModifierSpeed19X | ModifierSpeed20X |
		ModifierSpeed055X | ModifierSpeed065X | ModifierSpeed075X |
		ModifierSpeed085X | ModifierSpeed095X | ModifierSpeed105X |
		ModifierSpeed115X | ModifierSpeed125X | ModifierSpeed135X |
		ModifierSpeed145X | ModifierSpeed155X | ModifierSpeed165X |
		ModifierSpeed175X | ModifierSpeed185X | ModifierSpeed195X
)

func (m Modifier) hasRateModifiers() bool {
	return m&RateModifiers != 0
}
