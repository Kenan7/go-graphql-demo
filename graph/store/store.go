package store

import (
	"context"
	"go-graphql-tutorial-demo/graph/model"
	"net/http"
)

type Store struct {
	Todos []*model.Todo
}

func NewStore() *Store {
	todos := make([]*model.Todo, 0)

	return &Store{
		Todos: todos,
	}
}

func (s *Store) AddTodo(t *model.NewTodo) error {
	s.Todos = append(s.Todos, &model.Todo{
		ID:   "1",
		Text: t.Text,
		Done: false,
		User: &model.User{
			ID:   t.UserID,
			Name: "devtopics",
		},
	})
	return nil
}

type StoreKeyType string

var StoreKey StoreKeyType = "STORE"

//  WithStore middle - inject store into context
func WithStore(store *Store, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get context and add our store to context
		requestWithStore := r.WithContext(context.WithValue(r.Context(), StoreKey, store))
		next.ServeHTTP(w, requestWithStore)
	})
}

// GetDBFromContext - retrievers store from request context
func GetStoreFromContext(ctx context.Context) *Store {
	store, ok := ctx.Value(StoreKey).(*Store)
	if !ok {
		panic("Could not retrieve store from context")
	}
	return store
}
