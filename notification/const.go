package notification

const (
    // Basic notifications from the modules to the main service
	NotificationReady 			= "NotificationReady" // called when the module is ready to receive orders, i.e. after startup procedore, Publish to main
	NotificationUpdate 			= "NotificationUpdate" // called when the module has an update to share, i.e. update data or state, Publish to main
	NotificationShutdown        = "NotificationShutdown" // called when the module shuts down themself
)
