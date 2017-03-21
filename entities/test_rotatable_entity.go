package entities

import (
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/Yakud/mumiylovers/components"
)

type RotatableEntity struct {
	components.AnchorSpriteComponent
}

func NewRotatableEntity() *RotatableEntity {
	entity := &RotatableEntity{}
	entity.BasicEntity = ecs.NewBasic()
	entity.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{200, 200},
		Width:    100,
		Height:   71,
	}
	entity.SpriteSpaceComponent = entity.SpaceComponent

	texture, err := common.LoadedSprite("mummiy.png")
	if err != nil {
		log.Println("Unable to load texture: " + err.Error())
	}

	entity.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale:    engo.Point{1, 1},
	}
	entity.SetAnchor(engo.Point{50, 35})

	return entity
}
