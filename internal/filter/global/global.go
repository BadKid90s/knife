package global

import (
	"gateway/internal/filter"
)

var Filters []filter.OrderedFilter

func AddFilter(order int, f filter.Filter) {
	orderedFilter := filter.NewOrderedFilter(order, f)
	Filters = append(Filters, orderedFilter)
}
