package mpd

import (
	"testing"
)

func TestPlayId(t *testing.T) {
	client := createClient()
	client.Connect()
	err := client.PlayId(5)
	if err != nil {
		t.Errorf("PlayId() returned error: %v", err)
	}

	expectedFile := "5.mp3"
	data, err := client.CurrentSong()
	actualFile := data["file"]

	if expectedFile != actualFile {
		t.Errorf("Expected file: %s, got: %s", expectedFile, actualFile)
	}
}

func TestPlayPoisition(t *testing.T) {
	client := createClient()
	client.Connect()
	err := client.PlayPosition(4)
	if err != nil {
		t.Errorf("PlayPosition() returned error: %v", err)
	}

	expectedFile := "5.mp3"
	data, err := client.CurrentSong()
	actualFile := data["file"]

	if expectedFile != actualFile {
		t.Errorf("Expected file: %s, got: %s", expectedFile, actualFile)
	}
}

func TestPrevious(t *testing.T) {
	client := createClient()
	client.Connect()
	err := client.Previous()
	if err != nil {
		t.Errorf("Next() returned error: %v", err)
	}

	expectedFile := "9.mp3"
	data, err := client.CurrentSong()
	actualFile := data["file"]

	if expectedFile != actualFile {
		t.Errorf("Expected file: %s, got: %s", expectedFile, actualFile)
	}
}

func TestNext(t *testing.T) {
	client := createClient()
	client.Connect()
	err := client.Next()
	if err != nil {
		t.Errorf("Next() returned error: %v", err)
	}

	expectedFile := "2.mp3"
	data, err := client.CurrentSong()
	actualFile := data["file"]

	if expectedFile != actualFile {
		t.Errorf("Expected file: %s, got: %s", expectedFile, actualFile)
	}
}
