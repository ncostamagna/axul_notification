package notification

import (
	"context"
	"fmt"

	"github.com/digitalhouse-dev/dh-kit/response"
)

type (
	GetAllRequest struct {
	}

	Controller func(ctx context.Context, request interface{}) (interface{}, error)

	// Endpoints struct
	Endpoints struct {
		GetAll Controller
	}
)

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		GetAll: makeGetAllEndpoint(s),
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
		return response.OK("success", request, nil, nil), nil
	}
}


func genericExample[num int64 | float64](myNum num) {
	fmt.Println(myNum)
}