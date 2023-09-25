package token

import "time"

var CachedTokens map[string]time.Time

func ContainsToken(key string) bool {
	_, ok := CachedTokens[key]
	return ok
}
