package admin

import "smart-surveys/entities"

type UseCases interface {
	GetSurveysList() ([]Survey, error)
	GetSurvey(id uint64) (Survey, error)
	ApproveSurvey(id uint64) error
	RejectSurvey(id uint64) error
}

type Storage interface {
	SaveSurvey(s entities.Survey) error
	GetSurveyByID(id uint64) (entities.Survey, error)
	GetSurveys() ([]entities.Survey, error)
}

type Survey struct {
	ID          uint64
	Description string
	Condition   []entities.Condition
}

type admin struct {
	storage Storage
}

func NewUseCases() UseCases {
	return &admin{}
}
