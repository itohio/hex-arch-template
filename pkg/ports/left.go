package ports

type APIPort interface {
	SayHello(name string) string
}
