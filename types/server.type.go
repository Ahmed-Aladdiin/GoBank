package types

type Server struct {
	ListenAddr string
	DB         Storage
}

func NewServer(listenAddr string, store Storage) *Server {
	return &Server{
		ListenAddr: listenAddr,
		DB:         store,
	}
}
