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

type PointPayload struct {
	X int
	Y int
}

type AreaPayload struct {
	X int
	Y int
	Width int
	Height int
}

type OverlapingResponse struct {
	X int
	Y int
	Width int
	Height int
}

type RemovePayload struct {
	UUIID string
}