package globule

import "strings"

// Match reports whether name matches the shell file name pattern.
// The pattern syntax is:
//
//	pattern:
//		{ term }
//	term:
//		'?'         matches exactly 1 any non-Separator character
//		'*'         matches 0 or more non-Separator characters
//      '**'        matches 0 or more characters, including Separators
//      '[...]'     matches any single character of the ones listed between '[' and ']',
//                  except unescaped '-' which is a character range separator
//                  and unescaped '!' or '^' which invert the match
//      '[lo-hi]'   matches any single character c where lo <= c <= hi
//		c           matches character c except '(', ''
//		'\\' c      matches character c
//
//	character-range:
//		c           matches character c (c != '\\', '-', ']')
//		'\\' c      matches character c
//		lo '-' hi   matches character c for lo <= c <= hi
//
// Match requires pattern to match all of name, not just a substring.
// The only possible returned error is ErrBadPattern, when pattern
// is malformed.
//
// On Windows, escaping is disabled. Instead, '\\' is treated as
// path separator.
//

func Match(pattern, name string) (matched bool, err error) {
	// Empty pattern or name always means no match
	if pattern == "" || name == "" {
		return false, nil
	}

	// No special chars in pattern: only exact match
	if !hasMeta(pattern) {
		strings.TrimRight("/", pattern)
	}
	return false, nil
}

func scanChunk(pattern string) (star bool, chunk, rest string) {
	return false, "", ""
}

func matchChunk(chunk, s string) (rest string, ok bool, err error) {
	return "", false, nil
}

func getEsc(chunk string) (r rune, nchunk string, err error) {
	return rune(0), "", nil
}

func Glob(pattern string) (matches []string, err error) {
	return nil, nil
}

func glob(dir, pattern string, matches []string) (m []string, e error) {
	return nil, nil
}

func hasMeta(path string) bool {
	return strings.IndexAny(path, "?*[+@") >= 0
}
