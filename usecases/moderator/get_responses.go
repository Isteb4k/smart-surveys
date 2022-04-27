package moderator

import "smart-surveys/entities"

func (m *moderator) GetResponses(surveyID uint64) ([]entities.Answer, error) {
	answers, err := m.storage.GetResponsesBySurveyID(surveyID)
	if err != nil {
		return nil, err
	}

	return answers, nil
}
