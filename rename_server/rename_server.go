package main

import (
	"flag"
	"fmt"
	im "github.com/hangchan/sys-script/rename_server/internal/rename_imm"
)

var (
	oldServer = flag.String("oldServer", "old101", "the old server name")
	newServer = flag.String("newServer", "new101", "the new server name")
	oldIp     = flag.String("oldIp", "192.168.127.101", "the old ip address")
	newIp     = flag.String("newIp", "172.16.127.101", "the new ip address")
)

func main() {
	flag.Parse()

	//rt.RenameRacktable(oldServer, newServer)
	//fmt.Println(ic.RenameIcinga(oldServer, newServer))
	//co.RenameCobbler(oldServer, newServer)
	im.RenameImm(&oldServer, &newServer, &oldIp, &newIp)
}

// Replace entries in racktable

// Replace entries in icinga
