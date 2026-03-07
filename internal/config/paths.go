package config

import (
	"os"
	"path/filepath"
	"runtime"
)

// Paths holds all resolved filesystem paths for the application.
type Paths struct {
	ConfigFile   string // e.g. ~/.config/stream-pac/config.toml
	DataDir      string // e.g. ~/.local/share/stream-pac/
	OverlaysDir  string // DataDir/overlays/
	DatabaseFile string // DataDir/stream-pac.db
	LogFile      string // DataDir/stream-pac.log
}

// Resolve returns OS-specific application paths and creates any missing directories.
func Resolve() (*Paths, error) {
	panic("not implemented")
}

func configDir() (string, error) {
	base, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, "stream-pac"), nil
}

// dataDir returns the OS-specific user data directory for stream-pac.
// Linux: ~/.local/share/stream-pac
// Windows: %AppData%\stream-pac
// macOS: ~/Library/Application Support/stream-pac
func dataDir() (string, error) {
	switch runtime.GOOS {
	case "windows":
		base, err := os.UserConfigDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(base, "stream-pac"), nil
	case "darwin":
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, "Library", "Application Support", "stream-pac"), nil
	default:
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, ".local", "share", "stream-pac"), nil
	}
}
