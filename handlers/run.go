package handlers

import (
	"encoding/json"
	"math/big"
	"os/exec"
	"strings"
)

type RunArgs struct {
	Script    string                       `json:"script"`
	Balances  map[string]map[string]string `json:"balances,omitempty"`
	Variables map[string]string            `json:"variables,omitempty"`
	Metadata  map[string]string            `json:"metadata,omitempty"`
}

type Posting struct {
	Source      string  `json:"source"`
	Destination string  `json:"destination,omitempty"`
	Amount      big.Int `json:"amount,omitempty"`
	Asset       string  `json:"asset,omitempty"`
}

type RunCmdOutput struct {
	Postings []Posting         `json:"postings"`
	TxMeta   map[string]string `json:"txMeta"`
}

type RunResultOk struct {
	Ok    bool         `json:"ok"`
	Value RunCmdOutput `json:"value"`
}

type RunResultErr struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

func Run(args RunArgs) (any, error) {
	rawArgBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command("numscript", "run", "--output-format", "json", "--raw", string(rawArgBytes))

	var stdout strings.Builder
	var stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return RunResultErr{
			Ok:    false,
			Error: stderr.String(),
		}, nil
	}

	var runCmdOutput RunCmdOutput
	err = json.Unmarshal([]byte(stdout.String()), &runCmdOutput)
	if err != nil {
		return nil, err
	}

	return RunResultOk{Ok: true, Value: runCmdOutput}, nil
}
