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
package algebra_numeric_method

import (
	"errors"
	"fmt"
	"math"

	algebra_equation "github.com/FabioLuisBoni/go-algebra/equation"
	algebra_expression "github.com/FabioLuisBoni/go-algebra/equation/expression"
	algebra_frame "github.com/FabioLuisBoni/go-algebra/frame"
)

type LeastSquares struct {
	EquationSignature string

	Base []*algebra_expression.Expression
}

func NewLeastSquares(equationSignature string) *LeastSquares {
	return &LeastSquares{
		EquationSignature: equationSignature,
	}
}

func (method *LeastSquares) BaseOn(equation *algebra_equation.Equation) error {
	if equation.Expression.IsIndefiniteness() {
		return fmt.Errorf("cannot use indefinite equation as base for least squares")
	}

	if equation.Expression.Type == algebra_expression.ADDITION {
		method.Base = make([]*algebra_expression.Expression, 0, len(equation.Expression.Arguments))

		for _, branch := range equation.Expression.Arguments {
			method.Base = append(
				method.Base,
				branch.Clone(),
			)
		}

	} else {
		method.Base = []*algebra_expression.Expression{equation.Expression.Clone()}
	}

	return nil
}

func (method *LeastSquares) Solve(frame algebra_frame.Frame) (equation *algebra_equation.Equation, err error) {
	var pointQuantity int = len(frame.Point)
	var termsQuantity int = len(method.Base)

	if pointQuantity == 0 {
		return nil, fmt.Errorf("no points provided")
	}
	if termsQuantity == 0 {
		return nil, fmt.Errorf("no basis functions provided")
	}
	if pointQuantity < termsQuantity {
		return nil, fmt.Errorf("least squares requires at least same amount of points that equations terms: received %d points and %d equation terms", pointQuantity, termsQuantity)
	}

	X := make([][]float64, pointQuantity) // Build matrix X
	y := make([]float64, pointQuantity)   // Build vector y

	for i := range pointQuantity {
		X[i] = make([]float64, termsQuantity)
		for j := range termsQuantity {
			X[i][j] = method.Base[j].Execute(frame.Point[i].X)
		}
		y[i] = frame.Point[i].Y
	}

	XtX := make([][]float64, termsQuantity) // Compute XtX and Xty
	for i := range termsQuantity {
		XtX[i] = make([]float64, termsQuantity)
		for j := range termsQuantity {
			sum := 0.0
			for k := range pointQuantity {
				sum += X[k][i] * X[k][j]
			}
			XtX[i][j] = sum
		}
	}

	Xty := make([]float64, termsQuantity)
	for i := range termsQuantity {
		sum := 0.0
		for k := range pointQuantity {
			sum += X[k][i] * y[k]
		}
		Xty[i] = sum
	}

	coefficients, err := solveLinearSystem(XtX, Xty) // Solve linear system XtX * beta = Xty
	if err != nil {
		return nil, fmt.Errorf("error solving linear system: %w", err)
	}

	equation = algebra_equation.NewEquation(method.EquationSignature)
	equation.Expression.Sum(
		method.Base[0],
		method.Base[1:]...,
	)

	for i, branch := range equation.Expression.Arguments {
		branch.Multiply(
			algebra_expression.Float(coefficients[i]),
		)
	}

	return equation, nil
}

func solveLinearSystem(A [][]float64, b []float64) ([]float64, error) {
	n := len(b)

	for i := range n {
		maxRow := i
		for k := i + 1; k < n; k++ {
			if math.Abs(A[k][i]) > math.Abs(A[maxRow][i]) {
				maxRow = k
			}
		}

		A[i], A[maxRow] = A[maxRow], A[i]
		b[i], b[maxRow] = b[maxRow], b[i]

		if math.Abs(A[i][i]) < 1e-12 {
			return nil, errors.New("singular matrix")
		}

		for k := i + 1; k < n; k++ {
			factor := A[k][i] / A[i][i]
			for j := i; j < n; j++ {
				A[k][j] -= factor * A[i][j]
			}
			b[k] -= factor * b[i]
		}
	}

	x := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		sum := b[i]
		for j := i + 1; j < n; j++ {
			sum -= A[i][j] * x[j]
		}
		x[i] = sum / A[i][i]
	}

	return x, nil
}
