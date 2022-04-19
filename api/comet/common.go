package comet

import "maoim/api/protocal"

var (
	ProtoFinish         = &protocal.Proto{Op: protocal.OpProtoFinish}
	ProtoHeartBeatReply = &protocal.Proto{Op: protocal.OpHeartBeatReply}
)
