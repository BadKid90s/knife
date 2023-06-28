package predicate

type Definition struct {
	Name string            `config:"name"`
	Args map[string]string `config:"args"`
}
