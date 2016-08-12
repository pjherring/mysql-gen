package util

type StringParts []string

func (s StringParts) Each(f func(s string) string) {
	for i, p := range s {
		s[i] = f(p)
	}
}
