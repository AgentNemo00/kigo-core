package inquiry

import (
	"time"
)

type InquiryInformationPayload struct {
	Type 	int 
	Payload any
}

type InquiryRenderPayload struct {
	PositionX 	int // Position X
	PositionY 	int // Position Y

	Format 		string // Format choosen, e.g. raw, H264, H265, Jpeg
	Static 		bool // Is the render static or dynamic, e.g. a static image or a video stream, expecting 1 frame

	FPS 		int // FPS choosen for the video stream, e.g. 30, 60, must be bigger than 0
	Time 		time.Duration // Duration; FPS + Time = Animation length; optional zero value means until ringbuffer closes
}
