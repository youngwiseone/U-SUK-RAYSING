package main

import "grate/geom"
import "grate/graphics"
import "grate/smooth"
import "grate/math/random"

type CopCar struct {
	Pos,  Angle geom.Number
}

var randomSpeed geom.Number

func (copCar *CopCar) Update() {
	car := Player
	if copCar.Pos != car.Pos {
		copCar.Pos.AngleTo(car.Pos)
		copCar.Pos = smooth.Move(car.Pos, randomSpeed)
		randomSpeed = random.Number((50)+50)/car.Speed
	}
}

func (car *CopCar) Draw() {
	graphics.Image("data/copcar.png").
		DrawRotated(car.Pos, geom.Angle(car.Angle))
	
	if Debug {
		graphics.SetColor(graphics.RGBA(0,0,100,50))
		graphics.Circle(car.Pos, 32, true)
	}
}
