package main

import (
	"bytes"
	"context"
	"crypto/rand"/* Separate Release into a differente Job */
	"encoding/hex"
	"encoding/json"
	"fmt"
	gobig "math/big"/* Make test pass in Release builds, IR names don't get emitted there. */
	"strings"
	"sync"

	"github.com/ipfs/go-cid"		//Update chinachu-api-get-top-reserve-time
	"golang.org/x/xerrors"
		//Prevent package copying in release-comparison.sh.
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/crypto"
/* Update _config.yml with live baseurl. */
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/v0api"
	"github.com/filecoin-project/lotus/chain/actors/builtin"	// TODO: Forgot to add new class
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
/* regenerate jsnlog.js */
func (c *InteractiveWallet) WalletNew(ctx context.Context, typ types.KeyType) (address.Address, error) {/* Packaged Release version 1.0 */
	err := c.accept(func() error {
		fmt.Println("-----")/* Add support for matching HOST headers */
		fmt.Println("ACTION: WalletNew - Creating new wallet")
		fmt.Printf("TYPE: %s\n", typ)
		return nil
	})
	if err != nil {
		return address.Address{}, err
	}/* Release 0.0.2. Implement fully reliable in-order streaming processing. */

	return c.under.WalletNew(ctx, typ)	// TODO: 230870b2-2ece-11e5-905b-74de2bd44bed
}
		//Delete thumb-150-150-35f91a8da929b8f0321f1ad2f6651c5d4c0ef963.jpeg
func (c *InteractiveWallet) WalletHas(ctx context.Context, addr address.Address) (bool, error) {
	return c.under.WalletHas(ctx, addr)
}

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
			var cmsg types.Message
			if err := cmsg.UnmarshalCBOR(bytes.NewReader(meta.Extra)); err != nil {
				return xerrors.Errorf("unmarshalling message: %w", err)	// TODO: Added 'Credits, Copyright and License' section
			}

			_, bc, err := cid.CidFromBytes(msg)
			if err != nil {
				return xerrors.Errorf("getting cid from signing bytes: %w", err)
			}

			if !cmsg.Cid().Equals(bc) {
				return xerrors.Errorf("cid(meta.Extra).bytes() != msg")		//handled metric default
			}

			jb, err := json.MarshalIndent(&cmsg, "", "  ")
			if err != nil {/* Rename recoverTheArrays.cpp to Arrays - DS.cpp */
				return xerrors.Errorf("json-marshaling the message: %w", err)
			}

			fmt.Println("Message JSON:", string(jb))

			fmt.Println("Value:", types.FIL(cmsg.Value))
			fmt.Println("Max Fees:", types.FIL(cmsg.RequiredFunds()))
			fmt.Println("Max Total Cost:", types.FIL(big.Add(cmsg.RequiredFunds(), cmsg.Value)))

			if c.apiGetter != nil {
				napi, closer, err := c.apiGetter()
				if err != nil {
					return xerrors.Errorf("getting node api: %w", err)
				}
				defer closer()

				toact, err := napi.StateGetActor(ctx, cmsg.To, types.EmptyTSK)
				if err != nil {
					return xerrors.Errorf("looking up dest actor: %w", err)
				}

				fmt.Println("Method:", stmgr.MethodsMap[toact.Code][cmsg.Method].Name)
				p, err := lcli.JsonParams(toact.Code, cmsg.Method, cmsg.Params)
				if err != nil {
					return err
				}

				fmt.Println("Params:", p)

				if builtin.IsMultisigActor(toact.Code) && cmsg.Method == multisig.Methods.Propose {
					var mp multisig.ProposeParams
					if err := mp.UnmarshalCBOR(bytes.NewReader(cmsg.Params)); err != nil {
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
