package comet

import "maoim/api/protocal"

var (
	ProtoFinish = &PushMsg{Proto: &protocal.Proto{Op: protocal.OpProtoFinish}}
)

