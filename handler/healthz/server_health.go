package healthz

type ServerHealth struct {
}

func (s *ServerHealth) health() bool {
	return true
}
