package notification

const (
    // Basic notifications from the modules to the main service
	NotificationReady 			= "NotificationReady" // called when the module is ready to receive orders, i.e. after startup procedore, Publish to main
	NotificationUpdate 			= "NotificationUpdate" // called when the module has an update to share, i.e. update data or state, Publish to main
	NotificationInformation 	= "NotificationInformation" // called when the module has information to share, i.e. about itself or the environment, Publish to main
	NotificationRender 			= "NotificationRender" // called when the module has a render to share, i.e. a string to render, Publish to main
)
