package lox

type Expr interface {
	Accept(visitor Visitor) any
}

type Visitor interface {
	// VisitExprTernary(ternary Ternary) any
	VisitExprBinary(binary Binary) any
	VisitExprGrouping(grouping Grouping) any
	VisitExprLiteral(literal Literal) any
	VisitExprUnary(unary Unary) any
}

type Ternary struct {
	Condition Expr
	TrueExpr  Expr
	FalseExpr Expr
}

// func (t Ternary) Accept(visitor Visitor) any {
// 	return visitor.VisitExprTernary(t)
// }

type Binary struct {
	Left     Expr
	Operator Token
	Right    Expr
}

func (t Binary) Accept(visitor Visitor) any {
	return visitor.VisitExprBinary(t)
}

type Grouping struct {
	Expression Expr
}

func (t Grouping) Accept(visitor Visitor) any {
	return visitor.VisitExprGrouping(t)
}

type Literal struct {
	Value any
}

func (t Literal) Accept(visitor Visitor) any {
	return visitor.VisitExprLiteral(t)
}

type Unary struct {
	Operator Token
	Right    Expr
}

func (t Unary) Accept(visitor Visitor) any {
	return visitor.VisitExprUnary(t)
}
