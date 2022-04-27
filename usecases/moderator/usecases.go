package moderator

import "smart-surveys/entities"

type UseCases interface {
	CreateSurvey(s entities.Survey) error
	PublishSurvey(id uint64) error
	GetSurvey(id uint64) (entities.Survey, error)
	GetSurveysList() ([]entities.Survey, error)
	GetResponses(surveyID uint64) ([]entities.Answer, error)
}

type Storage interface {
	SaveSurvey(s entities.Survey) (entities.Survey, error)
	GetSurveyByID(id uint64) (entities.Survey, error)
	GetSurveys() ([]entities.Survey, error)
	GetResponsesBySurveyID(surveyID uint64) ([]entities.Answer, error)
	GetRespondentsByConditions(conditions []entities.Condition) ([]entities.User, error)
	UpdateUser(user entities.User) error
	UpdateSurvey(survey entities.Survey) error
}

type moderator struct {
	storage Storage
}

func NewUseCases(store Storage) UseCases {
	return &moderator{
		storage: store,
	}
}
