package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const Width = int32(160 * 3)
const Height = int32(144 * 3)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("go-sdl2-opengl-example-2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, Width, Height, sdl.WINDOW_OPENGL)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, 0)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, Width, Height)
	if err != nil {
		panic(err)
	}
	defer texture.Destroy()

	pixels, _, _ := texture.Lock(nil)

	for i := range pixels {
		pixels[i] = 180
	}

	texture.Unlock()

	renderer.Copy(texture, nil, nil)
	renderer.Present()

	running := true
	for running {
	}
}
