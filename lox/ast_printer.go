package lox

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type AstPrinter struct {
}

func (t AstPrinter) VisitExprTernary(ternary Ternary) any {
	return t.Parenthesize("?", ternary.Condition, ternary.TrueExpr, ternary.FalseExpr)
}

func (t AstPrinter) VisitExprBinary(binary Binary) any {
	return t.Parenthesize(binary.Operator.Lexeme, binary.Left, binary.Right)
}

func (t AstPrinter) VisitExprGrouping(grouping Grouping) any {
	return t.Parenthesize("group", grouping.Expression)
}

func (t AstPrinter) VisitExprLiteral(literal Literal) any {
	if literal.Value == nil {
		return "nil"
	}
	switch l := literal.Value.(type) {
	case float64:
		if math.Floor(l) == l {
			return fmt.Sprintf("%.1f", literal.Value)
		}
		return strconv.FormatFloat(l, 'g', -1, 64)
	default:
		return fmt.Sprint(l)
	}
}

func (t AstPrinter) VisitExprUnary(unary Unary) any {
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