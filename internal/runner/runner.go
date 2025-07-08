package runner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

func CLIRunTool(toolPath string, input map[string]interface{}) (map[string]interface{}, error) {
	inputBytes, _ := json.Marshal(input)

	cmd := exec.Command(toolPath)
	cmd.Stdin = bytes.NewReader(inputBytes)

	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	//setting timeout for 2 seconds for testing
	//can be added as input later with some validation
	//process is used to isolate tool to be run

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("start failed: %w", err)
	}

	done := make(chan error)
	go func() { done <- cmd.Wait() }()

	select {
	case <-time.After(2 * time.Second):
		_ = cmd.Process.Kill()
		return nil, fmt.Errorf("timeout")
	case err := <-done:
		if err != nil {
			return nil, fmt.Errorf("tool error: %v: %s", err, errBuf.String())
		}
	}

	var output map[string]interface{}
	if err := json.Unmarshal(outBuf.Bytes(), &output); err != nil {
		return nil, fmt.Errorf("parse error: %w", err)
	}

	return output, nil
}
