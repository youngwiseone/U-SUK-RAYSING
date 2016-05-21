package main

import "grate/geom"
import "grate/graphics"
import "grate/input"
import "grate/smooth"

type Car struct {
	Pos, Angle, Speed geom.Number
}

func (car *Car) Update() {
	if input.KeyIsDown(input.KeyW) || input.KeyIsDown(input.KeyUp) {
		car.Pos = smooth.Move(car.Pos, geom.Angle(car.Angle)*car.Speed)
	}
	if input.KeyIsDown(input.KeyA) || input.KeyIsDown(input.KeyLeft) {
		car.Angle = smooth.Move(car.Angle, -π)
	}
	if input.KeyIsDown(input.KeyD) || input.KeyIsDown(input.KeyRight) {
		car.Angle = smooth.Move(car.Angle, π)
	}
}

func (car *Car) Draw() {
	graphics.Image("data/car.png").
		DrawRotated(car.Pos, geom.Angle(car.Angle+π/2))
}
