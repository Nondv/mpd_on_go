package main

import (
	"encoding/json"
	"fmt"
	"github.com/nondv/mpd_on_go/mpd"
)

func main() {
	client := mpd.Client{Host: "raspberry", Port: 6600}
	err := client.Connect()

	if err != nil {
		fmt.Printf("can't connect")
		return
	}
	fmt.Printf("%+v\n\n", client)

	status, _ := client.Status()
	printMapAsJson(status)
	currentSong, _ := client.CurrentSong()
	printMapAsJson(currentSong)
	stats, _ := client.Stats()
	printMapAsJson(stats)
}

// Just print
func printMapAsJson(data map[string]string) {
	marshal, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n\n", marshal)
}
