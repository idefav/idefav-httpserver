package shutdown

import (
	"idefav-httpserver/handler/healthz"
	"log"
	"sort"
)

type GracefulComponent interface {
	GetName() string
	GetOrder() int
	DoClean()
}

var healthStat = true

type GracefulComponents []GracefulComponent

func (g GracefulComponents) Health() bool {
	return healthStat
}

func (g GracefulComponents) Len() int {
	return len(g)
}

func (g GracefulComponents) Less(i, j int) bool {
	return g[i].GetOrder() < g[j].GetOrder()
}

func (g GracefulComponents) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

var GracefulComponentList = GracefulComponents{}

func Add(component GracefulComponent) {
	GracefulComponentList = append(GracefulComponentList, component)
}

func RunShutdownClean() {
	healthStat = false
	sort.Sort(GracefulComponentList)
	for _, component := range GracefulComponentList {
		log.Println("Graceful component executing:" + component.GetName())
		component.DoClean()
	}
}

// Graceful shutdown set health to unhealthy
func init() {
	healthz.AddHealthIndicator("GracefulShutdown", GracefulComponentList)
}
