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

func (expression *Expression) IsConstant() bool {
	if expression.Cache.isCached(CACHE_IS_CONSTANT) {
		return expression.Cache.result(CACHE_IS_CONSTANT)
	}

	switch expression.Type {
	case INTEGER, FLOAT: // last leaf of the branch as a constant
		return expression.Cache.setRanResultPair(CACHE_IS_CONSTANT, true)

	case SYMBOL: // last leaf of the branch as a symbol
		switch expression.Name {
		case SYMBOL_EULER, SYMBOL_PI: // symbol that represents a constant
			return expression.Cache.setRanResultPair(CACHE_IS_CONSTANT, true)
		default: // symbol that represents a variable
			return expression.Cache.setRanResultPair(CACHE_IS_CONSTANT, false)
		}

	default:
		if len(expression.Arguments) == 0 { // reached here and have no arguments, should not even exist
			return expression.Cache.setRanResultPair(CACHE_IS_CONSTANT, false)
		}

		switch expression.Type { // special performance and evaluation cases
		case MULTIPLICATION:
			/*
				It may seem a redundant loop since a similar loop already exists forward in the code, but
				it is placed here for performance reasons.
				This Guaranteed that first of all, the IsZero() cases are evaluated first, since any 0
				render entire multiplication as 0.
				Otherwise it would unnecessarily compute IsAbsoluteOne() and IsEuler() over potentially complex
				recursion, since a latter 0 should return true anyway.
			*/
			for _, subExpression := range expression.Arguments {
				if subExpression.IsZero() {
					return expression.Cache.setRanResultPair(CACHE_IS_CONSTANT, true)
				}
			}

		case POWER:
			/*
				Power have special case needing to evaluated-compute base and exponent together, since:
					* 0^0 is mathematical mathematical indefiniteness.
					* x^0 equals to 1, a constant.
					* 0^x equals to 0, a constant.
				So, for performance reasons, this is a potential shortcut to avoid calling unnecessarily
				IsAbsoluteOne() and IsEuler().
				Even outside of performance scope, it treats 0^0 case isolated.
			*/
			if len(expression.Arguments) != 2 {
				return expression.Cache.setRanResultPair(CACHE_IS_CONSTANT, false)
			}

			var baseZero bool = expression.Arguments[0].IsZero()
			var exponentZero bool = expression.Arguments[1].IsZero()

			if baseZero && exponentZero {
				return expression.Cache.setRanResultPair(CACHE_IS_CONSTANT, false)
			}

			if baseZero || exponentZero {
				return expression.Cache.setRanResultPair(CACHE_IS_CONSTANT, true)
			}

			if expression.Arguments[0].IsAbsoluteOne() {
				return expression.Cache.setRanResultPair(CACHE_IS_CONSTANT, true)
			}
		}

		for _, subExpression := range expression.Arguments { // scraping for any non-constant
			switch expression.Type {
			case INTEGER, FLOAT: // last leaf of the branch as a constant
				continue

			case SYMBOL: // last leaf of the branch as a symbol
				switch expression.Name {
				case SYMBOL_EULER, SYMBOL_PI: // symbol that represents a constant
					continue
				default: // symbol that represents a variable
					return expression.Cache.setRanResultPair(CACHE_IS_CONSTANT, false)
				}

			case MULTIPLICATION, POWER: // IsZero() were already evaluated-computed due special cases
				if subExpression.IsAbsoluteOne() || subExpression.IsEuler() { // whole nested arguments context, simple constants
					continue
				}

			default: // no special cases for any other expression
				if subExpression.IsZero() || subExpression.IsAbsoluteOne() || subExpression.IsEuler() { // whole nested arguments context, simple constants
					continue
				}
			}

			if !subExpression.IsConstant() { // complex sub-expression that needs recursively check.
				return expression.Cache.setRanResultPair(CACHE_IS_CONSTANT, false)
			}
		}

		return expression.Cache.setRanResultPair(CACHE_IS_CONSTANT, true)
	}
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

	switch expression.Type {
	case INTEGER, FLOAT:
		return expression.Cache.setRanResultPair(
			CACHE_IS_SIGNAL_INVERTIBLE,
			expression.Value != nil && *expression.Value < 0,
		)

	case ADDITION:
		if len(expression.Arguments) == 0 {
			return expression.Cache.setRanResultPair(CACHE_IS_SIGNAL_INVERTIBLE, false)
		}

		var total float64
		for _, subExpression := range expression.Arguments {
			if !subExpression.IsConstant() {
				return expression.Cache.setRanResultPair(CACHE_IS_SIGNAL_INVERTIBLE, false)
			}

			total += subExpression.Execute(math.MaxInt)
		}

		return expression.Cache.setRanResultPair(
			CACHE_IS_SIGNAL_INVERTIBLE,
			total < 0,
		)

	case MULTIPLICATION:
		if len(expression.Arguments) == 0 {
			return expression.Cache.setRanResultPair(CACHE_IS_SIGNAL_INVERTIBLE, false)
		}

		var negative bool
		for _, argument := range expression.Arguments {
			if argument.IsSignalInvertible() {
				negative = !negative
			}
		}

		return expression.Cache.setRanResultPair(CACHE_IS_SIGNAL_INVERTIBLE, negative)

	case POWER:
		if len(expression.Arguments) != 2 {
			return expression.Cache.setRanResultPair(CACHE_IS_SIGNAL_INVERTIBLE, false)
		}

		var base *Expression = expression.Arguments[0]
		var exponent *Expression = expression.Arguments[1]

		if base.IsConstant() && exponent.IsOddInteger() {
			if base.Execute(math.MaxInt) < 0 {
				return expression.Cache.setRanResultPair(CACHE_IS_SIGNAL_INVERTIBLE, true)
			}
		}

		return expression.Cache.setRanResultPair(CACHE_IS_SIGNAL_INVERTIBLE, false)

	default:
		return expression.Cache.setRanResultPair(CACHE_IS_SIGNAL_INVERTIBLE, false)
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
*/
func (expression *Expression) IsNegative() (negative bool, applicable bool) {
	if expression.Cache.isCached(IS_NEGATIVE) {
		return expression.Cache.result(IS_NEGATIVE), expression.Cache.applicable(IS_NEGATIVE)
	}

	if expression.Type == EXPONENTIAL {
		if len(expression.Arguments) != 1 {
			return expression.Cache.setRanResultApplicableTrio(IS_NEGATIVE, false, false)
		}

		return expression.Cache.setRanResultApplicableTrio(IS_NEGATIVE, false, true)
	}

	if !expression.IsConstant() {
		return expression.Cache.setRanResultApplicableTrio(IS_NEGATIVE, false, false)
	}

	var result float64 = expression.Execute(math.MaxInt)
	if isApproximate(result, 0) {
		result = 0
	}

	return expression.Cache.setRanResultApplicableTrio(
		IS_NEGATIVE,
		result < 0,
		true,
	)
}

func (expression *Expression) IsEvenInteger() bool {
	if expression.Cache.isCached(CACHE_IS_EVEN_INTEGER) {
		return expression.Cache.result(CACHE_IS_EVEN_INTEGER)
	}

	if !expression.IsConstant() {
		return expression.Cache.setRanResultPair(CACHE_IS_EVEN_INTEGER, false)
	}

	var value float64 = expression.Execute(math.MaxInt)
	var rounded float64 = math.Round(value)

	if !isApproximate(value, rounded) {
		return expression.Cache.setRanResultPair(CACHE_IS_EVEN_INTEGER, false)
	}

	return expression.Cache.setRanResultPair(
		CACHE_IS_EVEN_INTEGER,
		int64(math.Abs(rounded))%2 == 0,
	)
}

func (expression *Expression) IsOddInteger() bool {
	if expression.Cache.isCached(CACHE_IS_ODD_INTEGER) {
		return expression.Cache.result(CACHE_IS_ODD_INTEGER)
	}

	if !expression.IsConstant() {
		return expression.Cache.setRanResultPair(CACHE_IS_ODD_INTEGER, false)
	}

	var value float64 = expression.Execute(math.MaxInt)
	var rounded float64 = math.Round(value)

	if !isApproximate(value, rounded) {
		return expression.Cache.setRanResultPair(CACHE_IS_ODD_INTEGER, false)
	}

	return expression.Cache.setRanResultPair(
		CACHE_IS_ODD_INTEGER,
		int64(math.Abs(rounded))%2 != 0,
	)
}

/*
Evaluate the expression to predict if its only possible result is 1 or -1.
If needed to evaluate a the signal as well, use it with expression.IsNegative().

Integer and Float just compare its absolute value standard approximation to 1.

Symbol is not applicable and always returns false.

Addition tries to decompose all sub-expressions into known values, then compare
the absolute result standard approximation to 1.

Multiplication checks if the only underlying value is 1, since only 1*1 equals
to 1 in multiplication cases.

Power and Exponential evaluates if 1^x, or 0^x when x != 0, the only possible
constant combinations with no mathematical indefiniteness.

Trigonometric functions check its periodic result if argument is a constant.

Logarithmic functions check is the argument is equal to its base.
*/
func (expression *Expression) IsAbsoluteOne() bool {
	if expression.Cache.isCached(CACHE_IS_ABSOLUTE_ONE) {
		return expression.Cache.result(CACHE_IS_ABSOLUTE_ONE)
	}

	if !expression.IsConstant() {
		return expression.Cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, false)
	}

	switch expression.Type {
	case INTEGER, FLOAT:
		return expression.Cache.setRanResultPair(
			CACHE_IS_ABSOLUTE_ONE,
			expression.Value != nil && isApproximate(math.Abs(*expression.Value), 1),
		)

	case ADDITION:
		if len(expression.Arguments) == 0 {
			return expression.Cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, false)
		}

		var total float64
		for _, subExpression := range expression.Arguments {
			total += subExpression.Execute(math.MaxInt)
		}

		return expression.Cache.setRanResultPair(
			CACHE_IS_ABSOLUTE_ONE,
			isApproximate(math.Abs(total), 1),
		)

	case MULTIPLICATION:
		if len(expression.Arguments) == 0 {
			return expression.Cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, false)
		}

		if expression.IsZero() {
			return expression.Cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, false)
		}

		var total float64 = 1
		for _, subExpression := range expression.Arguments {
			if subExpression.IsAbsoluteOne() {
				continue
			}

			total *= subExpression.Execute(math.MaxInt)
		}

		return expression.Cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, true)

	case POWER:
		if len(expression.Arguments) != 2 {
			return expression.Cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, false)
		}

		var base *Expression = expression.Arguments[0]
		var exponent *Expression = expression.Arguments[1]

		if base.IsAbsoluteOne() {
			return expression.Cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, true)
		}

		if exponent.IsZero() && !base.IsZero() {
			return expression.Cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, true)
		}

		return expression.Cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, false)

	case EXPONENTIAL:
		if len(expression.Arguments) != 1 {
			return expression.Cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, false)
		}

		if expression.Arguments[0].IsZero() {
			return expression.Cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, true)
		}

		return expression.Cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, false)

	case SIN:
		if len(expression.Arguments) != 1 {
			return expression.Cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, false)
		}

		return expression.Cache.setRanResultPair(
			CACHE_IS_ABSOLUTE_ONE,
			expression.Arguments[0].isSinOne(),
		)

	case COS:
		if len(expression.Arguments) != 1 {
			return expression.Cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, false)
		}

		return expression.Cache.setRanResultPair(
			CACHE_IS_ABSOLUTE_ONE,
			expression.Arguments[0].isCosineOne(),
		)

	case TAN:
		if len(expression.Arguments) != 1 {
			return expression.Cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, false)
		}

		return expression.Cache.setRanResultPair(
			CACHE_IS_ABSOLUTE_ONE,
			expression.Arguments[0].isTangentOne(),
		)

	case LOGARITHMIC:
		switch len(expression.Arguments) {
		case 1:
			return expression.Cache.setRanResultPair(
				CACHE_IS_ABSOLUTE_ONE,
				expression.Arguments[0].IsEuler(),
			)

		case 2:
			if expression.Arguments[0].IsEuler() && expression.Arguments[1].IsEuler() {
				return expression.Cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, true)
			}

			/*
				Any log that has its own base as x will result in 1.
				This means that any log may return 1 in a known constant, like log[10](10).
				But to this work in this recursive structure, a DeepEqual() method must be
					implemented to compare its base (expression.Arguments[1]) with the argument
					(expression.Arguments[0]).
				Basically, any log will be always equal to 1 if:
					* expression.Arguments[1].DeepEqual(expression.Arguments[0]) == true
			*/

			return expression.Cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, false)

		default:
			return expression.Cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, false)
		}

	default:
		return expression.Cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, false)
	}
}

/*
Evaluate the expression to predict if its only possible result is 0.
*/
func (expression *Expression) IsZero() bool {
	if expression.Cache.isCached(CACHE_IS_ZERO) {
		return expression.Cache.result(CACHE_IS_ZERO)
	}

	if !expression.IsConstant() {
		return expression.Cache.setRanResultPair(CACHE_IS_ZERO, false)
	}

	switch expression.Type {
	case INTEGER, FLOAT:
		return expression.Cache.setRanResultPair(
			CACHE_IS_ZERO,
			expression.Value != nil && isApproximate(*expression.Value, 0),
		)

	case ADDITION:
		if len(expression.Arguments) == 0 {
			return expression.Cache.setRanResultPair(CACHE_IS_ZERO, false)
		}

		var total float64
		for _, subExpression := range expression.Arguments {
			total += subExpression.Execute(math.MaxInt)
		}

		return expression.Cache.setRanResultPair(
			CACHE_IS_ZERO,
			isApproximate(total, 0),
		)

	case MULTIPLICATION:
		if len(expression.Arguments) == 0 {
			return expression.Cache.setRanResultPair(CACHE_IS_ZERO, false)
		}

		for _, subExpression := range expression.Arguments {
			if subExpression.IsZero() {
				return expression.Cache.setRanResultPair(CACHE_IS_ZERO, true)
			}
		}

		return expression.Cache.setRanResultPair(CACHE_IS_ZERO, false)

	case POWER:
		if len(expression.Arguments) != 2 {
			return expression.Cache.setRanResultPair(CACHE_IS_ZERO, false)
		}

		var base *Expression = expression.Arguments[0]
		var exponent *Expression = expression.Arguments[1]

		if base.IsZero() && !exponent.IsAbsoluteOne() {
			return expression.Cache.setRanResultPair(CACHE_IS_ZERO, true)
		}

		return expression.Cache.setRanResultPair(CACHE_IS_ZERO, false)

	case SIN:
		if len(expression.Arguments) != 1 {
			return expression.Cache.setRanResultPair(CACHE_IS_ZERO, false)
		}

		return expression.Cache.setRanResultPair(
			CACHE_IS_ZERO,
			expression.Arguments[0].isSinZero(),
		)

	case COS:
		if len(expression.Arguments) != 1 {
			return expression.Cache.setRanResultPair(CACHE_IS_ZERO, false)
		}

		return expression.Cache.setRanResultPair(
			CACHE_IS_ZERO,
			expression.Arguments[0].isCosineZero(),
		)

	case TAN:
		if len(expression.Arguments) != 1 {
			return expression.Cache.setRanResultPair(CACHE_IS_ZERO, false)
		}

		return expression.Cache.setRanResultPair(
			CACHE_IS_ZERO,
			expression.Arguments[0].isTangentZero(),
		)

	case LOGARITHMIC:
		switch len(expression.Arguments) {
		case 1, 2:
			return expression.Cache.setRanResultPair(
				CACHE_IS_ZERO,
				expression.Arguments[0].IsAbsoluteOne(),
			)
		default:
			return expression.Cache.setRanResultPair(CACHE_IS_ZERO, false)
		}

	default:
		return expression.Cache.setRanResultPair(CACHE_IS_ZERO, false)
	}
}

/*
Evaluate the expression to predict if its only possible result is e (euler number).
*/
func (expression *Expression) IsEuler() bool {
	if expression.Cache.isCached(CACHE_IS_EULER) {
		return expression.Cache.result(CACHE_IS_EULER)
	}

	if !expression.IsConstant() {
		return expression.Cache.setRanResultPair(CACHE_IS_EULER, false)
	}

	switch expression.Type {
	case SYMBOL:
		return expression.Cache.setRanResultPair(
			CACHE_IS_EULER,
			expression.Name == SYMBOL_EULER,
		)

	case FLOAT:
		if expression.Value == nil {
			return expression.Cache.setRanResultPair(CACHE_IS_EULER, false)
		}

		return expression.Cache.setRanResultPair(
			CACHE_IS_EULER,
			isApproximate(*expression.Value, math.E),
		)

	case ADDITION:
		if len(expression.Arguments) == 0 {
			return expression.Cache.setRanResultPair(CACHE_IS_EULER, false)
		}

		var total float64
		for _, subExpression := range expression.Arguments {
			total += subExpression.Execute(math.MaxInt)
		}

		return expression.Cache.setRanResultPair(
			CACHE_IS_EULER,
			isApproximate(total, math.E),
		)

	case MULTIPLICATION:
		if len(expression.Arguments) == 0 {
			return expression.Cache.setRanResultPair(CACHE_IS_EULER, false)
		}

		var total float64
		for _, subExpression := range expression.Arguments {
			total *= subExpression.Execute(math.MaxInt)
		}

		return expression.Cache.setRanResultPair(
			CACHE_IS_EULER,
			isApproximate(total, math.E),
		)

	case POWER:
		if len(expression.Arguments) != 2 {
			return expression.Cache.setRanResultPair(CACHE_IS_EULER, false)
		}

		var base *Expression = expression.Arguments[0]
		var exponent *Expression = expression.Arguments[1]

		if !base.IsEuler() {
			return expression.Cache.setRanResultPair(CACHE_IS_EULER, false)
		}

		if !exponent.IsAbsoluteOne() {
			return expression.Cache.setRanResultPair(CACHE_IS_EULER, false)
		}

		return expression.Cache.setRanResultPair(CACHE_IS_EULER, true)

	case LOGARITHMIC:
		switch len(expression.Arguments) {
		case 1:
			var subExpression *Expression = expression.Arguments[0]

			if subExpression.Type != POWER && len(subExpression.Arguments) != 2 {
				return expression.Cache.setRanResultPair(CACHE_IS_EULER, false)
			}

			var baseExpression *Expression = subExpression.Arguments[0]
			var exponentExpression *Expression = subExpression.Arguments[1]

			if baseExpression.IsEuler() && exponentExpression.IsEuler() {
				return expression.Cache.setRanResultPair(CACHE_IS_EULER, true)
			}

			return expression.Cache.setRanResultPair(CACHE_IS_EULER, true)

		case 2:
			if !expression.Arguments[1].IsEuler() {
				return expression.Cache.setRanResultPair(CACHE_IS_EULER, false)
			}

			if expression.Arguments[0].Type != POWER && len(expression.Arguments[0].Arguments) != 2 {
				return expression.Cache.setRanResultPair(CACHE_IS_EULER, false)
			}

			var baseExpression *Expression = expression.Arguments[0].Arguments[0]
			var exponentExpression *Expression = expression.Arguments[0].Arguments[1]

			if baseExpression.IsEuler() && exponentExpression.IsEuler() {
				return expression.Cache.setRanResultPair(CACHE_IS_EULER, true)
			}

			/*
				Any log that has its own base elevated to x power will result in x.
				This means that any log may return e in a known constant, like log[10](10^e).
				But to this work in this recursive structure, a DeepEqual() method should be
					implemented to compare its base (expression.Arguments[1]) with the nested
					power base (expression.Arguments[0].Arguments[0]).
				Basically, any log will be always equal to euler if both:
					* expression.Arguments[1].DeepEqual(expression.Arguments[0].Arguments[0]) == true
					* expression.Arguments[0].Arguments[1].IsEuler() == true
			*/

			return expression.Cache.setRanResultPair(CACHE_IS_EULER, false)

		default:
			return expression.Cache.setRanResultPair(CACHE_IS_EULER, false)
		}

	default:
		return expression.Cache.setRanResultPair(CACHE_IS_EULER, false)
	}
}

func (expression *Expression) isSinZero() bool {
	multiplier, found := piMultiplier(expression)
	if !found {
		return false
	}

	return periodicMatch(multiplier*math.Pi, 0, 2*math.Pi)
}

func (expression *Expression) isSinOne() bool {
	multiplier, found := piMultiplier(expression)
	if !found {
		return false
	}

	return periodicMatch(multiplier*math.Pi, math.Pi/2, math.Pi)
}

func (expression *Expression) isCosineZero() bool {
	multiplier, found := piMultiplier(expression)
	if !found {
		return false
	}

	return periodicMatch(multiplier*math.Pi, math.Pi/2, 2*math.Pi)
}

func (expression *Expression) isCosineOne() bool {
	multiplier, found := piMultiplier(expression)
	if !found {
		return false
	}

	return periodicMatch(multiplier*math.Pi, 0, math.Pi)
}

func (expression *Expression) isTangentZero() bool {
	multiplier, found := piMultiplier(expression)
	if !found {
		return false
	}

	return periodicMatch(multiplier*math.Pi, 0, 2*math.Pi)
}

func (expression *Expression) isTangentOne() bool {
	multiplier, found := piMultiplier(expression)
	if !found {
		return false
	}

	return periodicMatch(multiplier*math.Pi, math.Pi/4, math.Pi)
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

		return expression.Execute(math.MaxInt) / math.Pi, true
	}
}

func isApproximate(a float64, b float64) bool {
	return math.Abs(a-b) < APPROXIMATION_EPSILON
}

func periodicMatch(value float64, base float64, period float64) bool {
	var k float64 = math.Round((value - base) / period)
	var expected float64 = base + k*period

	return isApproximate(value, expected)
}
