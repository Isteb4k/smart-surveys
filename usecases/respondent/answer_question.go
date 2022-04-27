package respondent

import "smart-surveys/entities"

func (r *respondent) AnswerQuestion(surveyID, questionID uint64, answer string, options ...string) error {
	a := entities.Answer{
		UserID:     currentUserID,
		SurveyID:   surveyID,
		QuestionID: questionID,
		Answer:     answer,
		Options:    options,
	}

	// TODO validate question id and survey id

	err := r.storage.SaveAnswer(a)
	if err != nil {
		return err
	}

	return nil
}
