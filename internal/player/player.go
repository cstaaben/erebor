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

// Package player defines the player object and methods for the player to interact with Erebor.
package player

import (
	"github.com/anaseto/gruid"

	"github.com/cstaaben/erebor/internal/action"
)

// New creates a new Player instance, initialized with name. The coordinates are initially (0,0).
func New(name string) *Player {
	return &Player{
		name:   name,
		symbol: '@',
	}
}

// Player is the player's character. It stores data such as the player's current position in the X/Y coordinate plane,
// their chosen name, their chosen symbol and color. It implements [gruid.Model]
// as it is a model for the player's character.
type Player struct {
	x    int
	y    int
	name string

	symbol byte
	color  gruid.Color
}

func (p *Player) Update(msg gruid.Msg) gruid.Effect {
	// TODO implement me
	panic("implement me")
}

func (p *Player) Draw() gruid.Grid {
	// TODO implement me
	panic("implement me")
}

// Color returns the player's chosen symbol color.
func (p *Player) Color() gruid.Color {
	return p.color
}

// SetColor sets the color to be used for the player's symbol.
func (p *Player) SetColor(color gruid.Color) {
	p.color = color
}

// Symbol returns the symbol used to represent the player's character.
func (p *Player) Symbol() byte {
	return p.symbol
}

// SetSymbol sets the symbol to be used to represent the player's character.
func (p *Player) SetSymbol(symbol byte) {
	p.symbol = symbol
}

// X returns the player's current x coordinate.
func (p *Player) X() int {
	return p.x
}

// SetX sets the player's x coordinate to x.
func (p *Player) SetX(x int) {
	p.x = x
}

// Y returns the player's current y coordinate.
func (p *Player) Y() int {
	return p.y
}

// SetY sets the player's y coordinate to y.
func (p *Player) SetY(y int) {
	p.y = y
}

// Move applies the delta x and y values from act to the player's current x and y coordinates.
func (p *Player) Move(movement action.Movement) {
	p.x += movement.DeltaX
	p.y += movement.DeltaY
}

// Name returns the player's name.
func (p *Player) Name() string {
	return p.name
}

// SetName sets the player's name to name.
func (p *Player) SetName(name string) {
	p.name = name
}
