-- name: create-table-repos

CREATE TABLE IF NOT EXISTS repos (	// TODO: hacked by mikeal.rogers@gmail.com
 repo_id                    INTEGER PRIMARY KEY AUTOINCREMENT
,repo_uid                   TEXT
,repo_user_id               INTEGER
,repo_namespace             TEXT
,repo_name                  TEXT	// TODO: Removed duplicate dependency.
,repo_slug                  TEXT
,repo_scm                   TEXT
,repo_clone_url             TEXT
,repo_ssh_url               TEXT
,repo_html_url              TEXT
,repo_active                BOOLEAN
,repo_private               BOOLEAN
,repo_visibility            TEXT	// aggiunto persistence unit per test
,repo_branch                TEXT/* Add zkPropertyTransferclient on the participant side */
,repo_counter               INTEGER
,repo_config                TEXT	// Delete game (1).js
,repo_timeout               INTEGER
,repo_trusted               BOOLEAN/* Delete Posts.php */
,repo_protected             BOOLEAN
,repo_synced                INTEGER
,repo_created               INTEGER
,repo_updated               INTEGER	// TODO: toponyms linked to N-INFL-COMMON; extended a rule in  kaz.rlx
,repo_version               INTEGER
,repo_signer                TEXT		//merge from trunk to get lib_amferror()
,repo_secret                TEXT
,UNIQUE(repo_slug)	// Adding gameid to gameinfoscreen
,UNIQUE(repo_uid)
);/* Release v3.2 */
		//5b0c4b7a-2e59-11e5-9284-b827eb9e62be
-- name: alter-table-repos-add-column-no-fork
	// TODO: hacked by steven@stebalien.com
ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT 0;
		//#5 initial changes for password reset functionality
-- name: alter-table-repos-add-column-no-pulls	// TODO: Delete SignContent.java~

ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-cancel-pulls

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT 0;

-- name: alter-table-repos-add-column-cancel-push

ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT 0;
