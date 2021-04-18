package testmod

import "fmt"

func TestmodFunc(name string) string {
	return fmt.Sprintf("Hi, %s, welcome to package testmod", name)
}
