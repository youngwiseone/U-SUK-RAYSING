package main

import "grate/geom"
import "grate/graphics"
import "grate/smooth"
import "grate/math/random"

type CopCar struct {
	Pos,  Angle geom.Number
}

var randomSpeed geom.Number = 0.5

func (copCar *CopCar) Update() {
	car := Player
	copCar.Angle = copCar.Pos.AngleTo(car.Pos)
	copCar.Pos = smooth.Move(copCar.Pos, geom.Angle(copCar.Angle)*car.Speed*1.1)
}

func (car *CopCar) Draw() {
	graphics.Image("data/copCar.png").
		DrawRotated(car.Pos, geom.Angle(car.Angle))
	
	if Debug {
		graphics.SetColor(graphics.RGBA(0,0,100,50))
		graphics.Circle(car.Pos, 32, true)
	}
}
