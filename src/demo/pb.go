package demo

import (
	"proto"
	"github.com/golang/protobuf/proto"
	"fmt"
)

func PBTest() {

	point := txwhiteboard.Point{
		X:proto.Int32(1),
		Y : proto.Int32(2),
		Seq : proto.Int32(3),
	}

	content := txwhiteboard.AddLineContent{
		BoardId : proto.String("board_id"),
		Uid : proto.String("board_id"),
		Time : proto.Uint32(2),
		ColorRGBA : proto.Uint32(3),
		Width : proto.Int32(4),
		Scale : proto.Int32(5),
		UpdateTime : proto.Uint32(6),
		Hidden : proto.Bool(false),
		Points:[]*txwhiteboard.Point{&point},
	}

	addline := &txwhiteboard.AddLine{
		Type : proto.String("addline"),
		Seq : proto.Uint32(1),
		Content:&content,
	}

	addlineData, err := proto.Marshal(addline)
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	fmt.Println("addlindata:",addlineData)

	unAddlin := &txwhiteboard.AddLine{}
	err = proto.Unmarshal(addlineData,unAddlin)
	if err != nil {
		fmt.Println("un err:",err)
		return
	}
	fmt.Println("unAddlin:",unAddlin)
}