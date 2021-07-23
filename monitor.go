package cronitor

import "time"

type MonitorRequest struct {
	Monitors []Monitor `json:"monitors"`
}

type Monitor struct {
	Key               string      `json:"key"`
	Type              string      `json:"type"`
	Name              string      `json:"name,omitempty"`
	Assertions        []string    `json:"assertions,omitempty"`
	FailureTolerance  int         `json:"failure_tolerance,omitempty"`
	GraceSeconds      int         `json:"grace_seconds,omitempty"`
	Group             string      `json:"group,omitempty"`
	Note              string      `json:"note,omitempty"`
	Notify            []string    `json:"notify,omitempty"`
	Platform          string      `json:"platform,omitempty"`
	RealertInterval   string      `json:"realert_interval,omitempty"`
	Request           interface{} `json:"request,omitempty"`
	Schedule          string      `json:"schedule,omitempty"`
	ScheduleTolerance int         `json:"schedule_tolerance,omitempty"`
	Timezone          string      `json:"timezone,omitempty"`
	Tags              []string    `json:"tags,omitempty"`
	Metadata          string      `json:"metadata,omitempty"`

	// Read-Only attributes
	Created     time.Time `json:"created"`
	Disabled    bool      `json:"disabled"`
	Initialized bool      `json:"initialized"`
	Passing     bool      `json:"passing"`
	Paused      bool      `json:"paused"`
	Running     bool      `json:"running"`
	// TODO: The API is not returning status attribute like it described in docs
	Status string `json:"status"`
}
