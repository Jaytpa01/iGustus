package util

import "regexp"

var r *regexp.Regexp = regexp.MustCompile(`\s-\w{1}\s|\s--\w*\s*`)

func HasArgumentFlag(s string) bool {
	return r.MatchString(s)
}
