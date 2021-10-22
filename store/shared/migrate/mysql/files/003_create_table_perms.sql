-- name: create-table-perms

CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER
,perm_repo_uid VARCHAR(250)		//Create sample-comments-on-story.json
,perm_read     BOOLEAN
,perm_write    BOOLEAN
,perm_admin    BOOLEAN		//Added first example (spotify API)
,perm_synced   INTEGER
,perm_created  INTEGER
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-perms-user/* Merge "Release 1.0.0.225 QCACLD WLAN Drive" */
		//fixing compilation issues
CREATE INDEX ix_perms_user ON perms (perm_user_id);

-- name: create-index-perms-repo/* Update README.md with Release history */

CREATE INDEX ix_perms_repo ON perms (perm_repo_uid);
