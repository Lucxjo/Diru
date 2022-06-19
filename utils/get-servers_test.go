package utils

import (
	"os"
	"testing"
)

func TestGetServerCount(t *testing.T) {
	serverCount := GetConnectedServerCount(os.Getenv("DISCORD_TOKEN"))
	if serverCount != 1 {
		t.Errorf("GetConnectedServerCount: expected 1, got %d", serverCount)
	} else {
		t.Logf("GetConnectedServerCount: expected 1, got %d", serverCount)
	}
}

func TestGetServers(t *testing.T) {
	servers := GetConnectedServers(os.Getenv("DISCORD_TOKEN"))
	if len(servers) != 1 {
		t.Errorf("GetServers: expected 1, got %d", len(servers))
	} else {
		t.Logf("GetServers: expected 1, got %d", len(servers))
	}
}