package game

import (
	input "github.com/quasilyte/ebitengine-input"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gscene"
)

type Context struct {
	Input        *input.Handler
	InputSystem  input.System
	Loader       *resource.Loader
	windowHeight int
	windowWidth  int
}

func ChangeScene[T any](ctx *Context, c gscene.Controller[T]) {
	s := gscene.NewRootscene[T](c)
	ctx.scene = s
}

func (ctx *Context) CurrentScene() gscene.GameRunner{
	
}