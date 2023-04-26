/*
 * Copyright (c) 2023. Corbin Staaben - All Rights Reserved.
 *
 * This file is part of Erebor.
 *
 * Erebor is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
 *
 * Erebor is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty
 * of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License along with Erebor. If not,
 * see <https://www.gnu.org/licenses/>.
 */

// Package handler defines handlers for events.
package handler

import (
	"strings"

	"github.com/Joakker/tcod-go/tinput"

	"github.com/cstaaben/erebor/internal/action"
)

// Event handles user input events. It translates a given input code into an
// [action.Action] that can be interpreted by the game loop.
func Event(input tinput.Input) action.Action {
	switch input.GetVk() {
	case tinput.KeyEscape:
		return action.Escape
	case tinput.KeyLeft:
		return action.MovementLeft
	case tinput.KeyRight:
		return action.MovementRight
	case tinput.KeyUp:
		return action.MovementUp
	case tinput.KeyDown:
		return action.MovementDown
	case tinput.KeyChar:
		return wasdMovement(input.GetC())
	case tinput.KeyText:
		return wasdMovement(strings.ToLower(input.GetText())[0])
	default:
		return action.Undefined
	}
}

func wasdMovement(input byte) action.Action {
	switch input {
	case 'a':
		return action.MovementLeft
	case 'd':
		return action.MovementRight
	case 'w':
		return action.MovementUp
	case 's':
		return action.MovementDown
	default:
		return action.Undefined
	}
}
