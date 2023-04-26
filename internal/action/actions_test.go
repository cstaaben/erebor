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

package action_test

import (
	"testing"

	"github.com/cstaaben/erebor/internal/action"
)

// TestImplementsAction tests whether the actions defined in actions implement the [Action] interface.
func TestImplementsAction(_ *testing.T) {
	_ = action.Action(action.Movement{})
	_ = action.Action(action.Undefined)
	_ = action.Action(action.Escape)
}
