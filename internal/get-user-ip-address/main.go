package getuseripaddress

import (
	"os/exec"
	"strings"
)

func GetUserIpAddress() (string, error) {
	cmd := exec.Command("nslookup", "myip.opendns.com", "resolver1.opendns.com")
	cmd.Stdout = nil
	cmd.Stderr = nil
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	var stringOutPut = string(output)

	splitedString := strings.Split(stringOutPut, "Address: ")

	return splitedString[1], nil
}
