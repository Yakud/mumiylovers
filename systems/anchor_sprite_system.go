package systems

import (
	"engo.io/ecs"
	"github.com/Yakud/mumiylovers/components"
)

type AnchorSpriteSystem struct {
	entities []*components.AnchorSpriteComponent

	timeUpdate float32
}

func (t *AnchorSpriteSystem) Add(entity ...*components.AnchorSpriteComponent) {
	t.entities = append(t.entities, entity...)
}

// Remove is called whenever an Entity is removed from the World, in order to remove it from this sytem as well
func (t *AnchorSpriteSystem) Remove(basic ecs.BasicEntity) {
	var d int = -1
	for index, entity := range t.entities {
		if entity.ID() == basic.ID() {
			d = index
			break
		}
	}
	if d >= 0 {
		t.entities = append(t.entities[:d], t.entities[d+1:]...)
	}
}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (t *AnchorSpriteSystem) Update(dt float32) {
	t.timeUpdate += dt
	for _, entity := range t.entities {
		entity.Update()
	}
}
