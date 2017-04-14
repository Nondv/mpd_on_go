package mpd

import (
	"testing"
)

func TestSetVolume(t *testing.T) {
	client := createClient()
	client.Connect()

	err := client.SetVolume(-25)
	if err == nil {
		t.Error("Expected error on negative argument")
	}
	err = client.SetVolume(101)
	if err == nil {
		t.Error("Expected error on argument > 100")
	}

	value := 43
	err = client.SetVolume(value)
	if err != nil {
		t.Errorf("Error returned: %v", err)
	}
	currentVolume, _ := client.currentVolume()
	if currentVolume != value {
		t.Errorf("Wrong volume. Expected: %d, got: %d", value, currentVolume)
	}
}

func TestIncreaseVolume(t *testing.T) {
	client := createClient()
	client.Connect()

	baseVolume := 4
	client.SetVolume(baseVolume)
	diff := 13
	err := client.IncreaseVolume(diff)
	if err != nil {
		t.Errorf("Error returned: %v", err)
	}
	actualVolume, _ := client.currentVolume()
	expectedVolume := baseVolume + diff
	if actualVolume != expectedVolume {
		t.Errorf("Wrong volume. Expected: %d, got: %d", expectedVolume, actualVolume)
	}
}

func TestIncreaseVolumeWithOverflow(t *testing.T) {
	client := createClient()
	client.Connect()

	baseVolume := 99
	client.SetVolume(baseVolume)
	err := client.IncreaseVolume(43)
	if err != nil {
		t.Errorf("Error returned: %v", err)
	}
	actualVolume, _ := client.currentVolume()
	if actualVolume != 100 {
		t.Errorf("Wrong volume. Expected: %d, got: %d", 100, actualVolume)
	}
}

func TestDecreaseVolume(t *testing.T) {
	client := createClient()
	client.Connect()

	baseVolume := 25
	client.SetVolume(baseVolume)
	diff := 13
	err := client.DecreaseVolume(diff)
	if err != nil {
		t.Errorf("Error returned: %v", err)
	}
	actualVolume, _ := client.currentVolume()
	expectedVolume := baseVolume - diff
	if actualVolume != expectedVolume {
		t.Errorf("Wrong volume. Expected: %d, got: %d", expectedVolume, actualVolume)
	}
}

func TestDecreaseVolumeWithOverflow(t *testing.T) {
	client := createClient()
	client.Connect()

	baseVolume := 25
	client.SetVolume(baseVolume)
	err := client.DecreaseVolume(43)
	if err != nil {
		t.Errorf("Error returned: %v", err)
	}
	actualVolume, _ := client.currentVolume()
	if actualVolume != 0 {
		t.Errorf("Wrong volume. Expected: %d, got: %d", 0, actualVolume)
	}
}
