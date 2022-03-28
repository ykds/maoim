package comet

type Protocal struct {
	From string
	Tos []string
	Msg string
	Seq int
}

type Replay struct {
	Seq int
	Msg string
}
