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
	CACHE_IS_CONSTANT CacheKey = 1 + CACHE_STRING + iota*2
	CACHE_IS_ZERO
	CACHE_IS_ABSOLUTE_ONE
	CACHE_IS_EULER
	CACHE_IS_EVEN_INTEGER
	CACHE_IS_ODD_INTEGER
	CACHE_IS_SIGNAL_INVERTIBLE // Keep this as last entry to avoid unnecessary complexity related to next pattern of keys.
)

/*
Cache key from ExpressionCache.Evaluated, expect [Ran][Result][Applicable]
status trio (3 bits).
*/
const (
	IS_NEGATIVE CacheKey = 1 + CACHE_IS_SIGNAL_INVERTIBLE + iota*3 // Keep this as last entry to avoid unnecessary complexity related to next pattern of keys.
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
//nolint:unused
func (cache *ExpressionCache) clearCache() {
	cache.Evaluated = 0x0
	cache.String = ""
}
