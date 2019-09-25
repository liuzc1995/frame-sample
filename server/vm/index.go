package vm

//页面呈现内容
type IndexViewModel struct {
	BaseViewModel
}

type IndexViewModelOp struct{}

func (IndexViewModelOp) GetVM() IndexViewModel {
	v := IndexViewModel{}
	v.SetTitle("Homepage")
	return v
}
