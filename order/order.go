package order

// Messages send to the modules.

type Order struct {
	From 	string
	To 		string
	Order 	string
	Payload any
}

// Payloads for orders

type OrderStartUpPayload struct {
	QueuePosition int // Position of the module
}

type OrderRenderPayload struct {
	SizeX int // size of the output to render on
	SizeY int // size of the output to render on
}

type OrderUpdatePayload struct {
	Order 		string  	// detailed order
	OrderValue 	string 	// order value or information, e.g. what to update, in JSON format
}
