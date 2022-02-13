package view_model

type (
	Breadcrumbs struct {
		items []*BreadcrumbItem
	}

	BreadcrumbItem struct {
		Url   string
		Label string
	}
)

func NewBreadcrumbsViewModel(home string) *Breadcrumbs {
	items := []*BreadcrumbItem{
		{Url: "/", Label: home},
	}

	return &Breadcrumbs{
		items: items,
	}
}

func (vm *Breadcrumbs) AddItem(url, label string) *Breadcrumbs {
	vm.items = append(vm.items, &BreadcrumbItem{
		Url:   url,
		Label: label,
	})

	return vm
}

func (vm *Breadcrumbs) Items() []*BreadcrumbItem {
	return vm.items
}
