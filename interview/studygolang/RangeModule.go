package main

type localRange struct {
	Left  int
	Right int
	Next  *localRange
}

type RangeModule struct {
	Range localRange
}

func Constructor() RangeModule {
	return RangeModule{Range: localRange{Left: 0, Right: 0, Next: nil}}
}

func (this *RangeModule) AddRange(left int, right int) {
	if this.Range.Left == 0 && this.Range.Right == 0 {
		this.Range.Left = left
		this.Range.Right = right
	}

	//全部在左侧区域
	if this.Range.Left > right {
		newLocalRange := localRange{Left: left, Right: right, Next: &this.Range}
		this.Range = newLocalRange
		return
	}

	//部分在作则区域
	if this.Range.Left < right && this.Range.Right > right {
		this.Range.Left=left
		return
	}

	//所添加的区域已经包含
	if this.Range.Left < left && this.Range.Right > right {
		return
	}

	//部分在右侧区域
	
}

func (this *RangeModule) QueryRange(left int, right int) bool {

}

func (this *RangeModule) RemoveRange(left int, right int) {

}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */
