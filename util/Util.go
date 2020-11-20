package util

const (
	Exit      = "4"
	File      = "2"
	Message   = "1"
	PROTOCOL  = "tcp"
	PORT      = ":9999"
	Separator = "|"
	Space     = ": "
)

// Removes an element from the string list
func Remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
