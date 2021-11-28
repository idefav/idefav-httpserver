package warmup

import (
	"idefav-httpserver/handler/healthz"
	"log"
	"sort"
)

type Component struct {
	Name      string
	Order     int
	IgnoreErr bool
	Async     bool
	Proc      func() error
}

func (c Component) GetOrder() int {
	return c.Order
}
func (c Component) GetName() string {
	return c.Name
}

func (c Component) Exec() error {
	return c.Proc()
}

type Interface interface {
	GetName() string
	GetOrder() int
	Exec() error
}
type componentList []*Component

var warmupSuccess bool

func (c componentList) Health() bool {
	return warmupSuccess
}

func (c componentList) Len() int {
	return len(c)
}

func (c componentList) Less(i, j int) bool {
	return c[i].GetOrder() < c[j].GetOrder()
}

func (c componentList) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

var components componentList

func Add(name string, proc func() error) {
	components = append(components, &Component{Name: name, Proc: proc})
}

func RunWarmup() {
	healthz.AddHealthIndicator("Warmup", components)
	go func() {
		warmupSuccess = false
		// 从小到大排序
		sort.Sort(components)
		for _, component := range components {
			if component.Async {
				go func() {
					exeComponent(component)
				}()
			}
			exeComponent(component)
		}
		warmupSuccess = true
	}()
}

func exeComponent(component *Component) {
	err := component.Exec()
	if err != nil && !component.IgnoreErr {
		log.Fatalf("Warmup failed! %s", component.Name)
	}
}
