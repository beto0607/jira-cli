package utils

import "os/exec"

func OpenBrowser(url string) error {
	cmd := exec.Command("xdg-open", url)
	return cmd.Run()
}
