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
package algebra_type

type EquationTermType string

// Polynomial Family
const (
	CONSTANT  EquationTermType = "constant"
	LINEAR    EquationTermType = "linear"
	QUADRATIC EquationTermType = "quadratic"
	CUBIC     EquationTermType = "cubic"
)

// Trigonometric Family
const (
	SIN EquationTermType = "sin"
	COS EquationTermType = "cos"
	TAN EquationTermType = "tan"
)

// Logarithmic Family
const (
	LOGARITHMIC EquationTermType = "logarithmic"
)

// Unknown Family
const (
	OTHER EquationTermType = "other"
)

type SortingMethod string

const (
	X_HIGHER SortingMethod = "x_higher"
	Y_HIGHER SortingMethod = "y_higher"
	X_LOWER  SortingMethod = "x_lower"
	Y_LOWER  SortingMethod = "y_lower"
)
