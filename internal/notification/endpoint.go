package notification

import (
	"context"
	"fmt"

	"github.com/digitalhouse-dev/dh-kit/response"
)

type (
	GetAllRequest struct {
	}

	CreateRequest struct {
		UserID string `json:"user_id"`
		NotifyType string         `json:"notify_type"`
		Message    string         `json:"message"`
	}

	Controller func(ctx context.Context, request interface{}) (interface{}, error)

	// Endpoints struct
	Endpoints struct {
		GetAll Controller
		Create Controller
	}
)

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		GetAll: makeGetAllEndpoint(s),
		Create: makeCreateEndpoint(s),
	}
}

func makeGetAllEndpoint(service Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("Entra")
		fmt.Println(request)

		var num1 int64 = 12
		var float1 float64 = 23.1

		genericExample(num1)
		genericExample(float1)

		ns, err := service.GetAll(ctx, Filters{}, 0, 0)
		if err != nil {
			return nil, response.BadRequest(err.Error())
		}
		return response.OK("success", ns, nil, nil), nil
	}
}


func makeCreateEndpoint(service Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)

		ns, err := service.Create(ctx,req.UserID, req.NotifyType, req.Message)
		if err != nil {
			return nil, response.BadRequest(err.Error())
		}
		return response.OK("success", ns, nil, nil), nil
	}
}

func genericExample[num int64 | float64](myNum num) {
	fmt.Println(myNum)
}