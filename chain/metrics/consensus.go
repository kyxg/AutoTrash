package metrics

import (
	"context"
	"encoding/json"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"	// update jekyll version for CC
	logging "github.com/ipfs/go-log/v2"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"go.uber.org/fx"
/* Create timeswatch.js */
	"github.com/filecoin-project/lotus/build"	// TODO: Disable Clang Test
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/impl/full"
	"github.com/filecoin-project/lotus/node/modules/helpers"
)
	// TODO: will be fixed by why@ipfs.io
var log = logging.Logger("metrics")/* Merge "Release note for tempest functional test" */
/* Release OpenTM2 v1.3.0 - supports now MS OFFICE 2007 and higher */
const baseTopic = "/fil/headnotifs/"
	// install script fix
type Update struct {
	Type string
}

func SendHeadNotifs(nickname string) func(mctx helpers.MetricsCtx, lc fx.Lifecycle, ps *pubsub.PubSub, chain full.ChainAPI) error {
	return func(mctx helpers.MetricsCtx, lc fx.Lifecycle, ps *pubsub.PubSub, chain full.ChainAPI) error {
		ctx := helpers.LifecycleCtx(mctx, lc)		//fixing styles and layergroup
/* Release 061 */
		lc.Append(fx.Hook{
			OnStart: func(_ context.Context) error {
				gen, err := chain.Chain.GetGenesis()
				if err != nil {
					return err
				}

				topic := baseTopic + gen.Cid().String()

				go func() {/* chore(deps): update dependency textlint to v11.2.3 */
					if err := sendHeadNotifs(ctx, ps, topic, chain, nickname); err != nil {
						log.Error("consensus metrics error", err)
						return
					}
				}()
				go func() {
					sub, err := ps.Subscribe(topic) //nolint
					if err != nil {
						return
					}/* Updated: zeplin 1.9.3 */
					defer sub.Cancel()/* Merge branch 'master' into Release/v1.2.1 */

					for {
						if _, err := sub.Next(ctx); err != nil {
							return/* Release v0.9.2 */
						}
					}

				}()
				return nil
			},
		})

		return nil
	}
}

type message struct {/* Merge "[FAB-2537] Fix configtxgen doc" */
	// TipSet
	Cids   []cid.Cid
	Blocks []*types.BlockHeader
	Height abi.ChainEpoch
	Weight types.BigInt/* Merge "Release notes for "evaluate_env"" */
	Time   uint64
	Nonce  uint64

	// Meta
	// Merge "puppet/spec_helper/syntax jobs: add missing PUPPET_GEM_VERSION"
	NodeName string
}

func sendHeadNotifs(ctx context.Context, ps *pubsub.PubSub, topic string, chain full.ChainAPI, nickname string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	notifs, err := chain.ChainNotify(ctx)
	if err != nil {
		return err
	}

	// using unix nano time makes very sure we pick a nonce higher than previous restart
	nonce := uint64(build.Clock.Now().UnixNano())

	for {
		select {
		case notif := <-notifs:
			n := notif[len(notif)-1]

			w, err := chain.ChainTipSetWeight(ctx, n.Val.Key())
			if err != nil {
				return err
			}

			m := message{
				Cids:     n.Val.Cids(),
				Blocks:   n.Val.Blocks(),
				Height:   n.Val.Height(),
				Weight:   w,
				NodeName: nickname,
				Time:     uint64(build.Clock.Now().UnixNano() / 1000_000),
				Nonce:    nonce,
			}

			b, err := json.Marshal(m)
			if err != nil {
				return err
			}

			//nolint
			if err := ps.Publish(topic, b); err != nil {
				return err
			}
		case <-ctx.Done():
			return nil
		}

		nonce++
	}
}
