//WARNING! You will not understand the maths in this file, it makes heavy use of 2d geometry with a catch.
//The catch is, the roads rotate around the car. Yes. You read that right. (Don't judge). Youngwiseone knows!

package main

import "fmt"

import (
	"grate/geom"
	"grate/game"
	"grate/graphics"
	"grate/math/random"
	"grate/smooth"
	"grate/game/collision"
)

type RoadTile struct {
	Pos, Angle geom.Number
	Type int
}

//Doubly Linked List of roadtiles.
type Road struct {
	RoadTile //Inheritance.
	
	Previous *Road
	Next *Road
}

const Left, Right = false, true

var NextAngle = random.Number(90)*1/180*π + π/2
var Going = Right

func RandomRoadTile(base RoadTile) RoadTile {
	size := graphics.Image("data/roadtile-1.png").Size
	angle := base.Angle
	
	if base.Angle.F32()+ π/2 <= NextAngle.F32() {
		if Going == Right {
			angle += π/80
		} else {
			NextAngle = random.Number(90)*1/180*π + π/2
			Going = Right
		}
	} else if base.Angle.F32()+ π/2 > NextAngle.F32() {
		if Going == Left {
			angle -= π/80
		} else {
			NextAngle = -random.Number(90)*1/180*π + π/2
			Going = Left
		}
	}
	
	return RoadTile {
		Pos: base.Pos-size.Y()*geom.Angle(base.Angle)+5*y*geom.Angle(base.Angle),
		Type: (base.Type+1)%6,
		Angle: angle,
	}
}

func (road *Road) Init() {
	size := graphics.Image("data/roadtile-1.png").Size
	var track *Road = new(Road)
	track.Next = &Road{
			Previous: track,
			RoadTile: RoadTile{ Pos: 500 },
		}
	road.Next = track
	for i := size.Y(); i.Int() < game.Height().Int(); i += size.Y() {
		track.Next = &Road{
			Previous: track,
			RoadTile: RoadTile{ Pos:i + 500 },
		}
		track = track.Next
	}
}

var count = 0

func (road *Road) Travel(pos, speed, angle geom.Number) bool {
	size := graphics.Image("data/roadtile-1.png").Size
	
	var track = road
	var last *Road
	
	if track.Next.Next.Pos.Y().F32() > -size.X().F32() {
		track.Next.Next = &Road{
			Next: track.Next.Next,
			RoadTile: RandomRoadTile(track.Next.Next.RoadTile),
		}
	}
	
	count = 0
	
	for road != nil && road.Next != nil {
		last = road
		road = road.Next
		
		road.Pos  = smooth.Move(road.Pos,  (geom.Angle(angle+π/2)*speed).Y())
		
		roadLeftStart 	:= road.Pos-(size/2-10)*geom.Angle(road.Angle)
		roadLeftEnd   	:= road.Pos+(size.Y()/2-size.X()/2+10)*geom.Angle(road.Angle)
		roadRightStart 	:= road.Pos+(size/2-10)*geom.Angle(road.Angle)
		roadRightEnd   	:= road.Pos+(-size.Y()/2+size.X()/2-10)*geom.Angle(road.Angle)
		
		a := collision.LineToCircle(roadLeftStart, roadLeftEnd, pos, 32)
		b := collision.LineToCircle(roadRightStart, roadRightEnd, pos, 32)
		
		if a || b {
			return true
		}
		
		//WHAT THE HELL, HACK ALERT, DO NOT TRY AND FIGURE THIS OUT!
		if road.Pos.Y().Int()-size.X().Int() > game.Height().Int() && road.Pos.X().Int() != 0 {
			road.Next = track.Next.Next
			road.RoadTile = RandomRoadTile(track.Next.Next.RoadTile)
			
			track.Next.Next = road
			
			last.Next = nil
			track.Travel(0, 0, 0)
			return false
		}
	}
	
	return false
}

func (tile *RoadTile) Draw(n geom.Number) {
	if tile.Pos.X() == 0 {
		return
	}
	count++
	graphics.Image("data/roadtile-"+fmt.Sprint(tile.Type)+".png").
		DrawRotated(tile.Pos, geom.Angle(tile.Angle))
	
	if Debug {
		graphics.SetColor(graphics.Red)
		size := graphics.Image("data/roadtile-1.png").Size
		graphics.Line(tile.Pos-(size/2-10)*geom.Angle(tile.Angle), tile.Pos+(size.Y()/2-size.X()/2+10)*geom.Angle(tile.Angle))
		graphics.Line(tile.Pos+(size/2-10)*geom.Angle(tile.Angle), tile.Pos+(-size.Y()/2+size.X()/2-10)*geom.Angle(tile.Angle))
	}
}

func (road *Road) Draw(n geom.Number) {
	road.RoadTile.Draw(n)
	if road.Next != nil {
		road.Next.Draw(n+1)
	}
}
