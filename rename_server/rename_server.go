package main

import (
	"flag"
	im "github.com/hangchan/sys-script/rename_server/internal/rename_imm"
)

var (
	oldServer string
	newServer string
	oldIp     string
	newIp     string
)

func main() {
	flag.StringVar(&oldServer, "oldServer", "old101", "the old server name")
	flag.StringVar(&newServer, "newServer", "new101", "the new server name")
	flag.StringVar(&oldIp, "oldIp", "192.168.127.101", "the old ip address")
	flag.StringVar(&newIp, "newIp", "172.16.127.101", "the new ip address")
	flag.Parse()

	//rt.RenameRacktable(oldServer, newServer)
	//fmt.Println(ic.RenameIcinga(oldServer, newServer))
	//co.RenameCobbler(oldServer, newServer)
	im.RenameImm(oldIp, newIp, oldServer, newServer)
}

// Replace entries in racktable

// Replace entries in icinga
