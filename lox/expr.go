package lox

type Expr interface {
	Accept(visitor Visitor) interface{}
}

type Visitor interface {
	VisitExprBinary(binary Binary) interface{}
	VisitExprGrouping(grouping Grouping) interface{}
	VisitExprLiteral(literal Literal) interface{}
	VisitExprUnary(unary Unary) interface{}
}

type Binary struct {
	Left Expr
	Operator Token
	Right Expr
}

func (t Binary) Accept(visitor Visitor) interface{} {
	return visitor.VisitExprBinary(t)
}

type Grouping struct {
	Expression Expr
}

func (t Grouping) Accept(visitor Visitor) interface{} {
	return visitor.VisitExprGrouping(t)
}

type Literal struct {
	Value interface{}
}

func (t Literal) Accept(visitor Visitor) interface{} {
	return visitor.VisitExprLiteral(t)
}

type Unary struct {
	Operator Token
	Right Expr
}

func (t Unary) Accept(visitor Visitor) interface{} {
	return visitor.VisitExprUnary(t)
}

