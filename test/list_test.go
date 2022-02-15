package test

import (
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"github.com/simonkimi/catweb-parser/rpc"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"testing"
)

func TestList(t *testing.T) {
	buffer, err := ioutil.ReadFile("./buffer.tmp")
	if err != nil {
		t.Error(err)
		return
	}
	data := rpc.ExecSelector(buffer)
	rsp := &protobuf.RpcResponse{}
	_ = proto.Unmarshal(data, rsp)

	model := protobuf.ListRpcModel{}
	_ = proto.Unmarshal(rsp.Data, &model)

	for _, i := range model.Items {
		t.Log(i.Badges)
	}
}
