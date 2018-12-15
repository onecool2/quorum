// Copyright 2017 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

// puppeth is a command to assemble and maintain private networks.
package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"gopkg.in/urfave/cli.v1"
)

// main is just a boring entry point to set up the CLI app.
func main() {
	app := cli.NewApp()
	app.Name = "puppeth"
	app.Usage = "assemble and maintain private Ethereum networks"
	app.Flags = []cli.Flag{
                cli.IntFlag{
			Name:  "newgenesis",
			Usage: "Configure new genesis",
               	},
		cli.StringFlag{
			Name:  "network",
			Usage: "name of the network to administer",
               	},
                cli.StringFlag{
			Name:  "consensys",
			Usage: "Which consensus engine to use",
               	},
                cli.IntFlag{
			Name:  "period",
			Usage: "How many seconds should blocks take",
               	},
                cli.StringFlag{
			Name:  "sealaccount",
			Usage: "Which accounts are allowed to seal",
               	},
                cli.StringFlag{
			Name:  "prefundedaccount",
			Usage: "Which accounts should be pre-funded",
               	},
                cli.IntFlag{
			Name:  "networkID",
			Usage: "Specify your chain/network ID if you want an explicit one",
               	},
		cli.IntFlag{
			Name:  "loglevel",
			Value: 4,
			Usage: "log level to emit to the screen",
		},
	}
	app.Action = func(c *cli.Context) error {
		// Set up the logger to print everything and the random generator
		log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(c.Int("loglevel")), log.StreamHandler(os.Stdout, log.TerminalFormat(true))))
		rand.Seed(time.Now().UnixNano())

		// Start the wizard and relinquish control
		makeWizard(c.String("network"), c.String("consensys"), c.Uint64("period"), c.String("sealaccount"), c.String("prefundedaccount"), c.Uint64("networkID"), c.Uint64("newgenesis")).run()
		return nil
	}
	app.Run(os.Args)
}
