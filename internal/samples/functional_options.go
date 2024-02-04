package samples

/*
Sample code for functional options pattern, adapted from:

- https://youtu.be/MDy7JQN5MN4?si=6kiXru1fr-uMB7Iq
*/

type Server struct {
	Options ServerOptions
}

type ServerOptions struct {
	Id             string
	MaxConnections int
	Tls            bool
}

type ServerOption func(*ServerOptions)

func defaultOptions() ServerOptions {
	return ServerOptions{
		Id:             "default",
		MaxConnections: 10,
		Tls:            false,
	}
}

func WithId(id string) ServerOption {
	return func(s *ServerOptions) { s.Id = id }
}

func WithMaxConnections(maxConnections int) ServerOption {
	return func(s *ServerOptions) { s.MaxConnections = maxConnections }
}

func WithTls(tls bool) ServerOption {
	return func(s *ServerOptions) { s.Tls = tls }
}

func NewServer(options ...ServerOption) *Server {
	serverOptions := defaultOptions()
	for _, option := range options {
		option(&serverOptions)
	}
	return &Server{Options: serverOptions}
}
