package view_model

type (
	Pagination struct {
		CurrentPage  int64
		ItemsPerPage int64
		TotalItems   int64
		TotalPages   int64
		LinksSpread  int64
		TotalLinks   int64
		URLBuilder   PaginationURLBuilder
	}

	PaginationURLBuilder func(page int64) string
)

func NewPagination(currentPage, itemsPerPage, totalItems, linksSpread int64, urlBuilder PaginationURLBuilder) *Pagination {
	totalPages := totalItems / itemsPerPage
	if totalItems%itemsPerPage > 0 {
		totalPages++
	}

	totalLinks := linksSpread*2 + 1

	return &Pagination{
		CurrentPage:  currentPage,
		ItemsPerPage: itemsPerPage,
		TotalItems:   totalItems,
		TotalPages:   totalPages,
		LinksSpread:  linksSpread,
		TotalLinks:   totalLinks,
		URLBuilder:   urlBuilder,
	}
}
