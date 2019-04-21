package rename_icinga

import "testing"

func TestRenameCobbler(t *testing.T) {
	oldServer := "old101"
	newServer := "new101"

	result := RenameCobbler(oldServer, newServer)
	if result != newServer {
		t.Errorf("newServer is incorrect, got: %s, want: %s", result, newServer)
	}
}
