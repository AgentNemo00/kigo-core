package information

import (
	"time"
)

type ModulesPayload struct {
	Modules []ModuleInformation // Information about the active modules
}

type ModuleInformation struct {
	ID   string  // ID of the module
	Name string // Name of the module
	Ready bool   // Is the module ready to receive orders
	Changes []string // List of changes the module can receive
	Heartbeat time.Duration // Duration between each heartbeat, for debugging and logging purposes	
	LastHeartbeat time.Time // Last time the module sent a heartbeat, for debugging and logging purposes
}
