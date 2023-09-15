package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"movie/models"
)

//var Validator = validator.New()

func UserValidator(c *fiber.Ctx) error {
	//var errors []*models.IError
	body := new(models.User)
	c.BodyParser(&body)
	validate := validator.New()
	errs := validate.Var(body.Name, "required")
	if errs != nil {
		return c.Status(400).SendString("name is required")
	}
	errs = validate.Var(body.Email, "required,email")
	if errs != nil {
		return c.Status(400).SendString("email is invalid")
	}
	errs = validate.Var(body.Password, "required,min=8,max=20")
	if errs != nil {
		return c.Status(400).SendString("password is less than 8 or more than 20 characters")
	}

	//
	//err := Validator.Struct(body)
	//if err != nil {
	//	for _, err := range err.(validator.ValidationErrors) {
	//		var el models.IError
	//		el.Field = err.Field()
	//		el.Tag = err.Tag()
	//		el.Value = err.Param()
	//		errors = append(errors, &el)
	//	}
	//	return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("%v", errors))
	//}
	return c.Next()
}
