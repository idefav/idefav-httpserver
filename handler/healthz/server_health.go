package healthz

type ServerHealth struct {
}

func (s *ServerHealth) Health() bool {
	return true
}
