package respondent

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"smart-surveys/entities"
	"smart-surveys/transport/graphql/respondent/generated"
	"smart-surveys/transport/graphql/respondent/model"
)

func (r *mutationResolver) Answer(ctx context.Context, input *model.AnswerInput) (*string, error) {
	var answer string
	if input.Answer != nil {
		answer = *input.Answer
	}

	err := r.UseCases.AnswerQuestion(input.SurveyID, input.QuestionID, answer, input.Options...)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *queryResolver) Surveys(ctx context.Context) ([]*model.SurveyBlank, error) {
	surveys, err := r.UseCases.GetSurveyBlanks()
	if err != nil {
		return nil, err
	}

	sList := make([]*model.SurveyBlank, len(surveys))
	for i, s := range surveys {
		sList[i] = r.convertSurveyToResolver(s)
	}

	return sList, nil
}

func (r *queryResolver) Survey(ctx context.Context, id uint64) (*model.SurveyBlank, error) {
	surveyBlank, err := r.UseCases.GetSurveyBlank(id)
	if err != nil {
		return nil, err
	}

	return r.convertSurveyToResolver(surveyBlank), nil
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
func (r *Resolver) convertSurveyToResolver(s entities.SurveyBlank) *model.SurveyBlank {
	questions := make([]*model.QuestionBlank, len(s.Questions))

	for i, q := range s.Questions {
		questions[i] = &model.QuestionBlank{
			ID:       q.ID,
			Question: q.Question,
			Type:     q.Type,
			Options:  q.Options,
		}
	}

	return &model.SurveyBlank{
		SurveyID:    s.SurveyID,
		Description: s.Description,
		Questions:   questions,
	}
}
