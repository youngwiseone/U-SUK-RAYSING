package main

import "grate/geom"
import "grate/graphics"
import "grate/smooth"
import "grate/math/random"

type copCar struct {
	Pos,  Angle
}
var randomSpeed int

func update() {
	if copCar.Pos.X != car.Pos.X{
		copCar.angleTo(car.Pos)
		copCar.Pos.X = smooth.Move(car.Pos.X, randomSpeed)
		randomSpeed = random.Number((50)+50)/car.Speed
	}
	if copCar.Pos.Y != car.Pos.Y{
		copCar.Pos.Y = smooth.Move(car.Pos.Y, 10)
	}
}