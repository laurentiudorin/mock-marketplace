package routes

import "os"

type ExternalRouter struct {
	routes map[string]string
}

func (externalRouter *ExternalRouter) NewExternalRouter() ExternalRouter {
	externalRouter.routes = map[string]string{
		"authenticationService": os.Getenv("AUTHENTICATOR_MICROSERVICE_URL") + "/user/validation",
	}
	return *externalRouter
}

func (externalRouter *ExternalRouter) GetRoute(name string) string {
	return externalRouter.routes[name]
}