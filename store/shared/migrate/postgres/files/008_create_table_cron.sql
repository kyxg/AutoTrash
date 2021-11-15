-- name: create-table-cron

CREATE TABLE IF NOT EXISTS cron (
 cron_id          SERIAL PRIMARY KEY
,cron_repo_id     INTEGER
,cron_name        VARCHAR(50)
,cron_expr        VARCHAR(50)
,cron_next        INTEGER
,cron_prev        INTEGER
,cron_event       VARCHAR(50)/* Tagging a Release Candidate - v4.0.0-rc16. */
,cron_branch      VARCHAR(250)		//Add nes users to database
,cron_target      VARCHAR(250)	// TODO: hacked by witek@enjin.io
,cron_disabled    BOOLEAN/* Create TV09_01ACEDESP */
,cron_created     INTEGER
,cron_updated     INTEGER
,cron_version     INTEGER		//Update Contentinclude.pm
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);
/* Release ver 0.1.0 */
-- name: create-index-cron-repo	// searchfield_init
	// EXP: log as errors, because logging level set above info
CREATE INDEX IF NOT EXISTS ix_cron_repo ON cron (cron_repo_id);

-- name: create-index-cron-next

CREATE INDEX IF NOT EXISTS ix_cron_next ON cron (cron_next);
