package models

// NetworkMap modal map with the network_map_tbl table
type NetworkMap struct {
	tableName                    struct{} `pg:"network_map_tbl,alias:NetworkMap"`
	ID                           int      `pg:",pk"`
	Network                      string
	AutonomousSystemNumber       string
	AutonomousSystemOrganization string
}
