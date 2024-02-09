package stats

type StaticStat interface {
	SetHealth(newHealth int)
	GetHealth() int
	SetMana(newMana int)
	GetMana() int
	SetSpeed(newSpeed int)
	GetSpeed() int
}

type staticStat struct {
	health int
	mana   int
	speed  int
}

// NewStaticStat creates a new instance of StaticStat with initial values
func NewStaticStat(initialHealth, initialMana, initialSpeed int) StaticStat {
	return &staticStat{
		health: initialHealth,
		mana:   initialMana,
		speed:  initialSpeed,
	}
}

// SetHealth allows you to set the Health from outside the package
func (ss *staticStat) SetHealth(newHealth int) {
	ss.health = newHealth
}

// GetHealth allows you to get the Health from outside the package
func (ss *staticStat) GetHealth() int {
	return ss.health
}

// SetMana allows you to set the Mana from outside the package
func (ss *staticStat) SetMana(newMana int) {
	ss.mana = newMana
}

// GetMana allows you to get the Mana from outside the package
func (ss *staticStat) GetMana() int {
	return ss.mana
}

// SetSpeed allows you to set the Speed from outside the package
func (ss *staticStat) SetSpeed(newSpeed int) {
	ss.speed = newSpeed
}

// GetSpeed allows you to get the Speed from outside the package
func (ss *staticStat) GetSpeed() int {
	return ss.speed
}
