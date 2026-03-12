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
package algebra_expression

import (
	"math"
)

type Expression struct {
	Type ExpressionType `json:"type"`

	Value     *float64      `json:"value,omitempty"`
	Name      string        `json:"name,omitempty"`
	Arguments []*Expression `json:"args,omitempty"`

	Cache ExpressionCache `json:"-"`
}

func (expression *Expression) String() string {
	return algebraicString(expression)
}

func (expression *Expression) Sum(addition *Expression, others ...*Expression) *Expression {
	var additions []*Expression = append([]*Expression{addition}, others...)

	switch expression.Type {
	case "":
		if len(additions) < 2 {
			additions = append(additions, Int(0))
		}

		*expression = Expression{
			Type: ADDITION,

			Arguments: additions,
		}

	case ADDITION:
		expression.Arguments = append(
			expression.Arguments,
			additions...,
		)

	default:
		var snapshot *Expression = &Expression{
			Type:      expression.Type,
			Value:     expression.Value,
			Name:      expression.Name,
			Arguments: expression.Arguments,
		}

		*expression = Expression{
			Type: ADDITION,

			Arguments: append(
				[]*Expression{
					snapshot,
				},
				additions...,
			),
		}
	}

	return expression
}

func (expression *Expression) Subtract(subtraction *Expression, others ...*Expression) *Expression {
	var subtractions []*Expression = append([]*Expression{subtraction}, others...)

	for i := range subtractions {
		switch subtractions[i].Type {
		case INTEGER, FLOAT:
			*subtractions[i].Value = -*subtractions[i].Value

		default:
			subtractions[i] = Multiply(
				subtractions[i],
				Int(-1),
			)
		}
	}

	return expression.Sum(subtractions[0], subtractions[1:]...)
}

func (expression *Expression) Multiply(multiplier *Expression, others ...*Expression) *Expression {
	var multipliers []*Expression = append([]*Expression{multiplier}, others...)

	switch expression.Type {
	case "":
		if len(multipliers) < 2 {
			multipliers = append(multipliers, Int(1))
		}

		*expression = Expression{
			Type: MULTIPLICATION,

			Arguments: multipliers,
		}

	case MULTIPLICATION:
		expression.Arguments = append(
			expression.Arguments,
			multipliers...,
		)

	default:
		var snapshot *Expression = &Expression{
			Type:      expression.Type,
			Value:     expression.Value,
			Name:      expression.Name,
			Arguments: expression.Arguments,
		}

		*expression = Expression{
			Type: MULTIPLICATION,

			Arguments: append(
				[]*Expression{
					snapshot,
				},
				multipliers...,
			),
		}
	}

	return expression
}

func (expression *Expression) Divide(denominator *Expression, others ...*Expression) *Expression {
	divisors := []*Expression{
		Pow(
			denominator,
			Int(-1),
		),
	}

	for _, other := range others {
		divisors = append(
			divisors,
			Pow(
				other,
				Int(-1),
			),
		)
	}

	return expression.Multiply(divisors[0], divisors[1:]...)
}

func (expression *Expression) Execute(operator float64) (result float64) {
	if expression == nil {
		return math.NaN()
	}

	switch expression.Type {
	case INTEGER, FLOAT:
		if expression.Value == nil {
			return math.NaN()
		}
		return *expression.Value

	case SYMBOL:
		switch expression.Name {
		case "":
			return math.NaN()
		case "e":
			return math.E
		case "pi":
			return math.Pi
		}

		/*
			Workable only for 2d equations, and will ignore completely the true symbology.
			For compatibility with further dimensions, the expression parameter should be a slice of
				a structure linking a float value with a variable name (or expression name in the
				library	context)
			Some extra validations must happen before calling expression.Execute() due to the
				possibility of missing operators, since Execute() itself is not aware of what
				happen in its child or parent Execute() calling.
		*/
		return operator

	case ADDITION:
		for _, argument := range expression.Arguments {
			result += argument.Execute(operator)
		}
		return result

	case MULTIPLICATION:
		result = 1
		for _, argument := range expression.Arguments {
			result *= argument.Execute(operator)
		}
		return result

	case POWER:
		if len(expression.Arguments) != 2 {
			return math.NaN()
		}

		var base float64 = expression.Arguments[0].Execute(operator)
		var exponent float64 = expression.Arguments[1].Execute(operator)

		if base == 0 && exponent == 0 {
			return math.NaN()
		}

		return math.Pow(base, exponent)

	case SIN:
		if len(expression.Arguments) != 1 {
			return math.NaN()
		}
		return math.Sin(expression.Arguments[0].Execute(operator))

	case COS:
		if len(expression.Arguments) != 1 {
			return math.NaN()
		}
		return math.Cos(expression.Arguments[0].Execute(operator))

	case TAN:
		if len(expression.Arguments) != 1 {
			return math.NaN()
		}
		return math.Tan(expression.Arguments[0].Execute(operator))

	case LOGARITHMIC:
		switch len(expression.Arguments) {
		case 1:
			return math.Log(expression.Arguments[0].Execute(operator))

		case 2:
			var value float64 = expression.Arguments[0].Execute(operator)
			var base float64 = expression.Arguments[1].Execute(operator)

			if base <= 0 || base == 1 {
				return math.NaN()
			}

			return math.Log(value) / math.Log(base)

		default:
			return math.NaN()
		}

	case EXPONENTIAL:
		if len(expression.Arguments) != 1 {
			return math.NaN()
		}
		return math.Exp(expression.Arguments[0].Execute(operator))

	default:
		return math.NaN()
	}
}
