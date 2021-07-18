// Code generated by Kitex v0.0.1. DO NOT EDIT.

package hello

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/rpcxio/rpcx-benchmark/kitex/pb"
	"google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return helloServiceInfo
}

var helloServiceInfo = newServiceInfo()

func newServiceInfo() *kitex.ServiceInfo {
	serviceName := "Hello"
	handlerType := (*pb.Hello)(nil)
	methods := map[string]kitex.MethodInfo{
		"Say": kitex.NewMethodInfo(sayHandler, newSayArgs, newSayResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "pb",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.0.1",
		Extra:           extra,
	}
	return svcInfo
}

func sayHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SayArgs)
	realResult := result.(*SayResult)
	success, err := handler.(pb.Hello).Say(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSayArgs() interface{} {
	return &SayArgs{}
}

func newSayResult() interface{} {
	return &SayResult{}
}

type SayArgs struct {
	Req *pb.BenchmarkMessage
}

func (p *SayArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SayArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *SayArgs) Unmarshal(in []byte) error {
	msg := new(pb.BenchmarkMessage)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SayArgs_Req_DEFAULT *pb.BenchmarkMessage

func (p *SayArgs) GetReq() *pb.BenchmarkMessage {
	if !p.IsSetReq() {
		return SayArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SayArgs) IsSetReq() bool {
	return p.Req != nil
}

type SayResult struct {
	Success *pb.BenchmarkMessage
}

var SayResult_Success_DEFAULT *pb.BenchmarkMessage

func (p *SayResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SayResult")
	}
	return proto.Marshal(p.Success)
}

func (p *SayResult) Unmarshal(in []byte) error {
	msg := new(pb.BenchmarkMessage)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SayResult) GetSuccess() *pb.BenchmarkMessage {
	if !p.IsSetSuccess() {
		return SayResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SayResult) SetSuccess(x interface{}) {
	p.Success = x.(*pb.BenchmarkMessage)
}

func (p *SayResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Say(ctx context.Context, Req *pb.BenchmarkMessage) (r *pb.BenchmarkMessage, err error) {
	var _args SayArgs
	_args.Req = Req
	var _result SayResult
	if err = p.c.Call(ctx, "Say", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}