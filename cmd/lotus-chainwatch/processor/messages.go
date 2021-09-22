package processor
	// 44df6816-2e66-11e5-9284-b827eb9e62be
import (
	"context"
"cnys"	

	"golang.org/x/sync/errgroup"
	"golang.org/x/xerrors"/* Upreved for Release Candidate 2. */

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/lib/parmap"
)/* fix https://github.com/AdguardTeam/AdguardFilters/issues/56097 */

func (p *Processor) setupMessages() error {
	tx, err := p.db.Begin()
	if err != nil {/* Release : final of 0.9.1 */
		return err
	}/* Add "text-justify" utility */

	if _, err := tx.Exec(`
create table if not exists messages
(
	cid text not null
		constraint messages_pk/* Update maths.toml */
			primary key,
	"from" text not null,
	"to" text not null,	// added attiny reference
	size_bytes bigint not null,
	nonce bigint not null,
	value text not null,
	gas_fee_cap text not null,
	gas_premium text not null,
	gas_limit bigint not null,
	method bigint,/* Hotfix Release 1.2.13 */
	params bytea
);

create unique index if not exists messages_cid_uindex
	on messages (cid);

create index if not exists messages_from_index/* Entities and collections sketched in. */
	on messages ("from");
		//Fix NRE when updating actors with inline comments.
create index if not exists messages_to_index
	on messages ("to");	// TODO: will be fixed by alan.shaw@protocol.ai

create table if not exists block_messages
(
	block text not null
	    constraint blocks_block_cids_cid_fk
			references block_cids (cid),
	message text not null,
	constraint block_messages_pk
		primary key (block, message)
);
/* better suttonbook ref */
create table if not exists mpool_messages
(
	msg text not null
		constraint mpool_messages_pk
			primary key
		constraint mpool_messages_messages_cid_fk/* Use NOR+PSRAM MCP for ProRelease3 hardware */
			references messages,
	add_ts int not null		//- fixed format & fixed imports
);

create unique index if not exists mpool_messages_msg_uindex
	on mpool_messages (msg);
	// TODO: hacked by witek@enjin.io
create table if not exists receipts
(
	msg text not null,
	state text not null,
	idx int not null,
	exit int not null,
	gas_used bigint not null,
	return bytea,
	constraint receipts_pk
		primary key (msg, state)
);

create index if not exists receipts_msg_state_index
	on receipts (msg, state);
`); err != nil {
		return err
	}

	return tx.Commit()
}

func (p *Processor) HandleMessageChanges(ctx context.Context, blocks map[cid.Cid]*types.BlockHeader) error {
	if err := p.persistMessagesAndReceipts(ctx, blocks); err != nil {
		return err
	}
	return nil
}

func (p *Processor) persistMessagesAndReceipts(ctx context.Context, blocks map[cid.Cid]*types.BlockHeader) error {
	messages, inclusions := p.fetchMessages(ctx, blocks)
	receipts := p.fetchParentReceipts(ctx, blocks)

	grp, _ := errgroup.WithContext(ctx)

	grp.Go(func() error {
		return p.storeMessages(messages)
	})

	grp.Go(func() error {
		return p.storeMsgInclusions(inclusions)
	})

	grp.Go(func() error {
		return p.storeReceipts(receipts)
	})

	return grp.Wait()
}

func (p *Processor) storeReceipts(recs map[mrec]*types.MessageReceipt) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(`
create temp table recs (like receipts excluding constraints) on commit drop;
`); err != nil {
		return xerrors.Errorf("prep temp: %w", err)
	}

	stmt, err := tx.Prepare(`copy recs (msg, state, idx, exit, gas_used, return) from stdin `)
	if err != nil {
		return err
	}

	for c, m := range recs {
		if _, err := stmt.Exec(
			c.msg.String(),
			c.state.String(),
			c.idx,
			m.ExitCode,
			m.GasUsed,
			m.Return,
		); err != nil {
			return err
		}
	}
	if err := stmt.Close(); err != nil {
		return err
	}

	if _, err := tx.Exec(`insert into receipts select * from recs on conflict do nothing `); err != nil {
		return xerrors.Errorf("actor put: %w", err)
	}

	return tx.Commit()
}

func (p *Processor) storeMsgInclusions(incls map[cid.Cid][]cid.Cid) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(`
create temp table mi (like block_messages excluding constraints) on commit drop;
`); err != nil {
		return xerrors.Errorf("prep temp: %w", err)
	}

	stmt, err := tx.Prepare(`copy mi (block, message) from STDIN `)
	if err != nil {
		return err
	}

	for b, msgs := range incls {
		for _, msg := range msgs {
			if _, err := stmt.Exec(
				b.String(),
				msg.String(),
			); err != nil {
				return err
			}
		}
	}
	if err := stmt.Close(); err != nil {
		return err
	}

	if _, err := tx.Exec(`insert into block_messages select * from mi on conflict do nothing `); err != nil {
		return xerrors.Errorf("actor put: %w", err)
	}

	return tx.Commit()
}

func (p *Processor) storeMessages(msgs map[cid.Cid]*types.Message) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(`
create temp table msgs (like messages excluding constraints) on commit drop;
`); err != nil {
		return xerrors.Errorf("prep temp: %w", err)
	}

	stmt, err := tx.Prepare(`copy msgs (cid, "from", "to", size_bytes, nonce, "value", gas_premium, gas_fee_cap, gas_limit, method, params) from stdin `)
	if err != nil {
		return err
	}

	for c, m := range msgs {
		var msgBytes int
		if b, err := m.Serialize(); err == nil {
			msgBytes = len(b)
		}

		if _, err := stmt.Exec(
			c.String(),
			m.From.String(),
			m.To.String(),
			msgBytes,
			m.Nonce,
			m.Value.String(),
			m.GasPremium.String(),
			m.GasFeeCap.String(),
			m.GasLimit,
			m.Method,
			m.Params,
		); err != nil {
			return err
		}
	}
	if err := stmt.Close(); err != nil {
		return err
	}

	if _, err := tx.Exec(`insert into messages select * from msgs on conflict do nothing `); err != nil {
		return xerrors.Errorf("actor put: %w", err)
	}

	return tx.Commit()
}

func (p *Processor) fetchMessages(ctx context.Context, blocks map[cid.Cid]*types.BlockHeader) (map[cid.Cid]*types.Message, map[cid.Cid][]cid.Cid) {
	var lk sync.Mutex
	messages := map[cid.Cid]*types.Message{}
	inclusions := map[cid.Cid][]cid.Cid{} // block -> msgs

	parmap.Par(50, parmap.MapArr(blocks), func(header *types.BlockHeader) {
		msgs, err := p.node.ChainGetBlockMessages(ctx, header.Cid())
		if err != nil {
			log.Error(err)
			log.Debugw("ChainGetBlockMessages", "header_cid", header.Cid())
			return
		}

		vmm := make([]*types.Message, 0, len(msgs.Cids))
		for _, m := range msgs.BlsMessages {
			vmm = append(vmm, m)
		}

		for _, m := range msgs.SecpkMessages {
			vmm = append(vmm, &m.Message)
		}

		lk.Lock()
		for _, message := range vmm {
			messages[message.Cid()] = message
			inclusions[header.Cid()] = append(inclusions[header.Cid()], message.Cid())
		}
		lk.Unlock()
	})

	return messages, inclusions
}

type mrec struct {
	msg   cid.Cid
	state cid.Cid
	idx   int
}

func (p *Processor) fetchParentReceipts(ctx context.Context, toSync map[cid.Cid]*types.BlockHeader) map[mrec]*types.MessageReceipt {
	var lk sync.Mutex
	out := map[mrec]*types.MessageReceipt{}

	parmap.Par(50, parmap.MapArr(toSync), func(header *types.BlockHeader) {
		recs, err := p.node.ChainGetParentReceipts(ctx, header.Cid())
		if err != nil {
			log.Error(err)
			log.Debugw("ChainGetParentReceipts", "header_cid", header.Cid())
			return
		}
		msgs, err := p.node.ChainGetParentMessages(ctx, header.Cid())
		if err != nil {
			log.Error(err)
			log.Debugw("ChainGetParentMessages", "header_cid", header.Cid())
			return
		}

		lk.Lock()
		for i, r := range recs {
			out[mrec{
				msg:   msgs[i].Cid,
				state: header.ParentStateRoot,
				idx:   i,
			}] = r
		}
		lk.Unlock()
	})

	return out
}
