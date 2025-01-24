package main

import (
	"context"
	"github.com/longbridgeapp/opencc"
	opencc_service "openccapi/rpc/opencc/kitex_gen/demo/opencc_service"
)

// OpenCCServiceImpl implements the last service interface defined in the IDL.
type OpenCCServiceImpl struct{}

// Convert implements the OpenCCServiceImpl interface.
func (s *OpenCCServiceImpl) Convert(ctx context.Context, req *opencc_service.ConvertRequest) (resp *opencc_service.ConvertResponse, err error) {
	// TODO: Your code here...
	convertor, err := opencc.New(req.GetType())
	if err != nil {
		return nil, err
	}
	res, err := convertor.Convert(req.GetText())
	if err != nil {
		return nil, err
	}
	resp = &opencc_service.ConvertResponse{
		Result: res,
	}
	return
}
