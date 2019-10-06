package services

import "time"

type IndexService struct {
}

func (service *IndexService) Index() int64 {
	return time.Now().Unix()
}
