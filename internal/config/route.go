package config

type GatewayRoutesConfiguration struct {
	Routes []*RouteConfiguration
}

type PredicateConfiguration struct {
	Name string
	Args []string
}

type RouteConfiguration struct {
	Id                     string
	Uri                    string
	Order                  string
	PredicateConfiguration []*PredicateConfiguration
}
