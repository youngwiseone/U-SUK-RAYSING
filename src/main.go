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
var Cop = new(CopCar)
var RaceTrack = new(Road)

var Debug bool

func init() {
	game.Fullscreen = true
}

func load() {
	graphics.LoadImage("data/car.png")
	graphics.LoadImage("data/road.png")
	
	Player.Speed = 0
	Player.Pos = game.Height()/2 + 500
	
	Cop.Pos = game.Height() + 500
	
	graphics.SetBackgroundColor(graphics.DarkGreen)

	RaceTrack.Init()
	
	camera.Track(&Player.Pos)
	
	input.OnKeyPress(func(key int) {
		if key == input.KeyTab {
			Debug = !Debug
		}
		
		if game.Over && key == input.KeyR {
			RaceTrack = new(Road)
			RaceTrack.Init()
			
			Scene = nil
			
			Player = new(Car)
			Player.Pos = game.Height()/2 + 500
			game.Over = false
			camera.Track(&Player.Pos)
		}
	})
}

var DeathVec geom.Number

func update() {
	if input.KeyIsDown(input.KeyEscape) {
		game.End()
	}

	if game.Over {
		if Player.Speed.Int() > 100*6 {
			Player.Pos = smooth.Move(Player.Pos, geom.Angle(DeathVec)*Player.Speed/2)
			Player.Angle = smooth.Move(Player.Angle, 2*π*Player.Speed/500)
		}
		return
	}
	Player.Update()
	Cop.Update()
	
	UpdateScene(Player.Speed, Player.Angle)
	
	if RaceTrack.Travel(Player.Pos, Player.Speed, Player.Angle) {
		DeathVec = Player.Angle-π/2
		game.Over = true
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
		Cop.Draw()
		
		DrawScene()
	}
	camera.Stop()
	
	text.Print(0, "Kph:",Player.Speed.Int()/6)	
		
	
	if game.Over {
		text.PrintCenter("U SUK @ RAYSING!")
		text.Print(game.Width()-text.CurrentFont().Width("(Press R to restart!)"), "(Press R to restart!)")
		
	} else {
		text.Print(game.Width()-text.CurrentFont().Width("(Press Tab to debug!)"), "(Press Tab to debug!)")
	}
	
}
