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

// Package erebor defines the gameplay loop for Erebor.
package erebor

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/anaseto/gruid"
	gtcell "github.com/anaseto/gruid-tcell"
	tcell "github.com/gdamore/tcell/v2"

	"github.com/cstaaben/erebor/internal/config"
	"github.com/cstaaben/erebor/internal/game"
)

const defaultConfigFile = "$XDG_CONFIG_HOME/erebor/.config.yaml"

var (
	help           bool
	configFilepath string
)

func init() {
	// TODO[erebor#13]: rand.Seed() is deprecated

	// 	rand.Seed(time.Now().UnixNano())
	// 	rand.New(rand.NewSource(time.Now().UnixNano()))

	flag.BoolVar(&help, "help", false, "Display help")
	flag.StringVar(&configFilepath, "config", defaultConfigFile, "Path to the file containing configuration settings.")
}

// Run is the main execution function of erebor.
func Run() error {
	flag.Parse()
	if help {
		printHelp()
		return nil
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := preRun()
	if err != nil {
		return err
	}

	cfg, err := config.New(configFilepath)
	if err != nil {
		return fmt.Errorf("parsing configuration: %w", err)
	}

	g := game.NewGame(cfg)

	driver := gtcell.NewDriver(gtcell.Config{
		DisableMouse: true,
		StyleManager: new(styleManager),
	})
	app := gruid.NewApp(gruid.AppConfig{
		Driver: driver,
		Model:  g,
		Logger: log.Default(),
	})

	return app.Start(ctx)
}

type styleManager struct{}

func (s *styleManager) GetStyle(_ gruid.Style) tcell.Style {
	return tcell.Style{}.Normal()
}

// preRun is a collection of all setup logic before the gameplay loop begins.
func preRun() error {
	// if the specified config file doesn't exist, create it
	_, err := os.Stat(configFilepath)
	if errors.Is(err, os.ErrNotExist) {
		f, err := os.Create(configFilepath)
		if err != nil {
			return fmt.Errorf("creating config file: %w", err)
		}
		f.Close() // nolint:errcheck
	} else if err != nil {
		return fmt.Errorf("fetching config file info: %w", err)
	}

	return nil
}

func printHelp() {
	fmt.Println("help")
}

// func startGameplayLoop(ctx context.Context, root tcod.Console) {
// 	player := player.New("Player")
// 	player.SetX(10)
// 	player.SetY(10)
// 	for !tcod.WindowClosed() {
// 		root.PrintFrame(0, 0, 5, 5, true, "Erebor")
// 		root.PutChar(player.X(), player.Y(), player.Symbol(), tcod.BgNone)
// 		root.SetCellFg(player.X(), player.Y(), player.Color())
//
// 		var input tinput.Input
// 		select {
// 		case <-ctx.Done():
// 			return
// 		default:
// 			input = tinput.NewInput()
// 		}
// 		if input.Check() != tinput.EvKeyPress {
// 			continue
// 		}
//
// 		playerAction := handler.Event(input)
// 		switch playerAction {
// 		case action.Escape:
// 			return
// 		case action.Undefined:
// 			continue
// 		default:
// 			if playerAction.IsMovement() {
// 				player.Move(playerAction.(action.Movement))
// 			}
// 		}
//
// 		root.Clear()
// 		tcod.Flush()
// 	}
// }
