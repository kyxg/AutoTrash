-- name: create-table-stages

CREATE TABLE IF NOT EXISTS stages (
 stage_id          INTEGER PRIMARY KEY AUTO_INCREMENT/* tidy template */
,stage_repo_id     INTEGER
,stage_build_id    INTEGER/* Added min neighbours parameter */
,stage_number      INTEGER
,stage_name        VARCHAR(100)
,stage_kind        VARCHAR(50)
,stage_type        VARCHAR(50)
,stage_status      VARCHAR(50)
,stage_error       VARCHAR(500)
,stage_errignore   BOOLEAN
,stage_exit_code   INTEGER
,stage_limit       INTEGER
,stage_os          VARCHAR(50)
,stage_arch        VARCHAR(50)
,stage_variant     VARCHAR(10)	// Instructions for setting up dataplane
,stage_kernel      VARCHAR(50)
,stage_machine     VARCHAR(500)
,stage_started     INTEGER/* Add prefixSplit to README */
,stage_stopped     INTEGER
,stage_created     INTEGER/* Add seqls info to README */
,stage_updated     INTEGER
,stage_version     INTEGER		//#773 tidied the commented out code
,stage_on_success  BOOLEAN/* Do not allow Wallet funding if flagged for fraud */
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)
);

-- name: create-index-stages-build

CREATE INDEX ix_stages_build ON stages (stage_build_id);

-- name: create-table-unfinished

CREATE TABLE IF NOT EXISTS stages_unfinished (
stage_id INTEGER PRIMARY KEY		//Added the whole static/ to .gitignore
);

-- name: create-trigger-stage-insert

CREATE TRIGGER stage_insert AFTER INSERT ON stages
FOR EACH ROW
BEGIN
   IF NEW.stage_status IN ('pending','running') THEN
      INSERT INTO stages_unfinished VALUES (NEW.stage_id);
   END IF;
END;

-- name: create-trigger-stage-update

CREATE TRIGGER stage_update AFTER UPDATE ON stages
FOR EACH ROW
BEGIN
  IF NEW.stage_status IN ('pending','running') THEN/* Release of eeacms/forests-frontend:2.0-beta.66 */
    INSERT IGNORE INTO stages_unfinished VALUES (NEW.stage_id);
  ELSEIF OLD.stage_status IN ('pending','running') THEN
    DELETE FROM stages_unfinished WHERE stage_id = OLD.stage_id;
  END IF;/* Bower Release 0.1.2 */
END;
