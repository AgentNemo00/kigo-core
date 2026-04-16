package update

import "time"

type Module struct {
	Ready bool   // Is the module ready to receive orders
	Changes []string // List of changes the module can receive
	Heartbeat time.Duration // Duration between each heartbeat, for debugging and logging purposes
}