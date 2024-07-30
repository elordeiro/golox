package lox

import (
	"fmt"
	"strings"
)

type AstPrinter struct {
}

func (t AstPrinter) VisitExprBinary(binary Binary) interface{} {
	return t.Parenthesize(binary.Operator.Lexeme, binary.Left, binary.Right)
}

func (t AstPrinter) VisitExprGrouping(grouping Grouping) interface{} {
	return t.Parenthesize("group", grouping.Expression)
}

func (t AstPrinter) VisitExprLiteral(literal Literal) interface{} {
	if literal.Value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", literal.Value)
}

func (t AstPrinter) VisitExprUnary(unary Unary) interface{} {
	return t.Parenthesize(unary.Operator.Lexeme, unary.Right)
}

func (t AstPrinter) Parenthesize(name string, exprs ...Expr) string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(name)
	for _, expr := range exprs {
		sb.WriteString(" ")
		sb.WriteString(expr.Accept(t).(string))
	}
	sb.WriteString(")")
	return sb.String()
}

func (t AstPrinter) Print(expr Expr) string {
	return expr.Accept(t).(string)
}

func NewAstPrinter() AstPrinter {
	return AstPrinter{}
}

func PrintAst(expr Expr) string {
	return NewAstPrinter().Print(expr)
}
