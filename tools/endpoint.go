package tools

import (
    "context"
    "github.com/go-kit/kit/endpoint"
)

type Endpoint struct {
    CreateLang endpoint.Endpoint
    GetLang    endpoint.Endpoint
}

func CreateEndpoint(sv Service) Endpoint {
    return Endpoint {
        CreateLang: CreatelLangEndpoint(sv),
        GetLang:    GetlLangEndpoint(sv)
    }
}

func CreatelLangEndpoint(sv Service) endpoint.Endpoint {
	return func (ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateLangReq)
		s, err := sv.CreateLang(ctx, req.Name, req.Paradigm)
		return CreateLangResp{S: success}, err
	}
}

func GetlLangEndpoint(sv Service) endpoint.Endpoint {
	return func (ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetLangReq)
		name, err := sv.GetLang(ctx, req.Id)

		return GetLangResp {
			Name: name,
		}, err
	}
}
