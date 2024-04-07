package external

import (
	"fmt"
	"go-practice/domain/entity"
	"go-practice/usecase"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

type UserSlackNotification struct{}

func NewUserSlackNotification() usecase.IUserExtarnal {
	return &UserSlackNotification{}
}

func (usn *UserSlackNotification) SendRegisterd(user *entity.User) error {
	err := godotenv.Load()
	if err != nil {
		log.Print(err.Error())
		return err
	}

	api := slack.New(os.Getenv("SLACK_BOT_OAUTH_TOKEN"))

	_, _, err = api.PostMessage(os.Getenv("SLACK_CHANNEL_ID"), slack.MsgOptionText(fmt.Sprintf("ID: %d 名前: %s が登録されました。", user.Id, user.Name), false))
	if err != nil {
		log.Print(err.Error())
		return err
	}

	return nil
}
