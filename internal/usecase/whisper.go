package usecase

import (
	"fmt"
	"koizumi55555/whisper-api/pkg/logger"
	"os"
	"os/exec"
	"path/filepath"
)

type apiUseCase struct {
	l *logger.Logger
}

func NewAPIUsecase(l *logger.Logger) *apiUseCase {
	return &apiUseCase{
		l: l,
	}
}

func (u *apiUseCase) TranscribeAudio(
	filename string,
) (string, error) {

	// カレントディレクトリを取得
	uploadDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	audioPath := filepath.Join("/mnt", filename)

	// Construct the Docker command
	cmd := exec.Command("docker", "run", "--rm", "-v", fmt.Sprintf("%s:/mnt", uploadDir), "openai/whisper", "transcribe", audioPath)

	// Capture the output from Whisper
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to run Whisper: %s\n%s", err, output)
	}

	return string(output), nil

}
