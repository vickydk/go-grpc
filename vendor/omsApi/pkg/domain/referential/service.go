package referential

import (
	"database/sql"
	"omsApi/pkg/domain/referential/platform/mysql"
	"omsApi/pkg/utl/model"
)

type Service interface {
	List(*omsApi.ReferentialReq) ([]omsApi.ReferentialResp, error)
}

func New(db *sql.DB, rdb RDB) *Referntial {
	return &Referntial{db: db, rdb: rdb}
}

func Initialize(db *sql.DB) *Referntial {
	return New(db, mysql.NewReferential())
}

type Referntial struct {
	db  *sql.DB
	rdb RDB
}

type RDB interface {
	List(*sql.DB, *omsApi.ReferentialReq) ([]omsApi.ReferentialResp, error)
}
