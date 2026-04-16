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
	NumberOfModules 		int // number of modules called startup ; incrementing number
	MessageTo 				MessageTo // where to send the notifications and renders to
}

type MessageTo struct {
	Notification	 		string // ID to publish notification to ; should be itself
	Render					string // ID to publish render to ; should be one of the render wating routines
}

type OrderErrorPayload struct {
	Message string // error message to report to the module, e.g. when a module is not responding or has an unexpected response
}

type OrderInformationPayload struct {
	Type 				int 
	Payload 			any
}

type OrderChangePayload struct {
	Type 				string 
	Payload 			any
}

type OrderShutdownPayload struct {
	Reason string // reason for shutdown, e.g. when the main service is shutting down, Publish to main
}

