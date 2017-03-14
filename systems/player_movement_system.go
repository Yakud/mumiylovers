package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"github.com/Yakud/mumiylovers/entities"
)

// Input keys and move player
type PlayerMovementSystem struct {
	mumiy *entities.Mumiy
}

// Remove is called whenever an Entity is removed from the World, in order to remove it from this sytem as well
func (*PlayerMovementSystem) Remove(ecs.BasicEntity) {

}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (t *PlayerMovementSystem) Update(dt float32) {
	if engo.Input.Button("MoveUp").Down() {
		t.mumiy.Position.Subtract(engo.Point{0, t.mumiy.Speed() * 60 * dt})
	}
	if engo.Input.Button("MoveDown").Down() {
		t.mumiy.Position.Add(engo.Point{0, t.mumiy.Speed() * 60 * dt})
	}
	if engo.Input.Button("MoveLeft").Down() {
		t.mumiy.Position.Subtract(engo.Point{t.mumiy.Speed() * 60 * dt, 0})
	}
	if engo.Input.Button("MoveRight").Down() {
		t.mumiy.Position.Add(engo.Point{t.mumiy.Speed() * 60 * dt, 0})
	}

	if t.mumiy.Position.X < 0 {
		t.mumiy.Position.X = 0
	}

	if t.mumiy.Position.Y < 0 {
		t.mumiy.Position.Y = 0
	}

	if t.mumiy.Position.X+t.mumiy.Width > 1024 {
		t.mumiy.Position.X = 1024 - t.mumiy.Width
	}

	if t.mumiy.Position.Y+t.mumiy.Height > 768 {
		t.mumiy.Position.Y = 768 - t.mumiy.Height
	}
}

func NewPlayerMovementSystem(mumiy *entities.Mumiy) *PlayerMovementSystem {
	engo.Input.RegisterButton("MoveUp", engo.ArrowUp)
	engo.Input.RegisterButton("MoveDown", engo.ArrowDown)
	engo.Input.RegisterButton("MoveLeft", engo.ArrowLeft)
	engo.Input.RegisterButton("MoveRight", engo.ArrowRight)

	return &PlayerMovementSystem{
		mumiy: mumiy,
	}
}
