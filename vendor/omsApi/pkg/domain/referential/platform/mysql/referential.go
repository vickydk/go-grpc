package mysql

import (
	"database/sql"
	"fmt"
	omsApi "omsApi/pkg/utl/model"
	"omsApi/pkg/utl/structs"
	"strings"
)

func NewReferential() *Referential {
	return &Referential{}
}

type Referential struct {
	// bind is slice string to hold placeholder value
	bind []interface{}
}

func (u *Referential) List(db *sql.DB, req *omsApi.ReferentialReq) ([]omsApi.ReferentialResp, error) {
	var refs []omsApi.ReferentialResp
	var buf strings.Builder
	buf.WriteString("SELECT code, sub_code, field_name 'key', name FROM referential WHERE table_name = ? ")
	u.bind = []interface{}{req.TblName}
	if req.Code != 0 {
		buf.WriteString(" AND code = ?")
		u.bind = append(u.bind, req.Code)
	}
	if len(req.Key) > 0 {
		buf.WriteString(" AND field_name = ?")
		u.bind = append(u.bind, req.Key)
	}
	if len(req.Name) > 0 {
		buf.WriteString(" AND (lower(replace(name, ' ', ''))) = ?")
		u.bind = append(u.bind, strings.ToLower(strings.ReplaceAll(req.Name, " ", "")))
	}
	buf.WriteString(" order by field_name asc")
	fmt.Println("sql: ", buf.String())
	rows, err := db.Query(buf.String(), u.bind...)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var each omsApi.ReferentialResp
		structs.MergeRow(rows, &each)
		refs = append(refs, each)
	}

	return refs, nil
}
