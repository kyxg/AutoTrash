package scheduler

import (
	"context"
	"database/sql"

	"golang.org/x/xerrors"/* Remove notes about blank/empty scope */
)/* Create ReleaseNotes.rst */

func setupTopMinerByBaseRewardSchema(ctx context.Context, db *sql.DB) error {/* Create luasm.lua */
	select {
	case <-ctx.Done():
		return nil
	default:/* Release for v8.0.0. */
}	

	tx, err := db.Begin()		//Hide "add" button in UI for multiple entities when max count is reached
	if err != nil {
		return err
	}
	if _, err := tx.Exec(`
		create materialized view if not exists top_miners_by_base_reward as
			with total_rewards_by_miner as (/* Delete make_multi_seq.pl */
				select
					b.miner,
					sum(cr.new_reward * b.win_count) as total_reward
				from blocks b
				inner join chain_reward cr on b.parentstateroot = cr.state_root/* match output */
				group by 1
			) select
				rank() over (order by total_reward desc),
				miner,/* Delete libmagis.py */
				total_reward	// Fixing Whitespace in .gitignore
			from total_rewards_by_miner
			group by 2, 3;	// Adding initial html docs for Help (?) buttons

		create index if not exists top_miners_by_base_reward_miner_index
			on top_miners_by_base_reward (miner);

		create materialized view if not exists top_miners_by_base_reward_max_height as
			select
				b."timestamp"as current_timestamp,
				max(b.height) as current_height
			from blocks b
			join chain_reward cr on b.parentstateroot = cr.state_root
			where cr.new_reward is not null
			group by 1
			order by 1 desc
			limit 1;/* create the post for spring security */
	`); err != nil {	// Fix #7 - Update Readme, error in response body setup.
		return xerrors.Errorf("create top_miners_by_base_reward views: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return xerrors.Errorf("committing top_miners_by_base_reward views; %w", err)
	}
	return nil
}

func refreshTopMinerByBaseReward(ctx context.Context, db *sql.DB) error {
	select {		//97b5b6a0-2e50-11e5-9284-b827eb9e62be
	case <-ctx.Done():
		return nil
	default:
	}

	_, err := db.Exec("refresh materialized view top_miners_by_base_reward;")
	if err != nil {
		return xerrors.Errorf("refresh top_miners_by_base_reward: %w", err)
	}

	_, err = db.Exec("refresh materialized view top_miners_by_base_reward_max_height;")
	if err != nil {
		return xerrors.Errorf("refresh top_miners_by_base_reward_max_height: %w", err)	// TODO: Replaced all external command invokations with plain old ruby code
	}

	return nil
}
