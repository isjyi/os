package server

type PaginationQuery struct {
	PageIndex int `form:"page_index" binding:"omitempty"`
	PageSize  int `form:"page_size"  binding:"omitempty"`
}

func (s *PaginationQuery) GetPageIndex() int {
	if s.PageIndex <= 0 {
		s.PageIndex = 1
	}
	return s.PageIndex
}

func (s *PaginationQuery) GetPageSize() int {
	if s.PageSize <= 0 {
		s.PageSize = 10
	}
	return s.PageSize
}
