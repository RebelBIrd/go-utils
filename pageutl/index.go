package pageutl

func GetPageCount(pageSize, total int) (sizes int) {
	if total%pageSize == 0 {
		sizes = total / pageSize
	} else {
		sizes = total/pageSize + 1
	}
	return
}
