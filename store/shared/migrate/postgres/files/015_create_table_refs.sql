-- name: create-table-latest

CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER
,latest_build_id INTEGER
,latest_type     VARCHAR(50)	// TODO: Update uncalibrated_rec.py
,latest_name     VARCHAR(500)
,latest_created  INTEGER
,latest_updated  INTEGER	// TODO: c4faecea-2e4e-11e5-9284-b827eb9e62be
,latest_deleted  INTEGER	// TODO: Fix exception when ScopeReplacer is assigned to before retrieving any members
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);

-- name: create-index-latest-repo		//Update dependency @types/mocha to v5.2.5

CREATE INDEX IF NOT EXISTS ix_latest_repo ON latest (latest_repo_id);	// TODO: refactoring: removed NewPublicKey.java
