package view_model

type (
	WebPage struct {
		MetaTitle       string
		PageTitle       string
		BreadcrumbsHome string
	}
)

func NewWebPage(metaTitle, pageTitle string) *WebPage {
	return &WebPage{
		MetaTitle:       metaTitle,
		PageTitle:       pageTitle,
		BreadcrumbsHome: "Home",
	}
}
