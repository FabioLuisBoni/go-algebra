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
	"math"
	"testing"

	algebraic "github.com/FabioLuisBoni/go-algebra/equation/expression"
)

func TestExecuteExpression(t *testing.T) {
	t.Logf("testing Execute for Expression\n\n")
	var value_integer_1 int = 2
	var value_integer_2 int = 1
	var value_integer_pow int = -1
	var value_integer_log int = 10
	var value_float float64 = 1.1
	var x float64 = 5

	var expected float64 = float64(value_integer_1) +
		value_float*(x+float64(value_integer_2)) +
		math.Sin(math.Abs(float64(value_integer_pow))/x) +
		math.Cos(x) +
		math.Tan(x) +
		math.Exp(x) +
		math.Log(x) +
		math.Log10(x)

	var expression algebraic.Expression

	expression.Sum(
		algebraic.Int(value_integer_1),
		algebraic.Multiply(
			algebraic.Float(value_float),
			algebraic.Sum(
				algebraic.Symbol("x"),
				algebraic.Int(value_integer_2),
			),
		),
		algebraic.Sin(
			algebraic.Pow(
				algebraic.Symbol("x"),
				algebraic.Int(value_integer_pow),
			),
		),
		algebraic.Cos(
			algebraic.Symbol("x"),
		),
		algebraic.Tan(
			algebraic.Symbol("x"),
		),
		algebraic.Exp(
			algebraic.Symbol("x"),
		),
		algebraic.Ln(
			algebraic.Symbol("x"),
		),
		algebraic.Log(
			algebraic.Int(value_integer_log),
			algebraic.Symbol("x"),
		),
	)

	var result float64 = expression.Execute(x)

	if result != expected {
		t.Errorf("\nexpected: %.25f\ngot     : %.25f", expected, result)
	}

	t.Logf("finishing testing Execute for Expression\n\n")
}
