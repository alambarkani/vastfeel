package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"vastfeel-backend/internal/models"
	"vastfeel-backend/internal/services"
	"vastfeel-backend/utils"

	"github.com/gofiber/fiber/v2"
)


type UserController struct {
	Service services.UserService
	Validator *utils.XValidator
}

func NewUserController(service services.UserService, validator *utils.XValidator) *UserController{
	return &UserController{
		Service: service,
		Validator: validator,
	}
}

func (uc *UserController) CreateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(models.User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}

		if errs := uc.Validator.Validate(user); len(errs) > 0 && errs[0].Error {
			errMsgs := make([]string, 0)

			for _, err := range errs {
				errMsgs = append(errMsgs, fmt.Sprintf(
					"[%s]: '%v' | Needs to implement '%s'",
					err.FailedField,
					err.Value,
					err.Tag,
				))
			}

			return &fiber.Error{
				Code:    fiber.StatusBadRequest,
				Message: strings.Join(errMsgs, " and "),
			}
		}
		if user.Role.IsValidRole(){
			err := uc.Service.Create(user)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"success": false,
					"message": err.Error(),
				})
			}
			return c.Status(fiber.StatusCreated).JSON(fiber.Map{
				"success": true,
				"message": "User created successfully",
			})
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Invalid role",
			})
		}
	}
}

func (uc *UserController) GetAllUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := uc.Service.GetAllUser()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message":   err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message":   "Users retrieved successfully",
			"data":  users,
		})
	}
}


func (uc *UserController) GetUserById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message":   err.Error(),
			})
		}
		user, err := uc.Service.GetByID(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message":   err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message":   "User retrieved successfully",
			"data":  user,
		})
	}
}

func (uc *UserController) UpdateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(models.User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message":   err.Error(),
			})
		}

		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message":   err.Error(),
			})
		}
		user.ID = id
		err = uc.Service.Update(user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message":   err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message":   "User updated successfully",
		})
	}
}

func (uc *UserController) DeleteUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message":   err.Error(),
			})
		}
		err = uc.Service.Delete(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message":   err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message":   "User deleted successfully",
		})
	}
}
