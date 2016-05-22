package main

import (
	"grate/geom"
	"grate/graphics"
	"grate/graphics/text"
	"grate/game"
	"grate/smooth"
	
	"grate/game/camera"
)

var Scene []Scenery

type Scenery struct {
	Pos geom.Number 
	Image *graphics.ImageData
}

func AddToScene(Pos geom.Number, Image *graphics.ImageData) {
	Scene = append(Scene, Scenery{Pos:Pos, Image:Image})
}

func UpdateScene(speed, angle geom.Number) {
	for i, scenery := range Scene {
		if i < len(Scene) {
			Scene[i].Pos  = smooth.Move(scenery.Pos,  (geom.Angle(angle+π/2)*speed).Y())
			
			if scenery.Pos.Y().Int()-scenery.Image.Size.Y().Int() > game.Height().Int() {
				a := Scene
				a[i] = a[len(a)-1] 
				a = a[:len(a)-1]
				Scene = a
			}	
		}
	}
}

func DrawScene() {
	var count int
	for i := len(Scene)-1; i >= 0; i-- {
		Scene[i].Draw()
		count ++
	}
	camera.Stop()
	text.Print(300*y, count)
	camera.Start()
}

func (scenery *Scenery) Draw() {
	scenery.Image.DrawRotated(scenery.Pos, 0)
}
