package plugins

type RuntimeConfig struct {
	MaxInFLight       int    `json:"max_in_flight"`
	Canary            bool   `json:"canary"`
	CanaryNumHosts    int    `json:"canary_num_hosts"`
	CanaryTime        int    `json:"canary_time"`
	StopIfCanaryFails bool   `json:"stop_if_canary_fails"`
	SendStatusToES    bool   `json:"send_status_to_es"`
	ESHostPort        string `json:"es_host_port"`
	ESIndex           string `json:"es_index"`
	RunName           string `json:"run_name"` // goes into ESIndex
	Team              string `json:"team"`     // goes into ESIndex
	Operator          string `json:"operator"` // goes into ESIndex
}

func DefaultRuntimeConfig() RuntimeConfig {
	return RuntimeConfig{
		MaxInFLight:       5,
		Canary:            false,
		CanaryNumHosts:    1,
		CanaryTime:        15,
		StopIfCanaryFails: true,
		SendStatusToES:    false,
		ESHostPort:        "127.0.0.1:9200",
		ESIndex:           "es_index",
	}
}

type Runner interface {
  ExecuteCmd()
  RuntimeConfig() RuntimeConfig
}
