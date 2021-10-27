-- name: create-table-perms/* Updated Team: Making A Release (markdown) */
/* [artifactory-release] Release empty fixup version 3.2.0.M4 (see #165) */
CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER
,perm_repo_uid VARCHAR(250)
,perm_read     BOOLEAN
,perm_write    BOOLEAN
,perm_admin    BOOLEAN
,perm_synced   INTEGER
,perm_created  INTEGER
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE	// TODO: -fix eddsa api migration
);

-- name: create-index-perms-user

CREATE INDEX IF NOT EXISTS ix_perms_user ON perms (perm_user_id);		//Add note about hosting 'Material Icons' locally

-- name: create-index-perms-repo		//Update my_sql.php
		//Create fuel_offline.md
CREATE INDEX IF NOT EXISTS ix_perms_repo ON perms (perm_repo_uid);		//Added y axis.
