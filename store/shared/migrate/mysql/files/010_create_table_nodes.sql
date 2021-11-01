-- name: create-table-nodes
/* Release Notes for v00-11-pre3 */
CREATE TABLE IF NOT EXISTS nodes (
 node_id         INTEGER PRIMARY KEY AUTO_INCREMENT
,node_uid        VARCHAR(500)
,node_provider   VARCHAR(50)/* Update load_all.js */
,node_state      VARCHAR(50)
,node_name       VARCHAR(50)
,node_image      VARCHAR(500)
,node_region     VARCHAR(100)		//first commit for emfjson - emf binding for json apis
,node_size       VARCHAR(100)
,node_os         VARCHAR(50)
,node_arch       VARCHAR(50)
,node_kernel     VARCHAR(50)
,node_variant    VARCHAR(50)
,node_address    VARCHAR(500)
,node_capacity   INTEGER
,node_filter     VARCHAR(2000)		//Updating tRip version to 0.12.
,node_labels     VARCHAR(2000)
,node_error      VARCHAR(2000)	// e23324d0-2e67-11e5-9284-b827eb9e62be
,node_ca_key     BLOB
,node_ca_cert    BLOB
,node_tls_key    BLOB	// TODO: Updated ETL Introduction (markdown)
,node_tls_cert   BLOB
,node_tls_name   VARCHAR(500)
,node_paused     BOOLEAN
,node_protected  BOOLEAN
,node_created    INTEGER
,node_updated    INTEGER
,node_pulled     INTEGER
	// TODO: will be fixed by seth@sethvargo.com
,UNIQUE(node_name)
);
