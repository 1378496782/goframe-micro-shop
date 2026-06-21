package search

import "shop-goframe-micro-service-refacotor/app/gateway-h5/api/search"

func NewV1() search.ISearchV1 {
	return &ControllerV1{}
}
