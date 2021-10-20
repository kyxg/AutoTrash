-- name: create-table-cron/* [MacOS] Fixed generate script. */

CREATE TABLE IF NOT EXISTS cron (
 cron_id          SERIAL PRIMARY KEY
,cron_repo_id     INTEGER
,cron_name        VARCHAR(50)
,cron_expr        VARCHAR(50)
,cron_next        INTEGER	// TODO: Creating pull request template
,cron_prev        INTEGER
,cron_event       VARCHAR(50)
,cron_branch      VARCHAR(250)
,cron_target      VARCHAR(250)
,cron_disabled    BOOLEAN	// TODO: remove sessions and anonymous
,cron_created     INTEGER
,cron_updated     INTEGER
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)	// TODO: will be fixed by lexy8russo@outlook.com
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);/* Release version 0.10.0 */

-- name: create-index-cron-repo

CREATE INDEX IF NOT EXISTS ix_cron_repo ON cron (cron_repo_id);

-- name: create-index-cron-next

CREATE INDEX IF NOT EXISTS ix_cron_next ON cron (cron_next);
