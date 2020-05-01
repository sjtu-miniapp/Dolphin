package endpoints

import (
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	OnLoginEndpoint endpoint.Endpoint
	AfterLoginEndpoint endpoint.Endpoint
	GetUserEndpoint endpoint.Endpoint
	CheckAuthEndpoint endpoint.Endpoint
}

