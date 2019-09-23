package vm

type BaseViewModel struct {
	Title string //标题
}

//设置模板页标题
func (v *BaseViewModel) SetTitle(title string) {
	v.Title = title
}
