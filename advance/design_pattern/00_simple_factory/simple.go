package main

import "fmt"

// API is an interface
type API interface {
	Say(name string) string
}

// NewAPI return API interface by type
func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

type hiAPI struct{}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("hi, %s", name)
}

type helloAPI struct{}

func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("hello, %s", name)
}
