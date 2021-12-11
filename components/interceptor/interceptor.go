package interceptor

import (
	"net/http"
	"sort"
)

type Component struct {
	Name  string
	Order int
	Next  *Component
	Proc  func(writer http.ResponseWriter, request *http.Request) error
}

func (c Component) exec(writer http.ResponseWriter, request *http.Request) error {
	err := c.Proc(writer, request)
	if err != nil {
		return err
	}
	c.next(writer, request)
	return nil
}

func (c Component) next(writer http.ResponseWriter, request *http.Request) {
	if c.Next != nil {
		c.Next.next(writer, request)
	}
}

func (c Component) GetName() string {

	return c.Name
}

func (c Component) GetOrder() int {
	return c.Order
}

func (c Component) Exec(writer http.ResponseWriter, request *http.Request) error {
	return c.Proc(writer, request)
}

type Interface interface {
	GetName() string
	GetOrder() int
	Exec(writer http.ResponseWriter, request *http.Request) error
	next(writer http.ResponseWriter, request *http.Request)
}
type componentList []*Component

func (c componentList) Len() int {
	return len(c)
}

func (c componentList) Less(i, j int) bool {
	return c[i].Order < c[j].Order
}

func (c componentList) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

var components componentList

func Add(name string, order int, proc func(writer http.ResponseWriter, request *http.Request) error) {
	components = append(components, &Component{
		Name:  name,
		Order: order,
		Proc:  proc,
	})
	build()
}

func build() {
	sort.Sort(components)
	if components.Len() <= 0 {
		return
	}
	root := components[0]
	for i := 1; i < components.Len(); i++ {
		if root == nil {
			return
		}
		root.Next = components[i]
		root = root.Next
	}
	return
}

func Run(writer http.ResponseWriter, request *http.Request) error {

	if components.Len() <= 0 {
		return nil
	}
	root := components[0]
	return root.exec(writer, request)
}
