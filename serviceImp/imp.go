package serviceImp

import (
	"context"
	"gomicro-grpcExample/Services"
)

type service struct {}

func (s service) GetProd(ctx context.Context, req *Services.Req, rsp *Services.Rsp) error {
	req.Size+=1
	tests := make([]*Services.Test, 0)

	tests = append(tests, &Services.Test{
		Test1: 1,
		Test2: "2",
	})
	rsp.Data =tests
	return nil
}

func GetServer() Services.ProdServiceHandler {
	return service{}
}

