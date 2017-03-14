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

	common.RenderComponent
	common.SpaceComponent

	components.SpeedComponent
}

// Construct and init new mumiy
func NewMumiy(world *ecs.World) *Mumiy {
	mumiy := &Mumiy{BasicEntity: ecs.NewBasic()}
	mumiy.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{10, 10},
		Width:    100,
		Height:   71,
	}
	mumiy.SetSpeed(5)

	texture, err := common.LoadedSprite("mummiy.png")
	if err != nil {
		log.Println("Unable to load texture: " + err.Error())
	}

	mumiy.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale:    engo.Point{1, 1},
	}

	return mumiy
}
