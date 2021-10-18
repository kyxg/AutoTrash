-- name: alter-table-builds-add-column-deploy-id
		//Very basic searching on item's description/name. Can extend later
ALTER TABLE builds ADD COLUMN build_deploy_id INTEGER NOT NULL DEFAULT 0;
