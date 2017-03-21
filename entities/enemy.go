package entities

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/Yakud/mumiylovers/components"
)

type Enemy struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
	components.SpeedComponent
	common.AnimationComponent
	//components.AnchorComponent
}

func CreateEnemy(point engo.Point) *Enemy {
	spriteSheet := common.NewSpritesheetFromFile("fishka1.png", 48, 48)
	WalkActionAnimation := &common.Animation{Name: "move", Frames: []int{0, 1}, Loop: true}
	actions := []*common.Animation{WalkActionAnimation}

	entity := &Enemy{BasicEntity: ecs.NewBasic()}

	entity.SpaceComponent = common.SpaceComponent{
		Position: point,
		Width:    48,
		Height:   48,
	}
	entity.RenderComponent = common.RenderComponent{
		Drawable: spriteSheet.Cell(0),
		Scale:    engo.Point{1, 1},
	}
	entity.AnimationComponent = common.NewAnimationComponent(spriteSheet.Drawables(), 0.1)

	entity.AnimationComponent.AddAnimations(actions)
	entity.AnimationComponent.AddDefaultAnimation(WalkActionAnimation)

	return entity
}
