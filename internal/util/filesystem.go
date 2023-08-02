package util

import (
	"fmt"
	"os"
	"runtime"
)

func GetTempDir() (string, error) {
	switch runtime.GOOS {
	case "darwin": // macOS
		return os.TempDir(), nil
	case "windows":
		return os.TempDir(), nil
	case "linux":
		return "/tmp", nil
	default:
		return "", fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}
