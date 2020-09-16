package process

// Processes represents the arrays or processes
type Processes struct {
	Processes []Process `json:"processes"`
	Count     int64     `json:"count"`
}

// Process represents a single process
type Process struct {
	Commands      []Command `json:"commands"`
	LastRun       string    `json:"last_run"`
	LastRunStatus string    `json:"last_run_status"`
	Asleep        bool      `json:"asleep"`
	Logs          []Logs    `json:"logs"`
	ID            string    `json:"_id"`
	ProcessID     string    `json:"process_id"`
	ProcessName   string    `json:"process_name"`
	Developer     string    `json:"developer"`
	CreatedAt     string    `json:"createdAt"`
	UpdatedAt     string    `json:"updatedAt"`
	V             int       `json:"__v"`
}

// Command represents a single command
type Command struct {
	ID        string `json:"_id"`
	ProcessID string `json:"process"`
	Command   string `json:"command"`
	CommandID string `json:"command_id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Logs represents a single logs
type Logs struct {
	ID      string `json:"_id"`
	Process string `json:"process"`
	Node    string `json:"node"`
}
