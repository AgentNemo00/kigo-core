package order

import "github.com/AgentNemo00/kigo-core/notification"

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

