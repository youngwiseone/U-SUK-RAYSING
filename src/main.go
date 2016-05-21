package main

import "grate/graphics"
import "grate/graphics/text"
import "grate/game"
import "grate/input"

var Player = new(Car)

func init() {
	game.Fullscreen = true
}

func load() {
	graphics.LoadImage("data/car.png")
	Player.Speed = 500
	Player.Pos = 200 + 300*y
	
	graphics.SetBackgroundColor(graphics.Green)
}

func update() {
	if input.KeyIsDown(input.KeyEscape) {
		game.End()
	}
	Player.Update()
}

func draw() {
	text.PrintCenter("This will be a racing game! (Press ESC to Quit)")
	Player.Draw()
}
