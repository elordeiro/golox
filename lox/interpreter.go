package lox

import "fmt"

type Interpreter struct {
	lox *Lox
}

type RuntimeError struct {
	Token   Token
	Message string
}

func (r RuntimeError) Error() string {
	return r.Message
}

func NewInterpreter(lox *Lox) *Interpreter {
	return &Interpreter{lox: lox}
}

func (i *Interpreter) VisitExprBinary(binary Binary) any {
	left := i.evaluate(binary.Left)
	right := i.evaluate(binary.Right)

	switch binary.Operator.Type {
	case GREATER:
		err := i.checkNumberOperands(binary.Operator, left, right)
		if err != nil {
			return err
		}
		return left.(float64) > right.(float64)
	case GREATER_EQUAL:
		err := i.checkNumberOperands(binary.Operator, left, right)
		if err != nil {
			return err
		}
		return left.(float64) >= right.(float64)
	case LESS:
		err := i.checkNumberOperands(binary.Operator, left, right)
		if err != nil {
			return err
		}
		return left.(float64) < right.(float64)
	case LESS_EQUAL:
		err := i.checkNumberOperands(binary.Operator, left, right)
		if err != nil {
			return err
		}
		return left.(float64) <= right.(float64)
	case BANG_EQUAL:
		return !i.isEqual(left, right)
	case EQUAL_EQUAL:
		return i.isEqual(left, right)
	case MINUS:
		err := i.checkNumberOperands(binary.Operator, left, right)
		if err != nil {
			return err
		}
		return left.(float64) - right.(float64)
	case PLUS:
		// Left and Right must be numbers
		_, leftIsFloat := left.(float64)
		_, rightIsFloat := right.(float64)
		if leftIsFloat && rightIsFloat {
			return left.(float64) + right.(float64)
		}

		// Left and Right must be strings
		_, leftIsString := left.(string)
		_, rightIsString := right.(string)
		if leftIsString && rightIsString {
			return left.(string) + right.(string)
		}

		return RuntimeError{binary.Operator, "Operands must be two numbers or two strings"}
	case SLASH:
		err := i.checkNumberOperands(binary.Operator, left, right)
		if err != nil {
			return err
		}
		return left.(float64) / right.(float64)
	case STAR:
		err := i.checkNumberOperands(binary.Operator, left, right)
		if err != nil {
			return err
		}
		return left.(float64) * right.(float64)
	}

	// Unreachable
	return nil
}

func (i *Interpreter) VisitExprGrouping(grouping Grouping) any {
	return i.evaluate(grouping.Expression)
}

func (i *Interpreter) VisitExprLiteral(literal Literal) any {
	return literal.Value
}

func (i *Interpreter) VisitExprUnary(unary Unary) any {
	right := i.evaluate(unary.Right)

	switch unary.Operator.Type {
	case MINUS:
		err := i.checkNumberOperand(unary.Operator, right)
		if err != nil {
			return err
		}
		return -right.(float64)
	case BANG:
		return !i.isTruty(right)
	}
	return nil
}

func (i *Interpreter) evaluate(expr Expr) any {
	return expr.Accept(i)
}

func (i *Interpreter) isTruty(val any) bool {
	switch v := val.(type) {
	case nil:
		return false
	case bool:
		return v
	default:
		return true
	}
}

func (i *Interpreter) isEqual(val1, val2 any) bool {
	if val1 == nil && val2 == nil {
		return true
	}
	if val1 == nil {
		return false
	}
	return val1 == val2
}

func (i *Interpreter) stringify(obj any) string {
	if obj == nil {
		return "nil"
	}

	if val, ok := obj.(float64); ok {
		text := FormatNumber(val)
		if text[len(text)-2:] == ".0" {
			return text[:len(text)-2]
		}
		return text
	}

	return fmt.Sprintf("%v", obj)
}

func (i *Interpreter) checkNumberOperand(operator Token, operand any) error {
	if _, ok := operand.(float64); ok {
		return nil
	}
	return RuntimeError{operator, "Operand must be a number."}
}

func (i *Interpreter) checkNumberOperands(operator Token, left, right any) error {
	_, leftIsFloat := left.(float64)
	_, rightIsFloat := right.(float64)
	if leftIsFloat && rightIsFloat {
		return nil
	}
	return RuntimeError{operator, "Operands must be numbers."}
}

func (i *Interpreter) Interpret(expr Expr) {
	value := i.evaluate(expr)
	if v, ok := value.(error); ok {
		i.lox.RuntimeError(v.(RuntimeError))
		return
	}

	fmt.Println(i.stringify(value))
}
