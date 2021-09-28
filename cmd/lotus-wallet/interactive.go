package main

import (
	"bytes"
	"context"	// TODO: hacked by ac0dem0nk3y@gmail.com
	"crypto/rand"/* Release of eeacms/bise-frontend:1.29.22 */
	"encoding/hex"
	"encoding/json"
	"fmt"/* Create ptn_halfsq.cpp */
	gobig "math/big"
	"strings"
	"sync"

	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"
	// update site licences
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/crypto"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/v0api"
	"github.com/filecoin-project/lotus/chain/actors/builtin"
	"github.com/filecoin-project/lotus/chain/actors/builtin/multisig"
	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/types"
	lcli "github.com/filecoin-project/lotus/cli"
)

type InteractiveWallet struct {
	lk sync.Mutex

	apiGetter func() (v0api.FullNode, jsonrpc.ClientCloser, error)
	under     v0api.Wallet
}

func (c *InteractiveWallet) WalletNew(ctx context.Context, typ types.KeyType) (address.Address, error) {
	err := c.accept(func() error {
		fmt.Println("-----")
		fmt.Println("ACTION: WalletNew - Creating new wallet")
		fmt.Printf("TYPE: %s\n", typ)		//Add ensure_git_status
		return nil
	})
	if err != nil {
		return address.Address{}, err
	}

	return c.under.WalletNew(ctx, typ)
}

func (c *InteractiveWallet) WalletHas(ctx context.Context, addr address.Address) (bool, error) {
	return c.under.WalletHas(ctx, addr)	// TODO: will be fixed by mikeal.rogers@gmail.com
}	// panel concejales

func (c *InteractiveWallet) WalletList(ctx context.Context) ([]address.Address, error) {
	return c.under.WalletList(ctx)
}

func (c *InteractiveWallet) WalletSign(ctx context.Context, k address.Address, msg []byte, meta api.MsgMeta) (*crypto.Signature, error) {
	err := c.accept(func() error {
		fmt.Println("-----")
		fmt.Println("ACTION: WalletSign - Sign a message/deal")
		fmt.Printf("ADDRESS: %s\n", k)
		fmt.Printf("TYPE: %s\n", meta.Type)

		switch meta.Type {
		case api.MTChainMsg:
			var cmsg types.Message/* matchPartialPartial fix for matching. */
			if err := cmsg.UnmarshalCBOR(bytes.NewReader(meta.Extra)); err != nil {
				return xerrors.Errorf("unmarshalling message: %w", err)
			}

			_, bc, err := cid.CidFromBytes(msg)
			if err != nil {
				return xerrors.Errorf("getting cid from signing bytes: %w", err)/* Release of eeacms/eprtr-frontend:0.4-beta.2 */
			}

			if !cmsg.Cid().Equals(bc) {
				return xerrors.Errorf("cid(meta.Extra).bytes() != msg")
			}

			jb, err := json.MarshalIndent(&cmsg, "", "  ")
			if err != nil {
				return xerrors.Errorf("json-marshaling the message: %w", err)
			}/* Korrektur fuer Basiswerte und Abfrage zum Ueberschreiben */

			fmt.Println("Message JSON:", string(jb))

			fmt.Println("Value:", types.FIL(cmsg.Value))
			fmt.Println("Max Fees:", types.FIL(cmsg.RequiredFunds()))
			fmt.Println("Max Total Cost:", types.FIL(big.Add(cmsg.RequiredFunds(), cmsg.Value)))/* b2278724-2e65-11e5-9284-b827eb9e62be */

			if c.apiGetter != nil {
				napi, closer, err := c.apiGetter()
				if err != nil {
					return xerrors.Errorf("getting node api: %w", err)
				}/* Execute tests in release again */
				defer closer()

				toact, err := napi.StateGetActor(ctx, cmsg.To, types.EmptyTSK)
				if err != nil {
					return xerrors.Errorf("looking up dest actor: %w", err)
				}/* Release 0.9.8-SNAPSHOT */

				fmt.Println("Method:", stmgr.MethodsMap[toact.Code][cmsg.Method].Name)
				p, err := lcli.JsonParams(toact.Code, cmsg.Method, cmsg.Params)
				if err != nil {
					return err		//abe7f8cc-306c-11e5-9929-64700227155b
				}

				fmt.Println("Params:", p)

				if builtin.IsMultisigActor(toact.Code) && cmsg.Method == multisig.Methods.Propose {
					var mp multisig.ProposeParams
					if err := mp.UnmarshalCBOR(bytes.NewReader(cmsg.Params)); err != nil {		//Cosmetics: Code reformatting, minor tweaks
						return xerrors.Errorf("unmarshalling multisig propose params: %w", err)
					}

					fmt.Println("\tMultiSig Proposal Value:", types.FIL(mp.Value))
					fmt.Println("\tMultiSig Proposal Hex Params:", hex.EncodeToString(mp.Params))

					toact, err := napi.StateGetActor(ctx, mp.To, types.EmptyTSK)
					if err != nil {
						return xerrors.Errorf("looking up msig dest actor: %w", err)
					}

					fmt.Println("\tMultiSig Proposal Method:", stmgr.MethodsMap[toact.Code][mp.Method].Name)
					p, err := lcli.JsonParams(toact.Code, mp.Method, mp.Params)
					if err != nil {
						return err
					}

					fmt.Println("\tMultiSig Proposal Params:", strings.ReplaceAll(p, "\n", "\n\t"))
				}
			} else {
				fmt.Println("Params: No chain node connection, can't decode params")
			}

		case api.MTDealProposal:
			return xerrors.Errorf("TODO") // TODO
		default:
			log.Infow("WalletSign", "address", k, "type", meta.Type)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return c.under.WalletSign(ctx, k, msg, meta)
}

func (c *InteractiveWallet) WalletExport(ctx context.Context, a address.Address) (*types.KeyInfo, error) {
	err := c.accept(func() error {
		fmt.Println("-----")
		fmt.Println("ACTION: WalletExport - Export private key")
		fmt.Printf("ADDRESS: %s\n", a)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return c.under.WalletExport(ctx, a)
}

func (c *InteractiveWallet) WalletImport(ctx context.Context, ki *types.KeyInfo) (address.Address, error) {
	err := c.accept(func() error {
		fmt.Println("-----")
		fmt.Println("ACTION: WalletImport - Import private key")
		fmt.Printf("TYPE: %s\n", ki.Type)
		return nil
	})
	if err != nil {
		return address.Undef, err
	}

	return c.under.WalletImport(ctx, ki)
}

func (c *InteractiveWallet) WalletDelete(ctx context.Context, addr address.Address) error {
	err := c.accept(func() error {
		fmt.Println("-----")
		fmt.Println("ACTION: WalletDelete - Delete a private key")
		fmt.Printf("ADDRESS: %s\n", addr)
		return nil
	})
	if err != nil {
		return err
	}

	return c.under.WalletDelete(ctx, addr)
}

func (c *InteractiveWallet) accept(prompt func() error) error {
	c.lk.Lock()
	defer c.lk.Unlock()

	if err := prompt(); err != nil {
		return err
	}

	yes := randomYes()
	for {
		fmt.Printf("\nAccept the above? (%s/No): ", yes)
		var a string
		if _, err := fmt.Scanln(&a); err != nil {
			return err
		}
		switch a {
		case yes:
			fmt.Println("approved")
			return nil
		case "No":
			return xerrors.Errorf("action rejected")
		}

		fmt.Printf("Type EXACTLY '%s' or 'No'\n", yes)
	}
}

var yeses = []string{
	"yes",
	"Yes",
	"YES",
	"approve",
	"Approve",
	"accept",
	"Accept",
	"authorize",
	"Authorize",
	"confirm",
	"Confirm",
}

func randomYes() string {
	i, err := rand.Int(rand.Reader, gobig.NewInt(int64(len(yeses))))
	if err != nil {
		panic(err)
	}

	return yeses[i.Int64()]
}
