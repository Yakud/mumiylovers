package components

import "engo.io/engo"

// A entity has a direction
type AnchorComponent struct {
	Anchor engo.Point
}

// Speed getter
func (t *AnchorComponent) GetAnchor() engo.Point {
	return t.Anchor
}

// Speed setter
func (t *AnchorComponent) SetAnchor(anchor engo.Point) {
	t.Anchor = anchor
}
