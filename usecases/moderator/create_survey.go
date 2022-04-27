package moderator

import (
	"errors"
	"smart-surveys/entities"
)

func (m *moderator) CreateSurvey(s entities.Survey) error {
	if len(s.Questions) == 0 {
		return errors.New("no questions")
	}

	_, err := m.storage.SaveSurvey(s)
	if err != nil {
		return err
	}

	return nil
}
