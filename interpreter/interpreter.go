package interpreter

import "strings"

// InterpretString interpretes string, roses are red, sky is blue
func InterpretString(input string) (interface{}, error) {
	reader := strings.NewReader(input)
	lexer, err := NewLexer(reader)

	if err != nil {
		return nil, err
	}

	parser, err := NewParser(lexer)

	if err != nil {
		return nil, err
	}

	node, err := parser.Parse()

	if err != nil {
		return nil, err
	}

	result, err := node.Traverse()

	if err != nil {
		return nil, err
	}

	return result, nil
}
