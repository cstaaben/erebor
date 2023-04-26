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

// Package action defines the various actions the player can take.
package action

var (
	// MovementLeft corresponds to a left movement, based on the player's keybindings.
	MovementLeft = Movement{DeltaX: -1}
	// MovementRight corresponds to a right movement, based on the player's keybindings.
	MovementRight = Movement{DeltaX: 1}
	// MovementUp corresponds to an upwards movement, based on the player's keybindings.
	MovementUp = Movement{DeltaY: 1}
	// MovementDown corresponds to a downwards movement, based on the player's keybindings.
	MovementDown = Movement{DeltaY: -1}
	// Undefined represents an unknown action type.
	Undefined = undefined{}
	// Escape represents the escape key being pressed.
	Escape = escape{}
)

// Action is essentially a no-op interface to allow a common type to be surfaced for events.
type Action interface {
	IsMovement() bool
}

// Movement is the change in position from the player's current position.
type Movement struct {
	DeltaX int
	DeltaY int
}

func (_ Movement) IsMovement() bool { return true }

// escape is the action to exit Erebor.
type escape struct{}

func (_ escape) IsMovement() bool { return false }

type undefined struct{}

func (_ undefined) IsMovement() bool { return false }
