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

type PositionPayload struct {
	X int
	Y int
}

type AreaPayload struct {
	X int
	Y int
	Width int
	Height int
}