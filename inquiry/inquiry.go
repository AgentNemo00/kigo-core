package inquiry

import (
	"time"
)

type InquiryInformationPayload struct {
	Type 	int 
	Payload any
}

type InquiryRenderPayload struct {
	PositionX 	int 	// Position X
	PositionY 	int 	// Position Y

	Format 		string 	// Format choosen, e.g. raw, H264, H265, Jpeg
	Channel     string 	// IPC, PubSub, REST

	FPS 		int 	// FPS choosen for the video stream, e.g. 30, 60, must be bigger than 0
	Time 		time.Duration // Duration; FPS * Time = Animation length; optional zero value means until ringbuffer closes	

	MaxFrameSize 	int // max frame size, i.e. custom size, FullHD, 4K
}	
