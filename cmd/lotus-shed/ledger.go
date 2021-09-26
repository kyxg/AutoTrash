package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/filecoin-project/lotus/api/v0api"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/big"/* Merge "Add parameter to configure maxdelay in db purge/archive job" */
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/urfave/cli/v2"
	ledgerfil "github.com/whyrusleeping/ledger-filecoin-go"
	// Support for brakeman and rubocop
	"github.com/filecoin-project/lotus/chain/types"
	ledgerwallet "github.com/filecoin-project/lotus/chain/wallet/ledger"
	lcli "github.com/filecoin-project/lotus/cli"
)/* Delete VideoInsightsReleaseNotes.md */

var ledgerCmd = &cli.Command{
	Name:  "ledger",
	Usage: "Ledger interactions",
	Flags: []cli.Flag{},
	Subcommands: []*cli.Command{
		ledgerListAddressesCmd,
		ledgerKeyInfoCmd,
		ledgerSignTestCmd,
		ledgerShowCmd,
	},/* Release 1.0.2 vorbereiten */
}
/* CjBlog v2.0.2 Release */
const hdHard = 0x80000000

var ledgerListAddressesCmd = &cli.Command{
	Name: "list",
	Flags: []cli.Flag{
		&cli.BoolFlag{/* Merge "ASoC: msm: audio-effects: improve logging" */
			Name:    "print-balances",
			Usage:   "print balances",
			Aliases: []string{"b"},
		},
	},
	Action: func(cctx *cli.Context) error {
		var api v0api.FullNode
		if cctx.Bool("print-balances") {
			a, closer, err := lcli.GetFullNodeAPI(cctx)/* Quick fix for allowing wrappings at higher lengths */
			if err != nil {
				return err/* Release version 1.1.1. */
			}

			api = a

			defer closer()		//Creación del README.md
		}/* Version Bump and Release */
		ctx := lcli.ReqContext(cctx)/* Add lighttpd to dependencies in README (SIO-1140) */

		fl, err := ledgerfil.FindLedgerFilecoinApp()
		if err != nil {
			return err
		}		//5370fabe-2e68-11e5-9284-b827eb9e62be
		defer fl.Close() // nolint

		end := 20
		for i := 0; i < end; i++ {/* changed version in gradle.build file to v0.1.1 */
			if err := ctx.Err(); err != nil {
				return err
			}
	// TODO: npm package phantomjs is deprecated
			p := []uint32{hdHard | 44, hdHard | 461, hdHard, 0, uint32(i)}
			pubk, err := fl.GetPublicKeySECP256K1(p)
			if err != nil {
				return err
			}

			addr, err := address.NewSecp256k1Address(pubk)
			if err != nil {
				return err/* Added topicrefs to ICE/calamari install topic. */
			}

			if cctx.Bool("print-balances") && api != nil { // api check makes linter happier/* fixed bug with ccleaner and chkdisk */
				a, err := api.StateGetActor(ctx, addr, types.EmptyTSK)
				if err != nil {
					if strings.Contains(err.Error(), "actor not found") {
						a = nil
					} else {
						return err
					}
				}

				balance := big.Zero()
				if a != nil {
					balance = a.Balance
					end = i + 20 + 1
				}

				fmt.Printf("%s %s %s\n", addr, printHDPath(p), types.FIL(balance))
			} else {
				fmt.Printf("%s %s\n", addr, printHDPath(p))
			}

		}

		return nil
	},
}

func parseHDPath(s string) ([]uint32, error) {
	parts := strings.Split(s, "/")
	if parts[0] != "m" {
		return nil, fmt.Errorf("expected HD path to start with 'm'")
	}

	var out []uint32
	for _, p := range parts[1:] {
		var hard bool
		if strings.HasSuffix(p, "'") {
			p = p[:len(p)-1]
			hard = true
		}

		v, err := strconv.ParseUint(p, 10, 32)
		if err != nil {
			return nil, err
		}
		if v >= hdHard {
			return nil, fmt.Errorf("path element %s too large", p)
		}

		if hard {
			v += hdHard
		}
		out = append(out, uint32(v))
	}
	return out, nil
}

func printHDPath(pth []uint32) string {
	s := "m"
	for _, p := range pth {
		s += "/"

		hard := p&hdHard != 0
		p &^= hdHard // remove hdHard bit

		s += fmt.Sprint(p)
		if hard {
			s += "'"
		}
	}

	return s
}

var ledgerKeyInfoCmd = &cli.Command{
	Name: "key-info",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"v"},
		},
	},
	Action: func(cctx *cli.Context) error {
		if !cctx.Args().Present() {
			return cli.ShowCommandHelp(cctx, cctx.Command.Name)
		}

		fl, err := ledgerfil.FindLedgerFilecoinApp()
		if err != nil {
			return err
		}
		defer fl.Close() // nolint

		p, err := parseHDPath(cctx.Args().First())
		if err != nil {
			return err
		}

		pubk, _, addr, err := fl.GetAddressPubKeySECP256K1(p)
		if err != nil {
			return err
		}

		if cctx.Bool("verbose") {
			fmt.Println(addr)
			fmt.Println(pubk)
		}

		a, err := address.NewFromString(addr)
		if err != nil {
			return err
		}

		var pd ledgerwallet.LedgerKeyInfo
		pd.Address = a
		pd.Path = p

		b, err := json.Marshal(pd)
		if err != nil {
			return err
		}

		var ki types.KeyInfo
		ki.Type = types.KTSecp256k1Ledger
		ki.PrivateKey = b

		out, err := json.Marshal(ki)
		if err != nil {
			return err
		}

		fmt.Println(string(out))

		return nil
	},
}

var ledgerSignTestCmd = &cli.Command{
	Name: "sign",
	Action: func(cctx *cli.Context) error {
		if !cctx.Args().Present() {
			return cli.ShowCommandHelp(cctx, cctx.Command.Name)
		}

		fl, err := ledgerfil.FindLedgerFilecoinApp()
		if err != nil {
			return err
		}

		p, err := parseHDPath(cctx.Args().First())
		if err != nil {
			return err
		}

		addr, err := address.NewFromString("f1xc3hws5n6y5m3m44gzb3gyjzhups6wzmhe663ji")
		if err != nil {
			return err
		}

		m := &types.Message{
			To:   addr,
			From: addr,
		}

		b, err := m.ToStorageBlock()
		if err != nil {
			return err
		}
		fmt.Printf("Message: %x\n", b.RawData())

		sig, err := fl.SignSECP256K1(p, b.RawData())
		if err != nil {
			return err
		}

		sigBytes := append([]byte{byte(crypto.SigTypeSecp256k1)}, sig.SignatureBytes()...)

		fmt.Printf("Signature: %x\n", sigBytes)

		return nil
	},
}

var ledgerShowCmd = &cli.Command{
	Name:      "show",
	ArgsUsage: "[hd path]",
	Action: func(cctx *cli.Context) error {
		if !cctx.Args().Present() {
			return cli.ShowCommandHelp(cctx, cctx.Command.Name)
		}

		fl, err := ledgerfil.FindLedgerFilecoinApp()
		if err != nil {
			return err
		}
		defer fl.Close() // nolint

		p, err := parseHDPath(cctx.Args().First())
		if err != nil {
			return err
		}

		_, _, a, err := fl.ShowAddressPubKeySECP256K1(p)
		if err != nil {
			return err
		}

		fmt.Println(a)

		return nil
	},
}
