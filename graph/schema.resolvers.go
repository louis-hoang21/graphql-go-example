package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql-go-example/graph/generated"
	"graphql-go-example/graph/model"
	"strconv"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	var targetUser *model.User

	for _, user := range r.users {

		if user.ID == input.UserID {

			targetUser = user

			break

		}

	}

	if targetUser == nil {

		return nil, fmt.Errorf("user with id='%s' not found", input.UserID)

	}

	newTodo := &model.Todo{

		ID: strconv.Itoa(r.lastTodoId),

		Text: input.Text,

		Done: false,

		User: targetUser,
	}

	r.todos = append(r.todos, newTodo)

	r.lastTodoId++

	return newTodo, nil
}

/*
//query
	mutation {
	  createTodo(input:{text:"get flowers for Leela", userId:"1"}) {

		id

		text

		done

		user {

		  id

		  name

		}
	  },
	}
//response
{
  "data": {
    "createTodo": {
      "id": "3",
      "text": "get flowers for Leela",
      "done": false,
      "user": {
        "id": "1",
        "name": "fphilip"
      }
    }
  }
}
*/

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	newUser := &model.User{

		ID: strconv.Itoa(r.lastUserId),

		Name: input.Name,
	}

	r.users = append(r.users, newUser)

	r.lastUserId++

	return newUser, nil
}

/*
//query
	mutation {
	  createUser(input:{name:"brodriguez"}) {

		id

		name

	  }

	}
//response
{
  "data": {
    "createUser": {
      "id": "7",
      "name": "brodriguez"
    }
  }
}
*/

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
