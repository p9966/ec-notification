package configs

type RestFulService struct {
	Resources []Resources
}

type Resources struct {
	Resource Resource
}

type Resource struct {
	Name   string
	Path   string
	Method string
	Host   string
}
