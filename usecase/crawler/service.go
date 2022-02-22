package crawler

import (
	"github.com/d-jo/webcrawler/entity"
	entity "github.com/d-jo/webcrawler/entity"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) StartCrawling(sc *entity.StartCommand) *entity.GenericResponse {
	// TODO
	return &entity.GenericResponse{
		success: false,
		message: "not implemented",
	}
}

func (s *Service) StopCrawling(sc *entity.StopCommand) *entity.GenericResponse {
	// TODO
	return &entity.GenericResponse{
		success: false,
		message: "not implemented",
	}
}

func (s *Service) List(lc *entity.ListCommand) *entity.ListResponse {

}
