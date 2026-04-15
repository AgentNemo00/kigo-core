package order

const (
	// Basic orders from the main service to the modules
	OrderStartUp 	= "OrderStartUp" // called during startup procedore, after configuration is readed
	OrderError		= "OrderError" // called to report an error to the module, e.g. when a module is not responding or has an unexpected response
	OrderInformation = "OrderInformation" // called to request information from the module, e.g. when a module is not responding or has an unexpected response
	OrderChange		= "OrderChange" // called to update the module, e.g. when a module has an update to share, Publish to main
	OrderShutdown		= "OrderShutdown" // called to shutdown the module, e.g. when the main service is shutting down, Publish to main
)
