package entities

import (
	"strconv"
)

type Condition struct {
	Field    ConditionField
	Operator ConditionOperator
	Sample   string
}

func (c Condition) IsMatched(u User) (bool, error) {
	matched := true

	switch c.Field {
	case ConditionFieldAge:
		sampleAge, err := strconv.ParseUint(c.Sample, 10, 8)
		if err != nil {
			return false, err
		}

		switch c.Operator {
		case ConditionOperatorGreat:
			if !(u.Age > uint8(sampleAge)) {
				matched = false
				break
			}
		case ConditionOperatorLess:
			if !(u.Age < uint8(sampleAge)) {
				matched = false
				break
			}
		case ConditionOperatorEqual:
			if !(u.Age == uint8(sampleAge)) {
				matched = false
				break
			}
		}

	case ConditionFieldCity:
		if u.City != c.Sample {
			matched = false
			break
		}

	case ConditionFieldGender:
		genderSample := Gender(c.Sample)

		if u.Gender != genderSample {
			matched = false
			break
		}
	}

	return matched, nil
}

type ConditionField string

const (
	ConditionFieldAge    ConditionField = "age"
	ConditionFieldCity   ConditionField = "city"
	ConditionFieldGender ConditionField = "gender"
)

func (c ConditionField) String() string {
	return string(c)
}

type ConditionOperator string

func (c ConditionOperator) String() string {
	return string(c)
}

const (
	ConditionOperatorGreat ConditionOperator = "gt"
	ConditionOperatorLess  ConditionOperator = "lt"
	ConditionOperatorEqual ConditionOperator = "eq"
)
