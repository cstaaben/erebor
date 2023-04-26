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

package config

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"
)

func New(filepath string) (*Config, error) {
	fileBytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("reading config: %w", err)
	}

	c := new(Config)
	err = yaml.Unmarshal(fileBytes, c)
	if err != nil {
		return nil, fmt.Errorf("parsing config: %w", err)
	}

	return c, nil
}

// Config combines all configuration destinations into a single object.
type Config struct {
	User   User   `yaml:"user"`
	Window Window `yaml:"window"`
}

// User is the configuration settings specific to the user (e.g. keybindings)
type User struct {
	Input Input `yaml:"input"`
}

// Input is the configuration settings for the user's input preferences.
type Input struct {
	// KeyUp is the key used for moving the character or cursor up on the screen.
	KeyUp string `yaml:"key_up"`
	// KeyDown is the key used for moving the character or cursor down on the screen.
	KeyDown string `yaml:"key_down"`
	// KeyRight is the key used for moving the character or cursor right on the screen.
	KeyRight string `yaml:"key_right"`
	// KeyLeft is the key used for moving the character or cursor left on the screen.
	KeyLeft string `yaml:"key_left"`
	// KeyAction is the key used to perform actions.
	KeyAction string `yaml:"key_action"`
}

// Window is the configuration settings for the game window.
type Window struct {
	// Fullscreen enables displaying the window in fullscreen.
	// This overrides all other configurations relating to screen size.
	Fullscreen bool `yaml:"fullscreen"`
	// LockDimensionsToDisplayRatio enables locking the display dimensions to the chosen display ratio.
	// This cannot be combined with Fullscreen.
	LockDimensionsToDisplayRatio bool `yaml:"lock_display_ratio"`
	// DisplayRatio is the chosen display ratio. This cannot be combined with Fullscreen.
	DisplayRatio string `yaml:"display_ratio"`
	// ScreenHeight is the height of the screen. This cannot be combined with Fullscreen.
	ScreenHeight int64 `yaml:"screen_height"`
	// ScreenWidth is the width of the screen. This cannot be combined with Fullscreen.
	ScreenWidth int64 `yaml:"screen_width"`
}
