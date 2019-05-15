package rename_imm

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func RenameImm(oldIp string, newIp string, oldServer string, newServer string) {
	// check if asu64 exists
	cmd := "/opt/ibm/toolscenter/asu/asu64"
	asu64 := make(map[string]string)
	asu64 = map[string]string{
		"IMM.HostIPAddress1": "",
		"IMM.IMMInfo_Name":   "",
	}

	if _, err := os.Stat(cmd); err == nil {

		for k, _ := range asu64 {
			show := cmd + " show " + k + " | tail -1"
			showCmd := exec.Command("/bin/bash", "-c", show)
			out, _ := showCmd.Output()
			outArr := strings.Split(string(out), "=")
			asu64[outArr[0]] = outArr[1]
		}

		if oldIp != asu64["IMM.HostIPAddress1"] || oldServer != asu64["IMM.HostName1"] {
			fmt.Println("Old Hostname or IP Address doesn't match, exiting")
		} else {
			fmt.Println(cmd, "set", "IMM.HostIPAddress1", newIp)
			fmt.Println(cmd, "set", "IMM.IMMInfo_Name", newServer)
		}

	} else {
		fmt.Println("asu64 util does not exist!")
	}
}
