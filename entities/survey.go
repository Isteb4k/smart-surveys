package entities

type Survey struct {
	ID          uint64
	Description string
	Status      SurveyStatus
	Questions   []Question
	Conditions  []Condition
}

type SurveyStatus uint8

const (
	SurveyStatusNew SurveyStatus = iota
	SurveyStatusPublished
	SurveyStatusApproved
	SurveyStatusRejected
	SurveyStatusClosed
)

func (s SurveyStatus) String() string {
	return string(s)
}

type Question struct {
	ID         uint64
	Type       string
	Question   string
	Options    []string
	Conditions []Condition
}

type Answer struct {
	ID         uint64
	UserID     uint64
	SurveyID   uint64
	QuestionID uint64
	Answer     string
	Options    []string
}
