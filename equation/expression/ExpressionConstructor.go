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

type ExpressionType string

/*
ATOMIC FAMILY

Represent a leaf of the branch.
*/
const (
	INTEGER ExpressionType = "integer"
	FLOAT   ExpressionType = "float"
	SYMBOL  ExpressionType = "symbol"
)

/*
OPERATION FAMILY

Represent multiple repeated actions over another operation.
*/
const (
	ADDITION       ExpressionType = "add"
	MULTIPLICATION ExpressionType = "mul"
)

// FUNCTION FAMILY
const (
	POWER       ExpressionType = "pow"
	EXPONENTIAL ExpressionType = "exp"
	LOGARITHMIC ExpressionType = "log"
	SIN         ExpressionType = "sin"
	COS         ExpressionType = "cos"
	TAN         ExpressionType = "tan"
)

func Symbol(name string) *Expression {
	return &Expression{
		Type: SYMBOL,
		Name: name,
	}
}

func Int(integer int) *Expression {
	var value float64 = float64(integer)

	return &Expression{
		Type:  INTEGER,
		Value: &value,
	}
}

func Float(float float64) *Expression {
	var value float64 = float

	return &Expression{
		Type:  FLOAT,
		Value: &value,
	}
}

func Sum(operation1 *Expression, operation2 *Expression, others ...*Expression) *Expression {
	return &Expression{
		Type: ADDITION,

		Arguments: append(
			[]*Expression{
				operation1,
				operation2,
			},
			others...,
		),
	}
}

func Multiply(operation1 *Expression, operation2 *Expression, others ...*Expression) *Expression {
	return &Expression{
		Type: MULTIPLICATION,

		Arguments: append(
			[]*Expression{
				operation1,
				operation2,
			},
			others...,
		),
	}
}

func Pow(base *Expression, exponent *Expression) *Expression {
	return &Expression{
		Type: POWER,

		Arguments: []*Expression{
			base,
			exponent,
		},
	}
}

func Exp(exponent *Expression) *Expression {
	return &Expression{
		Type: EXPONENTIAL,

		Arguments: []*Expression{
			exponent,
		},
	}
}

func Ln(operation *Expression) *Expression {
	return &Expression{
		Type: LOGARITHMIC,

		Arguments: []*Expression{
			operation,
		},
	}
}

func Log(base *Expression, operation *Expression) *Expression {
	return &Expression{
		Type: LOGARITHMIC,

		Arguments: []*Expression{
			operation,
			base,
		},
	}
}

func Sin(operation *Expression) *Expression {
	return &Expression{
		Type: SIN,

		Arguments: []*Expression{
			operation,
		},
	}
}

func Cos(operation *Expression) *Expression {
	return &Expression{
		Type: COS,

		Arguments: []*Expression{
			operation,
		},
	}
}

func Tan(operation *Expression) *Expression {
	return &Expression{
		Type: TAN,

		Arguments: []*Expression{
			operation,
		},
	}
}
