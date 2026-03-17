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

/*
Defines the offset to be applied on the mask to align the results related to
this specific named key.
*/
type CacheKey int

/*
Cache key from ExpressionCache.Evaluated, expect only [Ran] status (1 bit).
*/
const (
	CACHE_STRING CacheKey = iota // Keep this as last entry to avoid unnecessary complexity related to next pattern of keys.
)

/*
Cache key from ExpressionCache.Evaluated, expect [Ran][Result] status pair (2
bits).
*/
const (
	CACHE_IS_MALFORMED_STRUCTURE CacheKey = 1 + CACHE_STRING + iota*2
	CACHE_IS_INDEFINITENESS
	CACHE_IS_CONSTANT
	CACHE_IS_ZERO
	CACHE_IS_ABSOLUTE_ONE
	CACHE_IS_EULER
	CACHE_IS_FRACTION
	CACHE_IS_INTEGER
	CACHE_IS_EVEN_INTEGER
	CACHE_IS_ODD_INTEGER
	CACHE_IS_SIGNAL_INVERTIBLE // Keep this as last entry to avoid unnecessary complexity related to next pattern of keys.
)

/*
Cache key from ExpressionCache.Evaluated, expect [Ran][Result][Applicable]
status trio (3 bits).
*/
const (
	CACHE_IS_NEGATIVE CacheKey = 1 + CACHE_IS_SIGNAL_INVERTIBLE + iota*3 // Keep this as last entry to avoid unnecessary complexity related to next pattern of keys.
)

/*
Masks for bitwise key-boolean cache.
*/
const (
	CACHE_MASK_RAN        uint64 = 0x1
	CACHE_MASK_RESULT     uint64 = 0x2
	CACHE_MASK_APPLICABLE uint64 = 0x4
)

type ExpressionCache struct {
	Evaluated uint64 // Binary key to check status and results of cached operations

	String string // Cached result of the String method.
}

func (cache *ExpressionCache) setString(str string) string {
	cache.setBits(CACHE_STRING, 1, CACHE_MASK_RAN)

	cache.String = str

	return str
}

func (cache *ExpressionCache) setRanResultPair(key CacheKey, result bool) bool {
	var status uint64 = CACHE_MASK_RAN
	if result {
		status |= CACHE_MASK_RESULT
	}
	cache.setBits(key, 2, status)

	return result
}

func (cache *ExpressionCache) setRanResultApplicableTrio(key CacheKey, result bool, applicable bool) (bool, bool) {
	var status uint64 = CACHE_MASK_RAN
	if result {
		status |= CACHE_MASK_RESULT
	}
	if applicable {
		status |= CACHE_MASK_APPLICABLE
	}
	cache.setBits(key, 3, status)

	return result, applicable
}

/*
Checks if a specific CacheKey has been evaluated.

The first bit relative to key is always CACHE_MASK_RAN.
*/
func (cache *ExpressionCache) isCached(key CacheKey) bool {
	return (cache.getBits(key, 1) & CACHE_MASK_RAN) != 0
}

/*
Retrieves the cached boolean result.
*/
func (cache *ExpressionCache) result(key CacheKey) (result bool) {
	return (cache.getBits(key, 2) & CACHE_MASK_RESULT) != 0
}

/*
Retrieves the cached boolean applicable.
*/
func (cache *ExpressionCache) applicable(key CacheKey) (applicable bool) {
	return (cache.getBits(key, 3) & CACHE_MASK_APPLICABLE) != 0
}

/*
Returns the raw bits at a specific key offset from Evaluated field.
*/
func (cache *ExpressionCache) getBits(key CacheKey, width int) uint64 {
	return (cache.Evaluated >> key) & ((1 << width) - 1)
}

/*
Updates a specific range of bits in the Evaluated field.
*/
func (cache *ExpressionCache) setBits(key CacheKey, width int, value uint64) {
	mask := uint64((1<<width)-1) << key
	cache.Evaluated = (cache.Evaluated & ^mask) | ((value << key) & mask)

}

/*
ClearCacheKey resets the whole cache.
*/
func (cache *ExpressionCache) clearCache() {
	cache.Evaluated = 0x0
	cache.String = ""
}

/*
As a malformed expression, any operation is false, since it has no mathematical meaning.
*/
func (cache *ExpressionCache) setMalformedStructure(result bool) bool {
	if result {
		cache.setRanResultPair(CACHE_IS_CONSTANT, false)
		cache.setRanResultPair(CACHE_IS_ZERO, false)
		cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, false)
		cache.setRanResultPair(CACHE_IS_EULER, false)
		cache.setRanResultPair(CACHE_IS_FRACTION, false)
		cache.setRanResultPair(CACHE_IS_INTEGER, false)
		cache.setRanResultPair(CACHE_IS_EVEN_INTEGER, false)
		cache.setRanResultPair(CACHE_IS_ODD_INTEGER, false)
		cache.setRanResultPair(CACHE_IS_SIGNAL_INVERTIBLE, false)

		cache.setRanResultApplicableTrio(CACHE_IS_NEGATIVE, false, false)

		cache.setRanResultPair(CACHE_IS_INDEFINITENESS, true)
	}

	return cache.setRanResultPair(CACHE_IS_MALFORMED_STRUCTURE, result)
}

/*
As a indefiniteness, any operation is false, since it has no mathematical meaning.
*/
func (cache *ExpressionCache) setIndefiniteness(result bool) bool {
	if result {
		cache.setRanResultPair(CACHE_IS_CONSTANT, false)
		cache.setRanResultPair(CACHE_IS_ZERO, false)
		cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, false)
		cache.setRanResultPair(CACHE_IS_EULER, false)
		cache.setRanResultPair(CACHE_IS_FRACTION, false)
		cache.setRanResultPair(CACHE_IS_INTEGER, false)
		cache.setRanResultPair(CACHE_IS_EVEN_INTEGER, false)
		cache.setRanResultPair(CACHE_IS_ODD_INTEGER, false)
		cache.setRanResultPair(CACHE_IS_SIGNAL_INVERTIBLE, false)

		cache.setRanResultApplicableTrio(CACHE_IS_NEGATIVE, false, false)
	}

	return cache.setRanResultPair(CACHE_IS_INDEFINITENESS, result)
}

func (cache *ExpressionCache) setConstant(result bool) bool {
	if !result {
		cache.setRanResultPair(CACHE_IS_MALFORMED_STRUCTURE, false)
		cache.setRanResultPair(CACHE_IS_INDEFINITENESS, false)
		cache.setRanResultPair(CACHE_IS_ZERO, false)
		cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, false)
		cache.setRanResultPair(CACHE_IS_EULER, false)
		cache.setRanResultPair(CACHE_IS_FRACTION, false)
		cache.setRanResultPair(CACHE_IS_INTEGER, false)
		cache.setRanResultPair(CACHE_IS_EVEN_INTEGER, false)
		cache.setRanResultPair(CACHE_IS_ODD_INTEGER, false)
	}

	return cache.setRanResultPair(CACHE_IS_CONSTANT, result)
}

func (cache *ExpressionCache) setZero(result bool) bool {
	if result {
		cache.setRanResultPair(CACHE_IS_MALFORMED_STRUCTURE, false)
		cache.setRanResultPair(CACHE_IS_INDEFINITENESS, false)
		cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, false)
		cache.setRanResultPair(CACHE_IS_EULER, false)
		cache.setRanResultPair(CACHE_IS_FRACTION, false)
		cache.setRanResultPair(CACHE_IS_ODD_INTEGER, false)

		cache.setRanResultPair(CACHE_IS_CONSTANT, true)
		cache.setRanResultPair(CACHE_IS_INTEGER, true)
		cache.setRanResultPair(CACHE_IS_EVEN_INTEGER, true)

		cache.setRanResultApplicableTrio(CACHE_IS_NEGATIVE, false, true)
	}

	return cache.setRanResultPair(CACHE_IS_ZERO, result)
}

func (cache *ExpressionCache) setAbsoluteOne(result bool) bool {
	if result {
		cache.setRanResultPair(CACHE_IS_MALFORMED_STRUCTURE, false)
		cache.setRanResultPair(CACHE_IS_INDEFINITENESS, false)
		cache.setRanResultPair(CACHE_IS_ZERO, false)
		cache.setRanResultPair(CACHE_IS_EULER, false)
		cache.setRanResultPair(CACHE_IS_FRACTION, false)
		cache.setRanResultPair(CACHE_IS_EVEN_INTEGER, false)

		cache.setRanResultPair(CACHE_IS_CONSTANT, true)
		cache.setRanResultPair(CACHE_IS_INTEGER, true)
		cache.setRanResultPair(CACHE_IS_ODD_INTEGER, true)
	}

	return cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, result)
}

func (cache *ExpressionCache) setEuler(result bool) bool {
	if result {
		cache.setRanResultPair(CACHE_IS_MALFORMED_STRUCTURE, false)
		cache.setRanResultPair(CACHE_IS_INDEFINITENESS, false)
		cache.setRanResultPair(CACHE_IS_ZERO, false)
		cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, false)
		cache.setRanResultPair(CACHE_IS_INTEGER, false)
		cache.setRanResultPair(CACHE_IS_EVEN_INTEGER, false)
		cache.setRanResultPair(CACHE_IS_ODD_INTEGER, false)

		cache.setRanResultPair(CACHE_IS_CONSTANT, true)
		cache.setRanResultPair(CACHE_IS_FRACTION, true)

		cache.setRanResultApplicableTrio(CACHE_IS_NEGATIVE, false, true)
	}

	return cache.setRanResultPair(CACHE_IS_EULER, result)
}

func (cache *ExpressionCache) setFraction(result bool) bool {
	if result {
		cache.setRanResultPair(CACHE_IS_MALFORMED_STRUCTURE, false)
		cache.setRanResultPair(CACHE_IS_INDEFINITENESS, false)
		cache.setRanResultPair(CACHE_IS_ZERO, false)
		cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, false)
		cache.setRanResultPair(CACHE_IS_INTEGER, false)
		cache.setRanResultPair(CACHE_IS_EVEN_INTEGER, false)
		cache.setRanResultPair(CACHE_IS_ODD_INTEGER, false)

		cache.setRanResultPair(CACHE_IS_CONSTANT, true)
	}

	return cache.setRanResultPair(CACHE_IS_FRACTION, result)
}

func (cache *ExpressionCache) setInteger(result bool) bool {
	if result {
		cache.setRanResultPair(CACHE_IS_MALFORMED_STRUCTURE, false)
		cache.setRanResultPair(CACHE_IS_INDEFINITENESS, false)
		cache.setRanResultPair(CACHE_IS_EULER, false)
		cache.setRanResultPair(CACHE_IS_FRACTION, false)

		cache.setRanResultPair(CACHE_IS_CONSTANT, true)
	}

	return cache.setRanResultPair(CACHE_IS_INTEGER, result)
}

func (cache *ExpressionCache) setEvenInteger(result bool) bool {
	if result {
		cache.setRanResultPair(CACHE_IS_MALFORMED_STRUCTURE, false)
		cache.setRanResultPair(CACHE_IS_INDEFINITENESS, false)
		cache.setRanResultPair(CACHE_IS_ABSOLUTE_ONE, false)
		cache.setRanResultPair(CACHE_IS_EULER, false)
		cache.setRanResultPair(CACHE_IS_FRACTION, false)
		cache.setRanResultPair(CACHE_IS_ODD_INTEGER, false)

		cache.setRanResultPair(CACHE_IS_CONSTANT, true)
		cache.setRanResultPair(CACHE_IS_INTEGER, true)
	}

	return cache.setRanResultPair(CACHE_IS_EVEN_INTEGER, result)
}

func (cache *ExpressionCache) setOddInteger(result bool) bool {
	if result {
		cache.setRanResultPair(CACHE_IS_MALFORMED_STRUCTURE, false)
		cache.setRanResultPair(CACHE_IS_INDEFINITENESS, false)
		cache.setRanResultPair(CACHE_IS_ZERO, false)
		cache.setRanResultPair(CACHE_IS_EULER, false)
		cache.setRanResultPair(CACHE_IS_FRACTION, false)
		cache.setRanResultPair(CACHE_IS_EVEN_INTEGER, false)

		cache.setRanResultPair(CACHE_IS_CONSTANT, true)
		cache.setRanResultPair(CACHE_IS_INTEGER, true)
	}

	return cache.setRanResultPair(CACHE_IS_ODD_INTEGER, result)
}

func (cache *ExpressionCache) setSignalInvertible(result bool) bool {
	if result {
		cache.setRanResultPair(CACHE_IS_MALFORMED_STRUCTURE, false)
		cache.setRanResultPair(CACHE_IS_INDEFINITENESS, false)
		cache.setRanResultPair(CACHE_IS_EULER, false)
	}

	return cache.setRanResultPair(CACHE_IS_SIGNAL_INVERTIBLE, result)
}

func (cache *ExpressionCache) setNegative(result bool, applicable bool) (bool, bool) {
	if result {
		applicable = true

		cache.setRanResultPair(CACHE_IS_MALFORMED_STRUCTURE, false)
		cache.setRanResultPair(CACHE_IS_INDEFINITENESS, false)
		cache.setRanResultPair(CACHE_IS_EULER, false)
	}

	return cache.setRanResultApplicableTrio(CACHE_IS_NEGATIVE, result, applicable)
}
