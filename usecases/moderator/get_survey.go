package moderator

import "smart-surveys/entities"

func (m *moderator) GetSurvey(id uint64) (entities.Survey, error) {
	survey, err := m.storage.GetSurveyByID(id)
	if err != nil {
		return entities.Survey{}, err
	}

	s := entities.Survey{
		ID:          survey.ID,
		Description: survey.Description,
		Questions:   survey.Questions,
		Conditions:  survey.Conditions,
	}

	return s, nil
}
