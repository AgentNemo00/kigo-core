# kigo-core

Pubsub shema. There is no constrains who can call them. Simple rules are, KiGo calls only `Oders`.

## Orders

Orders are messages sended by KiGo and modules to the modules. The modules should react to them.

- `OrderStartUp` is called after receiving the notification `NotificationReady` from the module and gives the module basic information about the current state of KiGo
- `OrderError` is called when an error occurs on the KiGo side which is caused by the module
- `OrderInformation` is called when an update is send which contains information about the general system
- `OrderChange` is called when an update is send to change the module
- `OrderShutdown` should free up the resources and shutdown

## Notifications

Notification are messages sended by the modules to KiGo. KiGo should react to them.

- `NotificationReady` is called when the module is ready to communicate
- `NotificationUpdate` is called when an update of informations is wanted   

### NotificationUpdate

Currently there are two update reasons:
- `ModulesInformation` is 0

## Informations

Informations are message which can be sended to modules to give them informations. They to not expect a reaction and are also attached to `OrderInformation`.

## Changes

Changes are messages which can be sended to modules to change the internal status. They are attached to `OrderChanges`. They cause a change in the state of the module which causes a redraw.

## Module lifecycle

TBD
