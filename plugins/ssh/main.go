package ssh

import (
	"errors"
	"fmt"

	"github.com/samitpal/run-ssh/plugins"
)

var (
	PrivateKey string
	User       string
)

type RunSsh struct {
	runId          int
	action         string
	rtc            plugins.RuntimeConfig
	Command        string   `json:"command"`
	CommandTimeout int      `json:"command_timeout"` // in secs
	Hosts          []string `json:"hosts"`
}

func (r RunSsh) ExecuteCmd() {
	fmt.Println("Hello ExecuteCmd")
}

func (r RunSsh) RuntimeConfig() plugins.RuntimeConfig {
	return r.rtc
}

func NewRunSsh() RunSsh {
	return RunSsh{
		action:         "ssh",
		CommandTimeout: 300,
		rtc:            plugins.DefaultRuntimeConfig(),
	}
}

func RunValidate(r RunSsh) (*RunSsh, error) {
	fmt.Println(r)
	if r.Command == "" {
		err := errors.New("Command Not Set")
		return nil, err
	}
	if r.rtc.Operator == "" {
		err := errors.New("Operator Not Set")
		return nil, err
	}
	if r.rtc.Team == "" {
		err := errors.New("Team Not Set")
		return nil, err
	}
	if len(r.Hosts) < 1 {
		err := errors.New("Hosts Not Set")
		return nil, err
	}
	return &r, nil
}
