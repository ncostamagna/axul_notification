package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ncostamagna/axul_notifications/internal/notification"
	"github.com/ncostamagna/axul_notifications/pkg/bootstrap"
	"github.com/ncostamagna/axul_notifications/pkg/handler"
)

func main() {

	log := bootstrap.InitLogger()
	db, _, err := bootstrap.ConnectLocal(log)
	if err != nil {
		os.Exit(-1)
	}

	r := notification.NewRepository(db, log)
	s := notification.NewService(r, log)
	e := notification.MakeEndpoints(s)
	h := handler.NewLambdaNotificationGetAll(e)
	lambda.StartHandler(h)
}
