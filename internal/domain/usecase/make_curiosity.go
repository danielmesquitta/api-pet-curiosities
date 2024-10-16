package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/danielmesquitta/api-pet-curiosities/internal/domain/errs"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent"
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/gpt"
)

type MakeCuriosityUseCase struct {
	gptProvider gpt.Provider
}

func NewMakeCuriosityUseCase(
	gptProvider gpt.Provider,
) *MakeCuriosityUseCase {
	return &MakeCuriosityUseCase{
		gptProvider: gptProvider,
	}
}

func (u *MakeCuriosityUseCase) Execute(
	ctx context.Context,
	breed string,
	curiosityTitles []string,
) (*ent.Curiosity, error) {
	const baseMessage = "Write a short title and interesting fact about the %s breed (max 3 lines), " +
		"and return in JSON format with the keys \"title\" and \"content\""

	message := fmt.Sprintf(baseMessage, breed)

	if len(curiosityTitles) > 0 {
		topics := strings.Join(curiosityTitles, ", ")

		message += fmt.Sprintf(
			". Avoid topics on %s",
			topics,
		)
	}

	completion, err := u.gptProvider.CreateChatCompletion(ctx, message)
	if err != nil {
		return nil, errs.New(err)
	}

	var curiosity ent.Curiosity
	if err := json.Unmarshal([]byte(completion), &curiosity); err != nil {
		return nil, errs.New(err)
	}

	return &curiosity, nil
}
