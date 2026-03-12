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
	"slices"
	"testing"

	algebraic "github.com/FabioLuisBoni/go-algebra/equation/expression"
)

func TestStringExpressionCaseBareInteger(t *testing.T) {
	t.Logf("testing String for Expression CASE BARE INTEGER\n\n")
	{
		var expected string = "1"

		var result string = algebraic.Int(1).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "-1"

		var result string = algebraic.Int(-1).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	t.Logf("finishing testing String for Expression CASE BARE INTEGER\n\n")
}

func TestStringExpressionCaseBareFloat(t *testing.T) {
	t.Logf("testing String for Expression CASE BARE FLOAT\n\n")
	{
		var expected string = "1.113"

		var result string = algebraic.Float(1.113).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "-1.113"

		var result string = algebraic.Float(-1.113).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "1"

		var result string = algebraic.Float(1.0).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "-1"

		var result string = algebraic.Float(-1.0).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	t.Logf("finishing testing String for Expression CASE BARE FLOAT\n\n")
}

func TestStringExpressionCaseBareVariable(t *testing.T) {
	t.Logf("testing String for Expression CASE BARE VARIABLE\n\n")
	{
		var expected string = "x"

		var result string = algebraic.Symbol("x").String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "pi"

		var result string = algebraic.Symbol("pi").String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	t.Logf("finishing testing String for Expression CASE BARE VARIABLE\n\n")
}

func TestStringExpressionCaseAdditionConstant(t *testing.T) {
	t.Logf("testing String for Expression CASE ADDITION CONSTANT\n\n")
	{
		var expected string = "3"

		var result string = algebraic.Sum(
			algebraic.Int(1),
			algebraic.Int(2),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "-1"

		var result string = algebraic.Sum(
			algebraic.Int(1),
			algebraic.Int(-2),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "2.3"

		var result string = algebraic.Sum(
			algebraic.Float(1.1),
			algebraic.Float(1.2),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "-0.75"

		var result string = algebraic.Sum(
			algebraic.Float(1.25),
			algebraic.Float(-2.0),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "3.1"

		var result string = algebraic.Sum(
			algebraic.Float(1.1),
			algebraic.Int(2),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "8.3" // Nested Sum

		var result string = algebraic.Sum(
			algebraic.Float(1.1),
			algebraic.Int(2),
			algebraic.Sum(
				algebraic.Float(2.2),
				algebraic.Int(3),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "" // Flat sum

		var result string = algebraic.Sum(
			algebraic.Float(1.1),
			algebraic.Int(2),
			algebraic.Float(-1.1),
			algebraic.Int(-2),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "" // Nested Sum

		var result string = algebraic.Sum(
			algebraic.Float(1.1),
			algebraic.Int(2),
			algebraic.Sum(
				algebraic.Float(-1.1),
				algebraic.Int(-2),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	t.Logf("finishing testing String for Expression CASE ADDITION CONSTANT\n\n")
}

func TestStringExpressionCaseAdditionSymbol(t *testing.T) {
	t.Logf("testing String for Expression CASE ADDITION SYMBOL\n\n")
	{
		var expected []string = []string{ // This is array to avoid messed order due to map random sorting
			"3x",
		}

		var result string = algebraic.Sum(
			algebraic.Symbol("x"),
			algebraic.Symbol("x"),
			algebraic.Symbol("x"),
		).String()

		if !slices.Contains(expected, result) {
			t.Errorf("\nexpected: %+v\ngot     : %s", expected, result)
		}
	}

	{
		var expected []string = []string{
			"2e +3x +pi",
			"2e +pi +3x",
			"3x +2e +pi",
			"3x +pi +2e",
			"pi +2e +3x",
			"pi +3x +2e",
		}

		var result string = algebraic.Sum(
			algebraic.Symbol("e"),
			algebraic.Symbol("x"),
			algebraic.Symbol("e"),
			algebraic.Symbol("pi"),
			algebraic.Symbol("x"),
			algebraic.Symbol("x"),
		).String()

		if !slices.Contains(expected, result) {
			t.Errorf("\nexpected: %+v\ngot     : %s", expected, result)
		}
	}

	{
		var expected []string = []string{
			"2e +3x +pi -0.75",
			"2e +pi +3x -0.75",
			"3x +2e +pi -0.75",
			"3x +pi +2e -0.75",
			"pi +2e +3x -0.75",
			"pi +3x +2e -0.75",
		}

		var result string = algebraic.Sum(
			algebraic.Symbol("e"),
			algebraic.Symbol("x"),
			algebraic.Symbol("e"),
			algebraic.Float(1.25),
			algebraic.Int(-2),
			algebraic.Symbol("x"),
			algebraic.Symbol("pi"),
			algebraic.Symbol("x"),
		).String()

		if !slices.Contains(expected, result) {
			t.Errorf("\nexpected: %+v\ngot     : %s", expected, result)
		}
	}

	{
		var expected []string = []string{
			"2e +3x +pi",
			"2e +pi +3x",
			"3x +2e +pi",
			"3x +pi +2e",
			"pi +2e +3x",
			"pi +3x +2e",
		}

		var result string = algebraic.Sum(
			algebraic.Symbol("e"),
			algebraic.Symbol("x"),
			algebraic.Symbol("e"),
			algebraic.Float(-2),
			algebraic.Int(2),
			algebraic.Symbol("pi"),
			algebraic.Symbol("x"),
			algebraic.Symbol("x"),
		).String()

		if !slices.Contains(expected, result) {
			t.Errorf("\nexpected: %+v\ngot     : %s", expected, result)
		}
	}

	t.Logf("finishing testing String for Expression CASE ADDITION SYMBOL\n\n")
}

func TestStringExpressionCaseAdditionZeroSituation(t *testing.T) {
	t.Logf("testing String for Expression CASE ADDITION ZERO SITUATION\n\n")
	{
		var expected string = "2"

		var result string = algebraic.Sum(
			algebraic.Int(2),
			algebraic.Int(0),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "2"

		var result string = algebraic.Sum(
			algebraic.Int(2),
			algebraic.Float(0),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "-2"

		var result string = algebraic.Sum(
			algebraic.Int(-2),
			algebraic.Int(0),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "-2"

		var result string = algebraic.Sum(
			algebraic.Int(-2),
			algebraic.Float(0),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "2.1"

		var result string = algebraic.Sum(
			algebraic.Float(2.1),
			algebraic.Int(0),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "2.1"

		var result string = algebraic.Sum(
			algebraic.Float(2.1),
			algebraic.Float(0),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "-2.1"

		var result string = algebraic.Sum(
			algebraic.Float(-2.1),
			algebraic.Int(0),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "-2.1"

		var result string = algebraic.Sum(
			algebraic.Float(-2.1),
			algebraic.Float(0),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "x"

		var result string = algebraic.Sum(
			algebraic.Symbol("x"),
			algebraic.Int(0),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "x"

		var result string = algebraic.Sum(
			algebraic.Symbol("x"),
			algebraic.Float(0),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "x"

		var result string = algebraic.Sum(
			algebraic.Symbol("x"),
			algebraic.Sin(
				algebraic.Int(0),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "x"

		var result string = algebraic.Sum(
			algebraic.Symbol("x"),
			algebraic.Ln(
				algebraic.Int(1),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "x"

		var result string = algebraic.Sum(
			algebraic.Symbol("x"),
			algebraic.Log(
				algebraic.Int(10),
				algebraic.Int(1),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	t.Logf("finishing testing String for Expression CASE ADDITION ZERO SITUATION\n\n")
}

func TestStringExpressionCaseMultiplyConstant(t *testing.T) {
	t.Logf("testing String for Expression CASE MULTIPLY CONSTANT\n\n")
	{
		var expected string = "6"

		var result string = algebraic.Multiply(
			algebraic.Int(2),
			algebraic.Int(3),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "-6"

		var result string = algebraic.Multiply(
			algebraic.Int(2),
			algebraic.Int(3),
			algebraic.Int(-1),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "8.75"

		var result string = algebraic.Multiply(
			algebraic.Float(2.5),
			algebraic.Float(3.5),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "-8.75"

		var result string = algebraic.Multiply(
			algebraic.Float(2.5),
			algebraic.Float(3.5),
			algebraic.Float(-1),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "-5"

		var result string = algebraic.Multiply(
			algebraic.Float(2.5),
			algebraic.Float(-2),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "5"

		var result string = algebraic.Multiply(
			algebraic.Float(2.5),
			algebraic.Int(-2),
			algebraic.Float(-1),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	t.Logf("finishing testing String for Expression CASE MULTIPLY CONSTANT\n\n")
}

func TestStringExpressionCaseAtomicFamilyMultiplyAggregatedSymbol(t *testing.T) {
	t.Logf("testing String for Expression CASE MULTIPLY SYMBOL\n\n")
	{
		var expected string = "x"

		var result string = algebraic.Multiply(
			algebraic.Symbol("x"),
			algebraic.Int(1),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "-x"

		var result string = algebraic.Multiply(
			algebraic.Symbol("x"),
			algebraic.Int(-1),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "x"

		var result string = algebraic.Multiply(
			algebraic.Symbol("x"),
			algebraic.Float(1),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "-x"

		var result string = algebraic.Multiply(
			algebraic.Symbol("x"),
			algebraic.Float(-1),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "6.2x"

		var result string = algebraic.Multiply(
			algebraic.Symbol("x"),
			algebraic.Float(3.1),
			algebraic.Int(2),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "-6.2x"

		var result string = algebraic.Multiply(
			algebraic.Symbol("x"),
			algebraic.Float(-3.1),
			algebraic.Int(2),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "x * e * x" // this is expected to change to power once its implemented

		var result string = algebraic.Multiply(
			algebraic.Symbol("x"),
			algebraic.Symbol("e"),
			algebraic.Symbol("x"),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "-(x * e * x)" // this is expected to change to power once its implemented

		var result string = algebraic.Multiply(
			algebraic.Symbol("x"),
			algebraic.Symbol("e"),
			algebraic.Multiply(
				algebraic.Symbol("x"),
				algebraic.Int(-1),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "x * e"

		var result string = algebraic.Multiply(
			algebraic.Symbol("x"),
			algebraic.Symbol("e"),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "-(x * e)"

		var result string = algebraic.Multiply(
			algebraic.Symbol("x"),
			algebraic.Multiply(
				algebraic.Symbol("e"),
				algebraic.Int(-1),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "x * e"

		var result string = algebraic.Multiply(
			algebraic.Symbol("x"),
			algebraic.Multiply(
				algebraic.Symbol("e"),
				algebraic.Int(-1),
			),
			algebraic.Int(-1),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "-6.2 * x * e"

		var result string = algebraic.Multiply(
			algebraic.Symbol("x"),
			algebraic.Multiply(
				algebraic.Symbol("e"),
				algebraic.Int(-1),
			),
			algebraic.Int(2),
			algebraic.Float(3.1),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "6.2 * x * e"

		var result string = algebraic.Multiply(
			algebraic.Symbol("x"),
			algebraic.Multiply(
				algebraic.Symbol("e"),
				algebraic.Int(-1),
			),
			algebraic.Int(2),
			algebraic.Float(-3.1),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	t.Logf("finishing testing String for Expression CASE MULTIPLY SYMBOL\n\n")
}

func TestStringExpressionCaseMultiplyZeroSituation(t *testing.T) {
	t.Logf("testing String for Expression CASE MULTIPLY ZERO SITUATION\n\n")
	var expected string = ""

	{
		var result string = algebraic.Multiply(
			algebraic.Int(2),
			algebraic.Int(0),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var result string = algebraic.Multiply(
			algebraic.Int(2),
			algebraic.Float(0),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var result string = algebraic.Multiply(
			algebraic.Int(-2),
			algebraic.Int(0),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var result string = algebraic.Multiply(
			algebraic.Int(-2),
			algebraic.Float(0),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var result string = algebraic.Multiply(
			algebraic.Float(2.1),
			algebraic.Int(0),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var result string = algebraic.Multiply(
			algebraic.Float(2.1),
			algebraic.Float(0),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var result string = algebraic.Multiply(
			algebraic.Float(-2.1),
			algebraic.Int(0),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var result string = algebraic.Multiply(
			algebraic.Float(-2.1),
			algebraic.Float(0),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var result string = algebraic.Multiply(
			algebraic.Symbol("x"),
			algebraic.Int(0),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var result string = algebraic.Multiply(
			algebraic.Symbol("x"),
			algebraic.Float(0),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var result string = algebraic.Multiply(
			algebraic.Symbol("x"),
			algebraic.Multiply(
				algebraic.Float(2.1),
				algebraic.Int(0),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var result string = algebraic.Multiply(
			algebraic.Symbol("x"),
			algebraic.Sin(
				algebraic.Int(0),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var result string = algebraic.Multiply(
			algebraic.Symbol("x"),
			algebraic.Cos(
				algebraic.Multiply(
					algebraic.Symbol("pi"),
					algebraic.Float(0.5),
				),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var result string = algebraic.Multiply(
			algebraic.Symbol("x"),
			algebraic.Tan(
				algebraic.Multiply(
					algebraic.Symbol("pi"),
					algebraic.Float(0),
				),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var result string = algebraic.Multiply(
			algebraic.Symbol("x"),
			algebraic.Ln(
				algebraic.Int(1),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var result string = algebraic.Multiply(
			algebraic.Symbol("x"),
			algebraic.Log(
				algebraic.Int(10),
				algebraic.Int(1),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	t.Logf("finishing testing String for Expression CASE MULTIPLY ZERO SITUATION\n\n")
}

func TestStringExpressionCaseNestedAdditionANdMultiplyZeroSituation(t *testing.T) {
	t.Logf("testing String for Expression CASE NESTED ADDITION AND MULTIPLY ZERO SITUATION\n\n")

	{
		var expected string = "x"

		var result string = algebraic.Sum(
			algebraic.Symbol("x"),
			algebraic.Multiply(
				algebraic.Int(1),
				algebraic.Sin(
					algebraic.Int(0),
				),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = ""

		var result string = algebraic.Multiply(
			algebraic.Symbol("x"),
			algebraic.Sum(
				algebraic.Int(0),
				algebraic.Sin(
					algebraic.Int(0),
				),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	t.Logf("finishing testing String for Expression CASE NESTED ADDITION AND MULTIPLY ZERO SITUATION\n\n")
}

func TestStringExpressionCaseMultiplyWithNestedAddition(t *testing.T) {
	t.Logf("testing String for Expression CASE MULTIPLY WITH NESTED ADDITION\n\n")
	{
		var expected string = "2(x +1)"

		var result string = algebraic.Multiply(
			algebraic.Int(2),
			algebraic.Sum(
				algebraic.Symbol("x"),
				algebraic.Int(1),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "-2(x +1)"

		var result string = algebraic.Multiply(
			algebraic.Int(-2),
			algebraic.Sum(
				algebraic.Symbol("x"),
				algebraic.Int(1),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "x +1"

		var result string = algebraic.Multiply(
			algebraic.Int(1),
			algebraic.Sum(
				algebraic.Symbol("x"),
				algebraic.Int(1),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "-(x +1)"

		var result string = algebraic.Multiply(
			algebraic.Int(-1),
			algebraic.Sum(
				algebraic.Symbol("x"),
				algebraic.Int(1),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "-6(x +1)"

		var result string = algebraic.Multiply(
			algebraic.Int(2),
			algebraic.Sum(
				algebraic.Symbol("x"),
				algebraic.Int(1),
			),
		).Multiply(
			algebraic.Int(-3),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "6(x +1)"

		var result string = algebraic.Multiply(
			algebraic.Int(-2),
			algebraic.Sum(
				algebraic.Symbol("x"),
				algebraic.Int(1),
			),
			algebraic.Int(-3),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "(a -2) * (x +1) * x"

		var result string = algebraic.Multiply(
			algebraic.Sum(
				algebraic.Symbol("a"),
				algebraic.Int(-2),
			),
			algebraic.Sum(
				algebraic.Symbol("x"),
				algebraic.Int(1),
			),
			algebraic.Symbol("x"),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "2 * (x +1) * x"

		var result string = algebraic.Multiply(
			algebraic.Int(2),
			algebraic.Sum(
				algebraic.Symbol("x"),
				algebraic.Int(1),
			),
			algebraic.Symbol("x"),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "2 * (x +1) * (x -1)"

		var result string = algebraic.Multiply(
			algebraic.Int(2),
			algebraic.Sum(
				algebraic.Symbol("x"),
				algebraic.Int(1),
			),
			algebraic.Sum(
				algebraic.Symbol("x"),
				algebraic.Int(-1),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "2 * (x +1) * (x -1) * x"

		var result string = algebraic.Multiply(
			algebraic.Int(2),
			algebraic.Sum(
				algebraic.Symbol("x"),
				algebraic.Int(1),
			),
			algebraic.Sum(
				algebraic.Symbol("x"),
				algebraic.Int(-1),
			),
			algebraic.Symbol("x"),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "(x +1) * (x -1) * x"

		var result string = algebraic.Multiply(
			algebraic.Int(2),
			algebraic.Sum(
				algebraic.Symbol("x"),
				algebraic.Int(1),
			),
			algebraic.Multiply(
				algebraic.Float(0.5),
				algebraic.Sum(
					algebraic.Symbol("x"),
					algebraic.Int(-1),
				),
			),
			algebraic.Symbol("x"),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	t.Logf("finishing testing String for Expression CASE MULTIPLY WITH NESTED ADDITION\n\n")
}

func TestStringExpressionCaseAdditionWithNestedMultiply(t *testing.T) {
	t.Logf("testing String for Expression CASE ADDITION WITH NESTED MULTIPLY\n\n")
	{
		var expected string = "x +2"

		var result string = algebraic.Sum(
			algebraic.Int(2),
			algebraic.Multiply(
				algebraic.Symbol("x"),
				algebraic.Int(1),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "-x +2"

		var result string = algebraic.Sum(
			algebraic.Int(2),
			algebraic.Multiply(
				algebraic.Symbol("x"),
				algebraic.Int(-1),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "2.1x +2"

		var result string = algebraic.Sum(
			algebraic.Int(2),
			algebraic.Multiply(
				algebraic.Symbol("x"),
				algebraic.Float(2.1),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "-2.1x +2"

		var result string = algebraic.Sum(
			algebraic.Int(2),
			algebraic.Multiply(
				algebraic.Symbol("x"),
				algebraic.Float(-2.1),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "(-2.1 * x * e) +x +2"

		var result string = algebraic.Sum(
			algebraic.Int(2),
			algebraic.Symbol("x"),
			algebraic.Multiply(
				algebraic.Symbol("x"),
				algebraic.Float(-2.1),
				algebraic.Symbol("e"),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	t.Logf("finishing testing String for Expression CASE ADDITION WITH NESTED MULTIPLY\n\n")
}

// func TestStringExpressionCaseAdditionAndMultiplyWithComplexNestedSituation(t *testing.T) {
// 	t.Logf("testing String for Expression CASE ADDITION AND MULTIPLY WITH COMPLEX NESTED SITUATION\n\n")
// 	{
// 		var expected string = "(-2.1 * x * (x +e +1)) +x +2"

// 		var result string = algebraic.Sum(
// 			algebraic.Int(2),
// 			algebraic.Symbol("x"),
// 			algebraic.Multiply(
// 				algebraic.Symbol("x"),
// 				algebraic.Float(-2.1),
// 				algebraic.Sum(
// 					algebraic.Int(1),
// 					algebraic.Symbol("x"),
// 					algebraic.Symbol("e"),
// 				),
// 			),
// 		).String()

// 		if result != expected {
// 			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
// 		}
// 	}

// 	{
// 		var expected string = "x +2"

// 		var result string = algebraic.Sum(
// 			algebraic.Int(2),
// 			algebraic.Symbol("x"),
// 			algebraic.Multiply(
// 				algebraic.Symbol("x"),
// 				algebraic.Cos(
// 					algebraic.Multiply(
// 						algebraic.Float(0.5),
// 						algebraic.Symbol("pi"),
// 					),
// 				),
// 				algebraic.Sum(
// 					algebraic.Int(1),
// 					algebraic.Symbol("x"),
// 					algebraic.Symbol("e"),
// 				),
// 			),
// 		).String()

// 		if result != expected {
// 			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
// 		}
// 	}

// 	t.Logf("finishing testing String for Expression CASE ADDITION AND MULTIPLY WITH COMPLEX NESTED SITUATION\n\n")
// }

func TestStringExpressionCaseBarePower(t *testing.T) {
	t.Logf("testing String for Expression CASE BARE POWER\n\n")
	{
		var expected string = "x^y"

		var result string = algebraic.Pow(
			algebraic.Symbol("x"),
			algebraic.Symbol("y"),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "x"

		var result string = algebraic.Pow(
			algebraic.Symbol("x"),
			algebraic.Int(1),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "x^2"

		var result string = algebraic.Pow(
			algebraic.Symbol("x"),
			algebraic.Int(2),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "x^2.1"

		var result string = algebraic.Pow(
			algebraic.Symbol("x"),
			algebraic.Float(2.1),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "1/x"

		var result string = algebraic.Pow(
			algebraic.Symbol("x"),
			algebraic.Int(-1),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "1/(x^3)"

		var result string = algebraic.Pow(
			algebraic.Symbol("x"),
			algebraic.Int(-3),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "1/(x^3.3)"

		var result string = algebraic.Pow(
			algebraic.Symbol("x"),
			algebraic.Float(-3.3),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "0^0"

		var result string = algebraic.Pow(
			algebraic.Cos(
				algebraic.Symbol("pi").Divide(
					algebraic.Int(2),
				),
			),
			algebraic.Multiply(
				algebraic.Int(0),
				algebraic.Symbol("x"),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "1"

		var result string = algebraic.Pow(
			algebraic.Cos(
				algebraic.Symbol("x"),
			),
			algebraic.Multiply(
				algebraic.Int(0),
				algebraic.Symbol("x"),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = ""

		var result string = algebraic.Pow(
			algebraic.Cos(
				algebraic.Symbol("pi").Divide(
					algebraic.Int(2),
				),
			),
			algebraic.Symbol("x"),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	t.Logf("finishing testing String for Expression CASE BARE POWER\n\n")
}

func TestStringExpressionCasePowerWithNestedAdditionAndMultiplyArguments(t *testing.T) {
	t.Logf("testing String for Expression CASE POWER WITH NESTED ADDITION AND MULTIPLY ARGUMENTS\n\n")
	{
		var expected string = "(-x)^y"

		var result string = algebraic.Pow(
			algebraic.Multiply(
				algebraic.Symbol("x"),
				algebraic.Int(-1),
			),
			algebraic.Symbol("y"),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "1/((-x)^y)"

		var result string = algebraic.Pow(
			algebraic.Multiply(
				algebraic.Symbol("x"),
				algebraic.Int(-1),
			),
			algebraic.Multiply(
				algebraic.Symbol("y"),
				algebraic.Int(-1),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "(x +1)^2"

		var result string = algebraic.Pow(
			algebraic.Sum(
				algebraic.Symbol("x"),
				algebraic.Int(1),
			),
			algebraic.Int(2),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "(-(x +1))^2"

		var result string = algebraic.Pow(
			algebraic.Multiply(
				algebraic.Int(-1),
				algebraic.Sum(
					algebraic.Symbol("x"),
					algebraic.Int(1),
				),
			),
			algebraic.Int(2),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "1/((x +1)^2)"

		var result string = algebraic.Pow(
			algebraic.Sum(
				algebraic.Symbol("x"),
				algebraic.Int(1),
			),
			algebraic.Int(-2),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "1/((-(x +1))^2)"

		var result string = algebraic.Pow(
			algebraic.Multiply(
				algebraic.Int(-1),
				algebraic.Sum(
					algebraic.Symbol("x"),
					algebraic.Int(1),
				),
			),
			algebraic.Int(-2),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "(x +1)^(x +2)"

		var result string = algebraic.Pow(
			algebraic.Sum(
				algebraic.Symbol("x"),
				algebraic.Int(1),
			),
			algebraic.Sum(
				algebraic.Symbol("x"),
				algebraic.Int(2),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "1/((x +1)^(x +2))"

		var result string = algebraic.Pow(
			algebraic.Sum(
				algebraic.Symbol("x"),
				algebraic.Int(1),
			),
			algebraic.Multiply(
				algebraic.Sum(
					algebraic.Symbol("x"),
					algebraic.Int(2),
				),
				algebraic.Int(-1),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "1/((-(x +1))^(x +2))"

		var result string = algebraic.Pow(
			algebraic.Multiply(
				algebraic.Int(-1),
				algebraic.Sum(
					algebraic.Symbol("x"),
					algebraic.Int(1),
				),
			),
			algebraic.Multiply(
				algebraic.Sum(
					algebraic.Symbol("x"),
					algebraic.Int(2),
				),
				algebraic.Int(-1),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "(2x)^2"

		var result string = algebraic.Pow(
			algebraic.Multiply(
				algebraic.Symbol("x"),
				algebraic.Int(2),
			),
			algebraic.Int(2),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "(-2x)^2"

		var result string = algebraic.Pow(
			algebraic.Multiply(
				algebraic.Symbol("x"),
				algebraic.Int(-2),
			),
			algebraic.Int(2),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "1/((2x)^2)"

		var result string = algebraic.Pow(
			algebraic.Multiply(
				algebraic.Symbol("x"),
				algebraic.Int(2),
			),
			algebraic.Int(-2),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "1/((-2x)^2)"

		var result string = algebraic.Pow(
			algebraic.Multiply(
				algebraic.Symbol("x"),
				algebraic.Int(-2),
			),
			algebraic.Int(-2),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "(2x)^(4x)"

		var result string = algebraic.Pow(
			algebraic.Multiply(
				algebraic.Symbol("x"),
				algebraic.Int(2),
			),
			algebraic.Multiply(
				algebraic.Symbol("x"),
				algebraic.Int(4),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "1/((2x)^(4x))"

		var result string = algebraic.Pow(
			algebraic.Multiply(
				algebraic.Symbol("x"),
				algebraic.Int(2),
			),
			algebraic.Multiply(
				algebraic.Symbol("x"),
				algebraic.Int(-4),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "1/((-2x)^(4x))"

		var result string = algebraic.Pow(
			algebraic.Multiply(
				algebraic.Symbol("x"),
				algebraic.Int(-2),
			),
			algebraic.Multiply(
				algebraic.Symbol("x"),
				algebraic.Int(-4),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "(2 * x * e)^2"

		var result string = algebraic.Pow(
			algebraic.Multiply(
				algebraic.Symbol("x"),
				algebraic.Int(2),
				algebraic.Symbol("e"),
			),
			algebraic.Int(2),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "(-2 * x * e)^2"

		var result string = algebraic.Pow(
			algebraic.Multiply(
				algebraic.Symbol("x"),
				algebraic.Int(-2),
				algebraic.Symbol("e"),
			),
			algebraic.Int(2),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "1/((2 * x * e)^2)"

		var result string = algebraic.Pow(
			algebraic.Multiply(
				algebraic.Symbol("x"),
				algebraic.Int(2),
				algebraic.Symbol("e"),
			),
			algebraic.Int(-2),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "1/((-2 * x * e)^2)"

		var result string = algebraic.Pow(
			algebraic.Multiply(
				algebraic.Symbol("x"),
				algebraic.Int(-2),
				algebraic.Symbol("e"),
			),
			algebraic.Int(-2),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "(2 * x * e)^(4 * y * e)"

		var result string = algebraic.Pow(
			algebraic.Multiply(
				algebraic.Symbol("x"),
				algebraic.Int(2),
				algebraic.Symbol("e"),
			),
			algebraic.Multiply(
				algebraic.Symbol("y"),
				algebraic.Int(4),
				algebraic.Symbol("e"),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "(-2 * x * e)^(4 * y * e)"

		var result string = algebraic.Pow(
			algebraic.Multiply(
				algebraic.Symbol("x"),
				algebraic.Int(-2),
				algebraic.Symbol("e"),
			),
			algebraic.Multiply(
				algebraic.Symbol("y"),
				algebraic.Int(4),
				algebraic.Symbol("e"),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "1/((2 * x * e)^(4 * y * e))"

		var result string = algebraic.Pow(
			algebraic.Multiply(
				algebraic.Symbol("x"),
				algebraic.Int(2),
				algebraic.Symbol("e"),
			),
			algebraic.Multiply(
				algebraic.Symbol("y"),
				algebraic.Int(-4),
				algebraic.Symbol("e"),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	{
		var expected string = "1/((-2 * x * e)^(4 * y * e))"

		var result string = algebraic.Pow(
			algebraic.Multiply(
				algebraic.Symbol("x"),
				algebraic.Int(-2),
				algebraic.Symbol("e"),
			),
			algebraic.Multiply(
				algebraic.Symbol("y"),
				algebraic.Int(-4),
				algebraic.Symbol("e"),
			),
		).String()

		if result != expected {
			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
		}
	}

	t.Logf("finishing testing String for Expression CASE POWER WITH NESTED ADDITION AND MULTIPLY ARGUMENTS\n\n")
}

// func TestStringExpressionCasePowerHandlingParenthesisWhenNestedZero(t *testing.T) {
// 	t.Logf("testing String for Expression CASE POWER HANDLING PARENTHESIS WHEN NESTED ZERO\n\n")

// 	{
// 		var expected string = "x^2"

// 		var result string = algebraic.Pow(
// 			algebraic.Multiply(
// 				algebraic.Symbol("x"),
// 				algebraic.Int(1),
// 			),
// 			algebraic.Int(2),
// 		).String()

// 		if result != expected {
// 			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
// 		}
// 	}

// 	{
// 		var expected string = "(-x)^2"

// 		var result string = algebraic.Pow(
// 			algebraic.Multiply(
// 				algebraic.Symbol("x"),
// 				algebraic.Int(-1),
// 			),
// 			algebraic.Int(2),
// 		).String()

// 		if result != expected {
// 			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
// 		}
// 	}

// 	{
// 		var expected string = "1/(x^2)"

// 		var result string = algebraic.Pow(
// 			algebraic.Multiply(
// 				algebraic.Symbol("x"),
// 				algebraic.Int(1),
// 			),
// 			algebraic.Int(-2),
// 		).String()

// 		if result != expected {
// 			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
// 		}
// 	}

// 	{
// 		var expected string = "1/((-x)^2)"

// 		var result string = algebraic.Pow(
// 			algebraic.Multiply(
// 				algebraic.Symbol("x"),
// 				algebraic.Int(-1),
// 			),
// 			algebraic.Int(-2),
// 		).String()

// 		if result != expected {
// 			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
// 		}
// 	}

// 	{
// 		var expected string = "x^y"

// 		var result string = algebraic.Pow(
// 			algebraic.Multiply(
// 				algebraic.Symbol("x"),
// 				algebraic.Int(1),
// 			),
// 			algebraic.Multiply(
// 				algebraic.Symbol("y"),
// 				algebraic.Int(1),
// 			),
// 		).String()

// 		if result != expected {
// 			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
// 		}
// 	}

// 	{
// 		var expected string = "(-x)^y"

// 		var result string = algebraic.Pow(
// 			algebraic.Multiply(
// 				algebraic.Symbol("x"),
// 				algebraic.Int(-1),
// 			),
// 			algebraic.Multiply(
// 				algebraic.Symbol("y"),
// 				algebraic.Int(1),
// 			),
// 		).String()

// 		if result != expected {
// 			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
// 		}
// 	}

// 	{
// 		var expected string = "1/(x^y)"

// 		var result string = algebraic.Pow(
// 			algebraic.Multiply(
// 				algebraic.Symbol("x"),
// 				algebraic.Int(1),
// 			),
// 			algebraic.Multiply(
// 				algebraic.Symbol("y"),
// 				algebraic.Int(-1),
// 			),
// 		).String()

// 		if result != expected {
// 			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
// 		}
// 	}

// 	{
// 		var expected string = "1/((-x)^y)"

// 		var result string = algebraic.Pow(
// 			algebraic.Multiply(
// 				algebraic.Symbol("x"),
// 				algebraic.Int(-1),
// 			),
// 			algebraic.Multiply(
// 				algebraic.Symbol("y"),
// 				algebraic.Int(-1),
// 			),
// 		).String()

// 		if result != expected {
// 			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
// 		}
// 	}

// 	{
// 		var expected string = "x^2"

// 		var result string = algebraic.Pow(
// 			algebraic.Sum(
// 				algebraic.Symbol("x"),
// 				algebraic.Int(0),
// 			),
// 			algebraic.Int(2),
// 		).String()

// 		if result != expected {
// 			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
// 		}
// 	}

// 	{
// 		var expected string = "x^2"

// 		var result string = algebraic.Pow(
// 			algebraic.Symbol("x"),
// 			algebraic.Sum(
// 				algebraic.Int(2),
// 				algebraic.Int(0),
// 			),
// 		).String()

// 		if result != expected {
// 			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
// 		}
// 	}

// 	{
// 		var expected string = "1/(x^2)"

// 		var result string = algebraic.Pow(
// 			algebraic.Sum(
// 				algebraic.Symbol("x"),
// 				algebraic.Int(0),
// 			),
// 			algebraic.Int(-2),
// 		).String()

// 		if result != expected {
// 			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
// 		}
// 	}

// 	{
// 		var expected string = "1/(x^2)"

// 		var result string = algebraic.Pow(
// 			algebraic.Symbol("x"),
// 			algebraic.Sum(
// 				algebraic.Int(-2),
// 				algebraic.Int(0),
// 			),
// 		).String()

// 		if result != expected {
// 			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
// 		}
// 	}

// 	{
// 		var expected string = "x^y"

// 		var result string = algebraic.Pow(
// 			algebraic.Sum(
// 				algebraic.Symbol("x"),
// 				algebraic.Int(0),
// 			),
// 			algebraic.Sum(
// 				algebraic.Symbol("y"),
// 				algebraic.Int(0),
// 			),
// 		).String()

// 		if result != expected {
// 			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
// 		}
// 	}

// 	{
// 		var expected string = "1/(x^y)"

// 		var result string = algebraic.Pow(
// 			algebraic.Sum(
// 				algebraic.Symbol("x"),
// 				algebraic.Int(0),
// 			),
// 			algebraic.Sum(
// 				algebraic.Multiply(
// 					algebraic.Symbol("y"),
// 					algebraic.Int(-1),
// 				),
// 				algebraic.Int(0),
// 			),
// 		).String()

// 		if result != expected {
// 			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
// 		}
// 	}

// 	{
// 		var expected string = "2^x"

// 		var result string = algebraic.Pow(
// 			algebraic.Int(2),
// 			algebraic.Sum(
// 				algebraic.Symbol("x"),
// 				algebraic.Int(0),
// 			),
// 		).String()

// 		if result != expected {
// 			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
// 		}
// 	}

// 	{
// 		var expected string = "2^x"

// 		var result string = algebraic.Pow(
// 			algebraic.Sum(
// 				algebraic.Int(2),
// 				algebraic.Int(0),
// 			),
// 			algebraic.Symbol("x"),
// 		).String()

// 		if result != expected {
// 			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
// 		}
// 	}

// 	{
// 		var expected string = "1/(2^x)"

// 		var result string = algebraic.Pow(
// 			algebraic.Int(2),
// 			algebraic.Sum(
// 				algebraic.Multiply(
// 					algebraic.Symbol("x"),
// 					algebraic.Int(-1),
// 				),
// 				algebraic.Int(0),
// 			),
// 		).String()

// 		if result != expected {
// 			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
// 		}
// 	}

// 	{
// 		var expected string = "1/(2^x)"

// 		var result string = algebraic.Pow(
// 			algebraic.Sum(
// 				algebraic.Int(2),
// 				algebraic.Int(0),
// 			),
// 			algebraic.Multiply(
// 				algebraic.Symbol("x"),
// 				algebraic.Int(-1),
// 			),
// 		).String()

// 		if result != expected {
// 			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
// 		}
// 	}

// 	{
// 		var expected string = "(-2)^x"

// 		var result string = algebraic.Pow(
// 			algebraic.Int(-2),
// 			algebraic.Sum(
// 				algebraic.Symbol("x"),
// 				algebraic.Int(0),
// 			),
// 		).String()

// 		if result != expected {
// 			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
// 		}
// 	}

// 	{
// 		var expected string = "(-2)^x"

// 		var result string = algebraic.Pow(
// 			algebraic.Sum(
// 				algebraic.Int(-2),
// 				algebraic.Int(0),
// 			),
// 			algebraic.Symbol("x"),
// 		).String()

// 		if result != expected {
// 			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
// 		}
// 	}

// 	{
// 		var expected string = "1/((-2)^x)"

// 		var result string = algebraic.Pow(
// 			algebraic.Int(-2),
// 			algebraic.Sum(
// 				algebraic.Multiply(
// 					algebraic.Symbol("x"),
// 					algebraic.Int(-1),
// 				),
// 				algebraic.Int(0),
// 			),
// 		).String()

// 		if result != expected {
// 			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
// 		}
// 	}

// 	{
// 		var expected string = "1/((-2)^x)"

// 		var result string = algebraic.Pow(
// 			algebraic.Sum(
// 				algebraic.Int(-2),
// 				algebraic.Int(0),
// 			),
// 			algebraic.Multiply(
// 				algebraic.Symbol("x"),
// 				algebraic.Int(-1),
// 			),
// 		).String()

// 		if result != expected {
// 			t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
// 		}
// 	}

// 	t.Logf("finishing testing String for Expression CASE POWER HANDLING PARENTHESIS WHEN NESTED ZERO\n\n")
// }
