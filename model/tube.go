package model

import (
	"github.com/ggvishnu29/horlix/serde"
)

const (
	REPLACE_DATA = iota
	APPEND_DATA
)

// Tube struct encapsulates all information specific to a tube
type Tube struct {
	ID                  string
	Lock                *Lock
	MsgMap              *MsgMap
	ReadyQueue          *ReadyQueue
	DelayedQueue        *DelayedQueue
	ReservedQueue       *ReservedQueue
	ReserveTimeoutInSec int64
	FuseSetting         *FuseSetting
	IsDeleted           bool
}

func NewTube(ID string, reserveTimeoutInSec int64, fuseSetting *FuseSetting) *Tube {
	t := &Tube{
		ID:                  ID,
		Lock:                &Lock{},
		MsgMap:              NewMsgMap(ID),
		ReadyQueue:          NewReadyQueue(ID),
		DelayedQueue:        NewDelayedQueue(ID),
		ReservedQueue:       NewReservedQueue(ID),
		ReserveTimeoutInSec: reserveTimeoutInSec,
		FuseSetting:         fuseSetting,
	}
	return t
}

func (t *Tube) SetDeleted(isDeleted bool, shouldTransLog bool) {
	t.IsDeleted = isDeleted
	if shouldTransLog {
		opr := serde.NewOperation(TUBE, SET_TUBE_DELETED_OPR, &t.ID, isDeleted)
		LogOpr(opr)
	}
}

type FuseSetting struct {
	Data int
}

func NewFuseSetting(data int) *FuseSetting {
	return &FuseSetting{
		Data: data,
	}
}

func (f *FuseSetting) Clone() *FuseSetting {
	return &FuseSetting{
		Data: f.Data,
	}
}
