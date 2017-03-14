package components

// A entity has a direction
type SpeedComponent struct {
	speed float32
}

// Speed getter
func (t *SpeedComponent) Speed() float32 {
	return t.speed
}

// Speed setter
func (t *SpeedComponent) SetSpeed(speed float32) {
	t.speed = speed
}

// Speed add
func (t *SpeedComponent) AddSpeed(speed float32) {
	t.speed += speed
}
