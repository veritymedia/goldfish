package utils

import (
	"fmt"
	"os"
	"path/filepath"
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

func DataDir() (string, error)	 {
	// Get the user's home directory from the environment variables
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// Determine the appropriate app files directory based on the operating system
	var appDir string
	switch runtime.GOOS {
	case "windows":
		appDir = filepath.Join(homeDir, "AppData", "Local", config.AppName)
	case "darwin":
		appDir = filepath.Join(homeDir, "Library", "Application Support", config.AppName)
	default: // For Linux and other Unix-like systems
		appDir = filepath.Join(homeDir, ".local", config.AppName)
	}

	fmt.Println("SUCCESS: App dir: ", appDir)
	return appDir, nil
}



func DataDir2() string {
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
		fmt.Println("Home path: ", homePath + pathSuffix + pathSeparator + config.AppName + pathSeparator)
		return homePath + pathSuffix + pathSeparator + config.AppName + pathSeparator
	}

}
