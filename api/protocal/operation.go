package protocal

const (
	// server 流转
	OpHeartBeat      = int32(2)
	OpHeartBeatReply = int32(3)
	OpProtoReady     = int32(4)
	OpProtoFinish    = int32(5)

	// server - client 流转
	OpRaw                      = int32(1)
	OpAck                      = int32(6)
	OpErr                      = int32(7)
	OpNewFriendShipApplyNotice = int32(8)
	OpNewFriendShipPassNotice  = int32(9)
)
