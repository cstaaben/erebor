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

// Package mapgen defines logic for generating maps.
package mapgen

import (
	"github.com/anaseto/gruid"
	"github.com/solarlune/dngn"
)

// const roomMargin = 10

// GenerateDngnMap uses [dngn] to generate a map, then converts that into a [rl.Grid] that can be used in Erebor.
func GenerateDngnMap(screenWidth int, screenHeight int) gruid.Grid {
	layout := dngn.NewLayout(screenWidth-8, screenHeight-8)
	selection := layout.Select() // full selection
	// mapCenter := layout.Center()

	opts := dngn.NewDefaultBSPOptions()
	opts.SplitCount = 100
	opts.MinimumRoomSize = 5

	_ = layout.GenerateBSP(opts)

	// start := rooms[0]
	//
	// for _, r := range rooms {
	// 	roomCenter := r.Center()
	//
	// 	if roomCenter.X > mapCenter.X-roomMargin && roomCenter.X < mapCenter.X+roomMargin && roomCenter.Y > mapCenter.Y-roomMargin && roomCenter.Y < mapCenter.Y+roomMargin {
	// 		start = r
	// 		break
	// 	}
	// }

	// // is it necessary to trim rooms more than 5 hops away? The dngn example has it,
	// but is that just for a more concise example when displayed?
	// for _, r := range rooms {
	// 	hops := r.CountHopsTo(start)
	// }

	// fill outer walls
	selection.Remove(selection.FilterByArea(1, 1, layout.Width-2, layout.Height-2)).Fill('x')
	// Add alternate floor tile
	selection.FilterByRune(' ').FilterByPercentage(0.1).Fill('.')

	grid := gruid.NewGrid(layout.Width, layout.Height)
	grid.Map(func(point gruid.Point, cell gruid.Cell) gruid.Cell {
		return gruid.Cell{
			Rune:  layout.Get(point.X, point.Y),
			Style: cell.Style,
		}
	})

	return grid
}
