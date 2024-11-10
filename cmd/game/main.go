package main

import (
	"log"
	assets "rogue_like/internal/assets"
	"rogue_like/internal/controls"
	"rogue_like/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
	audio "github.com/hajimehoshi/ebiten/v2/audio"
	input "github.com/quasilyte/ebitengine-input"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gmath"
)

type Player struct {
	pos gmath.Vec
	img *ebiten.Image
}

type myGame struct {
	ctx *game.Context
}



func main() {
	// Создаём загрузчик ресурсов
	loader := createLoader()
	// Регистрируем ресурсы изображений
	assets.RegisterResources(loader)

	g := &myGame{
		windowWidth:  800,
		windowHeight: 600,
		loader:       loader, // Добавляем загрузчик в структуру
	}

	ebiten.SetWindowSize(g.windowWidth, g.windowHeight)
	ebiten.SetWindowTitle("Ebitengine Quest")
	g.init()

	g.inputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})

	g.input = g.inputSystem.NewHandler(0, controls.DefaultKeymap)
	// Запускаем игру
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}


func (g *myGame) Update() error {
	g.ctx.InputSystem.Update()
	g.ctx.CurrentScene().Update()
    return nil
}

func (g *myGame) Draw(screen *ebiten.Image) {
	g.ctx.CurrentScene.Draw(screen)
}



func createLoader() *resource.Loader {
	sampleRate := 44100
	audioContext := audio.NewContext(sampleRate)
	loader := resource.NewLoader(audioContext)
	loader.OpenAssetFunc = assets.OpenAsset
	return loader
}
