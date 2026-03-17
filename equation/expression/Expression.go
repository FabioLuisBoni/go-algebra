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

func (expression *Expression) Solve(operator float64) (result float64) {
	if expression == nil {
		return math.NaN()
	}

	if expression.IsIndefiniteness() {
		return math.NaN()
	}

	switch expression.Type {
	case INTEGER, FLOAT:
		return *expression.Value

	case SYMBOL:
		switch expression.Name {
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
			Some extra validations must happen before calling expression.Solve() due to the
				possibility of missing operators, since Solve() itself is not aware of what
				happen in its child or parent Solve() calling.
		*/
		return operator

	case ADDITION:
		for _, branch := range expression.Arguments {
			result += branch.Solve(operator)
		}
		return result

	case MULTIPLICATION:
		result = 1
		for _, branch := range expression.Arguments {
			result *= branch.Solve(operator)
		}
		return result

	case POWER:
		return math.Pow(
			expression.Arguments[0].Solve(operator),
			expression.Arguments[1].Solve(operator),
		)

	case EXPONENTIAL:
		return math.Exp(expression.Arguments[0].Solve(operator))

	case SINE:
		return math.Sin(expression.Arguments[0].Solve(operator))

	case COSINE:
		return math.Cos(expression.Arguments[0].Solve(operator))

	case TANGENT:
		return math.Tan(expression.Arguments[0].Solve(operator))

	case LOGARITHMIC:
		if len(expression.Arguments) == 1 {
			return math.Log(expression.Arguments[0].Solve(operator))

		} else {
			return math.Log(expression.Arguments[0].Solve(operator)) / math.Log(expression.Arguments[1].Solve(operator))
		}

	default:
		return math.NaN()
	}
}

/*
Returns if the both expressions are equals.

The comparisons goes deep, caring about behavior rather than structure,
simplifying sub-trees to reach the answer.

For example:

	Pow(Symbol("x"), Int(0)) == Sin(Symbol("pi").Divide(Int(2)))
*/
func (expression *Expression) Equal(other *Expression) bool {
	if expression == nil || other == nil {
		return expression == other
	}

	if expression.IsConstant() && other.IsConstant() { // Cover all format that leads to same results, like x^0 == cos(0).
		return isApproximate(expression.Solve(SOLVE_CONSTANT_PLACEHOLDER), other.Solve(SOLVE_CONSTANT_PLACEHOLDER)) // being constants, the only possible true result is the Solve to be equal.
	}

	if expression.Type != other.Type { // If reached here, the types must be equal.
		if expression.isIdentityWrapper() {
			return expression.unwrapIdentity().Equal(other)
		}

		if other.isIdentityWrapper() {
			return other.unwrapIdentity().Equal(expression)
		}

		return false
	}

	switch expression.Type { // Since types are equal, switch is safe to be applied.
	case INTEGER, FLOAT: // Constant leafs needs to compare values.
		if (expression.Value != nil) != (other.Value != nil) { // Check if both are nil or non-nil.
			return false
		} else if expression.Value == nil { // Nil values actually mean something have gone wrong, it should not exist. But its not this method's responsibility to treat it, so... they're equal.
			return true
		}

		return expression.Value != nil && other.Value != nil && *expression.Value == *other.Value

	case SYMBOL: // Symbolic leafs needs to compare the symbol itself.
		return expression.Name == other.Name

	case ADDITION, MULTIPLICATION: // Additions and multiplications are commutative, so order of arguments does not matter and should be treated with special care.
		return expression.commutativeBehavioralEqual(other)

	default: // Any other type, have a strict pattern or arguments, and any ordered difference means completely different things, since previous validations were made.
		if len(expression.Arguments) != len(other.Arguments) {
			return false
		}

		for i := range expression.Arguments {
			if !expression.Arguments[i].Equal(other.Arguments[i]) {
				return false
			}
		}
	}

	return true
}

/*
Return a deep copied object of the expression.

The cache of every nested expression is cloned too.
*/
func (expression *Expression) Clone() *Expression {
	if expression == nil {
		return nil
	}

	var clone *Expression = &Expression{
		Type: expression.Type,
		Name: expression.Name,

		Cache: expression.Cache,
	}

	if expression.Value != nil {
		var val float64 = *expression.Value
		clone.Value = &val
	}

	if len(expression.Arguments) > 0 {
		clone.Arguments = make([]*Expression, len(expression.Arguments))
		for i, branch := range expression.Arguments {
			clone.Arguments[i] = branch.Clone()
		}
	}

	return clone
}

/*
Recursively clear the cache of pre-computing operations in the whole
expression.
*/
func (expression *Expression) ClearCache() {
	expression.Cache.clearCache()

	for _, branch := range expression.Arguments {
		branch.ClearCache()
	}
}

/*
Presents a simplified, human readable and GeoGebra friendly algebraic notation
formatted string.

The division '/' and root '√' will be completely omitted due to the
complexities and heavy computation for a reliable representation.

Instead, keep in mind that:
  - Negative exponents represents inverse values.
  - Fraction (float64) exponents are roots.
  - Divisions are multiplication by inverse values.
  - x^(-1.5) is analog to 1/√(x^3)
  - x^(-1.3478) is analog to something like 1/√(x^6739, 5000), and the reason
    to avoid this notation, not readable and breaks GeoGebra friendly policy.
*/
func (expression *Expression) String() string {
	return algebraicString(expression)
}

/*
Addition and multiplication have commutative behavior, and comparisons must be
extra careful about its content.

This method enforces that both Expression be equal and multiplication or
addition, then it re-arrange sub-trees, merging and comparing constant
sub-trees and comparing every variable sub-trees.
*/
func (expression *Expression) commutativeBehavioralEqual(other *Expression) bool {
	if expression.Type != other.Type {
		return false
	}
	if expression.Type != ADDITION && expression.Type != MULTIPLICATION {
		return false
	}

	constantExpression, treeExpression := expression.partitionArguments()
	constantOther, treeOther := other.partitionArguments()

	if !isApproximate(constantExpression, constantOther) {
		return false
	}

	if len(treeExpression) != len(treeOther) {
		return false
	}

	var matched []bool = make([]bool, len(treeOther))
	for _, subTreeExpression := range treeExpression {
		var found bool

		for j, subTreeOther := range treeOther {
			if !matched[j] && subTreeExpression.Equal(subTreeOther) {
				matched[j] = true
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}

/*
Simplify the arguments into merging constants together (independent of its
sub-tree structure) and enforcing only the variables sub-trees to be returned.
*/
func (expression *Expression) partitionArguments() (constant float64, trees []*Expression) {
	if expression.Type == MULTIPLICATION {
		constant = 1.0
	} else {
		constant = 0.0
	}

	for _, subExpression := range expression.Arguments {
		if subExpression.IsConstant() {
			if expression.Type == MULTIPLICATION {
				constant *= subExpression.Solve(SOLVE_CONSTANT_PLACEHOLDER)
			} else {
				constant += subExpression.Solve(SOLVE_CONSTANT_PLACEHOLDER)
			}

		} else {
			trees = append(trees, subExpression)
		}
	}

	return constant, trees
}

/*
Identify if the expression is a wrapper of irrelevant terms over a variable
For example:

	x + 0
	x + 1 - 1
	x * 1
	x * 0.5 * 2
*/
func (expression *Expression) isIdentityWrapper() bool {
	if expression.Type != ADDITION && expression.Type != MULTIPLICATION {
		return false
	}

	constant, trees := expression.partitionArguments()

	if expression.Type == ADDITION {
		return len(trees) == 1 && isApproximate(constant, 0)
	}

	return len(trees) == 1 && isApproximate(constant, 1)
}

/*
Helper to get the actual meaningful part of the identity wrapper.
*/
func (expression *Expression) unwrapIdentity() *Expression {
	_, trees := expression.partitionArguments()
	return trees[0]
}
