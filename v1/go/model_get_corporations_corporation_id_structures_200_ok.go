package esiv1

import "time"

/*
200 ok object */
type GetCorporationsCorporationIdStructures200Ok struct {
	/*
	 The Item ID of the structure */
	StructureId int64 `json:"structure_id,omitempty"`
	/*
	 The type id of the structure */
	TypeId int32 `json:"type_id,omitempty"`
	/*
	 ID of the corporation that owns the structure */
	CorporationId int32 `json:"corporation_id,omitempty"`
	/*
	 The solar system the structure is in */
	SystemId int32 `json:"system_id,omitempty"`
	/*
	 The id of the ACL profile for this citadel */
	ProfileId int32 `json:"profile_id,omitempty"`
	/*
	 Date on which the structure will run out of fuel */
	FuelExpires time.Time `json:"fuel_expires,omitempty"`
	/*
	 Contains a list of service upgrades, and their state */
	Services []GetCorporationsCorporationIdStructuresService `json:"services,omitempty"`
	/*
	 Date at which the structure entered it's current state */
	StateTimerStart time.Time `json:"state_timer_start,omitempty"`
	/*
	 Date at which the structure will move to it's next state */
	StateTimerEnd time.Time `json:"state_timer_end,omitempty"`
	/*
	 Date at which the structure will unanchor */
	UnanchorsAt time.Time `json:"unanchors_at,omitempty"`
	/*
	 This week's vulnerability windows, Monday is day 0 */
	CurrentVul []GetCorporationsCorporationIdStructuresCurrentVul `json:"current_vul,omitempty"`
	/*
	 Next week's vulnerability windows, Monday is day 0 */
	NextVul []GetCorporationsCorporationIdStructuresNextVul `json:"next_vul,omitempty"`
}