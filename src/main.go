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

var Debug bool

var ScrollingBackground geom.Number

func init() {
	game.Fullscreen = true
}

func load() {
	graphics.LoadImage("data/car.png")
	graphics.LoadImage("data/road.png")
	
	Player.Speed = 0
	Player.Pos = game.Height()/2 + 500
	
	graphics.SetBackgroundColor(graphics.Green)

	RaceTrack.Init()
	
	camera.Track(&Player.Pos)
	
	input.OnKeyPress(func(key int) {
		if key == input.KeyTab {
			Debug = !Debug
		}
	})
}

func update() {
	if input.KeyIsDown(input.KeyEscape) {
		game.End()
	}

	if game.Over {
		return
	}
	Player.Update()
	
	if RaceTrack.Travel(Player.Pos, Player.Speed, Player.Angle) {
		game.Over = true
	}
	
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
	text.Print(game.Width()-text.CurrentFont().Width("(Press Tab to debug!)"), "(Press Tab to debug!)")	
	
	if game.Over {
		text.PrintCenter("U SUK @ RAYSING!")
	}
	
}
