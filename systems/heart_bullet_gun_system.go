package systems

import (
	"math/rand"
	"sync"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/Yakud/mumiylovers/entities"
	"github.com/engoengine/math"
)

// Input keys and move player
type HeartBulletGunSystem struct {
	mumiy *entities.Mumiy
	world *ecs.World

	bullets      []*entities.HeartBullet
	bulletsPool  *sync.Pool
	bulletsMutex *sync.Mutex

	isReadyShot      bool
	bulletReloadTime float32
	timeFromLastShot float32
}

// Remove is called whenever an Entity is removed from the World, in order to remove it from this sytem as well
func (*HeartBulletGunSystem) Remove(ecs.BasicEntity) {

}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (t *HeartBulletGunSystem) Update(dt float32) {
	if t.isReadyShot && engo.Input.Button("Shot").Down() {
		t.CreateBullet()
		t.isReadyShot = false
	}

	if !t.isReadyShot {
		t.timeFromLastShot += dt
		if t.timeFromLastShot >= t.bulletReloadTime {
			t.isReadyShot = true
			t.timeFromLastShot -= t.bulletReloadTime
		}
	}

	for i := len(t.bullets) - 1; i >= 0; i-- {
		bullet := t.bullets[i]

		bullet.Position.Add(engo.Point{
			X: math.Cos(bullet.Direction()*math.Pi/180) * 60 * dt * bullet.Speed(),
			Y: math.Sin(bullet.Direction()*math.Pi/180) * 60 * dt * bullet.Speed(),
		})

		//bullet.Rotation += 60 * dt * bullet.Speed() * 5

		if bullet.Position.X > 1024 || bullet.Position.X < 0 ||
			bullet.Position.Y > 768 || bullet.Position.Y < 0 {
			t.bulletsMutex.Lock()
			t.DestroyBullet(i)
			t.bulletsMutex.Unlock()
		}
	}
}

func (t *HeartBulletGunSystem) CreateBullet() {
	bullet := t.bulletsPool.Get().(*entities.HeartBullet)

	bullet.Position.Set(
		t.mumiy.Pos.X+math.Cos(t.mumiy.Rotation*math.Pi/180)*50,
		t.mumiy.Pos.Y+math.Sin(t.mumiy.Rotation*math.Pi/180)*35,
	)

	bullet.SetDirection(t.mumiy.Rotation)
	bullet.SetSpeed(float32(6 + rand.Intn(10)))
	bullet.SetSpeed(15)

	t.bulletsMutex.Lock()
	t.bullets = append(t.bullets, bullet)
	t.bulletsMutex.Unlock()

	for _, system := range t.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&bullet.BasicEntity, &bullet.RenderComponent, &bullet.SpaceComponent)
		}
	}
}

func (t *HeartBulletGunSystem) DestroyBullet(index int) {
	bullet := t.bullets[index]
	t.bullets = append(t.bullets[:index], t.bullets[index+1:]...)

	t.bulletsPool.Put(bullet)

	for _, system := range t.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Remove(bullet.BasicEntity)
		}
	}
}

func CreateBulletInstance() *entities.HeartBullet {
	return entities.NewHeartBullet()
}

func NewHeartBulletGunSystem(world *ecs.World, mumiy *entities.Mumiy) *HeartBulletGunSystem {
	engo.Input.RegisterButton("Shot", engo.Space)

	return &HeartBulletGunSystem{
		mumiy: mumiy,
		world: world,

		bulletsMutex:     &sync.Mutex{},
		bulletReloadTime: 0.1,

		bulletsPool: &sync.Pool{
			New: func() interface{} { return CreateBulletInstance() },
		},
	}
}
