package services

type PingService struct {
}

func (service *PingService) Ping() string {
	return "Pong"
}
