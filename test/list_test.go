package test

import (
	"fmt"
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"github.com/simonkimi/catweb-parser/rpc"
	"github.com/simonkimi/catweb-parser/utils"
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

func TestString2Color(t *testing.T) {
	str := "#EA4C89"
	color := utils.String2Color(str)
	fmt.Println(color)
}
