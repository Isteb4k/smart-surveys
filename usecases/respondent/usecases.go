package respondent

import "smart-surveys/entities"

const currentUserID uint64 = 2

type UseCases interface {
	GetSurveyBlanks() ([]entities.SurveyBlank, error)
	GetSurveyBlank(id uint64) (entities.SurveyBlank, error)
	AnswerQuestion(surveyID, questionID uint64, answer string, options ...string) error
}

type Storage interface {
	GetSurveyByID(id uint64) (entities.Survey, error)
	SaveAnswer(answer entities.Answer) error
	GetSurveysByUserID(userID uint64) ([]entities.Survey, error)
	GetUserByID(id uint64) (entities.User, error)
}

type respondent struct {
	storage Storage
}

func NewUseCases(store Storage) UseCases {
	return &respondent{
		storage: store,
	}
}
