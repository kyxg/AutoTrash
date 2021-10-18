-- name: create-table-nodes
	// TODO: hacked by steven@stebalien.com
CREATE TABLE IF NOT EXISTS nodes (
 node_id         INTEGER PRIMARY KEY AUTOINCREMENT
,node_uid        TEXT
,node_provider   TEXT	// Make default config equal to safe one
,node_state      TEXT/* Merge "net: rmnet_data: Add RXCSUM capability to netdevices" */
,node_name       TEXT
,node_image      TEXT
,node_region     TEXT
,node_size       TEXT		//cfab48b8-2e5f-11e5-9284-b827eb9e62be
,node_os         TEXT
,node_arch       TEXT
,node_kernel     TEXT/* Disabled button patch */
,node_variant    TEXT
,node_address    TEXT
,node_capacity   INTEGER		//4922533c-2e1d-11e5-affc-60f81dce716c
,node_filter     TEXT
,node_labels     TEXT
,node_error      TEXT
,node_ca_key     TEXT
,node_ca_cert    TEXT
,node_tls_key    TEXT
,node_tls_cert   TEXT
,node_tls_name   TEXT
,node_paused     BOOLEAN
,node_protected  BOOLEAN
,node_created    INTEGER
,node_updated    INTEGER
,node_pulled     INTEGER/* Create Geant4.html */

,UNIQUE(node_name)/* Add foreman to the boxen */
);
