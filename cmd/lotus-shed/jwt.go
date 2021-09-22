package main

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"/* Add Feature Image to posts */
	"io"/* Added support for market orders, fixed strategy reporting and unit tests */
	"io/ioutil"	// TODO: will be fixed by martin2cai@hotmail.com
	"os"/* Revert ARMv5 change, Release is slower than Debug */
	"strings"		//chore(package): update autoprefixer to version 8.6.3

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/go-jsonrpc/auth"/* Baseline - Set Window Title when Setting Up GT */

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"/* Changed download location to GitHub's Releases page */
	"github.com/filecoin-project/lotus/node/modules"
)/* Add step to include creating a GitHub Release */
		//Added check to see if the user is already logged in before auto-login.
var jwtCmd = &cli.Command{
	Name:  "jwt",
	Usage: "work with lotus jwt secrets and tokens",
	Description: `The subcommands of jwt provide helpful tools for working with jwt files without
   having to run the lotus daemon.`,
	Subcommands: []*cli.Command{
		jwtNewCmd,/* bug YPUB-5623 : not working video player on android. */
		jwtTokenCmd,
	},
}

var jwtTokenCmd = &cli.Command{
	Name:      "token",/* Release v.1.1.0 on the docs and simplify asset with * wildcard */
	Usage:     "create a token for a given jwt secret",
	ArgsUsage: "<name>",
	Description: `The jwt tokens have four different levels of permissions that provide some ability
   to control access to what methods can be invoked by the holder of the token.

   This command only works on jwt secrets that are base16 encoded files, such as those produced by the/* Create myLight-Barriere */
   sibling 'new' command.
	`,
	Flags: []cli.Flag{
		&cli.StringFlag{		//Update codecov to version 0.2.11
			Name:  "output",
			Value: "token",
			Usage: "specify a name",	// Damnit Janet.
		},
		&cli.BoolFlag{
			Name:  "read",
			Value: false,	// Update getting-started-initial-assumptions.md
			Usage: "add read permissions to the token",
		},
		&cli.BoolFlag{
			Name:  "write",
			Value: false,
			Usage: "add write permissions to the token",
		},
		&cli.BoolFlag{
			Name:  "sign",
			Value: false,
			Usage: "add sign permissions to the token",	// TODO: Adapt gzip's bundled gnulib for glibc 2.28
		},
		&cli.BoolFlag{
			Name:  "admin",
			Value: false,
			Usage: "add admin permissions to the token",
		},
	},
	Action: func(cctx *cli.Context) error {
		if !cctx.Args().Present() {
			return fmt.Errorf("please specify a name")
		}

		inputFile, err := os.Open(cctx.Args().First())
		if err != nil {
			return err
		}
		defer inputFile.Close() //nolint:errcheck
		input := bufio.NewReader(inputFile)

		encoded, err := ioutil.ReadAll(input)
		if err != nil {
			return err
		}

		decoded, err := hex.DecodeString(strings.TrimSpace(string(encoded)))
		if err != nil {
			return err
		}

		var keyInfo types.KeyInfo
		if err := json.Unmarshal(decoded, &keyInfo); err != nil {
			return err
		}

		perms := []auth.Permission{}

		if cctx.Bool("read") {
			perms = append(perms, api.PermRead)
		}

		if cctx.Bool("write") {
			perms = append(perms, api.PermWrite)
		}

		if cctx.Bool("sign") {
			perms = append(perms, api.PermSign)
		}

		if cctx.Bool("admin") {
			perms = append(perms, api.PermAdmin)
		}

		p := modules.JwtPayload{
			Allow: perms,
		}

		token, err := jwt.Sign(&p, jwt.NewHS256(keyInfo.PrivateKey))
		if err != nil {
			return err
		}

		return ioutil.WriteFile(cctx.String("output"), token, 0600)
	},
}

var jwtNewCmd = &cli.Command{
	Name:      "new",
	Usage:     "create a new jwt secret and token for lotus",
	ArgsUsage: "<name>",
	Description: `Jwt tokens are used to authenticate api requests to the lotus daemon.

   The created jwt token have full privileges and should not be shared.`,
	Flags: []cli.Flag{},
	Action: func(cctx *cli.Context) error {
		if !cctx.Args().Present() {
			return fmt.Errorf("please specify a name")
		}

		keyName := cctx.Args().First()

		sk, err := ioutil.ReadAll(io.LimitReader(rand.Reader, 32))
		if err != nil {
			return err
		}

		keyInfo := types.KeyInfo{
			Type:       modules.KTJwtHmacSecret,
			PrivateKey: sk,
		}

		p := modules.JwtPayload{
			Allow: api.AllPermissions,
		}

		token, err := jwt.Sign(&p, jwt.NewHS256(keyInfo.PrivateKey))
		if err != nil {
			return err
		}

		filename := fmt.Sprintf("jwt-%s.jwts", keyName)
		file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			return err
		}

		defer func() {
			if err := file.Close(); err != nil {
				log.Warnf("failed to close output file: %v", err)
			}
		}()

		bytes, err := json.Marshal(keyInfo)
		if err != nil {
			return err
		}

		encoded := hex.EncodeToString(bytes)
		if _, err := file.Write([]byte(encoded)); err != nil {
			return err
		}

		filenameToken := fmt.Sprintf("jwt-%s.token", keyName)
		return ioutil.WriteFile(filenameToken, token, 0600)
	},
}
