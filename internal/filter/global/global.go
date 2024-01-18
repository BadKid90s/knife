package global

import (
	"gateway/internal/filter"
)

var Filters []filter.OrderedFilter

func AddFilter(name string, order int16, f filter.Filter) {
	orderedFilter := filter.NewOrderedFilter(name, order, f)
	Filters = append(Filters, orderedFilter)
}
