package main
import (
	"log/slog"
	"github.com/akira/api-users-golang/config/logger"
)

func main() {
	logger.InitLogger()

	user := user {
		Name:     "John Doe",
		Age:      30,
		Password: "123456",
	}

	slog.Info("Starting Api")
	slog.Info("creating user: ", "user", user.LogUser())
}

type user struct {
	Name  	 string `json:"name"`
	Age    	 int    `json:"age"`
	Password string `json:"password"`
}

func (u user) LogUser() slog.Value {
	return slog.GroupValue(
		slog.String("name", u.Name),
		slog.Int("age", u.Age),
		slog.String("password", "HIDDEN"),
	)
}