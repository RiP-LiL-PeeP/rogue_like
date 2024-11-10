package controls

import (
	input "github.com/quasilyte/ebitengine-input"
)

const (
	ActionNone input.Action = iota
	ActionMoveRight
	ActionMoveDown
	ActionMoveLeft
	ActionMoveUp

	ActionConfirm
	ActionRestart
)

var DefaultKeymap = input.Keymap{
	ActionMoveRight: {
		input.KeyRight,
		input.KeyD,
		input.KeyGamepadRight,
	},
	ActionMoveDown: {
		input.KeyDown,
		input.KeyS,
		input.KeyGamepadDown,
	},
	ActionMoveLeft: {
		input.KeyLeft,
		input.KeyA,
		input.KeyGamepadLeft,
	},
	ActionMoveUp: {
		input.KeyUp,
		input.KeyW,
		input.KeyGamepadUp,
	},

	ActionConfirm: {
		input.KeyEnter,
		input.KeyGamepadStart,
	},
	ActionRestart: {
		input.KeyWithModifier(input.KeyR, input.ModControl),
		input.KeyGamepadBack,
	},
}
