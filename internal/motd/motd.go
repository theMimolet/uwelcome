package motd

import (
	"math/rand"
	"os/exec"
	"uwelcome/internal/config"
)

func GetRandomMessage(cfg config.Config) string {

	messages := []string{}

	if len(cfg.Motd.Messages) > 0 {
		for _, msg := range cfg.Motd.Messages {
			messages = append(messages, msg)
		}
	}

	if len(cfg.Motd.Commands) > 0 {
		for _, msg := range cfg.Motd.Commands {
			out, _ := exec.Command(string(msg)).Output()
			messages = append(messages, string(out))
		}
	}

	return messages[rand.Intn(len(messages))]
}
