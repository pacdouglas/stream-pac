package config

import (
	"context"
	"time"
)

// Config is the full application configuration, loaded from config.toml.
type Config struct {
	Server    ServerConfig              `toml:"server"`
	Platforms map[string]PlatformConfig `toml:"platforms"`
	UI        UIConfig                  `toml:"ui"`
}

type ServerConfig struct {
	Port int    `toml:"port"` // default: 7777
	Host string `toml:"host"` // default: "127.0.0.1"
}

// PlatformConfig holds per-platform settings.
// Token is the static fallback; OAuth sessions are stored in the DB.
type PlatformConfig struct {
	Enabled      bool   `toml:"enabled"`
	Token        string `toml:"token"`         // static token (manual mode)
	Channel      string `toml:"channel"`       // Twitch: channel name
	ChatroomID   string `toml:"chatroom_id"`   // Kick: obtain via browser
	VideoID      string `toml:"video_id"`      // YouTube: live video ID
	PollInterval string `toml:"poll_interval"` // YouTube: e.g. "5s"
}

type UIConfig struct {
	Theme    string `toml:"theme"`    // default: "pacman"
	Language string `toml:"language"` // default: "pt-BR"
}

// Load reads and parses the TOML config file at path.
// Returns Default() if the file does not exist.
// Requires github.com/BurntSushi/toml.
func Load(path string) (*Config, error) {
	panic("not implemented")
}

// Save writes cfg to path in TOML format.
func Save(cfg *Config, path string) error {
	panic("not implemented")
}

// Default returns a Config populated with sensible defaults.
func Default() *Config {
	return &Config{
		Server: ServerConfig{Port: 7777, Host: "127.0.0.1"},
		UI:     UIConfig{Theme: "pacman", Language: "pt-BR"},
	}
}

// Watcher polls the config file and calls onChange when its content changes.
type Watcher struct {
	path     string
	interval time.Duration
}

// NewWatcher creates a Watcher that polls path every interval.
func NewWatcher(path string, interval time.Duration) *Watcher {
	return &Watcher{path: path, interval: interval}
}

// Watch starts watching. Calls onChange with the new config on each detected change.
// Blocks until ctx is cancelled.
func (w *Watcher) Watch(ctx context.Context, onChange func(*Config)) error {
	panic("not implemented")
}
