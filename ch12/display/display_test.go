package display

import (
	"os"
	"testing"
)

func Test(t *testing.T)  {
	Display("os.Stderr", os.Stderr)
}