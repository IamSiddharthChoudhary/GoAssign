package repository

import (
	"context"
	"time"

	"assignment/internal/db"
	"assignment/internal/models"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	q *db.Queries
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		q: db.New(pool),
	}
}

func (r *UserRepository) CreateUser(
	ctx context.Context,
	name string,
	dob time.Time,
) (*models.User, error) {

	u, err := r.q.CreateUser(ctx, db.CreateUserParams{
		Name: name,
		Dob: pgtype.Date{
			Time:  dob,
			Valid: true,
		},
	})
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:   int64(u.ID),
		Name: u.Name,
		DOB:  u.Dob.Time,
	}, nil
}

func (r *UserRepository) GetUserByID(
	ctx context.Context,
	id int64,
) (*models.User, error) {

	u, err := r.q.GetUserByID(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:   int64(u.ID),
		Name: u.Name,
		DOB:  u.Dob.Time,
	}, nil
}

func (r *UserRepository) UpdateUser(
	ctx context.Context,
	id int64,
	name string,
	dob time.Time,
) (*models.User, error) {

	u, err := r.q.UpdateUser(ctx, db.UpdateUserParams{
		ID:   int32(id),
		Name: name,
		Dob: pgtype.Date{
			Time:  dob,
			Valid: true,
		},
	})
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:   int64(u.ID),
		Name: u.Name,
		DOB:  u.Dob.Time,
	}, nil
}

func (r *UserRepository) DeleteUser(
	ctx context.Context,
	id int64,
) error {
	return r.q.DeleteUser(ctx, int32(id))
}

func (r *UserRepository) ListUsers(
	ctx context.Context,
) ([]models.User, error) {

	users, err := r.q.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]models.User, 0, len(users))
	for _, u := range users {
		result = append(result, models.User{
			ID:   int64(u.ID),
			Name: u.Name,
			DOB:  u.Dob.Time,
		})
	}

	return result, nil
}
