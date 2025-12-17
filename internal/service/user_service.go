package service

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/sanuj344/ainyx-go-user-api/db/sqlc"
	"github.com/sanuj344/ainyx-go-user-api/internal/models"
	"github.com/sanuj344/ainyx-go-user-api/internal/repository"
)

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, req models.CreateUserRequest) (*models.User, error) {
	parsedDOB, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return nil, err
	}

	// Convert time.Time → pgtype.Date
	dob := pgtype.Date{
		Time:  parsedDOB,
		Valid: true,
	}

	u, err := s.repo.Queries.CreateUser(ctx, sqlc.CreateUserParams{
		Name: req.Name,
		Dob:  dob,
	})
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:   int(u.ID),
		Name: u.Name,
		DOB:  u.Dob.Time,
	}, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	u, err := s.repo.Queries.GetUserByID(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	age := CalculateAge(u.Dob.Time)

	return &models.User{
		ID:   int(u.ID),
		Name: u.Name,
		DOB:  u.Dob.Time,
		Age:  age,
	}, nil
}

func (s *UserService) ListUsers(ctx context.Context) ([]models.User, error) {
	usersDB, err := s.repo.Queries.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	users := make([]models.User, 0, len(usersDB))

	for _, u := range usersDB {
		age := CalculateAge(u.Dob.Time)

		users = append(users, models.User{
			ID:   int(u.ID),
			Name: u.Name,
			DOB:  u.Dob.Time,
			Age:  age,
		})
	}

	return users, nil
}

func (s *UserService) UpdateUser(
	ctx context.Context,
	id int,
	req models.CreateUserRequest,
) (*models.User, error) {

	parsedDOB, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return nil, err
	}

	dob := pgtype.Date{
		Time:  parsedDOB,
		Valid: true,
	}

	u, err := s.repo.Queries.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   int32(id),
		Name: req.Name,
		Dob:  dob,
	})
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:   int(u.ID),
		Name: u.Name,
		DOB:  u.Dob.Time,
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	return s.repo.Queries.DeleteUser(ctx, int32(id))
}
