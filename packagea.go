package packagea

import "fmt"

func GetName(name string) string {
	msg := fmt.Sprintf("Hello %s", name)
	return msg
}
