/*
	/\\\\\\\\\\\\\\\                  /\\\\\\\\\\\\\         /\\\\\       /\\\\\     /\\\  /\\\\\\\\\\\
	\/\\\///////////                  \/\\\/////////\\\     /\\\///\\\    \/\\\\\\   \/\\\ \/////\\\///
	 \/\\\                             \/\\\       \/\\\   /\\\/  \///\\\  \/\\\/\\\  \/\\\     \/\\\
	  \/\\\\\\\\\\\                     \/\\\\\\\\\\\\\\   /\\\      \//\\\ \/\\\//\\\ \/\\\     \/\\\
	   \/\\\///////                      \/\\\/////////\\\ \/\\\       \/\\\ \/\\\\//\\\\/\\\     \/\\\
	    \/\\\                             \/\\\       \/\\\ \//\\\      /\\\  \/\\\ \//\\\/\\\     \/\\\
	     \/\\\                             \/\\\       \/\\\  \///\\\  /\\\    \/\\\  \//\\\\\\     \/\\\
	      \/\\\              /\\\           \/\\\\\\\\\\\\\/     \///\\\\\/     \/\\\   \//\\\\\  /\\\\\\\\\\\
	       \///              \///            \/////////////         \/////       \///     \/////  \///////////

	Created:    12 mar 2026
	Author:     F. Boni    Email:      fabioboni96@hotmail.com
	Repository: github.com/FabioLuisBoni/go-algebra

Copyright (c) 2026 Fabio Luis Boni - MIT License
*/
package algebra_equation

import (
	"fmt"

	algebra "github.com/FabioLuisBoni/go-algebra/equation/expression"
)

type Equation struct {
	Signature  string
	Expression *algebra.Expression
}

func NewEquation(signature string) *Equation {
	return &Equation{
		Signature:  signature,
		Expression: &algebra.Expression{},
	}
}

func (equation *Equation) SetExpression(expression *algebra.Expression) *Equation {
	equation.Expression = expression

	return equation
}

func (equation *Equation) Sum(expression *algebra.Expression, others ...*algebra.Expression) *Equation {
	var toSum []*algebra.Expression = make([]*algebra.Expression, len(others))
	for i, other := range others {
		toSum[i] = other.Clone()
	}

	equation.Expression.Sum(expression.Clone(), toSum...)

	return equation
}

func (equation *Equation) Subtract(expression *algebra.Expression, others ...*algebra.Expression) *Equation {
	var toSubtract []*algebra.Expression = make([]*algebra.Expression, len(others))
	for i, other := range others {
		toSubtract[i] = other.Clone()
	}

	equation.Expression.Subtract(expression.Clone(), toSubtract...)

	return equation
}

func (equation *Equation) Multiply(expression *algebra.Expression, others ...*algebra.Expression) *Equation {
	var toMultiply []*algebra.Expression = make([]*algebra.Expression, len(others))
	for i, other := range others {
		toMultiply[i] = other.Clone()
	}

	equation.Expression.Multiply(expression.Clone(), toMultiply...)

	return equation
}

func (equation *Equation) Divide(expression *algebra.Expression, others ...*algebra.Expression) *Equation {
	var toDivide []*algebra.Expression = make([]*algebra.Expression, len(others))
	for i, other := range others {
		toDivide[i] = other.Clone()
	}

	equation.Expression.Divide(expression.Clone(), toDivide...)

	return equation
}

func (equation *Equation) Solve(variable float64) (result float64) {
	return equation.Expression.Solve(variable)
}

func (equation *Equation) String() string {
	return fmt.Sprintf("%s = %s", equation.Signature, equation.Expression)
}

func (equation *Equation) IsMalformedStructure() bool {
	return equation.Expression.IsMalformedStructure()
}

func (equation *Equation) IsIndefiniteness() bool {
	return equation.Expression.IsIndefiniteness()
}

func (equation *Equation) IsConstant() bool {
	return equation.Expression.IsConstant()
}

func (equation *Equation) IsZero() bool {
	return equation.Expression.IsZero()
}

func (equation *Equation) IsAbsoluteOne() bool {
	return equation.Expression.IsAbsoluteOne()
}

func (equation *Equation) IsEuler() bool {
	return equation.Expression.IsEuler()
}

func (equation *Equation) IsFraction() bool {
	return equation.Expression.IsFraction()
}

func (equation *Equation) IsInteger() bool {
	return equation.Expression.IsInteger()
}

func (equation *Equation) IsEvenInteger() bool {
	return equation.Expression.IsEvenInteger()
}

func (equation *Equation) IsOddInteger() bool {
	return equation.Expression.IsOddInteger()
}
