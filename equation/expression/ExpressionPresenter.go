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

var isAlfaNumeric *regexp.Regexp = regexp.MustCompile(`([0-9][a-zA-Z]|[a-zA-Z][0-9])`)

func algebraicString(expression *Expression) (result string) {
	if expression == nil {
		return ""
	}

	if expression.Cache.isCached(CACHE_STRING) {
		return expression.Cache.String
	}

	if expression.IsMalformedStructure() {
		return expression.Cache.setString("malformed structure")
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

	case EXPONENTIAL:
		result = formatExp(expression)

	case SINE:
		result = formatSin(expression)

	case COSINE:
		result = formatCos(expression)

	case TANGENT:
		result = formatTan(expression)

	case ARCSINE:
		result = formatAsin(expression)

	case ARCCOSINE:
		result = formatAcos(expression)

	case ARCTANGENT:
		result = formatAtan(expression)

	case HYPERBOLIC_SINE:
		result = formatSinh(expression)

	case HYPERBOLIC_COSINE:
		result = formatCosh(expression)

	case HYPERBOLIC_TANGENT:
		result = formatTanh(expression)

	case HYPERBOLIC_ARCSINE:
		result = formatAsinh(expression)

	case HYPERBOLIC_ARCCOSINE:
		result = formatAcosh(expression)

	case HYPERBOLIC_ARCTANGENT:
		result = formatAtanh(expression)

	case LOGARITHMIC:
		result = formatLog(expression)
	}

	return expression.Cache.setString(result)
}

func formatNumber(expression *Expression) string {
	if expression.IsMalformedStructure() {
		return "NaN"
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
		return "NaN"
	}
}

func formatSymbol(expression *Expression) string {
	if expression.IsMalformedStructure() {
		return "NaN"
	}

	return expression.Name
}

func formatAddition(expression *Expression) (result string) {
	if expression.IsMalformedStructure() {
		return "NaN"
	}

	var branches []*Expression = flattenAddition(expression)

	var constantSum float64
	var symbolCoefficient map[string]float64 = make(map[string]float64)
	var branchesStrings []string

	for _, branch := range branches {
		if branch.IsConstant() && branch.Type != SYMBOL {
			constantSum += branch.Solve(SOLVE_CONSTANT_PLACEHOLDER) // Using IsConstant() guarantees Solve will always return the same value.
			continue
		}

		switch branch.Type {
		case SYMBOL:
			symbolCoefficient[branch.Name] += 1
			continue

		default: // Any non constant or symbol need to be evaluated the need of parenthesis
			switch branch.Type {
			case MULTIPLICATION:
				var branchString string = algebraicString(branch)

				if strings.Contains(branchString, " * ") { // If multiplication contains explicit '*' it means multiple impactful terms → needs parentheses
					branchesStrings = append(
						branchesStrings,
						fmt.Sprintf("+(%s)", branchString),
					)

				} else { // Single impactful term → behaves like atomic
					branchesStrings = append(
						branchesStrings,
						branchString,
					)
				}

			default: // Non-special cases → behaves like atomic
				branchesStrings = append(
					branchesStrings,
					fmt.Sprintf("+%s", algebraicString(branch)),
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
			branchesStrings = append(branchesStrings, fmt.Sprintf("%s%s", signal, name))
			continue
		}

		branchesStrings = append(branchesStrings, fmt.Sprintf("%s%s%s", signal, Float(coefficient), name))
	}

	if !isApproximate(constantSum, 0) {
		var signal string
		if constantSum > 0 {
			signal = "+"
		}

		branchesStrings = append(branchesStrings, fmt.Sprintf("%s%s", signal, Float(constantSum)))
	}

	if len(branchesStrings) == 0 {
		return ""
	}

	for i, term := range branchesStrings {
		if i == 0 {
			result = strings.TrimPrefix(term, "+")
			continue
		}

		result += " " + term
	}

	return result
}

func formatMultiplication(expression *Expression) (result string) {
	if expression.IsMalformedStructure() {
		return "NaN"
	}

	var branches []*Expression = flattenMultiplication(expression)

	for _, branch := range branches {
		if branch.IsZero() { // Using IsZero() because it recursively guarantees the whole branch evaluates to 0.
			return ""
		}
	}

	var isOverallNegative bool
	var numericAccumulator float64 = 1
	var impactful []*Expression

	for _, branch := range branches {
		if branch.IsSignalInvertible() { // Using IsSignalInvertible() because it recursively evaluates negativity including nested multiplications or constant sums.
			isOverallNegative = !isOverallNegative
		}

		if branch.IsAbsoluteOne() { // Ignore multiplicative identity. Using IsAbsoluteOne() because it already handles nested cases like x^0, exp(0), etc.
			continue
		}

		if branch.IsConstant() && branch.Type != SYMBOL { // Remove sign from numeric terms by absolute accumulation. Using IsConstant() ensures Solve is stable.
			var value float64 = branch.Solve(SOLVE_CONSTANT_PLACEHOLDER)
			numericAccumulator *= math.Abs(value)
			continue
		}

		impactful = append(impactful, branch)
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

	var branchesStrings []string
	for _, branch := range impactful {
		var branchString string = algebraicString(branch)

		if branch.Type == ADDITION && (len(impactful) > 1 || !isApproximate(numericAccumulator, 1) || isOverallNegative) {
			branchString = fmt.Sprintf("(%s)", branchString)
		}

		branchesStrings = append(branchesStrings, branchString)
	}
	var inner string = strings.Join(branchesStrings, " * ")

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
	if expression.IsMalformedStructure() {
		return "NaN"
	}

	var base *Expression = expression.Arguments[0]
	var exponent *Expression = expression.Arguments[1]

	if base.IsZero() && exponent.IsZero() {
		return "0^0"
	}

	if expression.IsConstant() {
		return algebraicString(Float(expression.Solve(SOLVE_CONSTANT_PLACEHOLDER)))
	}

	negative, _ := exponent.IsNegative()
	if exponent.IsAbsoluteOne() && !negative {
		return algebraicString(base)
	}

	var baseString string = algebraicString(base)
	var exponentString string = algebraicString(exponent)

	baseString = encloseInParenthesis(baseString)
	exponentString = encloseInParenthesis(exponentString)

	return fmt.Sprintf("%s^%s", baseString, exponentString)
}

func formatExp(expression *Expression) string {
	if expression.IsMalformedStructure() {
		return "NaN"
	}

	if expression.IsConstant() {
		return algebraicString(Float(expression.Solve(SOLVE_CONSTANT_PLACEHOLDER)))
	}

	negative, _ := expression.Arguments[0].IsNegative()
	if expression.Arguments[0].IsAbsoluteOne() && !negative {
		return "e"
	}

	return fmt.Sprintf("e^%s", encloseInParenthesis(algebraicString(expression.Arguments[0])))
}

func formatSin(expression *Expression) string {
	if expression.IsMalformedStructure() {
		return "NaN"
	}

	return fmt.Sprintf("sin(%s)", algebraicString(expression.Arguments[0]))
}

func formatCos(expression *Expression) string {
	if expression.IsMalformedStructure() {
		return "NaN"
	}

	return fmt.Sprintf("cos(%s)", algebraicString(expression.Arguments[0]))
}

func formatTan(expression *Expression) string {
	if expression.IsMalformedStructure() {
		return "NaN"
	}

	return fmt.Sprintf("sin(%s)", algebraicString(expression.Arguments[0]))
}

func formatAsin(expression *Expression) string {
	if expression.IsMalformedStructure() {
		return "NaN"
	}

	return fmt.Sprintf("asin(%s)", algebraicString(expression.Arguments[0]))
}

func formatAcos(expression *Expression) string {
	if expression.IsMalformedStructure() {
		return "NaN"
	}

	return fmt.Sprintf("acos(%s)", algebraicString(expression.Arguments[0]))
}

func formatAtan(expression *Expression) string {
	if expression.IsMalformedStructure() {
		return "NaN"
	}

	return fmt.Sprintf("asin(%s)", algebraicString(expression.Arguments[0]))
}

func formatSinh(expression *Expression) string {
	if expression.IsMalformedStructure() {
		return "NaN"
	}

	return fmt.Sprintf("sinh(%s)", algebraicString(expression.Arguments[0]))
}

func formatCosh(expression *Expression) string {
	if expression.IsMalformedStructure() {
		return "NaN"
	}

	return fmt.Sprintf("cosh(%s)", algebraicString(expression.Arguments[0]))
}

func formatTanh(expression *Expression) string {
	if expression.IsMalformedStructure() {
		return "NaN"
	}

	return fmt.Sprintf("sinh(%s)", algebraicString(expression.Arguments[0]))
}

func formatAsinh(expression *Expression) string {
	if expression.IsMalformedStructure() {
		return "NaN"
	}

	return fmt.Sprintf("asinh(%s)", algebraicString(expression.Arguments[0]))
}

func formatAcosh(expression *Expression) string {
	if expression.IsMalformedStructure() {
		return "NaN"
	}

	return fmt.Sprintf("acosh(%s)", algebraicString(expression.Arguments[0]))
}

func formatAtanh(expression *Expression) string {
	if expression.IsMalformedStructure() {
		return "NaN"
	}

	return fmt.Sprintf("asinh(%s)", algebraicString(expression.Arguments[0]))
}

func formatLog(expression *Expression) string {
	if expression.IsMalformedStructure() {
		return "NaN"
	}

	if len(expression.Arguments) == 1 {
		return fmt.Sprintf("ln(%s)", algebraicString(expression.Arguments[0]))
	}

	return fmt.Sprintf("log(%s, %s)", algebraicString(expression.Arguments[1]), algebraicString(expression.Arguments[0]))
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

func encloseInParenthesis(str string) string {
	if str == "" {
		return str
	}

	if isFullyEnclosedInParenthesis(str) {
		return str
	}

	if (strings.ContainsAny(str, "+-*") || isAlfaNumeric.MatchString(str)) && !strings.HasPrefix(str, "(-") {
		return fmt.Sprintf("(%s)", str)
	}

	return str
}

func isFullyEnclosedInParenthesis(str string) bool {
	if len(str) < 2 || str[0] != '(' || str[len(str)-1] != ')' {
		return false
	}

	var depth int = 0
	for i, character := range str {
		switch character {
		case '(':
			depth++
		case ')':
			depth--
			if depth == 0 && i != len(str)-1 {
				return false
			}
		}
	}

	return depth == 0
}
