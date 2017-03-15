package entities

import (
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/Yakud/mumiylovers/components"
)

// Is a player entity
type Mumiy struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
	components.SpeedComponent
	components.AnchorComponent

	Pos engo.Point
}

// Construct and init new mumiy
func NewMumiy(world *ecs.World) *Mumiy {
	mumiy := &Mumiy{BasicEntity: ecs.NewBasic()}
	mumiy.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{200, 200},
		Width:    100,
		Height:   71,
	}
	mumiy.Pos = mumiy.Position
	mumiy.SetSpeed(5)

	texture, err := common.LoadedSprite("mummiy.png")
	if err != nil {
		log.Println("Unable to load texture: " + err.Error())
	}

	mumiy.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale:    engo.Point{1, 1},
	}
	mumiy.SetAnchor(engo.Point{50, 35})

	return mumiy
}
