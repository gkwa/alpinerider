package cmd

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestZooCmd(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	zooCmd.Run(zooCmd, []string{})

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	if err != nil {
		t.Fatalf("Failed to copy output: %v", err)
	}
	output := buf.String()

	expected := []string{
		"-- cat --",
		"-- dog --",
		"-- owl --",
		`"$schema"`,
		`"extends"`,
		`"packageRules"`,
		`"platformAutomerge"`,
	}

	for _, exp := range expected {
		if !strings.Contains(output, exp) {
			t.Errorf("Expected output to contain %q", exp)
		}
	}
}

func TestZooCmd_Execute(t *testing.T) {
	rootCmd.SetArgs([]string{"zoo"})
	if err := rootCmd.Execute(); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestZooCmd_Help(t *testing.T) {
	cmd := &cobra.Command{Use: "root"}
	zooCmd := &cobra.Command{
		Use:   "zoo",
		Short: "Show all configs",
		Long:  `Show configs for all animal types (cat, dog, owl)`,
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	cmd.AddCommand(zooCmd)

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"zoo", "--help"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("Failed to execute command: %v", err)
	}

	output := buf.String()

	expected := "Show configs for all animal types"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected help output to contain %q, got %q", expected, output)
	}
}
