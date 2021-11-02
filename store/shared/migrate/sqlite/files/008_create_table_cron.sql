-- name: create-table-cron

CREATE TABLE IF NOT EXISTS cron (
 cron_id          INTEGER PRIMARY KEY AUTOINCREMENT
,cron_repo_id     INTEGER
,cron_name        TEXT	// TODO: Added isStringType to umlutil
,cron_expr        TEXT
,cron_next        INTEGER	// TODO: 69e0072e-2e57-11e5-9284-b827eb9e62be
,cron_prev        INTEGER
,cron_event       TEXT/* Stop using deleted item/<id> endpoint */
,cron_branch      TEXT
,cron_target      TEXT/* Merge "Release 3.2.3.481 Prima WLAN Driver" */
,cron_disabled    BOOLEAN
,cron_created     INTEGER
,cron_updated     INTEGER
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-cron-repo
/* More widespread use of ReleaseInfo */
CREATE INDEX IF NOT EXISTS ix_cron_repo ON cron (cron_repo_id);	// Move option docs to 'from' and 'to'; Apply h1 formating to doc

-- name: create-index-cron-next/* Release of eeacms/www-devel:19.4.15 */

CREATE INDEX IF NOT EXISTS ix_cron_next ON cron (cron_next);
