package main

import "grate/geom"
import "grate/graphics"
import "grate/graphics/text"
import "grate/input"
import "grate/smooth"

type Car struct {
	Pos,  Angle, Speed geom.Number
}

func (car *Car) Update() {
	if input.KeyIsDown(input.KeyW) || input.KeyIsDown(input.KeyUp) {
		car.Speed = smooth.Move(car.Speed, 200)
	}
	if input.KeyIsDown(input.KeyS) || input.KeyIsDown(input.KeyDown) {
		car.Speed = smooth.Move(car.Speed, -500)
		//Car should not go backwards.
		if car.Speed.Int() < 0 {
			car.Speed = 0
		}
	}
	if car.Speed == 0 {
		return
	}
	if input.KeyIsDown(input.KeyA) || input.KeyIsDown(input.KeyLeft) {
		car.Pos = smooth.Move(car.Pos, -car.Speed/2)
		car.Angle = smooth.Move(car.Angle, -π/2*car.Speed/100)
		if car.Angle.F32() < -π/10 {
			car.Angle = -π/10 
		}
	} else if input.KeyIsDown(input.KeyD) || input.KeyIsDown(input.KeyRight) {
		car.Pos = smooth.Move(car.Pos, car.Speed/2)
		car.Angle = smooth.Move(car.Angle, π/2*car.Speed/100)
		if car.Angle.F32() > π/10 {
			car.Angle = π/10
		}
	} else {
		if car.Angle.F32() > 0 {
			car.Angle = smooth.Move(car.Angle, -π/2*car.Speed/100)
			if car.Angle.F32() < 0 {
				car.Angle = 0
			} 
		} else if car.Angle.F32() < 0 {
			car.Angle = smooth.Move(car.Angle, π/2*car.Speed/100)
			if car.Angle.F32() > 0 {
				car.Angle = 0
			}
		}
	}
}

func (car *Car) Draw() {
	text.Print(0, "Kph:",car.Speed.Int()/6)	
	graphics.Image("data/car.png").
		DrawRotated(car.Pos, geom.Angle(car.Angle))
}
