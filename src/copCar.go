package main

import "grate/geom"
import "grate/graphics"
import "grate/smooth"
import "grate/math/random"
import "grate/time"

type copCar struct {
	Pos,  Angle
}
var randomSpeed int

func update() {
	time.Update()
	if copCar.Pos.Y != car.Pos.Y{
		copCar.Pos.Y = smooth.Move(car.Pos.Y, 10)
	}
	   time.Every(randomSpeed, func() {
        copCar.angleTo(car.Pos)
		copCar.Pos.X = smooth.Move(car.Pos.X, randomSpeed)
		randomSpeed = random.Number((10)/100)*car.Speed
        })
}