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

// Package game implements a controller for the game loop
package game

import (
	"github.com/anaseto/gruid"
	"github.com/anaseto/gruid/paths"
	"github.com/anaseto/gruid/rl"

	"github.com/cstaaben/erebor/internal/config"
	"github.com/cstaaben/erebor/internal/mapgen"
	"github.com/cstaaben/erebor/internal/player"
)

// Game is the controller for Erebor game loop.
// It contains references to all elements of the current game: the base grid the game is drawn upon, the player,
// the player's FOV, any enemies that are present, the current path, etc.
type Game struct {
	grid       gruid.Grid
	player     *player.Player
	path       []gruid.Point
	pathFinder *paths.PathRange
	// mapGrid    rl.Grid
	enemies []any
	fov     *rl.FOV
}

// NewGame returns an initialized Game. It uses config to determine screen size and bounds.
func NewGame(config *config.Config) *Game {
	grid := gruid.NewGrid(int(config.Window.ScreenWidth), int(config.Window.ScreenHeight))
	return &Game{
		grid:       grid,
		player:     player.New(""),
		path:       make([]gruid.Point, 0),
		pathFinder: paths.NewPathRange(grid.Range()),
		enemies:    make([]any, 0),
		fov:        rl.NewFOV(grid.Range()),
	}
}

// Update is the event handler for Game. It controls how messages from updates and/or interactions are processed.
func (g *Game) Update(msg gruid.Msg) gruid.Effect {
	switch msg := msg.(type) {
	case gruid.MsgInit:
		g.initialize()
	case gruid.MsgQuit:
		// TODO: confirm game exit before end
		// g.confirmExit()
		return gruid.End()
	case gruid.MsgKeyDown:
		g.handleKeyPress(msg)
	}

	return nil
}

// Draw defines how the game elements are drawn on the screen.
func (g *Game) Draw() gruid.Grid {
	return g.grid
}

func (g *Game) initialize() {
	screenSize := g.grid.Size()
	g.grid = mapgen.GenerateDngnMap(screenSize.X, screenSize.Y)
}

func (g *Game) confirmExit() { // nolint:unused
	// TODO: implement me
	panic("implement me")
}

func (g *Game) handleKeyPress(_ gruid.Msg) {
	// TODO: implement me
	panic("implement me")
}
