package main

import (
	"log"
	assets "rogue_like/internal/assets"

	"github.com/hajimehoshi/ebiten/v2"
	audio "github.com/hajimehoshi/ebiten/v2/audio"
	resource "github.com/quasilyte/ebitengine-resource"
)

func main() {
	// Создаём загрузчик ресурсов
	loader := createLoader()
	// Регистрируем ресурсы изображений
	assets.RegisterResources(loader)

	g := &myGame{
		windowWidth:  320,
		windowHeight: 240,
		loader:       loader, // Добавляем загрузчик в структуру
	}

	ebiten.SetWindowSize(g.windowWidth, g.windowHeight)
	ebiten.SetWindowTitle("Ebitengine Quest")

	// Запускаем игру
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

type myGame struct {
	windowWidth  int
	windowHeight int
	loader       *resource.Loader // Добавляем loader в структуру
}

func (g *myGame) Update() error {
	return nil
}

func (g *myGame) Draw(screen *ebiten.Image) {
	witch := g.loader.LoadImage(assets.ImageWitch).Data
	var options ebiten.DrawImageOptions
	screen.DrawImage(witch, &options)
}

func (g *myGame) Layout(w, h int) (int, int) {
	// Layout - тема для продвинутых, поэтому нам пока
	// достаточно считать, что screen size = window size.
	return g.windowWidth, g.windowHeight
}

func createLoader() *resource.Loader {
	sampleRate := 44100
	audioContext := audio.NewContext(sampleRate)
	loader := resource.NewLoader(audioContext)
	loader.OpenAssetFunc = assets.OpenAsset
	return loader
}
