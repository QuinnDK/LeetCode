package main

import (
	"fmt"
	"strings"
)

/*

 */
func simplifyPath(path string) string {
	ret := []string{}
	for _, v := range strings.Split(path, "/") {
		fmt.Println(v)
		switch v {
		case "":
			break
		case ".":
			break
		case "..":
			if l := len(ret); l > 0 {
				ret = ret[:l-1]
			}
		default:
			ret = append(ret, v)
		}
	}
	return "/" + strings.Join(ret, "/")
}
func main() {
	s := "/a/../../b/../c//.//"
	fmt.Println(simplifyPath(s))
	fmt.Println(strings.Join([]string{"xxx", "bbb", "aaa"}, "/"))
}
