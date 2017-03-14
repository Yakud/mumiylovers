package components

// A entity has a direction
type DirectionComponent struct {
	direction float32
}

// Direction getter
func (t *DirectionComponent) Direction() float32 {
	return t.direction
}

// Direction setter
func (t *DirectionComponent) SetDirection(direction float32) {
	t.direction = direction
}

// Direction add
func (t *DirectionComponent) AddDirection(direction float32) {
	t.direction += direction
}
