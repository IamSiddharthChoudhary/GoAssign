package handler

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"assignment/internal/models"
	"assignment/internal/repository"
	"assignment/internal/service"
)

type UserHandler struct {
	repo      *repository.UserRepository
	validator *validator.Validate
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		repo:      repo,
		validator: validator.New(),
	}
}

type createUserRequest struct {
	Name string `json:"name" validate:"required"`
	DOB  string `json:"dob"  validate:"required"`
}

type updateUserRequest struct {
	Name string `json:"name" validate:"required"`
	DOB  string `json:"dob"  validate:"required"`
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req createUserRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.validator.Struct(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid dob format (YYYY-MM-DD)")
	}

	user, err := h.repo.CreateUser(c.Context(), req.Name, dob)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	user, err := h.repo.GetUserByID(c.Context(), int64(id))
	if err != nil {
		return fiber.ErrNotFound
	}

	return c.JSON(fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.DOB,
		"age":  service.CalculateAge(user.DOB),
	})
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	var req updateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.validator.Struct(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid dob format (YYYY-MM-DD)")
	}

	user, err := h.repo.UpdateUser(c.Context(), int64(id), req.Name, dob)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.repo.DeleteUser(c.Context(), int64(id)); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.repo.ListUsers(c.Context())
	if err != nil {
		return fiber.ErrInternalServerError
	}

	resp := make([]models.User, 0, len(users))
	for _, u := range users {
		resp = append(resp, models.User{
			ID:   u.ID,
			Name: u.Name,
			DOB:  u.DOB,
			Age:  service.CalculateAge(u.DOB),
		})
	}

	return c.JSON(resp)
}
