package state

type KilledPlayer struct {
	amount int64
}

func NewKillPlayer() PlayerState {
	return &KilledPlayer{}
}

func (k *KilledPlayer) ComputeKills(amount int64) {
	k.amount = amount + 1
}

func (k *KilledPlayer) GetKills() int64 {
	return k.amount
}
