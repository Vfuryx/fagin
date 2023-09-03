package role

import (
	"fadmin/app/models/admin/domain/vo/role_vo"
	"github.com/go-playground/validator/v10"
)

func (r *AdminRole) SetType(t uint8) error {
	typeVo := new(role_vo.TypeVO)
	err := typeVo.Parse(t)
	if err != nil {
		return err
	}

	r.Type = t

	return nil
}

func (r *AdminRole) SetName(Name string) error {
	validate := validator.New()
	err := validate.Var(Name, "required,max=45")
	if err != nil {
		return err
	}

	r.Name = Name
	return err
}

func (r *AdminRole) SetKey(Key string) error {
	validate := validator.New()
	err := validate.Var(Key, "required,max=45")
	if err != nil {
		return err
	}

	r.Key = Key
	return err
}

func (r *AdminRole) SetRemark(Remark string) {
	r.Remark = Remark
}

func (r *AdminRole) SetSort(Sort uint) {
	r.Sort = Sort
}

func (r *AdminRole) SetStatus(Status uint8) {
	r.Status = Status
}
