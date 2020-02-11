package simple_factory

import "fmt"

type API interface {
	Say(name string) string
}


type hiApi struct {}


func (hiApi *hiApi) Say (name string) string {
	return fmt.Sprintf("Hi, %s", name)
}


type helloApi struct {}

func (hiApi *helloApi) Say (name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func NewApi(t int) API {
	switch  {
	case t == 1:
		return &hiApi{}
	case t == 2:
		return &helloApi{}
	}
	return nil
}
