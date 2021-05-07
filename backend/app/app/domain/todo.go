package domain

type Todo struct {
	TodoID  int64
	UserID  int64
	Title   string
	Content string
	Checked bool
}
