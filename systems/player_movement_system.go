package systems

import (
	"fmt"
	"math"

	"engo.io/ecs"
	"engo.io/engo"
	"github.com/Yakud/mumiylovers/entities"
	"github.com/Yakud/mumiylovers/utils"
)

// Input keys and move player
type PlayerMovementSystem struct {
	mumiy *entities.Mumiy

	lastCalcPosition float32
	lastPosition     engo.Point
}

// Remove is called whenever an Entity is removed from the World, in order to remove it from this sytem as well
func (*PlayerMovementSystem) Remove(ecs.BasicEntity) {

}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (t *PlayerMovementSystem) Update(dt float32) {
	isMove := false

	if engo.Input.Button("MoveUp").Down() {
		t.mumiy.Pos.Subtract(engo.Point{0, t.mumiy.Speed() * 60 * dt})
		isMove = true
	}
	if engo.Input.Button("MoveDown").Down() {
		t.mumiy.Pos.Add(engo.Point{0, t.mumiy.Speed() * 60 * dt})
		isMove = true
	}
	if engo.Input.Button("MoveLeft").Down() {
		t.mumiy.Pos.Subtract(engo.Point{t.mumiy.Speed() * 60 * dt, 0})
		isMove = true
	}
	if engo.Input.Button("MoveRight").Down() {
		t.mumiy.Pos.Add(engo.Point{t.mumiy.Speed() * 60 * dt, 0})
		isMove = true
	}

	if t.mumiy.Pos.X < 0 {
		t.mumiy.Pos.X = 0
	}

	if t.mumiy.Pos.Y < 0 {
		t.mumiy.Pos.Y = 0
	}

	if t.mumiy.Pos.X+t.mumiy.Width > 1024 {
		t.mumiy.Pos.X = 1024 - t.mumiy.Width
	}

	if t.mumiy.Pos.Y+t.mumiy.Height > 768 {
		t.mumiy.Pos.Y = 768 - t.mumiy.Height
	}

	if isMove {
		t.lastCalcPosition += dt
		if t.lastCalcPosition >= 0.001 {
			direction := utils.PointsDirection(t.lastPosition, t.mumiy.Pos)
			t.mumiy.Rotation = float32(direction)

			fmt.Println(direction)

			t.lastPosition = t.mumiy.Pos
			t.lastCalcPosition = 0
		}

	}

	//((xoffset)*cos(convert_dir_to_radian))-((yoffset)*sin(convert_dir_to_radian));
	//((yoffset)*cos(convert_dir_to_radian))+((xoffset)*sin(convert_dir_to_radian));
	sin, cos := math.Sincos(float64(t.mumiy.Rotation) * math.Pi / 180)
	t.mumiy.Position.X = t.mumiy.Pos.X - t.mumiy.GetAnchor().X*float32(cos) + t.mumiy.GetAnchor().Y*float32(sin)
	t.mumiy.Position.Y = t.mumiy.Pos.Y - t.mumiy.GetAnchor().Y*float32(cos) - t.mumiy.GetAnchor().X*float32(sin)
}

func (t *PlayerMovementSystem) Add(mumiy *entities.Mumiy) {
	t.mumiy = mumiy
}

func NewPlayerMovementSystem() *PlayerMovementSystem {
	return &PlayerMovementSystem{}
}
