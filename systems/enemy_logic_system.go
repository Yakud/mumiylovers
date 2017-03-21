package systems

import (
	"math"

	"engo.io/ecs"
	"engo.io/engo"
	"github.com/Yakud/mumiylovers/entities"
	"github.com/Yakud/mumiylovers/utils"
)

type EnemyLogicSystem struct {
	enemies []*entities.Enemy
	target  *engo.Point
}

// Remove is called whenever an Entity is removed from the World, in order to remove it from this sytem as well
func (*EnemyLogicSystem) Remove(ecs.BasicEntity) {

}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (t *EnemyLogicSystem) Update(dt float32) {
	for _, enemy := range t.enemies {
		if enemy.Position.PointDistance(*t.target) > 2 {
			direction := utils.PointsDirection(enemy.Position, *t.target)
			enemy.Position.Add(engo.Point{
				X: float32(math.Cos(direction*math.Pi/180)) * 60 * dt * 2,
				Y: float32(math.Sin(direction*math.Pi/180)) * 60 * dt * 2,
			})
			enemy.SpaceComponent.Rotation = float32(direction)
		}
	}
}

func (t *EnemyLogicSystem) Add(enemy *entities.Enemy) {
	t.enemies = append(t.enemies, enemy)
}

func (t *EnemyLogicSystem) AddTarget(target *engo.Point) {
	t.target = target
}
