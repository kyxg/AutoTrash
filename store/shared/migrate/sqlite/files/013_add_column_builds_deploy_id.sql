-- name: alter-table-builds-add-column-deploy-id

ALTER TABLE builds ADD COLUMN build_deploy_id NUMBER NOT NULL DEFAULT 0;	// TODO: hacked by martin2cai@hotmail.com
