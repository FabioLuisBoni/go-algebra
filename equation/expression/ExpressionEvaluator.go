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

const (
	SYMBOL_PI    string = "pi"
	SYMBOL_EULER string = "e"
)

const APPROXIMATION_EPSILON float64 = 1e-9
const SOLVE_CONSTANT_PLACEHOLDER float64 = 2

/*
Evaluates if the Expression object structure is valid for equation representation.
*/
func (expression *Expression) IsMalformedStructure() bool {
	if expression.Cache.isCached(CACHE_IS_MALFORMED_STRUCTURE) {
		return expression.Cache.result(CACHE_IS_MALFORMED_STRUCTURE)
	}
	switch expression.Type {
	case INTEGER, FLOAT:
		if expression.Value == nil {
			return expression.Cache.setMalformedStructure(true)
		}

	case SYMBOL:
		if expression.Name == "" {
			return expression.Cache.setMalformedStructure(true)
		}

	case ADDITION, MULTIPLICATION:
		if len(expression.Arguments) < 2 {
			return expression.Cache.setMalformedStructure(true)
		}

	case POWER:
		if len(expression.Arguments) != 2 {
			return expression.Cache.setMalformedStructure(true)
		}

	case LOGARITHMIC:
		if len(expression.Arguments) != 1 && len(expression.Arguments) != 2 {
			return expression.Cache.setMalformedStructure(true)
		}

	default:
		if len(expression.Arguments) != 1 {
			return expression.Cache.setMalformedStructure(true)
		}
	}

	for _, branch := range expression.Arguments {
		if branch == nil {
			return expression.Cache.setMalformedStructure(true)
		}

		if branch.IsMalformedStructure() {
			return expression.Cache.setMalformedStructure(true)
		}
	}

	return expression.Cache.setMalformedStructure(false)
}

/*
Evaluates if the whole branch could result in mathematical indefiniteness.

For protection and validation, it is considered indefiniteness if the arguments
of the expressions are not what the pattern expects.
*/
func (expression *Expression) IsIndefiniteness() bool {
	if expression.Cache.isCached(CACHE_IS_INDEFINITENESS) {
		return expression.Cache.result(CACHE_IS_INDEFINITENESS)
	}

	if expression.IsMalformedStructure() {
		return expression.Cache.setIndefiniteness(true)
	}

	switch expression.Type {
	case POWER:
		/*
			Power have special case needing to evaluated-compute base and exponent together,
			since:
				* 0^negative is mathematical indefiniteness. Represents a division by 0.
				* 0^0 is mathematical indefiniteness.
				* negative^fraction results in a complex number.
		*/

		exponentNegative, _ := expression.Arguments[1].IsNegative()
		if expression.Arguments[0].IsZero() && (expression.Arguments[1].IsZero() || exponentNegative) {
			return expression.Cache.setIndefiniteness(true)
		}

		baseNegative, _ := expression.Arguments[0].IsNegative()
		if baseNegative && expression.Arguments[1].IsFraction() {
			return expression.Cache.setIndefiniteness(true)
		}

	case TANGENT:
		if expression.Arguments[0].isTangentIndefinite() {
			return expression.Cache.setIndefiniteness(true)
		}

	case LOGARITHMIC:
		/*
			Logarithms have special case needing to evaluated-compute base, since:
				* Base 1 is mathematical indefiniteness.
				* Base 0 is mathematical indefiniteness.
				* Operator 0 is mathematical indefiniteness.
				* Base negative results in a complex number.
				* Operator negative results in a complex number.
		*/

		if expression.Arguments[0].IsZero() {
			return expression.Cache.setIndefiniteness(true)
		}

		var baseNegative bool
		if len(expression.Arguments) == 2 {
			if expression.Arguments[1].IsZero() || expression.Arguments[1].IsAbsoluteOne() {
				return expression.Cache.setIndefiniteness(true)
			}

			baseNegative, _ = expression.Arguments[1].IsNegative()
		}

		operatorNegative, _ := expression.Arguments[0].IsNegative()

		if baseNegative || operatorNegative {
			return expression.Cache.setIndefiniteness(true)
		}
	}

	for _, branch := range expression.Arguments {
		if branch.IsIndefiniteness() {
			return expression.Cache.setIndefiniteness(true)
		}
	}

	return expression.Cache.setIndefiniteness(false)
}

/*
Evaluate the expression to predict if its only possible result is a constant.
*/
func (expression *Expression) IsConstant() bool {
	if expression.Cache.isCached(CACHE_IS_CONSTANT) {
		return expression.Cache.result(CACHE_IS_CONSTANT)
	}

	if expression.IsIndefiniteness() {
		return expression.Cache.setConstant(false)
	}

	switch expression.Type {
	case INTEGER, FLOAT: // last leaf of the branch as a constant
		return expression.Cache.setConstant(true)

	case SYMBOL: // last leaf of the branch as a symbol
		switch expression.Name {
		case SYMBOL_EULER, SYMBOL_PI: // symbol that represents a constant
			return expression.Cache.setConstant(true)
		default: // symbol that represents a variable
			return expression.Cache.setConstant(false)
		}

	default:
		switch expression.Type { // special performance and evaluation cases
		case MULTIPLICATION:
			/*
				It may seem a redundant switch since a similar switch already exists forward in
				the code, but it is placed here for performance reasons.
				This guarantee that first of all, the IsZero() cases are evaluated first, since
				any 0 render entire multiplication as 0.
				Otherwise it would unnecessarily compute IsAbsoluteOne() and IsEuler() over
				potentially complex recursion, since a latter 0 should return true anyway.
			*/
			for _, subExpression := range expression.Arguments {
				if subExpression.IsZero() {
					_ = expression.Cache.setZero(true)
					return expression.Cache.setConstant(true)
				}
			}

		case POWER:
			/*
				Power have special case needing to evaluated-compute base and exponent together,
				since:
					* x^0 equals to 1, a constant.
					* 0^x equals to 0, a constant.
					* 1^x equals to 1, a constant.
				So, for performance reasons, this is a potential shortcut to avoid calling
				unnecessarily IsAbsoluteOne() and IsEuler().
			*/

			var baseZero bool = expression.Arguments[0].IsZero()
			var exponentZero bool = expression.Arguments[1].IsZero()

			if baseZero || exponentZero {
				if baseZero {
					_ = expression.Cache.setZero(true)
				} else {
					_ = expression.Cache.setAbsoluteOne(true)
					_, _ = expression.Cache.setNegative(false, true)
				}

				return expression.Cache.setConstant(true)
			}

			if expression.Arguments[0].IsAbsoluteOne() {
				_ = expression.Cache.setAbsoluteOne(true)
				return expression.Cache.setConstant(true)
			}

		case LOGARITHMIC:
			/*
				Logarithms have special case needing to evaluated-compute base, since:
					* Base equal to operator, unless hitting indefiniteness, is 1, a constant.
			*/

			if len(expression.Arguments) == 2 {
				if expression.Arguments[0].Equal(expression.Arguments[1]) {
					_ = expression.Cache.setAbsoluteOne(true)
					_, _ = expression.Cache.setNegative(false, true)
					return expression.Cache.setConstant(true)
				}

			} else {
				if expression.Arguments[0].IsEuler() {
					_ = expression.Cache.setAbsoluteOne(true)
					_, _ = expression.Cache.setNegative(false, true)
					return expression.Cache.setConstant(true)
				}
			}
		}

		for _, branch := range expression.Arguments { // scraping for any non-constant
			if !branch.IsConstant() {
				return expression.Cache.setConstant(false)
			}
		}

		return expression.Cache.setConstant(true)
	}
}

/*
Evaluate the expression to predict if its only possible result is 0.
*/
func (expression *Expression) IsZero() bool {
	if expression.Cache.isCached(CACHE_IS_ZERO) {
		return expression.Cache.result(CACHE_IS_ZERO)
	}

	if expression.IsIndefiniteness() {
		return expression.Cache.setZero(false)
	}

	if !expression.IsConstant() {
		return expression.Cache.setZero(false)
	}

	return expression.Cache.setZero(isApproximate(expression.Solve(SOLVE_CONSTANT_PLACEHOLDER), 0))
}

/*
Evaluate the expression to predict if its only possible result is 1 or -1.

If needed to evaluate a the signal as well, use it with expression.IsNegative().
*/
func (expression *Expression) IsAbsoluteOne() bool {
	if expression.Cache.isCached(CACHE_IS_ABSOLUTE_ONE) {
		return expression.Cache.result(CACHE_IS_ABSOLUTE_ONE)
	}

	if expression.IsIndefiniteness() {
		return expression.Cache.setAbsoluteOne(false)
	}

	if !expression.IsConstant() {
		return expression.Cache.setAbsoluteOne(false)
	}

	return expression.Cache.setAbsoluteOne(isApproximate(math.Abs(expression.Solve(SOLVE_CONSTANT_PLACEHOLDER)), 1))
}

/*
Evaluate the expression to predict if its only possible result is e (euler's number).
*/
func (expression *Expression) IsEuler() bool {
	if expression.Cache.isCached(CACHE_IS_EULER) {
		return expression.Cache.result(CACHE_IS_EULER)
	}

	if expression.IsIndefiniteness() {
		return expression.Cache.setEuler(false)
	}

	if !expression.IsConstant() {
		return expression.Cache.setEuler(false)
	}

	return expression.Cache.setEuler(isApproximate(expression.Solve(SOLVE_CONSTANT_PLACEHOLDER), math.E))
}

func (expression *Expression) IsFraction() bool {
	if expression.Cache.isCached(CACHE_IS_FRACTION) {
		return expression.Cache.result(CACHE_IS_FRACTION)
	}

	if !expression.IsConstant() {
		return expression.Cache.setFraction(false)
	}

	var value float64 = expression.Solve(SOLVE_CONSTANT_PLACEHOLDER)
	var rounded float64 = math.Round(value)

	return expression.Cache.setFraction(!isApproximate(value, rounded))
}

func (expression *Expression) IsInteger() bool {
	if expression.Cache.isCached(CACHE_IS_INTEGER) {
		return expression.Cache.result(CACHE_IS_INTEGER)
	}

	if !expression.IsConstant() {
		return expression.Cache.setInteger(false)
	}

	var value float64 = expression.Solve(SOLVE_CONSTANT_PLACEHOLDER)
	var rounded float64 = math.Round(value)

	return expression.Cache.setInteger(isApproximate(value, rounded))
}

func (expression *Expression) IsEvenInteger() bool {
	if expression.Cache.isCached(CACHE_IS_EVEN_INTEGER) {
		return expression.Cache.result(CACHE_IS_EVEN_INTEGER)
	}

	if !expression.IsConstant() {
		return expression.Cache.setEvenInteger(false)
	}

	var value float64 = expression.Solve(SOLVE_CONSTANT_PLACEHOLDER)
	var rounded float64 = math.Round(value)

	if !isApproximate(value, rounded) {
		return expression.Cache.setEvenInteger(false)
	}

	return expression.Cache.setEvenInteger(int64(math.Abs(rounded))%2 == 0)
}

func (expression *Expression) IsOddInteger() bool {
	if expression.Cache.isCached(CACHE_IS_ODD_INTEGER) {
		return expression.Cache.result(CACHE_IS_ODD_INTEGER)
	}

	if !expression.IsConstant() {
		return expression.Cache.setOddInteger(false)
	}

	var value float64 = expression.Solve(SOLVE_CONSTANT_PLACEHOLDER)
	var rounded float64 = math.Round(value)

	if !isApproximate(value, rounded) {
		return expression.Cache.setOddInteger(false)
	}

	return expression.Cache.setOddInteger(int64(math.Abs(rounded))%2 != 0)
}

/*
Evaluate the expression to predict if the whole expression may be adjusted by
signal inversion (multiplication by -1).
This is relevant since operators may be simplified and represented at an
overall context instead of spawning Multiply expressions all over the equation,
causing a performance downgrade.

Integer and Float expressions are leafs of larger branches, they just returns
if its value is < 0.

Symbol expression have no signal significance, and will always return false.

Power, Exponential, Sin, Cos, Tan and Log expressions are not evaluable, since
applicable signal inversion would be possible only over constant resulting
expressions, or not applicable at all.

Constant resulting expression simplification is IsConstant() method domain, and
probably should have been be called before IsSignalInvertible() for
consolidation over Int and Float expressions.
*/
func (expression *Expression) IsSignalInvertible() bool {
	if expression.Cache.isCached(CACHE_IS_SIGNAL_INVERTIBLE) {
		return expression.Cache.result(CACHE_IS_SIGNAL_INVERTIBLE)
	}

	if expression.IsIndefiniteness() {
		return expression.Cache.setSignalInvertible(false)
	}

	switch expression.Type {
	case INTEGER, FLOAT:
		return expression.Cache.setSignalInvertible(*expression.Value < 0)

	case ADDITION:
		var total float64
		for _, branch := range expression.Arguments {
			if !branch.IsConstant() {
				return expression.Cache.setSignalInvertible(false)
			}

			total += branch.Solve(SOLVE_CONSTANT_PLACEHOLDER)
		}

		return expression.Cache.setSignalInvertible(total < 0)

	case MULTIPLICATION:
		var negative bool
		for _, branch := range expression.Arguments {
			if branch.IsSignalInvertible() {
				negative = !negative
			}
		}

		return expression.Cache.setSignalInvertible(negative)

	case POWER:
		if expression.Arguments[0].IsConstant() && expression.Arguments[1].IsOddInteger() {
			if expression.Arguments[0].Solve(SOLVE_CONSTANT_PLACEHOLDER) < 0 {
				return expression.Cache.setSignalInvertible(true)
			}
		}

		return expression.Cache.setSignalInvertible(false)

	default:
		return expression.Cache.setSignalInvertible(false)
	}
}

/*
Evaluate the expression to predict if the whole expression always return
negative values. The 'applicable' returned value represents if there could be a
answer, or if its variable behavior should render the process inconclusive.

Sin, Cos, Tan and Log expressions can only be evaluated when its argument is a
constant.

Power is tricky, since f(x)^g(x) is not an algebraic function defined in the
real domain when it would represent a root (exponent different than integer).
This way, the only case it have negative values returning is negative base and
odd integer exponent.

Exponential, as e^f(x), have no negative result. Since it is actually a
sub-case of Power for euler constant as base, it already breaks the needed
negative base.

If negative returned value is true, applicable will always be true.
If negative returned value is false, applicable may be true or false.
*/
func (expression *Expression) IsNegative() (negative bool, applicable bool) {
	if expression.Cache.isCached(CACHE_IS_NEGATIVE) {
		return expression.Cache.result(CACHE_IS_NEGATIVE), expression.Cache.applicable(CACHE_IS_NEGATIVE)
	}

	if expression.IsIndefiniteness() {
		return expression.Cache.setNegative(false, false)
	}

	if expression.Type == EXPONENTIAL {
		return expression.Cache.setNegative(false, true)
	}

	if !expression.IsConstant() {
		return expression.Cache.setNegative(false, false)
	}

	var result float64 = expression.Solve(SOLVE_CONSTANT_PLACEHOLDER)
	if isApproximate(result, 0) {
		result = 0
	}

	return expression.Cache.setNegative(result < 0, true)
}

//nolint:unused
func (expression *Expression) isSineZero() bool {
	multiplier, found := piMultiplier(expression)
	if !found {
		return false
	}

	return periodicMatch(multiplier*math.Pi, 0, 2*math.Pi)
}

//nolint:unused
func (expression *Expression) isSineOne() bool {
	multiplier, found := piMultiplier(expression)
	if !found {
		return false
	}

	return periodicMatch(multiplier*math.Pi, math.Pi/2, math.Pi)
}

//nolint:unused
func (expression *Expression) isCosineZero() bool {
	multiplier, found := piMultiplier(expression)
	if !found {
		return false
	}

	return periodicMatch(multiplier*math.Pi, math.Pi/2, 2*math.Pi)
}

//nolint:unused
func (expression *Expression) isCosineOne() bool {
	multiplier, found := piMultiplier(expression)
	if !found {
		return false
	}

	return periodicMatch(multiplier*math.Pi, 0, math.Pi)
}

//nolint:unused
func (expression *Expression) isTangentZero() bool {
	multiplier, found := piMultiplier(expression)
	if !found {
		return false
	}

	return periodicMatch(multiplier*math.Pi, 0, 2*math.Pi)
}

//nolint:unused
func (expression *Expression) isTangentOne() bool {
	multiplier, found := piMultiplier(expression)
	if !found {
		return false
	}

	return periodicMatch(multiplier*math.Pi, math.Pi/4, math.Pi)
}

func (expression *Expression) isTangentIndefinite() bool {
	multiplier, found := piMultiplier(expression)
	if !found {
		return false
	}

	return math.Mod(multiplier-0.5, 1.0) == 0
}

func piMultiplier(expression *Expression) (value float64, applicable bool) {
	switch expression.Type {
	case SYMBOL:
		switch expression.Name {
		case "pi":
			return 1, true
		case "e":
			return math.E / math.Pi, true
		}

		return 0, false

	default:
		if !expression.IsConstant() {
			return 0, false
		}

		return expression.Solve(SOLVE_CONSTANT_PLACEHOLDER) / math.Pi, true
	}
}

func isApproximate(a float64, b float64) bool {
	return math.Abs(a-b) < APPROXIMATION_EPSILON
}

//nolint:unused
func periodicMatch(value float64, base float64, period float64) bool {
	var k float64 = math.Round((value - base) / period)
	var expected float64 = base + k*period

	return isApproximate(value, expected)
}
