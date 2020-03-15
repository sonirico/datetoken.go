package parser

import (
	"fmt"
	"strconv"

	"github.com/sonirico/datetoken.go/ast"
	"github.com/sonirico/datetoken.go/lexer"
	"github.com/sonirico/datetoken.go/token"
)

// Parser will construct the node-tree and the ast for a token
type Parser struct {
	lexer *lexer.Lexer

	errors []string

	curToken  token.Token
	peekToken token.Token
}

// New returns a new instance of Parser
func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{
		lexer:  lexer,
		errors: []string{},
	}
	parser.nextToken()
	return parser
}

func (p *Parser) addError(msg string) {
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) curTokenIs(expectedTokenType token.Type) bool {
	return expectedTokenType == p.curToken.Type
}

func (p *Parser) curError(tokenType token.Type) {
	msg := fmt.Sprintf("Expected current token to be of type '%s'. Got '%s' -> %s",
		tokenType, p.curToken.Type, p.curToken.Literal)
	p.addError(msg)
}

func (p *Parser) parseValueNode() *ast.ValueNode {
	return &ast.ValueNode{
		Token: p.curToken,
	}
}

func (p *Parser) parseArithmeticNode() *ast.ArithmeticNode {
	node := &ast.ArithmeticNode{
		Amount: 1,
		Token:  p.curToken,
		Sign:   p.curToken.Literal,
	}
	p.nextToken()
	if p.curTokenIs(token.Number) {
		parsedAmount, err := strconv.ParseInt(p.curToken.Literal, 10, 64)
		if err != nil {
			p.addError(fmt.Sprintf("unable to parse number %s", p.curToken.Literal))
		}
		node.Amount = parsedAmount
		p.nextToken()
	}
	if p.curTokenIs(token.Unit) {
		node.Unit = p.curToken.Literal
	} else {
		// TODO: disyuntive error
		p.curError(token.Number)
		p.curError(token.Unit)
	}
	return node
}

func (p *Parser) parseSnapNode() *ast.SnapNode {
	node := &ast.SnapNode{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}
	p.nextToken()
	if !(p.curTokenIs(token.Unit) || p.curTokenIs(token.Wd)) {
		p.addError(fmt.Sprintf("expected token to be Unit or Wd, got '%s' instead", p.curToken.Type))
		return nil
	}
	node.Unit = p.curToken.Literal
	return node
}

func (p *Parser) parseNode() ast.Node {
	switch p.curToken.Type {
	case token.Start:
		return p.parseValueNode()
	case token.Plus, token.Minus:
		return p.parseArithmeticNode()
	case token.SnapEnd, token.SnapStart:
		return p.parseSnapNode()
	}
	return nil
}

// Parse will iterate the tokens and build the AST
func (p *Parser) Parse() *ast.RootNode {
	root := ast.NewRootNode()
	p.nextToken()
	for p.curToken.Type != token.End {
		node := p.parseNode()
		if node != nil {
			root.AddNode(node)
		}
		p.nextToken()
	}
	return root
}

// Errors will returns a list of encountered errors during the parsing stage
func (p *Parser) Errors() []string {
	return p.errors
}
