package information

const (
	ModulesInformation = 1
)

type ModulesInformationPayload struct {
	Modules []ModuleInformation // Information about the active modules
}

type ModuleInformation struct {
	ID   string  // ID of the module
	Name string // Name of the module
	Ready bool   // Is the module ready to receive orders
	Changes []string // List of changes the module can receive
}
