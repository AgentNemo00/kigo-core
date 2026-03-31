package notification

const (
    // Basic notifications from the modules to the main service
	NotificationReady 	= "NotificationReady" // called when the module is ready to receive orders, i.e. after startup procedore, Publish to main
	NotificationUpdate 	= "NotificationUpdate" // called when the module has an update, i.e. update data or state, Publisch to main
	NotificationRender 	= "NotificationRender" // called when the module needs to render something, i.e. after update or order, Publish to main
)