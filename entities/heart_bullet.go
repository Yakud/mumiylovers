package entities

import (
	"log"

	"math/rand"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/Yakud/mumiylovers/components"
)

// Heart bullet
// todo delete bullet
type HeartBullet struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent

	components.SpeedComponent
	components.DirectionComponent
}

// Construct and init new mumiy
func NewHeartBullet() *HeartBullet {
	bullet := &HeartBullet{BasicEntity: ecs.NewBasic()}
	bullet.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{10, 10},
		Width:    100,
		Height:   71,
	}
	bullet.SetSpeed(7)

	var sprite *common.Texture
	var err error

	rnd := rand.Intn(2)
	if rnd == 0 {
		sprite, err = common.LoadedSprite("heart_bullet.png")
	}
	if rnd == 1 {
		sprite, err = common.LoadedSprite("heart_bullet_2.png")
	}
	if err != nil {
		log.Println("Unable to load texture: " + err.Error())
	}

	bullet.RenderComponent = common.RenderComponent{
		Drawable: sprite,
		Scale:    engo.Point{1, 1},
	}

	return bullet
}
