package parser

import (
	"github.com/sonirico/datetoken.go/ast"
	"github.com/sonirico/datetoken.go/lexer"
	"github.com/sonirico/datetoken.go/token"
)

type Parser struct {
	lexer *lexer.Lexer
}

func NewParser(lexer *lexer.Lexer) *Parser {
	return &Parser{
		lexer: lexer,
	}
}

func (p *Parser) ParseNode() ast.Node {
	return nil
}

func (p *Parser) Parse() *ast.RootNode {
	root := ast.NewRootNode()
	tok := p.lexer.NextToken()
	for tok.Type != token.End {
		node := p.ParseNode()
		if node != nil {
			root.AddNode(node)
		}
		tok = p.lexer.NextToken()
	}
	return root
}
