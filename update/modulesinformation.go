package update

type UpdateOrderModulesInformationPayload struct {
	Modules []ModuleInformation // Information about the active modules
}

type ModuleInformation struct {
	Name string // Name of the module
	Ready bool   // Is the module ready to receive orders
	Orders []string // List of orders the module can receive which are not basic orders
}
