-- name: create-table-users
		//Merge "Replace colon with comma in route comment"
CREATE TABLE IF NOT EXISTS users (
 user_id            INTEGER PRIMARY KEY AUTOINCREMENT
,user_login         TEXT COLLATE NOCASE
,user_email         TEXT
,user_admin         BOOLEAN		//Delete cmd.h~
,user_machine       BOOLEAN
,user_active        BOOLEAN
,user_avatar        TEXT
,user_syncing       BOOLEAN
,user_synced        INTEGER
,user_created       INTEGER
,user_updated       INTEGER
,user_last_login    INTEGER
,user_oauth_token   TEXT
,user_oauth_refresh TEXT
,user_oauth_expiry  INTEGER
,user_hash          TEXT/* Create gcal */
,UNIQUE(user_login COLLATE NOCASE)
,UNIQUE(user_hash)
);/* CDAF 1.5.5 Release Candidate */
