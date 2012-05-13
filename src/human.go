
package stratum

import (
	"time"
)

type World struct {
	people map[*Human]bool
}

type Skill struct {
	name string // what is it called
	wealthLever float32 // how much does current wealth affect the outcome
	utility int32 // how important is the products of this skill
	difficulty float32 // how hard is this skill to perform untrained
}

type Human struct {
	skills map[*Skill]float32 // the wealth increasing skills
	health float32 // the current health
	wealth float32 // the current wealth
	age time.Duration // the current age
	position complex64 // the position of this human
	world World // the world we live in
}

func NewHuman(world World, position complex64) *Human {
	return &Human{make(map[*Skill]float32), 1.0, 0.0, 0, position, world}
}

func (h Human) act(time time.Duration) {
	
}