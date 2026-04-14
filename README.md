# kigo-core

Pubsub shema. There is no constrains who can call them. Simple rules are, KiGo calls only `Oders`.

## Orders

Orders are messages sended by KiGo and modules to the modules. The modules should react to them.

- `OrderStartUp` is called after receiving the notification `NotificationReady` from the module 
- `OrderShutdown` is called to shutdown the module. i.e. when the amounts of reboots exceed the maximal amount of reboots allowed. Module is removed completly and needs to do the startup procedure again.
- `OrderReboot` is called when an error occurs on the KiGo side which is caused by the module.
- `OrderUpdate` is called to trigger an update on the module which should cause a update of the interal module state.
- `OrderRender` is called to trigger the module to render itself.

## Notifications

Notification are messages sended by the modules to KiGo. KiGo should react to them.

- `NotificationReady` is called when the module is ready to communicate
- `NotificationUpdate` is called when the module wants to trigger an update, the payload indicates what should be updated.
- `NotificationRender` is called when the module wants to trigger a render

### NotificationUpadate

Currently there are two update reasons:
- `Render` is 0
- `ModulesInformation` is 1

## Informations

Informations are message which can be sended to modules to give them informations. They to not expect a reaction and are also attached to `OrderUpdate`.

## Changes

Changes are messages which can be sended to modules to change the internal status. They are attached to `OrderUpdate`. They cause a change in the state of the module which causes a redraw.
