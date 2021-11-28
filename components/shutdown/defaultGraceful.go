package shutdown

type DefaultGracefulShutdownComponent struct {
	Name  string
	Order int
	Proc  func()
}

func (d DefaultGracefulShutdownComponent) GetName() string {
	return d.Name
}

func (d DefaultGracefulShutdownComponent) GetOrder() int {
	return d.Order
}

func (d DefaultGracefulShutdownComponent) DoClean() {
	d.Proc()
}
