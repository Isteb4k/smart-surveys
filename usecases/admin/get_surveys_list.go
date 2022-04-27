package admin

func (a *admin) GetSurveysList() ([]Survey, error) {
	surveys, err := a.storage.GetSurveys()
	if err != nil {
		return nil, err
	}

	preparedList := make([]Survey, len(surveys))
	for i, s := range surveys {
		preparedList[i].ID = s.ID
		preparedList[i].Description = s.Description
		preparedList[i].Condition = s.Conditions
	}

	return preparedList, nil
}
