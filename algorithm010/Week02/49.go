func groupAnagrams(strs []string) [][]string {
	// var sumMap map[int]int // key is row num,value is sum
	sumMap:=make(map[string]int)
	var rowCnt int=0
	var res [][]string
	for _,v:=range strs{
		cur:=orderString(v)
		if row,ok:=sumMap[cur];ok{
			res[row]=append(res[row],v)
			continue
		}else{
			sumMap[cur]=rowCnt
			if len(res)!=rowCnt+1{
				res=append(res,[]string{})
			}
			res[rowCnt]=append(res[rowCnt],v)
			rowCnt++
		}
	}
	return res
}

func orderString(s string)string{

	rs:=[]rune(s)

	sort.Slice(rs,func(a,b int)bool{return rs[a]<rs[b]})
	return string(rs)
}
