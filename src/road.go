//WARNING! You will not understand the maths in this file, it makes heavy use of 2d geometry with a catch.
//The catch is, the roads rotate around the car. Yes. You read that right. (Don't judge). Youngwiseone knows!

package main

import "fmt"

import (
	"grate/geom"
	"grate/game"
	"grate/graphics"
	"grate/math/random"
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

func RandomRoadTile(base RoadTile) RoadTile {
	size := graphics.Image("data/roadtile-1.png").Size
	angle := base.Angle+(random.Number(3)-2)*π/80
	return RoadTile {
		Pos: base.Pos-size.Y()*geom.Angle(base.Angle),
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

func (road *Road) Travel(speed geom.Number) {
	size := graphics.Image("data/roadtile-1.png").Size

	var track = road
	var last *Road
	
	if track.Next.Next.Pos.Y().F32() > 0 {
		track.Next.Next = &Road{
			Next: track.Next.Next,
			RoadTile: RoadTile{ Pos: track.Next.Next.Pos-size.Y() },
		}
	}
	
	count = 0
	
	for road != nil && road.Next != nil {
		last = road
		road = road.Next
		
		road.Pos += speed*y
		
		//WHAT THE HELL, HACK ALERT, DO NOT TRY AND FIGURE THIS OUT!
		if road.Pos.Y().Int()-road.Pos.X().Int() > game.Height().Int() && road.Pos.X().Int() > 0 {
			road.Next = track.Next.Next
			road.RoadTile = RandomRoadTile(track.Next.Next.RoadTile)
			
			track.Next.Next = road
			
			last.Next = nil
			track.Travel(0)
			return
		}
	}
}

func (tile *RoadTile) Draw(n geom.Number) {
	if tile.Pos.X() == 0 {
		return
	}
	count++
	graphics.Image("data/roadtile-"+fmt.Sprint(tile.Type)+".png").
		DrawRotated(tile.Pos, geom.Angle(tile.Angle))
}

func (road *Road) Draw(n geom.Number) {
	road.RoadTile.Draw(n)
	if road.Next != nil {
		road.Next.Draw(n+1)
	}
}
