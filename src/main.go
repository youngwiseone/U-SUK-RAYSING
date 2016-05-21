package main

import "grate/graphics"
import "grate/graphics/text"
import "grate/game"
import "grate/game/camera"
import "grate/input"

import (
	"grate/smooth"
	"grate/geom"
)

var Player = new(Car)
var RaceTrack = new(Road)

var ScrollingBackground geom.Number

func init() {
	game.Fullscreen = true
}

func load() {
	graphics.LoadImage("data/car.png")
	graphics.LoadImage("data/road.png")
	
	Player.Speed = 0
	Player.Pos = game.Height()/2 + 400
	
	graphics.SetBackgroundColor(graphics.Green)

	RaceTrack.Init()
	
	camera.Track(&Player.Pos)
}

func update() {
	if input.KeyIsDown(input.KeyEscape) {
		game.End()
	}
	Player.Update()
	RaceTrack.Travel(Player.Speed*game.DeltaTime(), Player.Angle)
	
	ScrollingBackground = smooth.Move(ScrollingBackground, 1i*Player.Speed)
	if (ScrollingBackground-game.Height()).Y().Int() >= 0 {
		ScrollingBackground = 0
	}
	
	camera.Update()
}

func draw() {
	//width := graphics.Image("data/road.png").Size.X()/2
	//graphics.Image("data/road.png").
		//Tile(width+ScrollingBackground-game.Height(), width+game.Height())
	//graphics.Image("data/road.png").
		//Tile(width+ScrollingBackground, width+game.Height())
	
	camera.Start()
	{
		RaceTrack.Draw(0)
	
		Player.Draw()
	}
	camera.Stop()
	
	text.Print(0, "Kph:",Player.Speed.Int()/6)	
		text.PrintCenter("This will be a racing game! (Press ESC to Quit)", count)
	
}
