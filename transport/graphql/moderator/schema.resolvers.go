package moderator

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"smart-surveys/entities"
	"smart-surveys/transport/graphql/moderator/generated"
	"smart-surveys/transport/graphql/moderator/model"
)

func (r *mutationResolver) PublishSurvey(ctx context.Context, surveyID uint64) (*string, error) {
	err := r.UseCases.PublishSurvey(surveyID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *mutationResolver) CreateSurvey(ctx context.Context, input model.CreateSurveyInput) (*string, error) {
	questions := make([]entities.Question, len(input.Questions))
	for i, q := range input.Questions {
		conditions := make([]entities.Condition, len(input.Conditions))
		for j, cond := range q.Conditions {
			conditions[j] = entities.Condition{
				Field:    entities.ConditionField(cond.Field),
				Operator: entities.ConditionOperator(cond.Operator),
				Sample:   cond.Sample,
			}
		}

		questions[i] = entities.Question{
			Question:   q.Question,
			Type:       q.Type,
			Options:    q.Options,
			Conditions: conditions,
		}
	}

	conditions := make([]entities.Condition, len(input.Conditions))
	for i, cond := range input.Conditions {
		conditions[i] = entities.Condition{
			Field:    entities.ConditionField(cond.Field),
			Operator: entities.ConditionOperator(cond.Operator),
			Sample:   cond.Sample,
		}
	}

	e := entities.Survey{
		Description: input.Description,
		Questions:   questions,
		Conditions:  conditions,
	}

	err := r.UseCases.CreateSurvey(e)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *queryResolver) Surveys(ctx context.Context) ([]*model.Survey, error) {
	surveys, err := r.UseCases.GetSurveysList()
	if err != nil {
		return nil, err
	}

	sList := make([]*model.Survey, len(surveys))
	for i, s := range surveys {
		sList[i] = r.convertSurveyToResolver(s)
	}

	return sList, nil
}

func (r *queryResolver) Survey(ctx context.Context, id uint64) (*model.Survey, error) {
	s, err := r.UseCases.GetSurvey(id)
	if err != nil {
		return nil, err
	}

	return r.convertSurveyToResolver(s), nil
}

func (r *queryResolver) Responses(ctx context.Context, surveyID uint64) ([]*model.Answer, error) {
	responses, err := r.UseCases.GetResponses(surveyID)
	if err != nil {
		return nil, err
	}

	models := make([]*model.Answer, len(responses))
	for i, response := range responses {
		models[i] = &model.Answer{
			ID:         response.ID,
			SurveyID:   response.SurveyID,
			QuestionID: response.QuestionID,
			UserID:     response.UserID,
			Answer:     &response.Answer,
			Options:    response.Options,
		}
	}

	return models, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *Resolver) convertSurveyToResolver(s entities.Survey) *model.Survey {
	questions := make([]*model.Question, len(s.Questions))
	for i, q := range s.Questions {
		conditions := make([]*model.Condition, len(q.Conditions))
		for j, c := range q.Conditions {
			conditions[j] = &model.Condition{
				Field:    c.Field.String(),
				Operator: c.Operator.String(),
				Sample:   c.Sample,
			}
		}

		questions[i] = &model.Question{
			ID:         q.ID,
			Question:   q.Question,
			Type:       q.Type,
			Options:    q.Options,
			Conditions: conditions,
		}
	}

	conditions := make([]*model.Condition, len(s.Conditions))
	for i, c := range s.Conditions {
		conditions[i] = &model.Condition{
			Field:    c.Field.String(),
			Operator: c.Operator.String(),
			Sample:   c.Sample,
		}
	}

	return &model.Survey{
		ID:          s.ID,
		Status:      model.SurveyStatus(s.Status),
		Description: s.Description,
		Questions:   questions,
		Conditions:  conditions,
	}
}
