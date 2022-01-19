package packet

import (
	"bytes"
	"testing"
)

func Test_readConnectPayloadFull(t *testing.T) {
	// Test Payload, should output following:
	// ClientID: cli
	// UserName: testo
	// Password: pesto
	payload := []byte{
		0, 3, 99, 108, 105, 0, 5, 116, 101, 115, 116, 111, 0, 5, 112, 101, 115, 116, 111,
	}
	r := bytes.NewReader(payload)
	res, err := readConnectPayload(r, len(payload))
	if err != nil {
		t.Fatalf("Error reading payload: %v", err)
	}
	t.Log("Result ConnectPayload", res)
	if res.ClientID != "cli" {
		t.Fatalf("Payload read incorrectly, expected 'cli', got '%s'", res.ClientID)
	}
	if res.Username != "testo" {
		t.Fatalf("Payload read incorrectly, expected 'testo', got '%s'", res.Username)
	}
	if res.Password != "pesto" {
		t.Fatalf("Payload read incorrectly, expected 'pesto', got '%s'", res.Password)
	}
}

func Test_readConnectPayloadClientIDOnly(t *testing.T) {
	// Test Payload, should output following:
	// ClientID: cli
	payload := []byte{
		0, 3, 99, 108, 105,
	}
	r := bytes.NewReader(payload)
	res, err := readConnectPayload(r, len(payload))
	if err != nil {
		t.Fatalf("Error reading payload: %v", err)
	}
	t.Log("Result ConnectPayload", res)
	if res.ClientID != "cli" {
		t.Fatalf("Payload read incorrectly, expected 'cli', got '%s'", res.ClientID)
	}
}

func Test_readConnectPayloadUnrealNoPassword(t *testing.T) {
	// Test Payload, should output following:
	// ClientID: cli
	payload := []byte{
		0, 3, 99, 108, 105, 0, 5, 116, 101, 115, 116, 111,
	}
	r := bytes.NewReader(payload)
	res, err := readConnectPayload(r, len(payload))
	if err != nil {
		t.Fatalf("Error reading payload: %v", err)
	}
	t.Log("Result ConnectPayload", res)
	if res.ClientID != "cli" {
		t.Fatalf("Payload read incorrectly, expected 'cli', got '%s'", res.ClientID)
	}
	if res.Username != "testo" {
		t.Fatalf("Payload read incorrectly, expected 'testo', got '%s'", res.Username)
	}
}