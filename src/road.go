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

func (road *Road) Travel(speed geom.Number) {
	size := graphics.Image("data/roadtile.png").Size

	var track = road
	var last *Road
	for road != nil && road.Next != nil {
		last = road
		road = road.Next
		
		road.Pos += speed*y
		
		if road.Pos.Y().Int() > game.Height().Int() {
			road.Next = track.Next.Next
			road.RoadTile = RoadTile{ Pos: track.Next.Next.Next.Pos-size.Y() }
			
			track.Next.Next = road
			
			last.Next = nil
			return
		}
	}
}

func (tile *RoadTile) Draw() {
	graphics.Image("data/roadtile.png").
		DrawRotated(tile.Pos, tile.Angle)
}

func (road *Road) Draw() {
	road.RoadTile.Draw()
	if road.Next != nil {
		road.Next.Draw()
	}
}
