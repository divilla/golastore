package pagination

type Pagination struct {
	currentPage  int64
	itemsPerPage int64
	totalItems   int64
	totalPages   int64
	halfLinks    int64
	totalLinks   int64
}

func New(currentPage, totalPages, halfLinks int64) *Pagination {
	totalLinks := halfLinks*2 + 1
	if totalPages < totalLinks {
		totalLinks = totalPages
	}

	return &Pagination{
		currentPage: currentPage,
		totalPages:  totalPages,
		halfLinks:   halfLinks,
		totalLinks:  totalLinks,
	}
}

func (p *Pagination) Links() []int64 {
	if p.totalPages == p.totalLinks || p.currentPage <= p.halfLinks {
		return p.makeLinks(1, p.totalLinks)
	}
	if p.currentPage >= p.totalPages-p.halfLinks {
		return p.makeLinks(p.totalPages-p.totalLinks+1, p.totalPages)
	}

	return p.makeLinks(p.currentPage-p.halfLinks, p.currentPage+p.halfLinks)
}

func (p *Pagination) ShowPreviousNext() bool {
	return p.totalPages > p.totalLinks
}

func (p *Pagination) ShowFirstLast() bool {
	return p.totalPages > p.totalLinks+p.halfLinks
}

func (p *Pagination) IsCurrent(page int64) bool {
	return p.currentPage == page
}

func (p *Pagination) First() (int64, bool) {
	return 1, p.currentPage > p.halfLinks+1
}

func (p *Pagination) Previous() (int64, bool) {
	return p.currentPage - 1, p.currentPage != 1
}

func (p *Pagination) Next() (int64, bool) {
	return p.currentPage + 1, p.currentPage != p.totalPages
}

func (p *Pagination) Last() (int64, bool) {
	return p.totalPages, p.currentPage < p.totalPages-p.halfLinks-1
}

func (p *Pagination) TotalPages() int64 {
	return p.totalPages
}

func (p *Pagination) makeLinks(from, to int64) []int64 {
	links := make([]int64, p.totalLinks)
	diff := to - from + 1
	for i := int64(0); i < diff; i++ {
		links[i] = i + from
	}

	return links
}
