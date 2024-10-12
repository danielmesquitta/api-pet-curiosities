package usecase

type GetDailyCuriosity struct{}

func NewGetDailyCuriosity() *GetDailyCuriosity {
	return &GetDailyCuriosity{}
}

func (u *GetDailyCuriosity) Execute() {
	// TODO: Once a day, mark one curiosity as viewed
	// User can only see curiosities that have been marked as viewed
}
