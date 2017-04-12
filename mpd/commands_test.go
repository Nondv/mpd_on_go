package mpd

import (
	"testing"
)

func createClient() Client {
	return Client{Host: "localhost", Port: 6789}
}

func TestConnection(t *testing.T) {
	client := createClient()
	err := client.Connect()
	if err != nil {
		t.Errorf("Connect() returned error: %v", err)
	}

	expectedVersion := "fake-version"
	if client.Version != expectedVersion {
		t.Errorf("Version incorrect. Expected: %s, got: %s", expectedVersion, client.Version)
	}
}

func TestCurrentSong(t *testing.T) {
	client := createClient()
	client.Connect()
	data, err := client.CurrentSong()

	if err != nil {
		t.Errorf("CurrentSong() returned error: %v", err)
	}

	expectations := [12][2]string{
		[2]string{"file", "1.mp3"},
		[2]string{"Id", "1"},
		[2]string{"Last-Modified", "2016-01-01"},
		[2]string{"Artist", "Fake-Artist"},
		[2]string{"AlbumArtist", "Fake-Album-Artist"},
		[2]string{"Title", "Song number 1"},
		[2]string{"Album", "Album number 1"},
		[2]string{"Track", "1"},
		[2]string{"Date", "2001"},
		[2]string{"Genre", "Rock"},
		[2]string{"Time", "240"},
		[2]string{"Pos", "1"},
	}

	for i := range expectations {
		e := expectations[i]
		if data[e[0]] != e[1] {
			t.Errorf("%s incorrect. Expected: %v, got: %v", e[0], e[1], data[e[0]])
		}
	}
}

// Commands like "stats", "status" and so on are pretty like CurrentSong.
// So I guess it's not neccessary to test them.
// But we can test some low-level functions.

func TestParseKeyValue(t *testing.T) {
	expectations := [][]string{
		[]string{"key: value", "key", "value"},
		[]string{"another key: value", "another key", "value"},
		[]string{"key: one : two : three", "key", "one : two : three"},
	}

	for i := range expectations {
		e := expectations[i]
		k, v := parseKeyValue(e[0])
		if k != e[1] || v != e[2] {
			t.Errorf("key-value incorrect. Expected: '%s'-'%s'. Got: '%s'-'%s'", e[1], e[2], k, v)
		}
	}
}

func TestParseAsMap(t *testing.T) {
	input := []string{"key1: value", "another key: value", "key2: one : two : three"}
	expectations := [][]string{
		[]string{"key1", "value"},
		[]string{"another key", "value"},
		[]string{"key2", "one : two : three"},
	}

	result := parseAsMap(input)

	for i := range expectations {
		e := expectations[i]
		if result[e[0]] != e[1] {
			t.Errorf("result[%s] incorrect. Expected: '%s'. Got: '%s'", e[0], e[1], result[e[0]])
		}
	}
}
