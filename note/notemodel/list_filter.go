package notemodel

import "fooddlv/common"

type ListParam struct {
	common.Paging `json:",inline"`
}
