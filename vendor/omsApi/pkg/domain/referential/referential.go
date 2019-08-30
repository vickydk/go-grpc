package referential

import (
	"fmt"
	"omsApi/pkg/utl/model"
)

func (o *Referntial) List(req *omsApi.ReferentialReq) ([]omsApi.ReferentialResp, error) {
	refs, err := o.rdb.List(o.db, req)
	if err != nil {
		fmt.Printf("err: %v", err)
		return refs, nil
	}
	return refs, nil
}