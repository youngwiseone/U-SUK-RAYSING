//WARNING! You will not understand the maths in this file, it makes heavy use of 2d geometry with a catch.
//The catch is, the roads rotate around the car. Yes. You read that right. (Don't judge). Youngwiseone knows!

package main

import (
	"grate/geom"
	"grate/game"
	"grate/graphics"
)

type RoadTile struct {
	Pos, Angle geom.Number
}

//Doubly Linked List of roadtiles.
type Road struct {
	RoadTile //Inheritance.
	
	Previous *Road
	Next *Road
}

func (road *Road) Init() {
	size := graphics.Image("data/roadtile.png").Size
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
	size := graphics.Image("data/roadtile.png").Size

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
		if road.Pos.Y().Int() > game.Height().Int() && road.Pos.X().Int() > 0 {
			road.Next = track.Next.Next
			road.RoadTile = RoadTile{ Pos: track.Next.Next.Pos-size.Y() }
			
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
	graphics.Image("data/roadtile.png").
		DrawRotated(tile.Pos, tile.Angle)
}

func (road *Road) Draw(n geom.Number) {
	road.RoadTile.Draw(n)
	if road.Next != nil {
		road.Next.Draw(n+1)
	}
}
