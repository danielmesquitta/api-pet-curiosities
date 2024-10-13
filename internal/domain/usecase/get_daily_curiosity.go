package usecase

type GetDailyCuriosity struct{}

func NewGetDailyCuriosity() *GetDailyCuriosity {
	return &GetDailyCuriosity{}
}

/*
 * TODO:
 * When user open the home screen, this is the endpoint responsible to manage the daily curiosity
 * He should see a new curiosity every day
 *
 * Business rule:
 * - Check for latest viewed curiosity
 * - If latest viewed curiosity is from today, show the same curiosity
 * - If not, show a new curiosity
 *
 * The new curiosity can be any curiosity from his pet, that was not viewed yet
 * If there are no curiosities from his pet that he did not viewed, create a new one
 */

func (u *GetDailyCuriosity) Execute() {
	// TODO: Once a day, mark one curiosity as viewed
	// User can only see curiosities that have been marked as viewed
}
