package viewmodel

type DeleteViewModel struct {
	Title      string
	Active     string
	Content    string
	Name       string
	ReturnPath string
}

func NewDeleteViewModel() DeleteViewModel {
	return DeleteViewModel{
		Title: "Meal Scheduler - Delete",
	}
}
