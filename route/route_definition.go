package route

import "gateway/handler/predicate"

type GatewayRoutesDefinition struct {
	Routes []*Definition `config:"routes"`
}

type Definition struct {
	Id         string                 `config:"id"`
	Uri        string                 `config:"uri"`
	Order      string                 `config:"order"`
	Predicates []predicate.Definition `config:"predicates"`
}
