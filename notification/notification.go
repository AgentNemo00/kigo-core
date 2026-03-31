package notification

// Messages send from the module

import (
	"time"
)	

type Notification struct {
	From 			string
	To 				string
	Notification	string
	Payload 		any
}

type NotificationReadyPayload struct {
	Duration time.Duration // Duration needed to be ready, when should order startup be called
	CallingInterval time.Duration // Interval in which the module should be updated without beeing called directly  TODO: needed ?
}

type NotificationUpdatePayload struct {
	Duration time.Duration // Duration needed to update, when should notification render be called
}

type NotificationRenderPayload struct {
	// where to render
	PositionX int
	PositionY int

	Object []byte // object to render
	Path string // path to the object to render, if object is empty, path should be used
}
