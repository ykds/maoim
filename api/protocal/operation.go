package protocal

const (
	OpRaw = int32(0)

	OpHeartBeat = int32(1)
	OpHeartBeatReply = int32(2)

	OpProtoReady = int32(3)
	OpProtoFinish = int32(4)
)
