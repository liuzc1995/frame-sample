package vm

type BaseViewModel struct {
	Title       string //标题
	CurrentUser string //当前登录用户
}

//设置模板页标题
func (v *BaseViewModel) SetTitle(title string) {
	v.Title = title
}
