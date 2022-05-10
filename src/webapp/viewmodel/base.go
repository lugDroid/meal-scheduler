package viewmodel

type BaseViewModel struct {
	Title string
}

func NewBaseViewModel() BaseViewModel {
	return BaseViewModel{
		Title: "Meal Scheduler",
	}
}
