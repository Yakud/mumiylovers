package main

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/Yakud/mumiylovers/entities"
	"github.com/Yakud/mumiylovers/systems"
)

type myScene struct{}

// Type uniquely defines your game type
func (*myScene) Type() string { return "mumiylovers" }

// Preload is called before loading any assets from the disk,
// to allow you to register / queue them
func (*myScene) Preload() {
	engo.Files.Load("mummiy.png", "heart_bullet.png", "heart_bullet_2.png", "fishka1.png")
}

var (
	WalkActionAnimation *common.Animation
	actions             []*common.Animation
)

// Setup is called before the main loop starts. It allows you
// to add entities and systems to your Scene.
func (t *myScene) Setup(world *ecs.World) {
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&systems.AnchorSpriteSystem{})

	rotatableEntityList := make([]*entities.RotatableEntity, 0)

	for i := 1; i <= 4; i++ {
		for j := 1; j <= 3; j++ {
			rotatableEntity := entities.NewRotatableEntity()
			rotatableEntity.SpaceComponent.Position.X = 200 * float32(i)
			rotatableEntity.SpaceComponent.Position.Y = 200 * float32(j)
			rotatableEntityList = append(rotatableEntityList, rotatableEntity)
		}
	}

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			for _, entity := range rotatableEntityList {
				sys.Add(&entity.BasicEntity, &entity.RenderComponent, &entity.SpriteSpaceComponent)
			}
		case *systems.AnchorSpriteSystem:
			for _, entity := range rotatableEntityList {
				sys.Add(&entity.AnchorSpriteComponent)
			}
		}
	}

	//world.AddSystem(&common.MouseSystem{})
	//world.AddSystem(&systems.PlayerMovementSystem{})
	//world.AddSystem(systems.NewHeartBulletGunSystem())
	//world.AddSystem(&common.AnimationSystem{})
	//world.AddSystem(&systems.EnemyLogicSystem{})

	//engo.Input.RegisterButton("MoveUp", engo.ArrowUp)
	//engo.Input.RegisterButton("MoveDown", engo.ArrowDown)
	//engo.Input.RegisterButton("MoveLeft", engo.ArrowLeft)
	//engo.Input.RegisterButton("MoveRight", engo.ArrowRight)
	//
	//world.AddSystem(&common.RenderSystem{})
	//world.AddSystem(&common.MouseSystem{})
	//world.AddSystem(&systems.PlayerMovementSystem{})
	//world.AddSystem(systems.NewHeartBulletGunSystem())
	//world.AddSystem(&common.AnimationSystem{})
	//world.AddSystem(&systems.EnemyLogicSystem{})
	//
	//mumiy := entities.NewMumiy(world)
	//
	//// todo enemy
	//hero := entities.CreateEnemy(engo.Point{200, 200})
	//hero2 := entities.CreateEnemy(engo.Point{450, 250})
	//hero3 := entities.CreateEnemy(engo.Point{500, 300})
	//
	//for _, system := range world.Systems() {
	//	switch sys := system.(type) {
	//	case *common.RenderSystem:
	//		sys.Add(&mumiy.BasicEntity, &mumiy.RenderComponent, &mumiy.SpaceComponent)
	//		sys.Add(&hero.BasicEntity, &hero.RenderComponent, &hero.SpaceComponent)
	//		sys.Add(&hero2.BasicEntity, &hero2.RenderComponent, &hero2.SpaceComponent)
	//		sys.Add(&hero3.BasicEntity, &hero3.RenderComponent, &hero3.SpaceComponent)
	//	case *systems.PlayerMovementSystem:
	//		sys.Add(mumiy)
	//	case *systems.HeartBulletGunSystem:
	//		sys.Add(world, mumiy)
	//	case *common.AnimationSystem:
	//		sys.Add(&hero.BasicEntity, &hero.AnimationComponent, &hero.RenderComponent)
	//		sys.Add(&hero2.BasicEntity, &hero2.AnimationComponent, &hero2.RenderComponent)
	//		sys.Add(&hero3.BasicEntity, &hero3.AnimationComponent, &hero3.RenderComponent)
	//	case *systems.EnemyLogicSystem:
	//		sys.Add(hero)
	//		sys.Add(hero2)
	//		sys.Add(hero3)
	//		sys.AddTarget(&mumiy.Pos)
	//	}
	//}

}

func main() {
	opts := engo.RunOptions{
		Title:  "Mumiy Lovers",
		Width:  1024,
		Height: 768,
		VSync:  true,
	}
	engo.Run(opts, &myScene{})
}
