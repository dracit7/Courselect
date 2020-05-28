package paginate

import (
	"github.com/dracit7/Courselect/setting"
)

// Paginator records paginating information.
type Paginator struct {
	Pages  []int
	First  bool
	Last   bool
	Prev   int
	Next   int
	Lastpg int
	URL    string
}

// MakePaginator returns a paginator which could be
// passed to the frontend.
func MakePaginator(url string, curpage int, total int) *Paginator {
	lastpage := total/setting.UI.Pagesize + 1
	pages := &Paginator{
		Pages:  make([]int, 0),
		First:  false,
		Last:   false,
		Prev:   curpage - 1,
		Next:   curpage + 1,
		Lastpg: lastpage,
		URL:    url,
	}

	// Is this page the first/last page?
	if curpage == 1 {
		pages.First = true
	}
	if curpage >= lastpage {
		pages.Last = true
	}

	// Generate all pages.
	for i := 0; i < setting.UI.Pagenum && i+curpage <= lastpage; i++ {
		pages.Pages = append(pages.Pages, i+curpage)
	}
	return pages
}
