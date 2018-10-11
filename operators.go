package connor

var opMap = map[string]Operator{}

// Operator instances are used by Connor to provide advanced query
// functionality.
type Operator interface {
	Name() string
	Evaluate(condition, data interface{}) (bool, error)
}

// Register allows you to add your own operators to Connor or override
// the built in operators if you wish.
func Register(op Operator) {
	opMap[op.Name()] = op
}

// Operators gets the list of all registered operators which can be
// used with Connor
func Operators() []string {
	names := []string{}

	for name := range opMap {
		names = append(names, name)
	}

	return names
}
