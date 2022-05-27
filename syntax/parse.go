package syntax

import (
	"errors"
	"fmt"
	"github.com/crossoverJie/gscript/token"
	"strconv"
)

// Parse syntax parsing.
// Program -> IntDeclaration | ExpressionStatement | AssignmentStatement
// IntDeclaration -> 'int' Id ( = Additive) '\n'
// ExpressionStatement -> Additive '\n'
// Additive -> Multiplicative ( (+ | -) Multiplicative)*
// Multiplicative -> Primary ((* | /) Primary)*
// Primary -> IntLiteral | Id | Additive
// AssignmentStatement -> Identifier = Additive
func Parse(script string) (*ASTNode, error) {
	tokenTypes := token.Tokenize(script)
	reader := NewTokenReader(tokenTypes)

	root := NewASTNode(Program, App)
	var err error
	for reader.Peek() != nil {
		child, _ := IntDeclare(reader)

		if child == nil {
			child, err = ExpressionStatement(reader)
		}

		if child == nil {
			child, err = AssignmentStatement(reader)
		}

		if child != nil {
			root.AddChild(child)
		} else {
			return nil, errors.New("syntax err: Invalid statement")
		}
	}

	//cal.PrintASTNode(root, "")

	return root, err
}

// ExpressionStatement -> 1+2*3 '\n'
func ExpressionStatement(reader *TokenReader) (*ASTNode, error) {
	pos := reader.GetPos()
	node, err := AdditiveLoop(reader)
	if err != nil {
		return nil, err
	}
	if node != nil {
		tokenType := reader.Peek()
		if tokenType != nil && tokenType.TokenType() == token.Enter {
			reader.Read()
		} else {
			node = nil
			reader.SetPos(pos) // reset pos
		}
	}
	return node, nil
}

// AssignmentStatement -> Identifier = Additive  parse: a = 10*2,
func AssignmentStatement(reader *TokenReader) (*ASTNode, error) {
	var node *ASTNode
	tokenType := reader.Peek()
	if tokenType == nil {
		return nil, errors.New("syntax err: invalid AssignmentStatement")
	}

	if tokenType.TokenType() == token.Identifier {
		tokenType = reader.Read()
		node = NewASTNode(AssignmentStmt, tokenType.Value())

		tokenType = reader.Peek()
		if tokenType == nil {
			return nil, errors.New("syntax err: invalid AssignmentStatement")
		}

		if tokenType.TokenType() == token.Assignment {
			tokenType = reader.Read()

			child1, err := AdditiveLoop(reader)
			if err != nil {
				return nil, err
			}
			node.AddChild(child1)

			// parse end
			tokenType = reader.Peek()
			if tokenType == nil || tokenType.TokenType() != token.Enter {
				return nil, errors.New("syntax err: invalid statement, miss enter")
			}
			tokenType = reader.Read()

		} else {
			reader.UnRead()
		}
	}
	return node, nil
}

// EvaluateWithRuntime deep recursion loop AST
func EvaluateWithRuntime(node *ASTNode, indent string) int {
	result := 0
	switch node.GetNodeType() {
	case Program:
		for _, astNode := range node.GetChildren() {
			result = EvaluateWithRuntime(astNode, indent)
		}
	case AdditiveType:
		ch01 := node.GetChildren()[0]
		val1 := EvaluateWithRuntime(ch01, indent+"\t")
		ch02 := node.GetChildren()[1]
		val2 := EvaluateWithRuntime(ch02, indent+"\t")

		if node.GetText() == "+" {
			result = val1 + val2
		} else {
			result = val1 - val2
		}
	case MultiplicativeType:
		ch01 := node.GetChildren()[0]
		val1 := EvaluateWithRuntime(ch01, indent+"\t")
		ch02 := node.GetChildren()[1]
		val2 := EvaluateWithRuntime(ch02, indent+"\t")
		if node.GetText() == "*" {
			result = val1 * val2
		} else {
			result = val1 / val2
		}
	case IntLiteralType:
		result, _ = strconv.Atoi(node.GetText())
	case IntDeclaration:
		// int a= 10
		varName := node.GetText()
		var value int
		if len(node.GetChildren()) > 0 {
			// int a = 45+2
			value = EvaluateWithRuntime(node.GetChildren()[0], indent+"\t")
		}
		runtime.Vars()[varName] = value
	case AssignmentStmt:
		// age =20
		value, ok := runtime.Vars()[node.GetText()]
		if !ok {
			fmt.Printf("syntax err: var %s not exit\n", node.GetText())
		}
		varName := node.GetText()
		if len(node.GetChildren()) > 0 {
			// int a = 45+2
			value = EvaluateWithRuntime(node.GetChildren()[0], indent+"\t")
		}
		runtime.Vars()[varName] = value
	case IdentifierType:
		// age, query age and assigment
		value, ok := runtime.Vars()[node.GetText()]
		if ok {
			result = value.(int)
		} else {
			fmt.Printf("syntax err: var %s not exit\n", node.GetText())
		}

	default:

	}

	if indent == "" {
		if node.GetNodeType() != Program {
			fmt.Println(result)
		}
	}
	return result
}

type Runtime struct {
	variables map[string]interface{}
	verbose   bool
}

func NewRuntime(verbose bool) *Runtime {
	return &Runtime{
		variables: make(map[string]interface{}),
		verbose:   verbose,
	}
}

func (r *Runtime) Vars() map[string]interface{} {
	return r.variables
}

func (r *Runtime) Verbose() bool {
	return r.verbose
}

var runtime *Runtime

func InitRuntime(verbose bool) {
	runtime = NewRuntime(verbose)
}

func GetRuntime() *Runtime {
	return runtime
}
