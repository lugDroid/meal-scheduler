package viewmodel

type DeleteViewModel struct {
	Title      string
	Active     string
	Content    string
	Name       string
	ReturnPath string
}

func NewDeleteViewModel(content string, name string, returnPath string) DeleteViewModel {
	return DeleteViewModel{
		Title:      "Meal Scheduler - Delete",
		Content:    content,
		Name:       name,
		ReturnPath: returnPath,
	}
}
