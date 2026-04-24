package notification

import (
	"time"
)
// Messages send from the module

type Notification struct {
	From 			string 	// ID from who the message is. Module ID
	To 				string	// ID to who the message is. KiGo ID
	Notification	string  // what notification to publish, e.g. ready, update, render
	Payload 		any
}

type NotificationReadyPayload struct {
	Duration 		time.Duration // Duration needed to be ready, when should order startup be called
	Name 			string // Human readable name of the module, for debugging and logging purposes  
	Changes         []string // List of changes that the module has, e.g. update, render, etc. For debugging and logging purposes
	Heartbeat 		time.Duration // Duration between each heartbeat, for debugging and logging purposes
}

type NotificationUpdatePayload struct {
	Type 	int 
	Payload any
}
