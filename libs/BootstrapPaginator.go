package libs

import (
	"fmt"
	"strings"
)

type BootstrapPaginator struct {
	Total int64
	CurrentPage int64
	Url string
	PerPage int64
	TotalPage int64
	Appends map[string]string
}

func (bp *BootstrapPaginator) ComputeTotalPage() {
	if mod := bp.Total % bp.PerPage; mod == 0 {
		bp.TotalPage  = bp.Total / bp.PerPage
	} else {
		bp.TotalPage = bp.Total / bp.PerPage + 1
	}
}

func (bp *BootstrapPaginator) Append(params map[string]string) {
	bp.Appends = params
}

func (bp *BootstrapPaginator) BuildUrl(Page int64) string {
	query := fmt.Sprintf(bp.Url + "?page=%d", Page)
	for index, item := range bp.Appends {
		query += "&" + index + "=" + item
	}
	query = strings.TrimSuffix(query, "&")
	return query
}

func (bp *BootstrapPaginator) Render() string {
	bp.ComputeTotalPage()

	html := "<ul class=\"pagination\">"
	html += bp.RenderPrevItem()

	if bp.TotalPage <= 8 {
		for i := int64(1); i <= bp.TotalPage; i++ {
			html += bp.RenderItem(i)
		}
	} else {
		html += bp.RenderItem(1)

		startPos := bp.CurrentPage - 3

		if startPos > 2 {
			html += bp.RenderUselessItem()
		}

		for i := startPos; i < bp.CurrentPage; i++ {
			if i < 2 {
				continue
			}
			html += bp.RenderItem(i)
		}

		html += bp.RenderItem(bp.CurrentPage)

		endPos := bp.CurrentPage + 3
		for i := bp.CurrentPage + 1; i <= endPos; i++ {
			if i > bp.TotalPage {
				break
			}
			html += bp.RenderItem(i)
		}

		if endPos < bp.TotalPage {
			html += bp.RenderUselessItem()
		}

		html += bp.RenderItem(bp.TotalPage)
	}

	html += bp.RenderNextItem()
	html += "</ul>"

	return html
}

func (bp *BootstrapPaginator) RenderItem(page int64) string {
	var active, liTag string
	if page == bp.CurrentPage {
		active = "active"
	}
	url := bp.BuildUrl(page)
	liTag = fmt.Sprintf("<li class=\"%s\"><a href=\"%s\">%d</a></li>", active, url, page)
	return liTag
}

func (bp *BootstrapPaginator) RenderUselessItem() string {
	return "<li>...</li>"
}

func (bp *BootstrapPaginator) RenderNextItem() string {
	page := bp.CurrentPage + 1
	var disabled, url, liTag string
	if page > bp.TotalPage {
		disabled = "disabled"
		url = "javascript:void(0)"
	} else {
		url = bp.BuildUrl(page)
	}
	liTag = fmt.Sprintf("<li class=\"%s\"><a aria-label=\"Next\" href=\"%s\"><span aria-hidden=\"true\">&raquo;</span></a></li>", disabled, url)
	return liTag
}

func (bp *BootstrapPaginator) RenderPrevItem() string {
	page := bp.CurrentPage - 1
	var disabled, url, liTag string
	if page < 1 {
		disabled = "disabled"
		url = "javascript:void(0)"
	} else {
		url = bp.BuildUrl(page)
	}
	liTag = fmt.Sprintf("<li class=\"%s\"><a aria-label=\"Previous\" href=\"%s\"><span aria-hidden=\"true\">&laquo;</span></a></li>", disabled, url)
	return liTag
}

func (bp *BootstrapPaginator) Instance(total int64, currentPage int64, perPage int64, url string) {
	if currentPage == 0 {
		currentPage = 1
	}
	if perPage == 0 {
		perPage = 10
	}
	bp.Total = total
	bp.CurrentPage = currentPage
	bp.PerPage = perPage
	bp.Url = url
	bp.ComputeTotalPage()
}