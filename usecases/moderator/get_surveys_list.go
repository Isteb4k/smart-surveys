package moderator

import (
	"smart-surveys/entities"
)

func (m *moderator) GetSurveysList() ([]entities.Survey, error) {
	surveys, err := m.storage.GetSurveys()
	if err != nil {
		return nil, err
	}

	preparedList := make([]entities.Survey, len(surveys))
	for i, s := range surveys {
		preparedList[i].ID = s.ID
		preparedList[i].Description = s.Description
		preparedList[i].Questions = s.Questions
		preparedList[i].Conditions = s.Conditions
	}

	return preparedList, nil
}
