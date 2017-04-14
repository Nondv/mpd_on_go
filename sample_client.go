package main

import (
	"fmt"
	"github.com/nondv/mpd_on_go/mpd"
	"github.com/urfave/cli"
	"os"
	"strconv"
)

func main() {
	var port int
	var address string

	app := cli.NewApp()
	app.Name = "gmpc"
	app.Usage = "MPD using mpd_on_go library sample"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "address, a",
			Value:       "localhost",
			Usage:       "MPD address",
			Destination: &address,
		},
		cli.IntFlag{
			Name:        "port, p",
			Value:       6600,
			Usage:       "MPD port",
			Destination: &port,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "play",
			Usage: "resume playing",
			Action: func(c *cli.Context) error {
				client, err := mpd.Connect(address, port)
				if err != nil {
					fmt.Println("Can't connect to MPD")
					return nil
				}

				client.Play()
				printStatus(client)
				return nil
			},
		},
		{
			Name:  "pause",
			Usage: "set on pause",
			Action: func(c *cli.Context) error {
				client, err := mpd.Connect(address, port)
				if err != nil {
					fmt.Println("Can't connect to MPD")
					return nil
				}

				client.Pause()
				printStatus(client)
				return nil
			},
		},
		{
			Name:  "next",
			Usage: "play next track",
			Action: func(c *cli.Context) error {
				client, err := mpd.Connect(address, port)
				if err != nil {
					fmt.Println("Can't connect to MPD")
					return nil
				}

				client.Next()
				printStatus(client)
				return nil
			},
		},
		{
			Name:  "previous",
			Usage: "play previous track",
			Action: func(c *cli.Context) error {
				client, err := mpd.Connect(address, port)
				if err != nil {
					fmt.Println("Can't connect to MPD")
					return nil
				}

				client.Previous()
				printStatus(client)
				return nil
			},
		},
		{
			Name:  "volume",
			Usage: "set volume",
			Action: func(c *cli.Context) error {
				arg := c.Args().First()
				value, err := strconv.ParseInt(arg, 10, 0)
				if err != nil {
					fmt.Println("Bad value")
					return nil
				}

				client, err := mpd.Connect(address, port)
				if err != nil {
					fmt.Println("Can't connect to MPD")
					return nil
				}

				client.SetVolume(int(value))
				printStatus(client)
				return nil
			},
		},
		{
			Name:  "status",
			Usage: "show status",
			Action: func(c *cli.Context) error {
				client, err := mpd.Connect(address, port)
				if err != nil {
					fmt.Println("Can't connect to MPD")
					return nil
				}

				printStatus(client)
				return nil
			},
		},
	}

	app.Run(os.Args)
}

func printStatus(client *mpd.Client) {
	status, _ := client.Status()
	currentSong, _ := client.CurrentSong()
	fmt.Printf("%s - %s\n", currentSong["Artist"], currentSong["Title"])
	fmt.Println(currentSong["Album"])
	fmt.Printf("volume: %s random: %s repeat: %s\n", status["volume"], status["random"], status["repeat"])
}
