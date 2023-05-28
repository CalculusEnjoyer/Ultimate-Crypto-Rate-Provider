package emails

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"storage.com/storage/emails/messages"
)

type Endpoints struct {
	AddEmail     endpoint.Endpoint
	GetAllEmails endpoint.Endpoint
}

func NewEndpointSet(svc StorageService) Endpoints {
	return Endpoints{
		AddEmail:     MakeAddEmailEndpoint(svc),
		GetAllEmails: MakeGetAllEmailsEndpoint(svc),
	}
}

func MakeGetAllEmailsEndpoint(svc StorageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		emails, err := svc.GetAllEmails()
		return emails, err
	}
}

func MakeAddEmailEndpoint(svc StorageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(messages.Email)
		err := svc.AddEmail(req)
		return nil, err
	}
}
