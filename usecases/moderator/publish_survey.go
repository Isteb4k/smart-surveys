package moderator

import (
	"errors"
	"smart-surveys/entities"
)

func (m *moderator) PublishSurvey(id uint64) error {
	survey, err := m.storage.GetSurveyByID(id)
	if err != nil {
		return err
	}

	if survey.Status != entities.SurveyStatusNew {
		return errors.New("unexpected survey status")
	}

	survey.Status = entities.SurveyStatusPublished

	// TODO move to admin use case
	survey.Status = entities.SurveyStatusApproved

	var respondents []entities.User
	respondents, err = m.storage.GetRespondentsByConditions(survey.Conditions)
	if err != nil {
		return err
	}

	for _, r := range respondents {
		r.Surveys = append(r.Surveys, survey.ID)

		err = m.storage.UpdateUser(r)
		if err != nil {
			return err
		}
	}

	err = m.storage.UpdateSurvey(survey)
	if err != nil {
		return err
	}

	return nil
}
