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
package test

import (
	"reflect"
	"testing"

	algebra_equation "github.com/FabioLuisBoni/go-algebra/equation"
	algebra_expression "github.com/FabioLuisBoni/go-algebra/equation/expression"
	algebra_frame "github.com/FabioLuisBoni/go-algebra/frame"
	algebra_numeric_method "github.com/FabioLuisBoni/go-algebra/numeric_method"
)

func TestBaseOnLeastSquares(t *testing.T) {
	t.Logf("testing BaseOn for LeastSquares\n\n")
	{ // CASE SUM
		var signature string = "f(x)"
		var value_integer_1 int = 2
		var value_integer_2 int = 1
		var value_integer_pow int = -1
		var value_integer_log int = 10
		var value_float float64 = 1.1

		var equation *algebra_equation.Equation = algebra_equation.NewEquation(signature).Sum(
			algebra_expression.Int(value_integer_1),
			algebra_expression.Multiply(
				algebra_expression.Float(value_float),
				algebra_expression.Sum(
					algebra_expression.Symbol("x"),
					algebra_expression.Int(value_integer_2),
				),
			),
			algebra_expression.Exp(
				algebra_expression.Symbol("x"),
			),
			algebra_expression.Sin(
				algebra_expression.Pow(
					algebra_expression.Symbol("x"),
					algebra_expression.Int(value_integer_pow),
				),
			),
			algebra_expression.Cos(
				algebra_expression.Symbol("x"),
			),
			algebra_expression.Tan(
				algebra_expression.Symbol("x"),
			),
			algebra_expression.Ln(
				algebra_expression.Symbol("x"),
			),

			algebra_expression.Log(
				algebra_expression.Int(value_integer_log),
				algebra_expression.Symbol("x"),
			),
		)

		var expected *algebra_numeric_method.LeastSquares = &algebra_numeric_method.LeastSquares{
			EquationSignature: signature,
			Base: []*algebra_expression.Expression{
				equation.Expression.Arguments[0].Clone(),
				equation.Expression.Arguments[1].Clone(),
				equation.Expression.Arguments[2].Clone(),
				equation.Expression.Arguments[3].Clone(),
				equation.Expression.Arguments[4].Clone(),
				equation.Expression.Arguments[5].Clone(),
				equation.Expression.Arguments[6].Clone(),
				equation.Expression.Arguments[7].Clone(),
			},
		}

		var method *algebra_numeric_method.LeastSquares = algebra_numeric_method.NewLeastSquares(signature)

		err := method.BaseOn(equation)
		if err != nil {
			t.Fatalf("error basing least squares on equation: %s", err)
		}

		for _, base := range method.Base {
			base.ClearCache()
		}

		if !reflect.DeepEqual(method, expected) {
			t.Fatalf("error with least squares\nexpected: %#v\ngot     : %#v", expected, method)
		}
	}

	t.Logf("finished testing BaseOn for LeastSquares\n\n")
}

func TestSolveLeastSquares(t *testing.T) {
	t.Logf("testing Solve for LeastSquares\n\n")
	var signature string = "f(x)"

	var method *algebra_numeric_method.LeastSquares = algebra_numeric_method.NewLeastSquares(signature)

	method.Base = []*algebra_expression.Expression{
		algebra_expression.Float(1),
		algebra_expression.Symbol("x"),
	}

	var expected *algebra_equation.Equation = algebra_equation.NewEquation(signature).SetExpression(
		algebra_expression.Sum(
			algebra_expression.Multiply(
				algebra_expression.Float(1),
				algebra_expression.Float(-4),
			),
			algebra_expression.Multiply(
				algebra_expression.Symbol("x"),
				algebra_expression.Float(2),
			),
		),
	)

	var frame algebra_frame.Frame = algebra_frame.Frame{
		Point: []algebra_frame.Point{
			{
				X: 5,
				Y: 6,
			},
			{
				X: 6,
				Y: 8,
			},
		},
	}

	result, err := method.Solve(frame)
	if err != nil {
		t.Fatalf("error at least squares: %s", err)
	}

	result.Expression.ClearCache()

	if !result.Expression.Equal(expected.Expression) {
		t.Fatalf("\nexpected: %+v\ngot     : %+v", expected, result)
	}

	t.Logf("finished testing Solve for LeastSquares\n\n")
}
