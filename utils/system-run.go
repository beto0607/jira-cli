package utils

import "os/exec"

func OpenBrowser(url string, browser string) error {
	if len(browser) > 0 {
		cmd := exec.Command(browser, url)
		return cmd.Run()
	}
	cmd := exec.Command("xdg-open", url)
	return cmd.Run()
}
