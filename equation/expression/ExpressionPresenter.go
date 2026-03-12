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
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var signalOmittedMultiplication = regexp.MustCompile(`([0-9][a-zA-Z]|[a-zA-Z][0-9])`)

func algebraicString(expression *Expression) (result string) {
	if expression == nil {
		return ""
	}

	if expression.Cache.isCached(CACHE_STRING) {
		return expression.Cache.String
	}

	switch expression.Type {
	case INTEGER, FLOAT:
		result = formatNumber(expression)

	case SYMBOL:
		result = formatSymbol(expression)

	case ADDITION:
		result = formatAddition(expression)

	case MULTIPLICATION:
		result = formatMultiplication(expression)

	case POWER:
		result = formatPower(expression)

	case SIN:
		result = formatSin(expression)

	case COS:
		result = formatCos(expression)

	case TAN:
		result = formatTan(expression)

	case LOGARITHMIC:
		result = formatLog(expression)

	case EXPONENTIAL:
		result = formatExp(expression)

	default:
		result = ""
	}

	return expression.Cache.setString(result)
}

func formatNumber(expression *Expression) string {
	if expression.Value == nil {
		return ""
	}

	var value float64 = *expression.Value
	switch expression.Type {
	case INTEGER:
		return strconv.FormatInt(int64(*expression.Value), 10)

	case FLOAT:
		if isApproximate(value, math.Round(value)) {
			return strconv.FormatInt(int64(math.Round(value)), 10)
		}

		return strconv.FormatFloat(value, 'f', -1, 64)

	default:
		return ""
	}
}

func formatSymbol(expression *Expression) string {
	return expression.Name
}

func formatAddition(expression *Expression) (result string) {
	if len(expression.Arguments) == 0 {
		return ""
	}

	var subExpressions []*Expression = flattenAddition(expression)

	var constantSum float64
	var symbolCoefficient map[string]float64 = make(map[string]float64)
	var subStrings []string

	for _, subExpression := range subExpressions {
		if subExpression.IsConstant() && subExpression.Type != SYMBOL {
			constantSum += subExpression.Execute(math.MaxInt) // Using IsConstant() guarantees Execute will always return the same value.
			continue
		}

		switch subExpression.Type {
		case SYMBOL:
			symbolCoefficient[subExpression.Name] += 1
			continue

		default: // Any non constant or symbol need to be evaluated the need of parenthesis
			switch subExpression.Type {
			case MULTIPLICATION:
				var subString string = algebraicString(subExpression)

				if strings.Contains(subString, " * ") { // If multiplication contains explicit '*' it means multiple impactful terms → needs parentheses
					subStrings = append(
						subStrings,
						fmt.Sprintf("+(%s)", subString),
					)

				} else { // Single impactful term → behaves like atomic
					subStrings = append(
						subStrings,
						subString,
					)
				}

			default: // Non-special cases → behaves like atomic
				subStrings = append(
					subStrings,
					fmt.Sprintf("+%s", algebraicString(subExpression)),
				)
			}
		}
	}

	for name, coefficient := range symbolCoefficient {
		if isApproximate(coefficient, 0) {
			continue
		}

		var signal string
		if coefficient > 0 {
			signal = "+"
		} else {
			signal = "-"
		}

		if isApproximate(coefficient, 1) || isApproximate(coefficient, -1) {
			subStrings = append(subStrings, fmt.Sprintf("%s%s", signal, name))
			continue
		}

		subStrings = append(subStrings, fmt.Sprintf("%s%s%s", signal, Float(coefficient), name))
	}

	if !isApproximate(constantSum, 0) {
		var signal string
		if constantSum > 0 {
			signal = "+"
		}

		subStrings = append(subStrings, fmt.Sprintf("%s%s", signal, Float(constantSum)))
	}

	if len(subStrings) == 0 {
		return ""
	}

	for i, term := range subStrings {
		if i == 0 {
			result = strings.TrimPrefix(term, "+")
			continue
		}

		result += " " + term
	}

	return result
}

func formatMultiplication(expression *Expression) (result string) {
	if len(expression.Arguments) == 0 {
		return ""
	}

	var subExpressions []*Expression = flattenMultiplication(expression)

	for _, subExpression := range subExpressions {
		if subExpression.IsZero() { // Using IsZero() because it recursively guarantees the whole branch evaluates to 0.
			return ""
		}
	}

	var isOverallNegative bool
	var numericAccumulator float64 = 1
	var impactful []*Expression

	for _, subExpression := range subExpressions {
		if subExpression.IsSignalInvertible() { // Using IsSignalInvertible() because it recursively evaluates negativity including nested multiplications or constant sums.
			isOverallNegative = !isOverallNegative
		}

		if subExpression.IsAbsoluteOne() { // Ignore multiplicative identity. Using IsAbsoluteOne() because it already handles nested cases like x^0, exp(0), etc.
			continue
		}

		if subExpression.IsConstant() && subExpression.Type != SYMBOL { // Remove sign from numeric terms by absolute accumulation. Using IsConstant() ensures Execute is stable.
			var value float64 = subExpression.Execute(math.MaxInt)
			numericAccumulator *= math.Abs(value)
			continue
		}

		impactful = append(impactful, subExpression)
	}

	if isApproximate(numericAccumulator, 0) {
		return ""
	}

	if len(impactful) == 0 { // If only numeric
		var value float64 = numericAccumulator
		if isOverallNegative {
			value = -value
		}

		return algebraicString(Float(value))
	}

	var subStrings []string
	for _, subExpression := range impactful {
		var subString string = algebraicString(subExpression)

		if subExpression.Type == ADDITION && (len(impactful) > 1 || !isApproximate(numericAccumulator, 1) || isOverallNegative) {
			subString = fmt.Sprintf("(%s)", subString)
		}

		subStrings = append(subStrings, subString)
	}
	var inner string = strings.Join(subStrings, " * ")

	var signal string
	if isOverallNegative {
		signal = "-"
	}

	if !isApproximate(numericAccumulator, 1) { // Numeric accumulated != 1 handling
		if len(impactful) == 1 { // Single impactful element → allow compact form
			return fmt.Sprintf("%s%s%s", signal, Float(numericAccumulator), inner)
		}

		return fmt.Sprintf("%s%s * %s", signal, Float(numericAccumulator), inner) // Multiple impactful elements → explicit multiplication
	}

	if isOverallNegative && len(impactful) > 1 {
		return fmt.Sprintf("%s(%s)", signal, inner)
	}

	return fmt.Sprintf("%s%s", signal, inner)
}

func formatPower(expression *Expression) string {
	if len(expression.Arguments) != 2 {
		return ""
	}

	var base *Expression = expression.Arguments[0]
	var exponent *Expression = expression.Arguments[1]

	if base.IsZero() && exponent.IsZero() {
		return "0^0"
	}

	if base.IsZero() {
		return ""
	}

	if exponent.IsZero() {
		return "1"
	}

	if expression.IsConstant() {
		return algebraicString(Float(expression.Execute(math.MaxInt)))
	}

	if exponent.IsAbsoluteOne() {
		negative, applicable := exponent.IsNegative()
		if negative && applicable {
			return fmt.Sprintf("1/%s", algebraicString(base))
		}

		return algebraicString(base)
	}

	var baseString string = algebraicString(base)
	var exponentString string = algebraicString(exponent)

	if strings.HasPrefix(baseString, "-") {
		baseString = fmt.Sprintf("(%s)", baseString)

	} else {
		switch base.Type {
		case ADDITION, MULTIPLICATION:
			if powerNeedsParenthesis(baseString) {
				baseString = fmt.Sprintf("(%s)", baseString)
			}
		}
	}

	switch exponent.Type {
	case ADDITION:
		var needParenthesis bool = powerNeedsParenthesis(exponentString)
		fmt.Printf("exponent powerNeedsParenthesis: %t\n", needParenthesis)

		if needParenthesis {
			exponentString = fmt.Sprintf("(%s)", exponentString)
		}

	case MULTIPLICATION:
		if exponent.IsSignalInvertible() {
			exponentString = exponentString[1:]
		}

		var needParenthesis bool = powerNeedsParenthesis(exponentString)
		fmt.Printf("exponent powerNeedsParenthesis: %t\n", needParenthesis)

		if needParenthesis {
			exponentString = fmt.Sprintf("(%s)", exponentString)
		}
	}

	if exponent.IsSignalInvertible() {
		return fmt.Sprintf("1/(%s^%s)", baseString, strings.TrimPrefix(exponentString, "-"))
	}

	return fmt.Sprintf("%s^%s", baseString, exponentString)
}

func formatSin(expression *Expression) string {
	fmt.Printf("%T\n", expression)
	return ""
}

func formatCos(expression *Expression) string {
	fmt.Printf("%T\n", expression)
	return ""
}

func formatTan(expression *Expression) string {
	fmt.Printf("%T\n", expression)
	return ""
}

func formatLog(expression *Expression) string {
	fmt.Printf("%T\n", expression)
	return ""
}

func formatExp(expression *Expression) string {
	fmt.Printf("%T\n", expression)
	return ""
}

func flattenAddition(expression *Expression) (result []*Expression) {
	for _, subExpression := range expression.Arguments {
		switch subExpression.Type {
		case ADDITION:
			result = append(result, flattenAddition(subExpression)...)
		default:
			result = append(result, subExpression)
		}
	}

	return result
}

func flattenMultiplication(expression *Expression) (result []*Expression) {
	for _, subExpression := range expression.Arguments {
		switch subExpression.Type {
		case MULTIPLICATION:
			result = append(result, flattenMultiplication(subExpression)...)
		default:
			result = append(result, subExpression)
		}
	}

	return result
}

func powerNeedsParenthesis(s string) bool {
	if strings.HasPrefix(s, "-(") {
		s = s[1:]
	}

	if !strings.ContainsAny(strings.TrimPrefix(s, "-"), "+-*/") {
		if !signalOmittedMultiplication.MatchString(s) { // 1. If it's a single leaf (no spaces or operators), no parens needed
			return false
		}
	}

	if strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")") { // 2. Check if it's already fully wrapped in parentheses
		inner := s[1 : len(s)-1]

		// Walk through the string. If balance hits 0 before the end,
		// it means the outer parens aren't actually a single wrapper.
		// Example: "(x+1)*(y+2)" -> balance hits 0 at the middle.
		var balance int = 0
		var isFullyEnclosed bool = true
		for _, char := range inner {
			switch char {
			case '(':
				balance++
			case ')':
				balance--
			}

			if balance < 0 { // More closing than opening
				isFullyEnclosed = false
				break
			}
		}

		if isFullyEnclosed && balance == 0 { // If we finished the loop and balance is 0, it was fully wrapped
			return false
		}
	}

	return true // 3. Otherwise, if it has operators, it needs them
}
