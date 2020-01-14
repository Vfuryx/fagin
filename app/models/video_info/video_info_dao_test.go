package video_info

import (
	"fagin/pkg/db"
	"testing"
)

func TestMain(m *testing.M) {
	// 每次运行都会清空一下表
	db.ORM.Exec("truncate video_infos")
	m.Run()
	db.Close()
}

func TestUserDaoFlow(t *testing.T) {
	t.Run("store", TestStore)
	t.Run("find", TestFindById)
	t.Run("all", TestAll)
	t.Run("query", TestQuery)
	t.Run("update", TestUpdate)
	t.Run("destroy", TestDestroy)
}

func TestAll(t *testing.T) {
	videoInfo, err := Dao().All([]string{"*"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(videoInfo)
}

func TestDestroy(t *testing.T) {
	err := Dao().Destroy(1)
	if err != nil {
		t.Fatal(err)
	}
	var videoInfo VideoInfo
	err = videoInfo.Dao().FindById(1, []string{"*"})
	if err == nil {
		t.Fatal(err)
	}
	t.Log(err)
}

func TestFindById(t *testing.T) {
	videoInfo := New()
	err := videoInfo.Dao().FindById(1, []string{"*"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(videoInfo)
}

func TestQuery(t *testing.T) {
	var videoInfos []VideoInfo
	params := map[string]interface{}{
		//"id": 2,
		"status": 1,
	}
	columns := []string{"*"}
	err := Dao().Query(params, columns, nil).Find(&videoInfos)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(videoInfos)
}

func TestStore(t *testing.T) {
	videoInfo := VideoInfo{
		AuthorId:    1,
		Title:       "测试视频标题",
		Duration:    "12:00",
		Description: "....",
		Status:      1,
	}
	err := Dao().Create(&videoInfo)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(videoInfo)
}

func TestUpdate(t *testing.T) {
	err := Dao().Update(1, map[string]interface{}{
		"Description": "测试详情",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestPaginator(t *testing.T) {
	var videoInfos []VideoInfo
	params := map[string]interface{}{
		"status": 1,
	}
	columns := []string{"*"}
	p := db.Paginator{
		CurrentPage: 1,
		PerPage:     2,
	}
	err := Dao().Query(params, columns, nil).Paginator(&videoInfos, &p)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(videoInfos)
	t.Log(p)
}
