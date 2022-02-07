package view_model

import "strconv"

const (
	itemsPerPage int64 = 30
	linkSpread   int64 = 4
)

type (
	PaginationVM struct {
		CurrentPage  int64
		ItemsPerPage int64
		TotalItems   int64
		LinkSpread   int64
		BaseUrl      string
	}

	PageLinkVM struct {
		IntPage  int64
		StrPage  string
		Url      string
		Current  bool
		Disabled bool
	}
)

func NewPagination(currentPage, totalItems int64) *PaginationVM {
	if currentPage < 1 {
		currentPage = 1
	}
	return &PaginationVM{
		CurrentPage:  currentPage,
		ItemsPerPage: itemsPerPage,
		TotalItems:   totalItems,
		LinkSpread:   linkSpread,
	}
}

func (p *PaginationVM) TotalPages() int64 {
	total := p.TotalItems / p.ItemsPerPage
	if p.TotalItems%p.ItemsPerPage > 0 {
		total++
	}

	return total
}

func (p *PaginationVM) TotalLinks() int64 {
	totalPages := p.TotalPages()
	totalLinks := p.LinkSpread*2 + 1
	if totalLinks > totalPages {
		return totalPages
	}
	return totalLinks
}

func (p *PaginationVM) Url(page string) string {
	return p.BaseUrl + page
}

func (p *PaginationVM) HideSideNavigation() bool {
	return p.TotalPages() <= p.TotalLinks()
}

func (p *PaginationVM) ShowSideNavigation() bool {
	return p.TotalPages() > p.TotalLinks()
}

func (p *PaginationVM) Link(page int64) PageLinkVM {
	strPage := strconv.FormatInt(page, 10)
	current := page == p.CurrentPage
	return PageLinkVM{
		IntPage: page,
		StrPage: strPage,
		Url:     p.Url(strPage),
		Current: current,
	}
}

func (p *PaginationVM) FirstLink() PageLinkVM {
	link := p.Link(1)
	if p.CurrentPage == 1 {
		link.Disabled = true
		link.Current = false
	}

	return link
}

func (p *PaginationVM) PreviousLink() PageLinkVM {
	link := p.Link(p.CurrentPage - 1)
	if p.CurrentPage == 1 {
		link.Disabled = true
	}

	return link
}

func (p *PaginationVM) NumberedLinks() []PageLinkVM {
	from := int64(1)
	totalPages := p.TotalPages()
	if p.CurrentPage >= totalPages-p.LinkSpread {
		from = totalPages - p.TotalLinks() + 1
	} else if p.CurrentPage > p.LinkSpread+1 {
		from = p.CurrentPage - p.LinkSpread
	}

	totalLinks := p.TotalLinks()
	links := make([]PageLinkVM, totalLinks)
	for i := int64(0); i < totalLinks; i++ {
		links[i] = p.Link(from + i)
	}

	return links
}

func (p *PaginationVM) NextLink() PageLinkVM {
	link := p.Link(p.CurrentPage + 1)
	if p.CurrentPage == p.TotalPages() {
		link.Disabled = true
	}

	return link
}

func (p *PaginationVM) LastLink() PageLinkVM {
	link := p.Link(p.TotalPages())
	if p.CurrentPage == p.TotalPages() {
		link.Disabled = true
		link.Current = false
	}

	return link
}

func (p *PaginationVM) Limit() int64 {
	return p.ItemsPerPage
}

func (p *PaginationVM) Offset() int64 {
	return (p.CurrentPage - 1) * p.ItemsPerPage
}
