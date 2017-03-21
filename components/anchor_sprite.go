package components

import (
	"math"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type AnchorSpriteComponent struct {
	AnchorComponent

	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent

	SpriteSpaceComponent common.SpaceComponent

	lastBodyPosition *engo.Point
	lastBodyScale    *engo.Point
}

func (t *AnchorSpriteComponent) Update() {
	if t.lastBodyPosition == nil {
		t.lastBodyPosition = &engo.Point{0, 0}
	}

	if t.lastBodyScale == nil {
		t.lastBodyScale = &engo.Point{0, 0}
	}

	needUpdate := t.SpriteSpaceComponent.Rotation != t.Rotation ||
		t.lastBodyPosition.X != t.Position.X ||
		t.lastBodyPosition.Y != t.Position.Y ||
		t.lastBodyScale.X != t.Scale.X ||
		t.lastBodyScale.Y != t.Scale.Y

	// Update only if is object changed
	if needUpdate {
		sin, cos := math.Sincos(float64(t.Rotation) * math.Pi / 180)

		t.SpriteSpaceComponent.Position.X = t.Position.X - t.Anchor.X*t.Scale.X*float32(cos) + t.Anchor.Y*t.Scale.Y*float32(sin)
		t.SpriteSpaceComponent.Position.Y = t.Position.Y - t.Anchor.Y*t.Scale.Y*float32(cos) - t.Anchor.X*t.Scale.X*float32(sin)
		t.SpriteSpaceComponent.Rotation = t.Rotation

		t.lastBodyPosition.X = t.Position.X
		t.lastBodyPosition.Y = t.Position.Y
		t.lastBodyScale.X = t.Scale.X
		t.lastBodyScale.Y = t.Scale.Y
	}
}
