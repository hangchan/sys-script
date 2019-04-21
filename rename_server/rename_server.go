package main

import (
	"fmt"
	co "github.com/hangchan/sys-script/rename_server/internal/rename_cobbler"
	ic "github.com/hangchan/sys-script/rename_server/internal/rename_icinga"
	rt "github.com/hangchan/sys-script/rename_server/internal/rename_racktable"
)

var (
	oldServer string = "old101.example.com"
	newServer string = "new101.example.com"
)

func main() {
	fmt.Println("test")
	rt.RenameRacktable(oldServer, newServer)
	fmt.Println(ic.RenameIcinga(oldServer, newServer))
	co.RenameCobbler(oldServer, newServer)
}

// Get old and new server name for arg
func getArg(oldServer string, newServer string) string {
	// fmt.Printf("old:", oldServer, "new:", newServer)
	return newServer
}

// Replace entries in racktable

// Replace entries in icinga
