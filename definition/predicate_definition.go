package definition

type PredicateDefinition struct {
	Name string            `yaml:"name"`
	Args map[string]string `yaml:"args"`
}
