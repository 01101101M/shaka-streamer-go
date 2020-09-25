package internal

import (
	"fmt"
	"os/exec"
	"regexp"

	"github.com/blang/semver/v4"
)

func CheckVersion(name string, cmd []string, minVersion string) (err error) {
	v, _ := semver.Make(minVersion)
	path, _ := exec.LookPath(cmd[0])
	out, _ := exec.Command(path, cmd...).Output()
	rx := regexp.MustCompile(`[0-9]+(?:\.[0-9]+)+`)
	groups := rx.FindSubmatch(out)
	if len(groups) > 0 {
		version, _ := semver.Make(string(groups[0]))
		fmt.Println(name, version)
		if version.Compare(v) != 1 {
			return fmt.Errorf("%[1]s out of date ! Please install version %[2]s or higher of %[1]s.", name, minVersion)

		} else {
			return
		}
	}
	return fmt.Errorf("%[1]s not found! Please install version %[2]s or higher of %[1]s.", name, minVersion)
}
