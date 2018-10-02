package lib

import (
	"os"
	"strings"

	"github.com/mgenware/go-packagex/stringsx"
)

const homeEnv = "$HOME"

func EvaluatePath(s string) string {
	if strings.HasPrefix(s, "~/") {
		s = homeEnv + "/" + stringsx.SubStringFromStart(s, 2)
	} else if s == "~" {
		s = homeEnv
	}
	return os.ExpandEnv(s)
}
