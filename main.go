package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowWidth  = 800
	windowHeight = 600
)

var (
	isGameRunning = false
	renderer      *sdl.Renderer
	window        *sdl.Window
	playerX       int32
	playerY       int32
)

func setup() {
	playerX, playerY = 0, 0
}

func initializeWindow() bool {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		log.Fatalf("Error initialzing SDL %s", err)
		return false
	}

	window, err := sdl.CreateWindow("", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, windowWidth, windowHeight, sdl.WINDOW_BORDERLESS)
	if err != nil {
		log.Fatalf("Error creating window %s", err)
	}

	renderer, err = sdl.CreateRenderer(window, -1, 0)
	if err != nil {
		log.Fatalf("Error creating renderer %s", err)
	}
	renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)

	return true
}

func update() {
	playerX += 1
	playerY += 1
}

func render() {
	err := renderer.SetDrawColor(0, 0, 0, 255)
	if err != nil {
		log.Fatalf("Error setting draw color %s", err)
	}
	renderer.Clear()

	err = renderer.SetDrawColor(255, 255, 0, 255)
	if err != nil {
		log.Fatalf("Error setting draw color %s", err)
	}
	rect := sdl.Rect{X: playerX, Y: playerY, W: 20, H: 20}
	err = renderer.FillRect(&rect)
	if err != nil {
		log.Fatalf("Eror rendering rectangle %s", err)
	}

	renderer.Present()
}

func processInput() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			isGameRunning = false
		case *sdl.KeyboardEvent:
			if t.Keysym.Sym == sdl.K_ESCAPE {
				isGameRunning = false
			}
		}
	}
}

func main() {
	isGameRunning = initializeWindow()

	for isGameRunning {
		processInput()
		update()
		render()
	}
}
