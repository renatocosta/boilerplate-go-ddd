package state

type PlayerState interface {
	ComputeKills(amount int64)
	GetKills() int64
}
