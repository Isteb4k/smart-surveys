package entities

type SurveyBlank struct {
	Description string
	SurveyID    uint64
	Questions   []QuestionBlank
}

type QuestionBlank struct {
	ID       uint64
	Type     string
	Question string
	Options  []string
}
