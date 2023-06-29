package definition

type PredicateDefinition struct {
	Name string            `config:"name"`
	Args map[string]string `config:"args"`
}
