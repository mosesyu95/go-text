package go_text

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func IsText(fp string) (bool, error) {
	var (
		outPut []byte
		err    error
		f      os.FileInfo
	)
	f, err = os.Stat(fp)
	if err != nil {
		return false, err
	}
	if f.IsDir() {
		return false, fmt.Errorf("this is a directory : %s ", fp)
	}
	outPut, err = exec.Command("file", "-b", "--mime-type", fp).Output()
	if err != nil {
		return false, err
	}
	if strings.HasPrefix(string(outPut), "text/") {
		return true, err
	}
	return false, nil
}
