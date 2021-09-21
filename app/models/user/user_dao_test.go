package user

import (
	"fagin/pkg/db"
	"testing"
)

func TestMain(m *testing.M) {
	// 每次运行都会清空一下表
	db.ORM().Exec("truncate users")
	m.Run()
	db.Close()
}

func TestUserDaoFlow(t *testing.T) {
	t.Run("store", TestStore)
	t.Run("find", TestFindById)
	t.Run("all", TestAll)
	t.Run("query", TestQuery)
	t.Run("paginator", TestPaginator)
	t.Run("update", TestUpdate)
	t.Run("destroy", TestDestroy)
}

func TestFindById(t *testing.T) {
	user := New()
	err := user.Dao().FindById(1, []string{"*"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user)
}

func TestAll(t *testing.T) {
	users, err := NewDao().All([]string{"*"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(users)
}

func TestStore(t *testing.T) {
	user := User{
		Username: "ade",
		Password: "123",
		Status:   1,
	}

	err := NewDao().Create(&user)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user)
}

func TestUpdate(t *testing.T) {
	err := NewDao().Update(1, map[string]interface{}{
		"user_name": "",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDestroy(t *testing.T) {
	err := NewDao().Destroy(1)
	if err != nil {
		t.Fatal(err)
	}
	var user User
	err = user.Dao().FindById(1, []string{"*"})
	if err == nil {
		t.Fatal(err)
	}
	t.Log(user)
}

func TestQuery(t *testing.T) {
	var users []User
	params := map[string]interface{}{
		//"id": 2,
		//"user_name": "ade",
		"status": 1,
	}
	columns := []string{"*"}
	err := NewDao().Query(params, columns, nil).Find(&users)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(users)
}

func TestPaginator(t *testing.T) {
	var users []User
	params := map[string]interface{}{
		//"id": 2,
		//"user_name": "ade",
		"status": 1,
	}
	columns := []string{"*"}
	p := db.NewPaginator(1, 2)
	err := NewDao().Query(params, columns, nil).Paginate(&users, p)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(users)
	t.Log(p)
}
