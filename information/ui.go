package information

type UIPayload struct {
	Channels 	[]string
	Formats 	[]string
}

type ScreenPayload struct {
	Width	int
	Height	int
	MaxFPS  int
}
