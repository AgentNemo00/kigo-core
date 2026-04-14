package notification

// Messages send from the module

import (
	"time"
)	

type Notification struct {
	From 			string 	// ID from who the message is. Module ID
	FromName 		string	// Human readable name of the module, for debugging and logging purposes
	To 				string	// ID to who the message is. KiGo ID
	Notification	string  // what notification to publish, e.g. ready, update, render
	Payload 		any
}

type NotificationReadyPayload struct {
	Duration 		time.Duration // Duration needed to be ready, when should order startup be called
	CallingInterval time.Duration // Interval in which the module should be updated without beeing called directly ; polling mode
	Changes 		[]string // Allowed changes to the module, e.g. change format
	
}

type NotificationUpdatePayload struct {
	Payload 	int // what should be updated
}

type NotificationRenderPayload struct {
	// where to render
	PositionX 	int
	PositionY 	int

	Payload 	string	// mmap name
}
