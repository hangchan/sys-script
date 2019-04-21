package rename_icinga

import "testing"

func TestRenameIcinga(t *testing.T) {
	oldServer := "old101"
	newServer := "new101"

	result := RenameIcinga(oldServer, newServer)
	if result != newServer {
		t.Errorf("newServer is incorrect, got: %s, want: %s", result, newServer)
	}
}
