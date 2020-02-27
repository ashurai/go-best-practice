package myStrFunc

import "strings"

func ConCat(st []string) string {
	s := st[0]

	for _, v := range st[1:] {
		s += " "
		s += v
	}

	return s

}

func Join(st []string) string {
	return strings.Join(st, " ")
}
