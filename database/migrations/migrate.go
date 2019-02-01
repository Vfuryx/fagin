package migrations

import (
	"fagin/app/models"
)

var ORM = models.ORM

func Init()  {

	create_user_2018_01_23_00001()
	//delete_user_col_phone_2018_01_23_00002()
}
