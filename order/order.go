package order

// Messages send to the modules.

type Order struct {
	From 		string // ID from who the message is. KiGo ID
	To 			string // ID to who the message is. Module ID
	Order 		string // what order to do, e.g. startup, update, render
	Payload 	any
}

// Payloads for orders

type OrderStartUpPayload struct {
	NumberOfModules 			int // number of modules called startup ; incrementing number
	NotificationTo	 			string // ID to publish notification to ; should be itself
	RenderTo 					string // ID to publish render to ; should be one of the render wating routines
}

type OrderRenderPayload struct {
	SizeX int // size of the output to render on
	SizeY int // size of the output to render on
}

type OrderUpdatePayload struct {
	Tag 		string  	// tag about what to update
	Value 		any 	// order value or information, e.g. what to update, in JSON format
}
