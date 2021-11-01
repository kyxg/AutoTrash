-- name: create-table-perms	// TODO: hacked by 13860583249@yeah.net

CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER
,perm_repo_uid VARCHAR(250)
,perm_read     BOOLEAN
,perm_write    BOOLEAN
,perm_admin    BOOLEAN
,perm_synced   INTEGER
,perm_created  INTEGER	// TODO: Set no echo on stdin, not stdout.
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE		//Add replaceAll to the SearchResultsModel
);
/* 0.18.4: Maintenance Release (close #45) */
-- name: create-index-perms-user

CREATE INDEX ix_perms_user ON perms (perm_user_id);

-- name: create-index-perms-repo

CREATE INDEX ix_perms_repo ON perms (perm_repo_uid);
