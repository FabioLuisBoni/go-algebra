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
	"math/bits"
	"testing"

	algebraic "github.com/FabioLuisBoni/go-algebra/equation/expression"
)

func TestIsConstantExpression(t *testing.T) {
	t.Logf("testing IsConstant for Expression\n\n")
	var key uint64 = uint64(algebraic.CACHE_IS_CONSTANT)
	var mask uint64 = 0x3 << key

	{ // TRUE CASES
		t.Logf("testing IsConstant for Expression CASE TRUE\n\n")
		var expectedStatus uint64 = (algebraic.CACHE_MASK_RAN | algebraic.CACHE_MASK_RESULT) << key

		{ // LEAF CASES
			{
				var expression *algebraic.Expression = algebraic.Int(1)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Float(1.1)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Symbol("e")

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Symbol("pi")

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // SUM CASES
			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Int(1),
					algebraic.Int(1),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Float(1.1),
					algebraic.Float(1.1),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Float(1.1),
					algebraic.Int(1),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // MULTIPLY CASES
			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Symbol("x"),
					algebraic.Int(0),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Symbol("x"),
					algebraic.Float(0),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Int(1),
					algebraic.Int(1),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Float(1.1),
					algebraic.Float(1.1),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Float(1.1),
					algebraic.Int(1),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // POWER CASES
			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Int(0),
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Symbol("x"),
					algebraic.Int(0),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Float(0),
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Symbol("x"),
					algebraic.Float(0),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Symbol("e"),
					algebraic.Symbol("e"),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Symbol("pi"),
					algebraic.Symbol("pi"),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Symbol("e"),
					algebraic.Symbol("pi"),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Symbol("pi"),
					algebraic.Symbol("e"),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Int(2),
					algebraic.Int(2),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Float(2.1),
					algebraic.Float(2.1),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Int(2),
					algebraic.Float(2.1),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Float(2.1),
					algebraic.Int(2),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Int(1),
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Float(1),
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // FUNCTION CASES
			{
				var expression *algebraic.Expression = algebraic.Sin(
					algebraic.Int(2),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sin(
					algebraic.Float(2.1),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sin(
					algebraic.Symbol("pi"),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Cos(
					algebraic.Int(2),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Cos(
					algebraic.Float(2.1),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Cos(
					algebraic.Symbol("pi"),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Tan(
					algebraic.Int(2),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Tan(
					algebraic.Float(2.1),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Tan(
					algebraic.Symbol("pi"),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Log(
					algebraic.Int(10),
					algebraic.Int(1),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Ln(
					algebraic.Int(1),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Ln(
					algebraic.Symbol("e"),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Exp(
					algebraic.Symbol("e"),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Exp(
					algebraic.Symbol("pi"),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Exp(
					algebraic.Int(2),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Exp(
					algebraic.Float(2.1),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // NESTED CASES
			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Pow(
						algebraic.Symbol("x"),
						algebraic.Int(0),
					),
					algebraic.Cos(
						algebraic.Symbol("pi").Divide(
							algebraic.Int(2),
						),
					),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Pow(
						algebraic.Symbol("x"),
						algebraic.Int(0),
					),
					algebraic.Cos(
						algebraic.Symbol("pi").Divide(
							algebraic.Int(2),
						),
					),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Pow(
						algebraic.Symbol("x"),
						algebraic.Int(0),
					),
					algebraic.Int(2),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Int(1),
					algebraic.Cos(
						algebraic.Symbol("pi").Divide(
							algebraic.Int(2),
						),
					),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Pow(
						algebraic.Symbol("x"),
						algebraic.Int(0),
					),
					algebraic.Cos(
						algebraic.Symbol("pi").Divide(
							algebraic.Int(2),
						),
					),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Cos(
						algebraic.Symbol("pi").Divide(
							algebraic.Int(2),
						),
					),
					algebraic.Pow(
						algebraic.Symbol("x"),
						algebraic.Int(0),
					),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Symbol("x"),
					algebraic.Cos(
						algebraic.Symbol("pi").Divide(
							algebraic.Int(2),
						),
					),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Symbol("x"),
					algebraic.Cos(
						algebraic.Symbol("pi").Divide(
							algebraic.Pow(
								algebraic.Int(2),
								algebraic.Int(1),
							),
						),
					),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sin(
					algebraic.Symbol("pi").Divide(
						algebraic.Int(2),
					),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Cos(
					algebraic.Symbol("pi").Divide(
						algebraic.Int(2),
					),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Tan(
					algebraic.Symbol("pi").Divide(
						algebraic.Int(2),
					),
				)

				var result bool = expression.IsConstant()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		t.Logf("finishing testing IsConstant for Expression CASE TRUE\n\n")
	}

	{ // FALSE CASES
		t.Logf("testing IsConstant for Expression CASE FALSE\n\n")
		var expectedStatus uint64 = algebraic.CACHE_MASK_RAN << key

		{ // LEAF CASES
			{
				var expression *algebraic.Expression = algebraic.Symbol("x")

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // SUM CASES
			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Symbol("x"),
					algebraic.Int(1),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Symbol("x"),
					algebraic.Float(1.1),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Symbol("x"),
					algebraic.Symbol("e"),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Symbol("x"),
					algebraic.Symbol("pi"),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // MULTIPLY CASES
			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Symbol("x"),
					algebraic.Int(1),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Symbol("x"),
					algebraic.Float(1.1),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Symbol("x"),
					algebraic.Symbol("e"),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Symbol("x"),
					algebraic.Symbol("pi"),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // POWER CASES
			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Int(0),
					algebraic.Int(0),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Float(0),
					algebraic.Int(0),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Int(0),
					algebraic.Float(0),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Float(0),
					algebraic.Float(0),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Symbol("x"),
					algebraic.Float(1.1),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Symbol("x"),
					algebraic.Int(1),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Float(1.1),
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Float(2),
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // FUNCTION CASES
			{
				var expression *algebraic.Expression = algebraic.Sin(
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Cos(
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Tan(
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Log(
					algebraic.Int(10),
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Ln(
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Exp(
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // NESTED CASES
			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Symbol("x"),
					algebraic.Cos(
						algebraic.Int(2),
					),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Symbol("x"),
					algebraic.Cos(
						algebraic.Int(2),
					),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Symbol("x"),
					algebraic.Cos(
						algebraic.Int(2),
					),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sin(
					algebraic.Sin(
						algebraic.Symbol("x"),
					),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Cos(
					algebraic.Cos(
						algebraic.Symbol("x"),
					),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Tan(
					algebraic.Tan(
						algebraic.Symbol("x"),
					),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Log(
					algebraic.Int(10),
					algebraic.Log(
						algebraic.Int(10),
						algebraic.Symbol("x"),
					),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Ln(
					algebraic.Ln(
						algebraic.Symbol("x"),
					),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Exp(
					algebraic.Exp(
						algebraic.Symbol("x"),
					),
				)

				var result bool = expression.IsConstant()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		t.Logf("finishing testing IsConstant for Expression CASE FALSE\n\n")
	}

	t.Logf("finishing testing IsConstant for Expression\n\n")
}

func TestIsSignalInvertibleExpression(t *testing.T) {
	t.Logf("testing IsSignalInvertible for Expression\n\n")
	var key uint64 = uint64(algebraic.CACHE_IS_SIGNAL_INVERTIBLE)
	var mask uint64 = 0x3 << key

	{ // TRUE CASES
		var expectedStatus uint64 = (algebraic.CACHE_MASK_RAN | algebraic.CACHE_MASK_RESULT) << key

		t.Logf("testing IsSignalInvertible for Expression CASE TRUE\n\n")
		{ // LEAF CASES
			{
				var expression *algebraic.Expression = algebraic.Int(-1)

				var result bool = expression.IsSignalInvertible()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Float(-1.1)

				var result bool = expression.IsSignalInvertible()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // SUM CASES
			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Int(1),
					algebraic.Int(-2),
				)

				var result bool = expression.IsSignalInvertible()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Int(1),
					algebraic.Float(-2.1),
				)

				var result bool = expression.IsSignalInvertible()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Int(-100),
					algebraic.Symbol("e"),
				)

				var result bool = expression.IsSignalInvertible()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Int(-100),
					algebraic.Symbol("pi"),
				)

				var result bool = expression.IsSignalInvertible()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // MULTIPLY CASES
			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Int(-2),
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsSignalInvertible()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Float(-2.1),
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsSignalInvertible()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Symbol("x"),
					algebraic.Float(-2.1),
					algebraic.Float(-3.1),
					algebraic.Int(-4),
				)

				var result bool = expression.IsSignalInvertible()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // POWER CASES
			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Int(-1),
					algebraic.Int(3),
				)

				var result bool = expression.IsSignalInvertible()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Float(-1),
					algebraic.Int(3),
				)

				var result bool = expression.IsSignalInvertible()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Int(-1),
					algebraic.Float(3),
				)

				var result bool = expression.IsSignalInvertible()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Float(-1),
					algebraic.Float(3),
				)

				var result bool = expression.IsSignalInvertible()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // NESTED CASES
			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Multiply(
						algebraic.Int(-2),
						algebraic.Float(3),
					),
					algebraic.Pow(
						algebraic.Int(2),
						algebraic.Float(2.1),
					),
				)

				var result bool = expression.IsSignalInvertible()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Pow(
						algebraic.Int(-2),
						algebraic.Float(3),
					),
					algebraic.Sum(
						algebraic.Int(2),
						algebraic.Float(1),
					),
				)

				var result bool = expression.IsSignalInvertible()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Multiply(
						algebraic.Int(-2),
						algebraic.Float(3),
					),
					algebraic.Sum(
						algebraic.Int(2),
						algebraic.Float(1),
					),
				)

				var result bool = expression.IsSignalInvertible()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		t.Logf("finishing testing IsSignalInvertible for Expression CASE TRUE\n\n")
	}

	{ // FALSE CASES
		var expectedStatus uint64 = algebraic.CACHE_MASK_RAN << key

		t.Logf("testing IsSignalInvertible for Expression CASE FALSE\n\n")
		{ // LEAF CASES
			{
				var expression *algebraic.Expression = algebraic.Int(1)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Float(1.1)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Symbol("x")

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Symbol("e")

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Symbol("pi")

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // SUM CASES
			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Int(-1),
					algebraic.Int(2),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Int(-1),
					algebraic.Float(2.1),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Int(-100),
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Int(-1),
					algebraic.Symbol("e"),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Int(-1),
					algebraic.Symbol("pi"),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // MULTIPLY CASES
			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Int(2),
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Float(2.1),
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Symbol("x"),
					algebraic.Float(-2.1),
					algebraic.Float(3.1),
					algebraic.Int(-4),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // POWER CASES
			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Int(-1),
					algebraic.Int(2),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Int(-1),
					algebraic.Int(-2),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Float(-1),
					algebraic.Int(2),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Float(-1),
					algebraic.Int(-2),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Int(-1),
					algebraic.Float(2),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Int(-1),
					algebraic.Float(-2),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Float(-1),
					algebraic.Float(2),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Float(-1),
					algebraic.Float(-2),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Symbol("x"),
					algebraic.Int(2),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Symbol("x"),
					algebraic.Float(2),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Symbol("x"),
					algebraic.Float(-2),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Symbol("x"),
					algebraic.Int(-2),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Symbol("x"),
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // FUNCTION CASES
			{
				var expression *algebraic.Expression = algebraic.Sin(
					algebraic.Int(-1),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sin(
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Cos(
					algebraic.Int(-1),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Cos(
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Tan(
					algebraic.Int(-1),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Tan(
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Log(
					algebraic.Int(10),
					algebraic.Int(-1),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Log(
					algebraic.Int(10),
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Ln(
					algebraic.Int(-1),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Ln(
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Exp(
					algebraic.Int(-1),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Exp(
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // NESTED CASES
			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Multiply(
						algebraic.Int(-2),
						algebraic.Float(3),
					),
					algebraic.Pow(
						algebraic.Int(2),
						algebraic.Float(6.1),
					),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Sum(
						algebraic.Int(-3),
						algebraic.Float(1),
					),
					algebraic.Pow(
						algebraic.Int(-2),
						algebraic.Float(3),
					),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Multiply(
						algebraic.Int(-2),
						algebraic.Float(3),
					),
					algebraic.Sum(
						algebraic.Int(2),
						algebraic.Float(2),
					),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Multiply(
						algebraic.Int(-2),
						algebraic.Symbol("x"),
					),
					algebraic.Sum(
						algebraic.Int(2),
						algebraic.Float(2),
					),
				)

				var result bool = expression.IsSignalInvertible()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		t.Logf("finishing testing IsSignalInvertible for Expression CASE FALSE\n\n")
	}

	t.Logf("finishing testing IsSignalInvertible for Expression\n\n")
}

func TestIsNegativeExpression(t *testing.T) {
	t.Logf("testing IsNegative for Expression\n\n")
	var key uint64 = uint64(algebraic.IS_NEGATIVE)
	var mask uint64 = 0x7 << key

	{ // TRUE CASES
		t.Logf("testing IsNegative for Expression CASE TRUE\n\n")
		var expectedStatus uint64 = (algebraic.CACHE_MASK_RAN | algebraic.CACHE_MASK_RESULT | algebraic.CACHE_MASK_APPLICABLE) << key

		{ // LEAF CASES
			{
				var expression *algebraic.Expression = algebraic.Int(-1)

				result, applicable := expression.IsNegative()

				if result != true {
					t.Errorf("\nresult expected: %t\nresult got     : %t", true, result)
				}
				if applicable != true {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", true, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Float(-1.1)

				result, applicable := expression.IsNegative()

				if result != true {
					t.Errorf("\nresult expected: %t\nresult got     : %t", true, result)
				}
				if applicable != true {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", true, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // NESTED CASES
			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Symbol("e"),
					algebraic.Int(-1),
				)

				result, applicable := expression.IsNegative()

				if result != true {
					t.Errorf("\nresult expected: %t\nresult got     : %t", true, result)
				}
				if applicable != true {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", true, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Symbol("pi"),
					algebraic.Int(-1),
				)

				result, applicable := expression.IsNegative()

				if result != true {
					t.Errorf("\nresult expected: %t\nresult got     : %t", true, result)
				}
				if applicable != true {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", true, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Multiply(
						algebraic.Pow(
							algebraic.Symbol("pi"),
							algebraic.Int(2),
						),
						algebraic.Int(-1),
					),
					algebraic.Int(2),
				)

				result, applicable := expression.IsNegative()

				if result != true {
					t.Errorf("\nresult expected: %t\nresult got     : %t", true, result)
				}
				if applicable != true {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", true, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Pow(
						algebraic.Symbol("pi"),
						algebraic.Int(2),
					),
					algebraic.Int(-1),
				)

				result, applicable := expression.IsNegative()

				if result != true {
					t.Errorf("\nresult expected: %t\nresult got     : %t", true, result)
				}
				if applicable != true {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", true, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Pow(
						algebraic.Int(-2),
						algebraic.Int(3),
					),
					algebraic.Int(3),
				)

				result, applicable := expression.IsNegative()

				if result != true {
					t.Errorf("\nresult expected: %t\nresult got     : %t", true, result)
				}
				if applicable != true {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", true, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sin(
					algebraic.Pow(
						algebraic.Symbol("pi"),
						algebraic.Int(2),
					),
				)

				result, applicable := expression.IsNegative()

				if result != true {
					t.Errorf("\nresult expected: %t\nresult got     : %t", true, result)
				}
				if applicable != true {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", true, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Cos(
					algebraic.Pow(
						algebraic.Symbol("pi"),
						algebraic.Int(2),
					),
				)

				result, applicable := expression.IsNegative()

				if result != true {
					t.Errorf("\nresult expected: %t\nresult got     : %t", true, result)
				}
				if applicable != true {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", true, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Tan(
					algebraic.Multiply(
						algebraic.Int(-1),
						algebraic.Pow(
							algebraic.Symbol("pi"),
							algebraic.Int(2),
						),
					),
				)

				result, applicable := expression.IsNegative()

				if result != true {
					t.Errorf("\nresult expected: %t\nresult got     : %t", true, result)
				}
				if applicable != true {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", true, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Log(
					algebraic.Int(10),
					algebraic.Multiply(
						algebraic.Float(0.000001),
						algebraic.Pow(
							algebraic.Symbol("pi"),
							algebraic.Int(2),
						),
					),
				)

				result, applicable := expression.IsNegative()

				if result != true {
					t.Errorf("\nresult expected: %t\nresult got     : %t", true, result)
				}
				if applicable != true {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", true, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Ln(
					algebraic.Multiply(
						algebraic.Float(0.000001),
						algebraic.Pow(
							algebraic.Symbol("pi"),
							algebraic.Int(2),
						),
					),
				)

				result, applicable := expression.IsNegative()

				if result != true {
					t.Errorf("\nresult expected: %t\nresult got     : %t", true, result)
				}
				if applicable != true {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", true, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		t.Logf("finishing testing IsNegative for Expression CASE TRUE\n\n")
	}

	{ // FALSE CASES
		t.Logf("testing IsNegative for Expression CASE FALSE\n\n")
		var expectedStatus uint64 = (algebraic.CACHE_MASK_RAN | algebraic.CACHE_MASK_APPLICABLE) << key

		{ // LEAF CASES
			{

				var expression *algebraic.Expression = algebraic.Int(1)

				result, applicable := expression.IsNegative()

				if result != false {
					t.Errorf("\nresult expected: %t\nresult got     : %t", false, result)
				}
				if applicable != true {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", true, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expectedStatus uint64 = (algebraic.CACHE_MASK_RAN | algebraic.CACHE_MASK_APPLICABLE) << key

				var expression *algebraic.Expression = algebraic.Float(1.1)

				result, applicable := expression.IsNegative()

				if result != false {
					t.Errorf("\nresult expected: %t\nresult got     : %t", false, result)
				}
				if applicable != true {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", true, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Symbol("e")

				result, applicable := expression.IsNegative()

				if result != false {
					t.Errorf("\nresult expected: %t\nresult got     : %t", false, result)
				}
				if applicable != true {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", true, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Symbol("pi")

				result, applicable := expression.IsNegative()

				if result != false {
					t.Errorf("\nresult expected: %t\nresult got     : %t", false, result)
				}
				if applicable != true {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", true, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // NESTED CASES
			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Pow(
						algebraic.Symbol("pi"),
						algebraic.Int(2),
					),
					algebraic.Int(2),
				)

				result, applicable := expression.IsNegative()

				if result != false {
					t.Errorf("\nresult expected: %t\nresult got     : %t", false, result)
				}
				if applicable != true {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", true, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Pow(
						algebraic.Int(-2),
						algebraic.Int(-3),
					),
					algebraic.Int(-1),
				)

				result, applicable := expression.IsNegative()

				if result != false {
					t.Errorf("\nresult expected: %t\nresult got     : %t", false, result)
				}
				if applicable != true {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", true, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Pow(
						algebraic.Int(-2),
						algebraic.Int(3),
					),
					algebraic.Int(2),
				)

				result, applicable := expression.IsNegative()

				if result != false {
					t.Errorf("\nresult expected: %t\nresult got     : %t", false, result)
				}
				if applicable != true {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", true, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Exp(
					algebraic.Pow(
						algebraic.Int(2),
						algebraic.Symbol("x"),
					),
				)

				result, applicable := expression.IsNegative()

				if result != false {
					t.Errorf("\nresult expected: %t\nresult got     : %t", false, result)
				}
				if applicable != true {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", true, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sin(
					algebraic.Multiply(
						algebraic.Int(-1),
						algebraic.Pow(
							algebraic.Symbol("pi"),
							algebraic.Int(2),
						),
					),
				)

				result, applicable := expression.IsNegative()

				if result != false {
					t.Errorf("\nresult expected: %t\nresult got     : %t", false, result)
				}
				if applicable != true {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", true, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Cos(
					algebraic.Multiply(
						algebraic.Int(2),
						algebraic.Pow(
							algebraic.Symbol("pi"),
							algebraic.Int(2),
						),
					),
				)

				result, applicable := expression.IsNegative()

				if result != false {
					t.Errorf("\nresult expected: %t\nresult got     : %t", false, result)
				}
				if applicable != true {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", true, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Tan(
					algebraic.Pow(
						algebraic.Symbol("pi"),
						algebraic.Int(2),
					),
				)

				result, applicable := expression.IsNegative()

				if result != false {
					t.Errorf("\nresult expected: %t\nresult got     : %t", false, result)
				}
				if applicable != true {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", true, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Log(
					algebraic.Int(10),
					algebraic.Pow(
						algebraic.Symbol("pi"),
						algebraic.Int(2),
					),
				)

				result, applicable := expression.IsNegative()

				if result != false {
					t.Errorf("\nresult expected: %t\nresult got     : %t", false, result)
				}
				if applicable != true {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", true, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Ln(
					algebraic.Pow(
						algebraic.Symbol("pi"),
						algebraic.Int(2),
					),
				)

				result, applicable := expression.IsNegative()

				if result != false {
					t.Errorf("\nresult expected: %t\nresult got     : %t", false, result)
				}
				if applicable != true {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", true, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		t.Logf("finishing testing IsNegative for Expression CASE FALSE\n\n")
	}

	{ // NOT APPLICABLE CASES
		t.Logf("testing IsNegative for Expression CASE NOT APPLICABLE\n\n")
		var expectedStatus uint64 = (algebraic.CACHE_MASK_RAN) << key

		{ // LEAF CASES
			{
				var expression *algebraic.Expression = algebraic.Symbol("x")

				result, applicable := expression.IsNegative()

				if result != false {
					t.Errorf("\nresult expected: %t\nresult got     : %t", false, result)
				}
				if applicable != false {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", false, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // NESTED CASES
			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Int(2),
					algebraic.Symbol("x"),
				)

				result, applicable := expression.IsNegative()

				if result != false {
					t.Errorf("\nresult expected: %t\nresult got     : %t", false, result)
				}
				if applicable != false {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", false, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Int(2),
					algebraic.Symbol("x"),
				)

				result, applicable := expression.IsNegative()

				if result != false {
					t.Errorf("\nresult expected: %t\nresult got     : %t", false, result)
				}
				if applicable != false {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", false, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Symbol("x"),
					algebraic.Int(2),
				)

				result, applicable := expression.IsNegative()

				if result != false {
					t.Errorf("\nresult expected: %t\nresult got     : %t", false, result)
				}
				if applicable != false {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", false, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sin(
					algebraic.Pow(
						algebraic.Symbol("x"),
						algebraic.Int(2),
					),
				)

				result, applicable := expression.IsNegative()

				if result != false {
					t.Errorf("\nresult expected: %t\nresult got     : %t", false, result)
				}
				if applicable != false {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", false, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Cos(
					algebraic.Pow(
						algebraic.Symbol("x"),
						algebraic.Int(2),
					),
				)

				result, applicable := expression.IsNegative()

				if result != false {
					t.Errorf("\nresult expected: %t\nresult got     : %t", false, result)
				}
				if applicable != false {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", false, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Tan(
					algebraic.Pow(
						algebraic.Symbol("x"),
						algebraic.Int(2),
					),
				)

				result, applicable := expression.IsNegative()

				if result != false {
					t.Errorf("\nresult expected: %t\nresult got     : %t", false, result)
				}
				if applicable != false {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", false, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Log(
					algebraic.Int(10),
					algebraic.Pow(
						algebraic.Symbol("x"),
						algebraic.Int(2),
					),
				)

				result, applicable := expression.IsNegative()

				if result != false {
					t.Errorf("\nresult expected: %t\nresult got     : %t", false, result)
				}
				if applicable != false {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", false, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Ln(
					algebraic.Pow(
						algebraic.Symbol("x"),
						algebraic.Int(2),
					),
				)

				result, applicable := expression.IsNegative()

				if result != false {
					t.Errorf("\nresult expected: %t\nresult got     : %t", false, result)
				}
				if applicable != false {
					t.Errorf("\napplicable expected: %t\napplicable got     : %t", false, applicable)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		t.Logf("finishing testing IsNegative for Expression CASE NOT APPLICABLE\n\n")
	}

	t.Logf("finishing testing IsNegative for Expression\n\n")
}

func TestIsEvenIntegerExpression(t *testing.T) {
	t.Logf("testing IsEvenInteger for Expression\n\n")
	var key uint64 = uint64(algebraic.CACHE_IS_EVEN_INTEGER)
	var mask uint64 = 0x3 << key

	{ // TRUE CASES
		var expectedStatus uint64 = (algebraic.CACHE_MASK_RAN | algebraic.CACHE_MASK_RESULT) << key

		t.Logf("testing IsEvenInteger for Expression CASE TRUE\n\n")
		{ // LEAF CASES
			{
				var expression *algebraic.Expression = algebraic.Int(2)

				var result bool = expression.IsEvenInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Int(-2)

				var result bool = expression.IsEvenInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Float(2)

				var result bool = expression.IsEvenInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Float(-2)

				var result bool = expression.IsEvenInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // ADDITION CASES
			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Int(1),
					algebraic.Int(3),
				)

				var result bool = expression.IsEvenInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Int(-1),
					algebraic.Int(3),
				)

				var result bool = expression.IsEvenInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Pow(
						algebraic.Int(7),
						algebraic.Int(3),
					),
					algebraic.Multiply(
						algebraic.Cos(
							algebraic.Int(0),
						),
						algebraic.Int(3),
					),
				)

				var result bool = expression.IsEvenInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Pow(
						algebraic.Int(-7),
						algebraic.Int(3),
					),
					algebraic.Multiply(
						algebraic.Cos(
							algebraic.Int(0),
						),
						algebraic.Int(3),
					),
				)

				var result bool = expression.IsEvenInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // MULTIPLY CASES
			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Int(2),
					algebraic.Int(3),
				)

				var result bool = expression.IsEvenInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Int(-2),
					algebraic.Int(3),
				)

				var result bool = expression.IsEvenInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Pow(
						algebraic.Int(7),
						algebraic.Int(3),
					),
					algebraic.Sum(
						algebraic.Cos(
							algebraic.Int(0),
						),
						algebraic.Int(3),
					),
				)

				var result bool = expression.IsEvenInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Pow(
						algebraic.Int(-7),
						algebraic.Int(3),
					),
					algebraic.Sum(
						algebraic.Cos(
							algebraic.Int(0),
						),
						algebraic.Int(3),
					),
				)

				var result bool = expression.IsEvenInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // POWER CASES
			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Int(2),
					algebraic.Int(3),
				)

				var result bool = expression.IsEvenInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Int(-2),
					algebraic.Int(3),
				)

				var result bool = expression.IsEvenInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Multiply(
						algebraic.Int(7),
						algebraic.Int(2),
					),
					algebraic.Sum(
						algebraic.Cos(
							algebraic.Int(0),
						),
						algebraic.Int(3),
					),
				)

				var result bool = expression.IsEvenInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Multiply(
						algebraic.Int(-7),
						algebraic.Int(2),
					),
					algebraic.Sum(
						algebraic.Cos(
							algebraic.Int(0),
						),
						algebraic.Int(3),
					),
				)

				var result bool = expression.IsEvenInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // FUNCTION CASES
			{
				var expression *algebraic.Expression = algebraic.Exp(
					algebraic.Float(0.6931471804287),
				)

				var result bool = expression.IsEvenInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sin(
					algebraic.Pow(
						algebraic.Int(0),
						algebraic.Symbol("x"),
					),
				)

				var result bool = expression.IsEvenInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Cos(
					algebraic.Symbol("pi").Divide(
						algebraic.Int(2),
					),
				)

				var result bool = expression.IsEvenInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Tan(
					algebraic.Int(0),
				)

				var result bool = expression.IsEvenInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Log(
					algebraic.Int(10),
					algebraic.Pow(
						algebraic.Int(10),
						algebraic.Int(2),
					),
				)

				var result bool = expression.IsEvenInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Ln(
					algebraic.Float(7.3890560989306),
				)

				var result bool = expression.IsEvenInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		t.Logf("finishing testing IsEvenInteger for Expression CASE TRUE\n\n")
	}

	{ // FALSE CASES
		var expectedStatus uint64 = algebraic.CACHE_MASK_RAN << key

		t.Logf("testing IsEvenInteger for Expression CASE FALSE\n\n")
		{ // LEAF CASES
			{
				var expression *algebraic.Expression = algebraic.Int(1)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Int(-1)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Float(2.1)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Float(-2.1)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Symbol("x")

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Symbol("e")

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Symbol("pi")

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // ADDITION CASES
			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Int(2),
					algebraic.Int(3),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Int(2),
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Multiply(
						algebraic.Symbol("x"),
						algebraic.Int(5),
					),
					algebraic.Int(3),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // MULTIPLY CASES
			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Int(3),
					algebraic.Int(3),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Int(3),
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Multiply(
						algebraic.Symbol("x"),
						algebraic.Int(5),
					),
					algebraic.Int(3),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // POWER CASES
			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Int(3),
					algebraic.Int(3),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Int(2),
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Multiply(
						algebraic.Symbol("x"),
						algebraic.Int(5),
					),
					algebraic.Int(3),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // FUNCTION CASES
			{
				var expression *algebraic.Expression = algebraic.Exp(
					algebraic.Int(3),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Exp(
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Exp(
					algebraic.Pow(
						algebraic.Symbol("x"),
						algebraic.Int(2),
					),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sin(
					algebraic.Symbol("pi").Divide(
						algebraic.Int(2),
					),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sin(
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sin(
					algebraic.Pow(
						algebraic.Symbol("x"),
						algebraic.Int(2),
					),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Cos(
					algebraic.Int(0),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Cos(
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Cos(
					algebraic.Pow(
						algebraic.Symbol("x"),
						algebraic.Int(2),
					),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Tan(
					algebraic.Symbol("pi").Divide(
						algebraic.Int(4),
					),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Tan(
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Tan(
					algebraic.Pow(
						algebraic.Symbol("x"),
						algebraic.Int(2),
					),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Log(
					algebraic.Int(10),
					algebraic.Int(10),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Log(
					algebraic.Int(10),
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Log(
					algebraic.Int(10),
					algebraic.Pow(
						algebraic.Symbol("x"),
						algebraic.Int(2),
					),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Ln(
					algebraic.Symbol("e"),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Ln(
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Ln(
					algebraic.Pow(
						algebraic.Symbol("x"),
						algebraic.Int(2),
					),
				)

				var result bool = expression.IsEvenInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		t.Logf("finishing testing IsEvenInteger for Expression CASE FALSE\n\n")
	}

	t.Logf("finishing testing IsEvenInteger for Expression\n\n")
}

func TestIsOddIntegerExpression(t *testing.T) {
	t.Logf("testing IsOddInteger for Expression\n\n")
	var key uint64 = uint64(algebraic.CACHE_IS_ODD_INTEGER)
	var mask uint64 = 0x3 << key

	{ // TRUE CASES
		var expectedStatus uint64 = (algebraic.CACHE_MASK_RAN | algebraic.CACHE_MASK_RESULT) << key

		t.Logf("testing IsOddInteger for Expression CASE TRUE\n\n")
		{ // LEAF CASES
			{
				var expression *algebraic.Expression = algebraic.Int(-3)

				var result bool = expression.IsOddInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Int(3)

				var result bool = expression.IsOddInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Float(-3)

				var result bool = expression.IsOddInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Float(3)

				var result bool = expression.IsOddInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // ADDITION CASES
			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Int(1),
					algebraic.Int(4),
				)

				var result bool = expression.IsOddInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Int(-1),
					algebraic.Int(4),
				)

				var result bool = expression.IsOddInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Pow(
						algebraic.Int(7),
						algebraic.Int(3),
					),
					algebraic.Multiply(
						algebraic.Cos(
							algebraic.Int(0),
						),
						algebraic.Int(4),
					),
				)

				var result bool = expression.IsOddInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Pow(
						algebraic.Int(-7),
						algebraic.Int(3),
					),
					algebraic.Multiply(
						algebraic.Cos(
							algebraic.Int(0),
						),
						algebraic.Int(4),
					),
				)

				var result bool = expression.IsOddInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // MULTIPLY CASES
			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Int(3),
					algebraic.Int(5),
				)

				var result bool = expression.IsOddInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Int(-3),
					algebraic.Int(5),
				)

				var result bool = expression.IsOddInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Pow(
						algebraic.Int(7),
						algebraic.Int(3),
					),
					algebraic.Sum(
						algebraic.Cos(
							algebraic.Int(0),
						),
						algebraic.Int(4),
					),
				)

				var result bool = expression.IsOddInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Pow(
						algebraic.Int(-7),
						algebraic.Int(3),
					),
					algebraic.Sum(
						algebraic.Cos(
							algebraic.Int(0),
						),
						algebraic.Int(4),
					),
				)

				var result bool = expression.IsOddInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // POWER CASES
			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Int(3),
					algebraic.Int(5),
				)

				var result bool = expression.IsOddInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Int(-3),
					algebraic.Int(5),
				)

				var result bool = expression.IsOddInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Multiply(
						algebraic.Int(7),
						algebraic.Int(3),
					),
					algebraic.Sum(
						algebraic.Cos(
							algebraic.Int(0),
						),
						algebraic.Int(4),
					),
				)

				var result bool = expression.IsOddInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Multiply(
						algebraic.Int(-7),
						algebraic.Int(3),
					),
					algebraic.Sum(
						algebraic.Cos(
							algebraic.Int(0),
						),
						algebraic.Int(4),
					),
				)

				var result bool = expression.IsOddInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // FUNCTION CASES
			{
				var expression *algebraic.Expression = algebraic.Exp(
					algebraic.Int(0),
				)

				var result bool = expression.IsOddInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sin(
					algebraic.Symbol("pi").Divide(
						algebraic.Int(2),
					),
				)

				var result bool = expression.IsOddInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Cos(
					algebraic.Int(0),
				)

				var result bool = expression.IsOddInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Tan(
					algebraic.Symbol("pi").Divide(
						algebraic.Int(4),
					),
				)

				var result bool = expression.IsOddInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Log(
					algebraic.Int(10),
					algebraic.Int(10),
				)

				var result bool = expression.IsOddInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Ln(
					algebraic.Symbol("e"),
				)

				var result bool = expression.IsOddInteger()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		t.Logf("finishing testing IsOddInteger for Expression CASE TRUE\n\n")
	}

	{ // FALSE CASES
		var expectedStatus uint64 = algebraic.CACHE_MASK_RAN << key

		t.Logf("testing IsOddInteger for Expression CASE FALSE\n\n")
		{ // LEAF CASES
			{
				var expression *algebraic.Expression = algebraic.Int(2)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Int(-2)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Float(3.1)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Float(-3.1)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Symbol("x")

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Symbol("e")

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Symbol("pi")

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // ADDITION CASES
			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Int(1),
					algebraic.Int(3),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Int(2),
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sum(
					algebraic.Multiply(
						algebraic.Symbol("x"),
						algebraic.Int(5),
					),
					algebraic.Int(3),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // MULTIPLY CASES
			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Int(3),
					algebraic.Int(4),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Int(3),
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Multiply(
					algebraic.Multiply(
						algebraic.Symbol("x"),
						algebraic.Int(5),
					),
					algebraic.Int(3),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // POWER CASES
			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Int(2),
					algebraic.Int(4),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Int(2),
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Pow(
					algebraic.Multiply(
						algebraic.Symbol("x"),
						algebraic.Int(5),
					),
					algebraic.Int(3),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // FUNCTION CASES
			{
				var expression *algebraic.Expression = algebraic.Exp(
					algebraic.Int(3),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Exp(
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Exp(
					algebraic.Pow(
						algebraic.Symbol("x"),
						algebraic.Int(2),
					),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sin(
					algebraic.Int(0),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sin(
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Sin(
					algebraic.Pow(
						algebraic.Symbol("x"),
						algebraic.Int(2),
					),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Cos(
					algebraic.Symbol("pi").Divide(
						algebraic.Int(2),
					),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Cos(
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Cos(
					algebraic.Pow(
						algebraic.Symbol("x"),
						algebraic.Int(2),
					),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Tan(
					algebraic.Int(0),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Tan(
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Tan(
					algebraic.Pow(
						algebraic.Symbol("x"),
						algebraic.Int(2),
					),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Log(
					algebraic.Int(10),
					algebraic.Int(100),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Log(
					algebraic.Int(10),
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Log(
					algebraic.Int(10),
					algebraic.Pow(
						algebraic.Symbol("x"),
						algebraic.Int(2),
					),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Ln(
					algebraic.Int(1),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Ln(
					algebraic.Symbol("x"),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Ln(
					algebraic.Pow(
						algebraic.Symbol("x"),
						algebraic.Int(2),
					),
				)

				var result bool = expression.IsOddInteger()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		t.Logf("finishing testing IsOddInteger for Expression CASE FALSE\n\n")
	}

	t.Logf("finishing testing IsOddInteger for Expression\n\n")
}

func TestIsAbsoluteOneExpression(t *testing.T) {
	t.Logf("testing IsAbsoluteOne for Expression\n\n")
	var key uint64 = uint64(algebraic.CACHE_IS_ABSOLUTE_ONE)
	var mask uint64 = 0x3 << key

	{ // TRUE CASES
		var expectedStatus uint64 = (algebraic.CACHE_MASK_RAN | algebraic.CACHE_MASK_RESULT) << key

		t.Logf("testing IsAbsoluteOne for Expression CASE TRUE\n\n")
		{ // LEAF CASES
			{
				var expression *algebraic.Expression = algebraic.Int(1)

				var result bool = expression.IsAbsoluteOne()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Int(-1)

				var result bool = expression.IsAbsoluteOne()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Float(1)

				var result bool = expression.IsAbsoluteOne()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Float(-1)

				var result bool = expression.IsAbsoluteOne()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ // ADDITION
			// {
			// 	var expression *algebraic.Expression = algebraic.Sum()

			// 	var result bool = expression.IsAbsoluteOne()

			// 	if result != true {
			// 		t.Errorf("\nexpected: %t\ngot     : %t", true, result)
			// 	}

			// 	if (expression.Cache.Evaluated & mask) != expectedStatus {
			// 		t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
			// 			bits.Len64(uint64(mask>>key)), expectedStatus>>key,
			// 			bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
			// 		)
			// 	}
			// }
		}

		t.Logf("finishing testing IsAbsoluteOne for Expression CASE TRUE\n\n")
	}

	{ // FALSE CASES
		var expectedStatus uint64 = algebraic.CACHE_MASK_RAN << key

		t.Logf("testing IsAbsoluteOne for Expression CASE FALSE\n\n")
		{ // LEAF CASES
			{
				var expression *algebraic.Expression = algebraic.Int(2)

				var result bool = expression.IsAbsoluteOne()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Int(-2)

				var result bool = expression.IsAbsoluteOne()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Float(1.1)

				var result bool = expression.IsAbsoluteOne()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Float(-1.1)

				var result bool = expression.IsAbsoluteOne()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Symbol("x")

				var result bool = expression.IsAbsoluteOne()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Symbol("e")

				var result bool = expression.IsAbsoluteOne()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Symbol("pi")

				var result bool = expression.IsAbsoluteOne()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ //

		}

		t.Logf("finishing testing IsAbsoluteOne for Expression CASE FALSE\n\n")
	}

	t.Logf("finishing testing IsAbsoluteOne for Expression\n\n")
}

func TestIsZeroExpression(t *testing.T) {
	t.Logf("testing IsZero for Expression\n\n")
	var key uint64 = uint64(algebraic.CACHE_IS_ZERO)
	var mask uint64 = 0x3 << key

	{ // TRUE CASES
		var expectedStatus uint64 = (algebraic.CACHE_MASK_RAN | algebraic.CACHE_MASK_RESULT) << key

		t.Logf("testing IsZero for Expression CASE TRUE\n\n")
		{ // LEAF CASES
			{
				var expression *algebraic.Expression = algebraic.Int(0)

				var result bool = expression.IsZero()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Float(0)

				var result bool = expression.IsZero()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ //

		}

		t.Logf("finishing testing IsZero for Expression CASE TRUE\n\n")
	}

	{ // FALSE CASES
		var expectedStatus uint64 = algebraic.CACHE_MASK_RAN << key

		t.Logf("testing IsZero for Expression CASE FALSE\n\n")
		{ // LEAF CASES
			{
				var expression *algebraic.Expression = algebraic.Int(1)

				var result bool = expression.IsZero()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Float(0.1)

				var result bool = expression.IsZero()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Float(-0.1)

				var result bool = expression.IsZero()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Symbol("x")

				var result bool = expression.IsZero()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Symbol("e")

				var result bool = expression.IsZero()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Symbol("pi")

				var result bool = expression.IsZero()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ //

		}

		t.Logf("finishing testing IsZero for Expression CASE FALSE\n\n")
	}

	t.Logf("finishing testing IsZero for Expression\n\n")
}

func TestIsEulerExpression(t *testing.T) {
	t.Logf("testing IsEuler for Expression\n\n")
	var key uint64 = uint64(algebraic.CACHE_IS_EULER)
	var mask uint64 = 0x3 << key

	{ // TRUE CASES
		var expectedStatus uint64 = (algebraic.CACHE_MASK_RAN | algebraic.CACHE_MASK_RESULT) << key

		t.Logf("testing IsEuler for Expression CASE TRUE\n\n")
		{ // LEAF CASES
			{
				var expression *algebraic.Expression = algebraic.Float(math.E)

				var result bool = expression.IsEuler()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Symbol("e")

				var result bool = expression.IsEuler()

				if result != true {
					t.Errorf("\nexpected: %t\ngot     : %t", true, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ //

		}

		t.Logf("finishing testing IsEuler for Expression CASE TRUE\n\n")
	}

	{ // FALSE CASES
		var expectedStatus uint64 = algebraic.CACHE_MASK_RAN << key

		t.Logf("testing IsEuler for Expression CASE FALSE\n\n")
		{ // LEAF CASES
			{
				var expression *algebraic.Expression = algebraic.Int(2)

				var result bool = expression.IsEuler()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Float(2.718)

				var result bool = expression.IsEuler()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Symbol("x")

				var result bool = expression.IsEuler()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}

			{
				var expression *algebraic.Expression = algebraic.Symbol("pi")

				var result bool = expression.IsEuler()

				if result != false {
					t.Errorf("\nexpected: %t\ngot     : %t", false, result)
				}

				if (expression.Cache.Evaluated & mask) != expectedStatus {
					t.Errorf("error with pre-computed cache.\nbits expected: %0*b\nbits got     : %0*b",
						bits.Len64(uint64(mask>>key)), expectedStatus>>key,
						bits.Len64(uint64(mask>>key)), (expression.Cache.Evaluated&mask)>>key,
					)
				}
			}
		}

		{ //

		}

		t.Logf("finishing testing IsEuler for Expression CASE FALSE\n\n")
	}

	t.Logf("finishing testing IsEuler for Expression\n\n")
}
