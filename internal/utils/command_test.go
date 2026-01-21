package utils

import (
	"strings"
	"testing"
)

func Test_ExecCommand(t *testing.T) {
	output, err := ExecCommand("echo", "hello world")
	if err != nil {
		t.Fatalf("ExecCommand failed: %v", err)
	}
	expected := "hello world\n"
	if output != expected {
		t.Fatalf("Expected output %q, got %q", expected, output)
	}
}

func Test_ExecCommand_WithPipe(t *testing.T) {
	// Test pipe command: echo hello world | grep hello
	output, err := ExecPipeCommand("echo hello world | grep hello")
	if err != nil {
		t.Fatalf("ExecPipeCommand failed: %v", err)
	}
	if !strings.Contains(output, "hello") {
		t.Fatalf("Expected output to contain 'hello', got %q", output)
	}
}

func Test_ExecCommand_WithMultiplePipes(t *testing.T) {
	// Test multiple pipes: seq 1 5 | grep -v 3 | wc -l
	// This should output 4 lines (1, 2, 4, 5)
	output, err := ExecPipeCommand("seq 1 5 | grep -v 3 | wc -l")
	if err != nil {
		t.Fatalf("ExecPipeCommand with multiple pipes failed: %v", err)
	}
	output = strings.TrimSpace(output)
	if output != "4" {
		t.Fatalf("Expected output '4', got %q", output)
	}
}

func Test_ExecCommand_ComplexPipe(t *testing.T) {
	// Test ls | head | wc -l
	output, err := ExecPipeCommand("ls -1 | head -n 3 | wc -l")
	if err != nil {
		t.Fatalf("ExecPipeCommand with complex pipe failed: %v", err)
	}
	output = strings.TrimSpace(output)
	// Should output 3 or less (depending on how many files exist)
	if output == "" {
		t.Fatalf("Expected non-empty output, got empty string")
	}
}
