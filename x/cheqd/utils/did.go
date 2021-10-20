package utils

import (
	"regexp"
	"strings"
)

var DidForbiddenSymbolsRegexp, _ = regexp.Compile(`^[^#?&/\\]+$`)

func SplitDidUrlIntoDidAndFragment(didUrl string) (string, string) {
	fragments := strings.Split(didUrl, "#")
	return fragments[0], fragments[1]
}

func IsDidFragment(prefix string, didUrl string) bool {
	if !strings.Contains(didUrl, "#") {
		return false
	}

	if didUrl[0] == '#' {
		return true
	}

	did, _ := SplitDidUrlIntoDidAndFragment(didUrl)
	return IsDid(prefix, did)
}

func IsFullDidFragment(prefix string, didUrl string) bool {
	if !strings.Contains(didUrl, "#") {
		return false
	}

	did, _ := SplitDidUrlIntoDidAndFragment(didUrl)
	return IsDid(prefix, did)
}

func ResolveId(did string, methodId string) string {
	result := methodId

	methodDid, methodFragment := SplitDidUrlIntoDidAndFragment(methodId)
	if len(methodDid) == 0 {
		result = did + "#" + methodFragment
	}

	return result
}

func ArrayContainsNotDid(prefix string, array []string) (bool, int) {
	for i, did := range array {
		if IsNotDid(prefix, did) {
			return true, i
		}
	}

	return false, 0
}

func ArrayContainsNotDidFragment(prefix string, array []string) (bool, int) {
	for i, did := range array {
		if !IsDidFragment(prefix, did) {
			return true, i
		}
	}

	return false, 0
}

func IsNotDid(prefix string, did string) bool {
	return !IsDid(prefix, did)
}

func IsDid(prefix string, did string) bool {
	if len(did) == 0 {
		return false
	}

	if !DidForbiddenSymbolsRegexp.MatchString(did) {
		return false
	}

	return strings.HasPrefix(did, prefix)
}
