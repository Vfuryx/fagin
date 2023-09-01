package user

import "github.com/go-playground/validator/v10"

func (user *AdminUser) setUsername(name string) error {
	validate := validator.New()
	err := validate.Var(name, "required,max=45")
	if err != nil {
		return err
	}

	user.Username = name

	return err
}

func (user *AdminUser) setNickName(name string) error {
	validate := validator.New()
	err := validate.Var(name, "required,max=45")
	if err != nil {
		return err
	}

	user.Username = name

	return err
}
