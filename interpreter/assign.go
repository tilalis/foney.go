package interpreter

// Assign represents assign operation
type Assign struct {
	Symbol *Symbol
	Value  AST
}

// Traverse traverse assign
func (a *Assign) Traverse() (interface{}, error) {
	name := a.Symbol.Token.Value.(string)
	value, err := a.Value.Traverse()

	if err != nil {
		return nil, err
	}

	globalSymbolTable.Set(name, value)

	return value, nil
}
