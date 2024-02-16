package stats

type StaticStat interface {
	SetHealth(newHealth int)
	GetHealth() int
	SetMana(newMana int)
	GetMana() int
	SetSpeed(newSpeed float32)
	GetSpeed() float32
}

type staticStat struct {
	health int
	mana   int
	speed  float32
}

// NewStaticStat creates a new instance of StaticStat with initial values
func NewStaticStat(initialHealth, initialMana int ,initialSpeed float32) StaticStat {
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
func (ss *staticStat) SetSpeed(newSpeed float32) {
	ss.speed = newSpeed
}

// GetSpeed allows you to get the Speed from outside the package
func (ss *staticStat) GetSpeed() float32 {
	return ss.speed
}
