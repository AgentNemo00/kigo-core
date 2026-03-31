package order

const (
	// Basic orders from the main service to the modules
	OrderStartUp 	= "OrderStartUp" // called during startup procedore, after configuration is readed
	OrderShutdown 	= "OrderShutdown" // called to shutdown the module, i.e. idle mode 
	OrderReboot  	= "OrderReboot" // called to reboot the module, i.e. unexpected error or response
	OrderUpdate  	= "OrderUpdate" // called to update the module, i.e. update data or state
	OrderRender  	= "OrderRender" // called to render the module, i.e. draw the module on the output
)