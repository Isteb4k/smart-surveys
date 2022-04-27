package admin

func (a *admin) GetSurvey(id uint64) (Survey, error) {
	survey, err := a.storage.GetSurveyByID(id)
	if err != nil {
		return Survey{}, err
	}

	s := Survey{
		ID:          survey.ID,
		Description: survey.Description,
		Condition:   survey.Conditions,
	}

	return s, nil
}
