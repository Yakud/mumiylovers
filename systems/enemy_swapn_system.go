package systems

import (
	"math/rand"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/Yakud/mumiylovers/entities"
)

type EnemySpawnSystem struct {
	spawnTime        float32
	spawnTimeCurrent float32

	world *ecs.World
}

func (t *EnemySpawnSystem) Add(world *ecs.World) {
	t.world = world
}

// Remove is called whenever an Entity is removed from the World, in order to remove it from this sytem as well
func (*EnemySpawnSystem) Remove(ecs.BasicEntity) {

}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (t *EnemySpawnSystem) Update(dt float32) {
	t.spawnTimeCurrent += dt
	if t.spawnTimeCurrent >= t.spawnTime {
		e := entities.CreateEnemy(engo.Point{float32(rand.Intn(1024)), float32(rand.Intn(768))})

		for _, system := range t.world.Systems() {
			switch sys := system.(type) {
			case *common.RenderSystem:
				sys.Add(&e.BasicEntity, &e.RenderComponent, &e.SpaceComponent)
			case *common.AnimationSystem:
				sys.Add(&e.BasicEntity, &e.AnimationComponent, &e.RenderComponent)
			case *EnemyLogicSystem:
				sys.Add(e)
			}
		}
	}
}
