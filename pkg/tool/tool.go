package tool

//Find function returns the index of the object finding in a string array.
func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}

func IsLogin(v string) bool {
	if v == "y" {
		return true
	} else {
		return false
	}
}
