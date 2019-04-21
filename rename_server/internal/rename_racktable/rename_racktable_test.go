package rename_racktable

import "testing"

func TestRenameRacktable(t *testing.T) {
	oldServer := "old101"
	newServer := "new101"

	result := RenameRacktable(oldServer, newServer)
	if result != newServer {
		t.Errorf("newServer is incorrect, got: %s, want: %s", result, newServer)
	}
}
