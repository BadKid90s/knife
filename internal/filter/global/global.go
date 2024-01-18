package global

import (
	"gateway/internal/filter"
)

var Filters []filter.OrderedFilter

func AddFilter(order int16, f filter.Filter) {
	orderedFilter := filter.NewOrderedFilter(order, f)
	Filters = append(Filters, orderedFilter)
}
