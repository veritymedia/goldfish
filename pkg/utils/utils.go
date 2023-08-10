package utils

import (
	"os"
	"runtime"

	"github.com/veritymedia/goldfish/pkg/config"
)

func CheckIsFirstLaunch() bool {
	return true
}

/*
  NOTE: How to resolve paths of where to store sqlite
   Linux: Resolves to $XDG_DATA_HOME or $HOME/.local/share.
   macOS: Resolves to $HOME/Library/Application Support.
   Windows: Resolves to {FOLDERID_RoamingAppData}.
*/

func DataDir() string {
	var (
		pathSeparator string = "/"
		homePathVar   string
		pathSuffix    string
	)

	switch os := runtime.GOOS; os {
	case "windows":
		homePathVar = "{FOLDERID_RoamingAppData}"
		pathSuffix = ""
		pathSeparator = "\\"
	case "darwin":
		homePathVar = "HOME"
		pathSuffix = "/Library/Application Support"
	case "linux":
		homePathVar = "HOME"
		pathSuffix = "/.local/share"
	default:
		homePathVar = ""
		pathSuffix = ""
	}

	if homePathVar == "" {
		return ""
	}
	homePath, ok := os.LookupEnv(homePathVar)
	if !ok {
		return ""
	} else {
		// TODO: May have to define compile time os based config. For separator
		return homePath + pathSuffix + pathSeparator + config.AppName
	}

}
