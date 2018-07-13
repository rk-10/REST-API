package tests

import (
	"testing"
	. "github.com/rk-10/REST-API/dao"
)

func TestDbConn(t *testing.T)  {
	var dao = MoviesDAO{}
	dao.Server = "localhost:27017"
	dao.Database = "test"
	if dao.Connect() == false {
		t.Fail()
	}
}


