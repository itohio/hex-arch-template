package ports

type DbPort interface {
	GetRandomGreeting() string
	GetGreetings() []string
}
