package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseConfigBytesRemoteManagementBaseURL(t *testing.T) {
	cfg, errParse := ParseConfigBytes([]byte(`remote-management:
  allow-remote: true
  base-url: "https://proxy.example.com"
`))
	if errParse != nil {
		t.Fatalf("ParseConfigBytes() error = %v", errParse)
	}

	if cfg.RemoteManagement.BaseURL != "https://proxy.example.com" {
		t.Fatalf("RemoteManagement.BaseURL = %q, want https://proxy.example.com", cfg.RemoteManagement.BaseURL)
	}
}

func TestManagementPanelAutoUpdateDisabledByDefault(t *testing.T) {
	parsed, errParse := ParseConfigBytes([]byte("remote-management:\n  allow-remote: true\n"))
	if errParse != nil {
		t.Fatalf("ParseConfigBytes() error = %v", errParse)
	}
	if !parsed.RemoteManagement.DisableAutoUpdatePanel {
		t.Fatal("ParseConfigBytes() left management panel auto-update enabled by default")
	}

	configPath := filepath.Join(t.TempDir(), "config.yaml")
	if errWrite := os.WriteFile(configPath, []byte("remote-management:\n  allow-remote: true\n"), 0o600); errWrite != nil {
		t.Fatalf("write config: %v", errWrite)
	}
	loaded, errLoad := LoadConfigOptional(configPath, false)
	if errLoad != nil {
		t.Fatalf("LoadConfigOptional() error = %v", errLoad)
	}
	if !loaded.RemoteManagement.DisableAutoUpdatePanel {
		t.Fatal("LoadConfigOptional() left management panel auto-update enabled by default")
	}
}
