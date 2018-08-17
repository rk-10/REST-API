package tests

import (
	"testing"
	. "github.com/rk-10/REST-API/dao"
	. "github.com/rk-10/REST-API/config"
)

func TestDbConn(t *testing.T)  {
	var config = Config{}
	var dao = MoviesDAO{}
	dao.Server = config.Server
	dao.Database = config.Database
	if dao.Connect() == false {
		t.Fail()
	}
}

