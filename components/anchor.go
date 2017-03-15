package components

import "engo.io/engo"

// A entity has a direction
type AnchorComponent struct {
	anchor engo.Point
}

// Speed getter
func (t *AnchorComponent) Anchor() engo.Point {
	return t.anchor
}

// Speed setter
func (t *AnchorComponent) SetAnchor(anchor engo.Point) {
	t.anchor = anchor
}
