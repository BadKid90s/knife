package definition

type GatewayRoutesDefinition struct {
	Routes []*RouteDefinition `yaml:"routes"`
}

type RouteDefinition struct {
	Id         string                 `yaml:"id"`
	Uri        string                 `yaml:"uri"`
	Order      string                 `yaml:"order"`
	Predicates []*PredicateDefinition `yaml:"predicates"`
}
