package stats

type StaticStat struct {
	Health int
	Mana   int
	Speed  int
}

// NewStaticStat creates a new instance of StaticStat with initial values
func NewStaticStat(initialHealth, initialMana, initialSpeed int) *StaticStat {
	return &StaticStat{
		Health: initialHealth,
		Mana:   initialMana,
		Speed:  initialSpeed,
	}
}

// SetHealth allows you to set the Health from outside the package
func (ss *StaticStat) SetHealth(newHealth int) {
	ss.Health = newHealth
}

// GetHealth allows you to get the Health from outside the package
func (ss *StaticStat) GetHealth() int {
	return ss.Health
}

// SetMana allows you to set the Mana from outside the package
func (ss *StaticStat) SetMana(newMana int) {
	ss.Mana = newMana
}

// GetMana allows you to get the Mana from outside the package
func (ss *StaticStat) GetMana() int {
	return ss.Mana
}

// SetSpeed allows you to set the Speed from outside the package
func (ss *StaticStat) SetSpeed(newSpeed int) {
	ss.Speed = newSpeed
}

// GetSpeed allows you to get the Speed from outside the package
func (ss *StaticStat) GetSpeed() int {
	return ss.Speed
}
