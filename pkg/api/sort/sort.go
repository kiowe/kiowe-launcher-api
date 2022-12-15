package sort

import (
	"fmt"
	"strings"
)

type Option struct {
	Column string
	Order  string
}

func GetSortOptions(queryParams map[string]string) (sortOption Option) {

	//sortOption := SortOption{}
	//for key, value := range queryParams {
	//	if key == "sort_by" || key == "sort_order" {
	//		sortOptions = append(sortOptions, SortOption{Column: key, Order: value})
	//	}
	//}
	if col, _ := queryParams["sort_by"]; col != "" {
		if order, _ := queryParams["sort_order"]; order == "asc" || order == "desc" {
			sortOption.Column = col
			sortOption.Order = strings.ToUpper(order)
		} else {
			sortOption.Column = col
			sortOption.Order = "ASC"
		}
	}

	return
}

func EnrichQueryWithSort(query string, option Option) string {
	if option.Column != "" {
		query = fmt.Sprintf("%s ORDER BY %s %s", query, option.Column, option.Order)
	}
	return query
}
