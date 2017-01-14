package swaggerServer

import "time"
var _ time.Time

/* 
The defending corporation or alliance that declared this war, only contains either corporation_id or alliance_id */
type GetWarsWarIdDefender struct {
/*
	 Alliance ID if and only if the defender is an alliance */
	alliance_id int32 `json:"alliance_id,omitempty"`
/*
	 Corporation ID if and only if the defender is a corporation */
	corporation_id int32 `json:"corporation_id,omitempty"`
/*
	 ISK value of ships the defender has killed */
	isk_destroyed float32 `json:"isk_destroyed,omitempty"`
/*
	 The number of ships the defender has killed */
	ships_killed int32 `json:"ships_killed,omitempty"`
}
