package engine

import (
	"fmt"
	"formula/core"
	"strconv"
)

type Parser struct {
	lex      *Lexer
	curr     Token
	debugger *Debugger
	depth    int
}

func NewParser(input string, debugger *Debugger) *Parser {
	p := &Parser{
		lex:      NewLexer(input, debugger),
		debugger: debugger,
	}
	p.next()
	return p
}

func (p *Parser) next() {
	p.curr = p.lex.NextToken()
	for p.curr.Type == TokenComment {
		p.curr = p.lex.NextToken()
	}
}

func (p *Parser) expect(t TokenType) {
	if p.curr.Type != t {
		panic(fmt.Sprintf("expected %v got %v", t, p.curr))
	}

	p.next()
}

func (p *Parser) ParseExpr() core.Expr {
	switch p.curr.Type {
	case TokenIdent:
		name := p.curr.Value
		p.next()
		p.expect(TokenLParen)

		var args core.Exprs
		if p.curr.Type != TokenRParen {
			for {
				p.depth++
				args = append(args, p.ParseExpr())
				p.depth--
				if p.curr.Type == TokenComma {
					p.next()
					continue
				}
				break
			}
		}

		p.expect(TokenRParen)
		eFunc := &core.Func{Name: name, Args: args}
		if p.debugger != nil {
			p.debugger.Log(PhaseParser, p.depth, fmt.Sprintf("Exit Func: %s (args=%d)", eFunc.Name, len(eFunc.Args)))
		}
		return eFunc
	case TokenNumber:
		val := p.curr.Value
		num, _ := strconv.ParseFloat(val, 64)
		p.next()
		literal := &core.Literal{Value: num}
		if p.debugger != nil {
			p.debugger.Log(PhaseParser, p.depth, fmt.Sprintf("Literal: %v", literal))
		}
		return literal
	case TokenString:
		val := p.curr.Value
		p.next()
		literal := &core.Literal{Value: val}
		if p.debugger != nil {
			p.debugger.Log(PhaseParser, p.depth, fmt.Sprintf("Literal: %v", literal))
		}
		return literal
	case TokenTrue:
		p.next()
		literal := &core.Literal{Value: true}
		if p.debugger != nil {
			p.debugger.Log(PhaseParser, p.depth, fmt.Sprintf("Literal: %v", literal))
		}
		return literal
	case TokenFalse:
		p.next()
		literal := &core.Literal{Value: false}
		if p.debugger != nil {
			p.debugger.Log(PhaseParser, p.depth, fmt.Sprintf("Literal: %v", literal))
		}
		return literal
	default:
		panic(fmt.Sprintf("unexpected token %v", p.curr))
	}
}
