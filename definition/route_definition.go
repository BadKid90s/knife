package definition

type GatewayRoutesDefinition struct {
	Routes []*RouteDefinition `config:"routes"`
}

type RouteDefinition struct {
	Id         string                 `config:"id"`
	Uri        string                 `config:"uri"`
	Order      string                 `config:"order"`
	Predicates []*PredicateDefinition `config:"predicates"`
}
