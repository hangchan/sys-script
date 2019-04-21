package rename_icinga

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var (
	icingaServer string = os.Getenv("icingaServer")
	hostFile     string = os.Getenv("icingaHostFile")
)

func RenameIcinga(oldServer string, newServer string) string {
	newIp := "10.10.10.10"
	newType := "rancher_ui"
	newEnv := "staging"
	var retOut []string
	out, err := exec.Command("ssh", icingaServer, "grep -A4", oldServer, hostFile).Output()
	if err != nil {
		log.Fatal(err)
	}

	outStr := string(out)
	outArr := strings.Split(outStr, "\n")

	for _, v := range outArr {
		reHost := regexp.MustCompile(`^object\sHost\s"(.*?)"\s{$`)
		reDN := regexp.MustCompile(`^\s\sdisplay_name\s=\s"(.*?)"$`)
		reAddr := regexp.MustCompile(`^\s\saddress\s=\s"(.*?)"$`)
		reIm := regexp.MustCompile(`^\s\simport\s"(.*?)"$`)
		reEnv := regexp.MustCompile(`^\s\svars.environment\s=\s"(.*?)"$`)
		repStr := "$1"
		switch {
		case reHost.MatchString(v):
			out := reHost.ReplaceAllString(v, repStr)
			retOut = append(retOut, fmt.Sprintf("%s", out))
		case reDN.MatchString(v):
			out := reDN.ReplaceAllString(v, repStr)
			retOut = append(retOut, fmt.Sprintf("%s", out))
		case reAddr.MatchString(v):
			out := reAddr.ReplaceAllString(v, repStr)
			retOut = append(retOut, fmt.Sprintf("%s", out))
		case reIm.MatchString(v):
			out := reIm.ReplaceAllString(v, repStr)
			retOut = append(retOut, fmt.Sprintf("%s", out))
		case reEnv.MatchString(v):
			out := reEnv.ReplaceAllString(v, repStr)
			retOut = append(retOut, fmt.Sprintf("%s", out))
		}
	}

	hostTemplate := fmt.Sprintf(`object Host "%s" {
  display_name = "%s"
  address = "%s"
  import "%s"
  vars.environment = "%s"
}`, newServer, newServer, newIp, newType, newEnv)
	return hostTemplate

}
