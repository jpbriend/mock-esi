package esiv1

/*
200 ok object */
type GetUniverseSchematicsSchematicIdOk struct {
	/*
	 schematic_name string */
	SchematicName string `json:"schematic_name,omitempty"`
	/*
	 Time in seconds to process a run */
	CycleTime int32 `json:"cycle_time,omitempty"`
}