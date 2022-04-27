package respondent

import (
	"smart-surveys/entities"
)

func (r *respondent) GetSurveyBlanks() ([]entities.SurveyBlank, error) {
	user, err := r.storage.GetUserByID(currentUserID)
	if err != nil {
		return nil, err
	}

	surveys, err := r.storage.GetSurveysByUserID(user.ID)
	if err != nil {
		return nil, err
	}

	blanks := make([]entities.SurveyBlank, len(surveys))
	for i, s := range surveys {
		var questions []entities.QuestionBlank
		for _, q := range s.Questions {
			matched := true

			for _, c := range q.Conditions {
				matched, err = c.IsMatched(user)
				if err != nil {
					return nil, err
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

		blanks[i] = entities.SurveyBlank{
			Description: s.Description,
			SurveyID:    s.ID,
			Questions:   questions,
		}
	}

	return blanks, nil
}
