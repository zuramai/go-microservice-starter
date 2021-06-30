package container

// This container will be passed on every usecase and wire the application together
type Container interface {
	// BuildUseCase creates concrete type for every use case
	// For example, BuildUseCase will be used in containerhelper for create RegistrationUseCase, ListUserUseCase, etc
	// BuildUseCase passes code of use case that set up in config
	BuildUseCase(code string) (interface{}, error)
	Put(code string, value interface{})
	Get(code string) (interface{}, bool)
}
