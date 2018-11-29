package utils

type Page struct {
	PageNo     int
	PagePrev     int
	PageNext   int
	PageSize   int
	TotalPage  int
	TotalCount int
	FirstPage  bool
	LastPage   bool
	List       interface{}
}

//参数是:总条数,第几页,总页数,movie构成的列表
func PageUtil(count int, pageNo int, pageSize int, list interface{}) Page {
	tp := count / pageSize
	if count % pageSize > 0 {
		tp = count / pageSize + 1
	}
	return Page{PageNo: pageNo,PagePrev:pageNo-1 ,PageNext:pageNo+1,PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: pageNo == 1, LastPage: pageNo == tp, List: list}
}

