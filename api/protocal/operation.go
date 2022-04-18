package protocal

const (
	OpRaw = int32(1)

	OpHeartBeat      = int32(2)
	OpHeartBeatReply = int32(3)

	OpProtoReady  = int32(4)
	OpProtoFinish = int32(5)

	OpAck = int32(6)

	OpErr = int32(7)
)
