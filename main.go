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
	engo.Files.Load("mummiy.png", "heart_bullet.png", "heart_bullet_2.png")
}

// Setup is called before the main loop starts. It allows you
// to add entities and systems to your Scene.
func (*myScene) Setup(world *ecs.World) {
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.MouseSystem{})

	mumiy := entities.NewMumiy(world)

	world.AddSystem(systems.NewPlayerMovementSystem(mumiy))
	world.AddSystem(systems.NewHeartBulletGunSystem(world, mumiy))

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&mumiy.BasicEntity, &mumiy.RenderComponent, &mumiy.SpaceComponent)
		}
	}
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
