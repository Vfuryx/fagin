package migrations

import (
	"fagin/database"
	_ "fagin/database"
)

var ORM = database.ORM

func init()  {

	create_user_2018_01_23_00001()
	//delete_user_col_phone_2018_01_23_00002()
}
