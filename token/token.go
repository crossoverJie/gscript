package token

import "fmt"

type Type string

const (
	Init       string = "Init"
	Identifier        = "Identifier"
	GT                = "GT"
	GE                = "GE"
	IntLiteral        = "IntLiteral"
	Float             = "Float"
	Int               = "Int"
	Int1              = "Int1"
	Int2              = "Int2"
	Int3              = "Int3"
	Assignment        = "Assignment"
	Plus              = "Plus"
	Minus             = "Minus"
	Star              = "Star"
	Slash             = "Slash"
	LeftParen         = "LeftParen"
	RightParen        = "RightParen"
	Enter             = "Enter"
)

type TokenType struct {
	tokenType string
	value     string
}

func (t *TokenType) TokenType() string {
	return t.tokenType
}

func (t *TokenType) Value() string {
	return t.value
}

// Tokenize Lexical analysis
func Tokenize(str string) []*TokenType {
	bytes := []byte(str)
	var result []*TokenType
	var values []byte
	status := Init
	for _, b := range bytes {
		switch status {
		case Init:
			status, values = InitStatus(b, values)
			break
		case Identifier:
			if isLetter(b) {
				values = append(values, b)
			} else {
				t := &TokenType{
					tokenType: Identifier,
					value:     string(values),
				}
				result = append(result, t)
				values = nil

				status, values = InitStatus(b, values)
				break
			}
		case Int1:
			if b == 'n' {
				status = Int2
				values = append(values, b)
			} else if isLetter(b) || isDigit(b) {
				status = Identifier
				values = append(values, b)
			} else {
				status, values = InitStatus(b, values)
			}
		case Int2:
			if b == 't' {
				status = Int3
				values = append(values, b)
			} else if isLetter(b) || isDigit(b) {
				status = Identifier
				values = append(values, b)
			} else {
				status, values = InitStatus(b, values)
			}
		case Int3:
			if b == ' ' {
				status = Int
			} else {
				status = Identifier
				values = append(values, b)
			}
		case Int:
			t := &TokenType{
				tokenType: Int,
				value:     string(values),
			}
			result = append(result, t)
			values = nil

			status, values = InitStatus(b, values)
			break
		case Assignment:
			t := &TokenType{
				tokenType: Assignment,
				value:     string(values),
			}
			result = append(result, t)
			values = nil

			status, values = InitStatus(b, values)
		case GT:
			if b == '=' { // =
				status = GE
				values = append(values, b)
				t := &TokenType{
					tokenType: GE,
					value:     string(values),
				}
				result = append(result, t)
				values = nil
				break
			} else {
				t := &TokenType{
					tokenType: GT,
					value:     string(values),
				}
				result = append(result, t)
				values = nil
				status, values = InitStatus(b, values)
				break
			}
		case GE:
			status, values = InitStatus(b, values)
			break
		case IntLiteral:
			if b == '.' {
				values = append(values, b)
				status = Float
				break
			}
			if isDigit(b) {
				values = append(values, b)
			} else {
				t := &TokenType{
					tokenType: IntLiteral,
					value:     string(values),
				}
				result = append(result, t)
				values = nil
				status, values = InitStatus(b, values)
				break
			}
		case Float:
			if b == '.' {
				fmt.Println("invalid float")
				return nil
			}
			if isDigit(b) {
				values = append(values, b)
			} else {
				t := &TokenType{
					tokenType: Float,
					value:     string(values),
				}
				result = append(result, t)
				values = nil
				status, values = InitStatus(b, values)
				break
			}
		case Plus:
			t := &TokenType{
				tokenType: Plus,
				value:     string(values),
			}
			result = append(result, t)
			values = nil
			status, values = InitStatus(b, values)
		case Minus:
			t := &TokenType{
				tokenType: Minus,
				value:     string(values),
			}
			result = append(result, t)
			values = nil
			status, values = InitStatus(b, values)
		case Star:
			t := &TokenType{
				tokenType: Star,
				value:     string(values),
			}
			result = append(result, t)
			values = nil
			status, values = InitStatus(b, values)
		case Slash:
			t := &TokenType{
				tokenType: Slash,
				value:     string(values),
			}
			result = append(result, t)
			values = nil
			status, values = InitStatus(b, values)
		case LeftParen:
			t := &TokenType{
				tokenType: LeftParen,
				value:     string(values),
			}
			result = append(result, t)
			values = nil
			status, values = InitStatus(b, values)
		case RightParen:
			t := &TokenType{
				tokenType: RightParen,
				value:     string(values),
			}
			result = append(result, t)
			values = nil
			status, values = InitStatus(b, values)
		case Enter:
			t := &TokenType{
				tokenType: Enter,
				value:     string(values),
			}
			result = append(result, t)
			values = nil
			status, values = InitStatus(b, values)
		}
	}

	// read last one.
	if len(values) > 0 {
		t := &TokenType{
			tokenType: status,
			value:     string(values),
		}
		result = append(result, t)
	}

	//for _, tokenType := range result {
	//	log.Printf("%s \t %s", tokenType.tokenType, tokenType.value)
	//}

	return result
}

// InitStatus finite state machine
func InitStatus(b byte, values []byte) (string, []byte) {
	if isLetter(b) {
		values = append(values, b)
		if b == 'i' {
			return Int1, values
		}
		return Identifier, values
	}

	if b == '>' {
		values = append(values, b)
		return GT, values
	}
	if b == '=' {
		values = append(values, b)
		return Assignment, values
	}
	if isDigit(b) {
		values = append(values, b)
		return IntLiteral, values
	}

	if b == '+' {
		values = append(values, b)
		return Plus, values
	}
	if b == '-' {
		values = append(values, b)
		return Minus, values
	}
	if b == '*' {
		values = append(values, b)
		return Star, values
	}
	if b == '/' {
		values = append(values, b)
		return Slash, values
	}
	if b == '(' {
		values = append(values, b)
		return LeftParen, values
	}
	if b == ')' {
		values = append(values, b)
		return RightParen, values
	}
	if b == '\n' {
		values = append(values, b)
		return Enter, values
	}

	return Init, values
}

func isLetter(b byte) bool {
	return b >= 65 && b <= 122
}

func isDigit(b byte) bool {
	return b >= 48 && b <= 57
}
