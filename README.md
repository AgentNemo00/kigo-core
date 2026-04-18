# kigo-core

Pubsub shema for communication with `KiGo` and `KiGoUI`.

## Communication

### Orders

Orders are messages sended by KiGo and modules to the modules. The modules should react to them.

- `OrderStartUp` is called after receiving the notification `NotificationReady` from the module and gives the module basic information about the current state of KiGo
- `OrderError` is called when an error occurs on the KiGo side which is caused by the module
- `OrderInformation` is called when an update is send which contains information about the general system
- `OrderChange` is called when an update is send to change the module
- `OrderShutdown` should free up the resources and shutdown

### Notifications

Notification are messages sended by the modules to `KiGo`.

- `NotificationReady` is called when the module is ready to communicate
- `NotificationInformation` is called when an information is required
- `NotificationUpdate` is called when an update is happened

- `NotificationRender` is called to render. Send to `KiGoUI`

#### NotificationUpdate

Currently reasons for an update:
- `Config` is 0 - change module configuration

### Informations

Informations are message which can be sended to modules to give them informations. They to not expect a reaction and are also attached to `OrderInformation`.
Currently reasons for an update:
- `Modules` is 0 - receive all modules information
- `Module` is 1 - receive information about the module with the given name or ID

### Changes

Changes are messages which can be sended to modules to change the internal status. They are attached to `OrderChanges`. They cause a change in the state of the module which causes a redraw.

## Module lifecycle

Modules are in a initiating period before they send `NotificationReady`. They can choose a heartbeat which is smaller than 24 hours or leave it empty. If empty skip everything, we assume a module which do not draw anything to `KiGoUI`
`KiGo` response with `OrderStartUp`. From this point on a constant heartbeat is been send to make sure the module is still alive. `KiGo`and `KiGoUI` share the module informations. The heartbeat is responded by initiating communication with `KiGo` or `KiGoUI`. If the heartbeat is not responded a `OrderShutdown` is been called and, if the module has drawn, the widget of the module is been removed.
`KiGo` accepts `NotificationInformation` and `NotificationUpdate`. `KiGoUI` accepts `NotificationRender`. The loop of heartbeat required the module to do anything, may it be to render something or any logic, instead of just aquiring resources.

## KiGoUI - TBD

Modules on the device can communication with `KiGoUI` is via pubsub. This keeps the latence low. Remote devices need to communicate via REST, which is implemented on Phase 5.
`KiGoUI` is sharing with `KiGo` a list of `ModuleInformation`. `NotificationRender` is send to `KiGoUI`. For the communication with `KiGoUI` we need a number of communications to make sure the module has the exact parameters to create the image which should be render to the screen. `KiGoUI` has the informations about the screen size, the already drawn modules positions, the modules informations and the max supported refresh rate (fps). The modules needs infomationen about the screen size and the max refresh rate. The module provides than the position it should be drawn to, the images at the fps choosen. Drawn is also an update (like calling `NotificationInformation` or `NotificationUpdate`).

Handshake protocol:

Module -> UI; I want to render on position XY with this size with this FPS.
UI -> create ringbuffer mmap with space for 2 frames.
UI -> Module; shares reference about ringbuffer mmap. Refreshes `ModuleInformation`. TBD
Moudle -> Send frame via ringbuffer until finish condition is met.
UI -> Draws frame on screen
