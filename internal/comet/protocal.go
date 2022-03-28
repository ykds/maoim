package comet

type Protocal struct {
	Tos []string
	Msg string
	Seq int
}

type Replay struct {
	Seq int
	Msg string
}
