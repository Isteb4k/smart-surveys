package storage

import (
	"errors"
	"smart-surveys/entities"
	"strconv"
	"time"
)

type Client struct {
	lastSurveyID uint64
	lastAnswerID uint64
	surveys      []entities.Survey
	answers      []entities.Answer
	users        []entities.User
	survey       []entities.User
}

func NewClient() *Client {
	users := []entities.User{
		{
			ID:        0,
			FirstName: "First Name",
			LastName:  "First Name",
			Email:     "first@mail.ru",
			Age:       18,
			City:      "Tomsk",
			Gender:    entities.Male,
			CreatedAt: time.Now(),
		},
		{
			ID:        1,
			FirstName: "Second Name",
			LastName:  "Second Name",
			Email:     "second@mail.ru",
			Age:       29,
			City:      "Tomsk",
			Gender:    entities.Male,
			CreatedAt: time.Now(),
		},
		{
			ID:        2,
			FirstName: "Third Name",
			LastName:  "Third Name",
			Email:     "third@mail.ru",
			Age:       26,
			City:      "Tomsk",
			Gender:    entities.Male,
			CreatedAt: time.Now(),
		},
	}

	return &Client{
		lastSurveyID: 0,
		lastAnswerID: 0,
		surveys:      nil,
		users:        users,
	}
}

func (c *Client) GetSurveys() ([]entities.Survey, error) {
	return c.surveys, nil
}

func (c *Client) GetSurveyByID(id uint64) (entities.Survey, error) {
	for _, s := range c.surveys {
		if s.ID == id {
			return s, nil
		}
	}

	return entities.Survey{}, errors.New("survey not found")
}

func (c *Client) SaveSurvey(s entities.Survey) (entities.Survey, error) {
	s.ID = c.lastSurveyID + 1
	c.lastSurveyID++

	var questionID uint64 = 0
	for i, _ := range s.Questions {
		s.Questions[i].ID = questionID
		questionID++
	}

	c.surveys = append(c.surveys, s)

	return s, nil
}

func (c *Client) SaveAnswer(answer entities.Answer) error {
	answer.ID = c.lastAnswerID + 1
	c.lastAnswerID++
	c.answers = append(c.answers, answer)

	return nil
}

func (c *Client) GetResponsesBySurveyID(surveyID uint64) ([]entities.Answer, error) {
	var answers []entities.Answer
	for _, a := range c.answers {
		if a.SurveyID == surveyID {
			answers = append(answers, a)
		}
	}

	return answers, nil
}

func (c *Client) GetRespondentsByConditions(conditions []entities.Condition) ([]entities.User, error) {
	var respondents []entities.User

	for _, u := range c.users {
		matched := true
		for _, cond := range conditions {
			switch cond.Field {
			case entities.ConditionFieldAge:
				sampleAge, err := strconv.ParseUint(cond.Sample, 10, 8)
				if err != nil {
					return nil, err
				}

				switch cond.Operator {
				case entities.ConditionOperatorGreat:
					if !(u.Age > uint8(sampleAge)) {
						matched = false
						break
					}
				case entities.ConditionOperatorLess:
					if !(u.Age < uint8(sampleAge)) {
						matched = false
						break
					}
				case entities.ConditionOperatorEqual:
					if !(u.Age == uint8(sampleAge)) {
						matched = false
						break
					}
				}

			case entities.ConditionFieldCity:
				if u.City != cond.Sample {
					matched = false
					break
				}

			case entities.ConditionFieldGender:
				genderSample := entities.Gender(cond.Sample)

				if u.Gender != genderSample {
					matched = false
					break
				}
			}
		}

		if matched {
			respondents = append(respondents, u)
		}
	}

	return respondents, nil
}

func (c *Client) UpdateUser(user entities.User) error {
	for i, u := range c.users {
		if u.ID == user.ID {
			c.users[i] = user
			return nil
		}
	}

	return errors.New("user not found to update")
}

func (c *Client) UpdateSurvey(survey entities.Survey) error {
	for i, s := range c.surveys {
		if s.ID == survey.ID {
			c.surveys[i] = survey
			return nil
		}
	}

	return errors.New("survey not found to update")
}

func (c *Client) GetUserByID(id uint64) (entities.User, error) {
	for _, u := range c.users {
		if u.ID == id {
			return u, nil
		}
	}

	return entities.User{}, errors.New("user not found")
}

func (c *Client) GetSurveysByUserID(id uint64) ([]entities.Survey, error) {
	var user *entities.User
	for _, u := range c.users {
		if u.ID == id {
			user = &u
			break
		}
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	var surveys []entities.Survey

	for _, s := range c.surveys {
		for _, surveyID := range user.Surveys {
			if s.ID == surveyID {
				surveys = append(surveys, s)
				break
			}
		}
	}

	return surveys, nil
}
