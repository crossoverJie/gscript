package syntax

import (
	"errors"
	"github.com/crossoverJie/gscript/token"
	"log"
)

// ASTNode AST tree
type ASTNode struct {
	parent   *ASTNode
	children []*ASTNode
	nodeType ASTNodeType
	text     string
}

type ASTNodeType string

// AST tree: node type.
const (
	Program            ASTNodeType = "Program"
	IntDeclaration     ASTNodeType = "IntDeclaration"
	AssignmentStmt     ASTNodeType = "AssignmentStmt"
	PrimaryType        ASTNodeType = "Primary"
	MultiplicativeType ASTNodeType = "Multiplicative"
	AdditiveType       ASTNodeType = "Additive"
	IdentifierType     ASTNodeType = "Identifier"
	IntLiteralType     ASTNodeType = "IntLiteral"
)

func NewASTNode(nodeType ASTNodeType, text string) *ASTNode {
	return &ASTNode{nodeType: nodeType, text: text}
}

func (a *ASTNode) AddChild(child *ASTNode) {
	a.children = append(a.children, child)
	child.parent = a
}

func (a *ASTNode) GetNodeType() ASTNodeType {
	return a.nodeType
}
func (a *ASTNode) GetText() string {
	return a.text
}

func (a *ASTNode) GetChildren() []*ASTNode {
	return a.children
}

func (a *ASTNode) Print(indent string) {
	log.Printf("%s %s %s", indent, a.GetNodeType(), a.GetText())
	for _, astNode := range a.GetChildren() {
		astNode.Print(indent + "\t")
	}
}

// TokenReader simulate token reader stream.
type TokenReader struct {
	tokens []*token.TokenType
	pos    uint64
}

func NewTokenReader(tokens []*token.TokenType) *TokenReader {
	return &TokenReader{tokens: tokens, pos: 0}
}

// Peek Pre read.
func (t *TokenReader) Peek() *token.TokenType {
	if int(t.pos) >= len(t.tokens) {
		return nil
	}
	return t.tokens[t.pos]
}

// Read read from token stream.
func (t *TokenReader) Read() *token.TokenType {
	if int(t.pos) >= len(t.tokens) {
		return nil
	}
	tokenType := t.tokens[t.pos]
	t.pos += 1
	return tokenType
}

func (t *TokenReader) UnRead() {
	if t.pos > 0 {
		t.pos -= 1
	}
}

func (t *TokenReader) GetPos() uint64 {
	return t.pos
}

func (t *TokenReader) SetPos(pos uint64) {
	if pos >= 0 && int(pos) < len(t.tokens) {
		t.pos = pos
	}
}

// IntDeclare 整形声明
func IntDeclare(tokenReader *TokenReader) (*ASTNode, error) {
	var node *ASTNode
	tokenType := tokenReader.Peek()
	if tokenType == nil || tokenType.TokenType() != token.Int {
		return nil, errors.New("syntax err: invalid statement")
	}
	tokenType = tokenReader.Read()

	// 解析标识符
	tokenType = tokenReader.Peek()
	if tokenType == nil {
		return nil, errors.New("invalid statement, miss Identifier")
	}
	tokenType = tokenReader.Read()
	// 加入 AST 节点
	node = NewASTNode(IntDeclaration, tokenType.Value())

	// 解析 Assignment=
	tokenType = tokenReader.Peek()
	if tokenType == nil || tokenType.TokenType() != token.Assignment {
		return nil, errors.New("syntax err: invalid statement, miss Assignment")
	}
	tokenType = tokenReader.Read()
	// 匹配后续的表达式
	child, err := AdditiveLoop(tokenReader)
	if err != nil {
		return nil, err
	}
	node.AddChild(child)

	// 解析最后的结束符
	tokenType = tokenReader.Peek()
	if tokenType == nil || tokenType.TokenType() != token.Enter {
		return nil, errors.New("syntax err: invalid statement, miss enter")
	}
	tokenType = tokenReader.Read()
	// 加入 AST

	return node, nil
}

// AdditiveLoop 加法循环解析，解决结合性问题
func AdditiveLoop(tokenReader *TokenReader) (*ASTNode, error) {
	child1, err := MultiplicativeLoop(tokenReader)
	node := child1
	if err != nil {
		return nil, err
	}
	for {
		// 应用 add' 规则
		tokenType := tokenReader.Peek()
		if tokenType != nil && (tokenType.TokenType() == token.Plus || tokenType.TokenType() == token.Minus) {
			tokenType = tokenReader.Read() //消耗加号

			child2, err := MultiplicativeLoop(tokenReader) // 读取下一个节点
			if err != nil {
				return nil, err
			}
			node = NewASTNode(AdditiveType, tokenType.Value())
			node.AddChild(child1)
			node.AddChild(child2)
			child1 = node
		} else {
			break
		}
	}

	if node != nil {
		return node, nil
	} else {
		return nil, errors.New("syntax err, invalid additive")
	}

}

// MultiplicativeLoop 乘法解析，解决结合性问题
func MultiplicativeLoop(tokenReader *TokenReader) (*ASTNode, error) {
	child1, err := Primary(tokenReader) // 先做一次基础解析，只会往后解析一位
	node := child1
	if err != nil {
		return nil, err
	}
	for {
		tokenType := tokenReader.Peek()
		if tokenType != nil && (tokenType.TokenType() == token.Star || tokenType.TokenType() == token.Slash) {
			tokenType = tokenReader.Read()

			child2, err := Primary(tokenReader)
			if err != nil {
				return nil, err
			}
			node = NewASTNode(MultiplicativeType, tokenType.Value())
			node.AddChild(child1) // 新节点在顶层，保证了结合性
			node.AddChild(child2)
			child1 = node
		} else {
			break
		}
	}
	if node != nil {
		return node, nil
	} else {
		return nil, errors.New("syntax err, invalid multiplicative ")
	}
}

// Primary Analyze primary  1+2->1;
func Primary(tokenReader *TokenReader) (*ASTNode, error) {
	var node *ASTNode
	tokenType := tokenReader.Peek() // 预读
	if tokenType == nil {
		return nil, errors.New("syntax err, invalid primary")
	}

	if tokenType.TokenType() == token.IntLiteral {
		// 整形字面量
		tokenType = tokenReader.Read() // 真正读取
		node = NewASTNode(IntLiteralType, tokenType.Value())
		return node, nil
	} else if tokenType.TokenType() == token.Identifier {
		// 标识符
		tokenType = tokenReader.Read() // 真正读取
		node = NewASTNode(IdentifierType, tokenType.Value())
		return node, nil
	} else if tokenType.TokenType() == token.LeftParen {
		// 左括号，需要重头解析加法表达式
		tokenType = tokenReader.Read()
		var err error
		node, err = AdditiveLoop(tokenReader)
		if err != nil {
			return nil, err
		}
		tokenType = tokenReader.Peek()
		if tokenType != nil && tokenType.TokenType() == token.RightParen {
			tokenType = tokenReader.Read()
		} else {
			return nil, errors.New("syntax err: miss right parenthesis")
		}

	}
	if node != nil {
		return node, nil
	} else {
		return nil, errors.New("syntax err, invalid primary")
	}

}
