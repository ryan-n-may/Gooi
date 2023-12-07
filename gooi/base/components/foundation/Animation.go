package foundation

import (
	time "time"
)

type Animation struct {
	animation_function []func()
	animation_speed time.Duration
}

func NewAnimation(functions []func(), time time.Duration) *Animation {
	var animation = Animation{}
	animation.animation_function = functions
	animation.animation_speed = time
	return &animation
}

func (animation *Animation) RunAnimation(alive *bool){
	var kill = func(alive *bool) {
		*alive = false
		return
	}
	if !(*alive) {
		*alive = true
		kill(alive)
		for _, function := range animation.animation_function {
			function()
			time.Sleep(animation.animation_speed * time.Millisecond)
		}
	}
}

