// +build windows

package tail

import (
	"github.com/simon4545/tail/winfile"
	"os"
)

func OpenFile(name string) (file *os.File, err error) {
	return winfile.OpenFile(name, os.O_RDONLY, 0)
}
