package foodmodel

import "fooddlv/common"

type ListParam struct {
	common.Paging `json:",inline"`
	*ListFilter   `json:",inline"`
}

type ListFilter struct {
}

