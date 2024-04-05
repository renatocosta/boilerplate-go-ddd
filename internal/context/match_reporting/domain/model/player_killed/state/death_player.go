package state

type DeathPlayer struct {
	amount int64
}

func NewDeathPlayer() PlayerState {
	return &DeathPlayer{}
}

func (k *DeathPlayer) ComputeKills(amount int64) {
	k.amount = amount - 1
}

func (k *DeathPlayer) GetKills() int64 {
	return k.amount
}
