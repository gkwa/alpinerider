package cmd

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestMonkeyCmd(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	monkeyCmd.Run(monkeyCmd, []string{})

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	if err != nil {
		t.Fatalf("Failed to copy output: %v", err)
	}
	output := buf.String()

	expected := []string{
		`"$schema"`,
		`"extends"`,
		`"packageRules"`,
		`"platformAutomerge"`,
		`"postUpdateOptions"`,
	}

	for _, exp := range expected {
		if !strings.Contains(output, exp) {
			t.Errorf("Expected output to contain %q", exp)
		}
	}
}
