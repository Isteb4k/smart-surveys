package respondent

import (
	"errors"
	"smart-surveys/entities"
)

func (r *respondent) GetSurveyBlank(id uint64) (entities.SurveyBlank, error) {
	user, err := r.storage.GetUserByID(currentUserID)
	if err != nil {
		return entities.SurveyBlank{}, err
	}

	var exists bool
	for _, availableSurveyID := range user.Surveys {
		if availableSurveyID == id {
			exists = true
			break
		}
	}

	if !exists {
		return entities.SurveyBlank{}, errors.New("survey not found")
	}

	survey, err := r.storage.GetSurveyByID(id)
	if err != nil {
		return entities.SurveyBlank{}, err
	}

	var questions []entities.QuestionBlank
	for _, q := range survey.Questions {
		matched := true

		for _, c := range q.Conditions {
			matched, err = c.IsMatched(user)
			if err != nil {
				return entities.SurveyBlank{}, err
			}

			if !matched {
				break
			}
		}

		if matched {
			questions = append(questions, entities.QuestionBlank{
				ID:       q.ID,
				Type:     q.Type,
				Question: q.Question,
				Options:  q.Options,
			})
		}
	}

	s := entities.SurveyBlank{
		SurveyID:    survey.ID,
		Description: survey.Description,
		Questions:   questions,
	}

	return s, nil
}
