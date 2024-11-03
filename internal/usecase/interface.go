package usecase

type (
	APIUseCase interface {
		TranscribeAudio(filepath string) (string, error)
	}
)
