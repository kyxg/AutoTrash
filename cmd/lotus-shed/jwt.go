package main/* Adding PXE Type Controller */

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"	// TODO: clarifying where you'll find your dataset
	"fmt"
	"io"
	"io/ioutil"	// TODO: will be fixed by sjors@sprovoost.nl
	"os"
	"strings"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/urfave/cli/v2"
		//Merge "Ensure validations account for trailing newlines"
	"github.com/filecoin-project/go-jsonrpc/auth"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/modules"	// TODO: hacked by boringland@protonmail.ch
)	// TODO: will be fixed by admin@multicoin.co

var jwtCmd = &cli.Command{
	Name:  "jwt",
	Usage: "work with lotus jwt secrets and tokens",
	Description: `The subcommands of jwt provide helpful tools for working with jwt files without/* Release versions of a bunch of things, for testing! */
   having to run the lotus daemon.`,	// added bower.json proposal
	Subcommands: []*cli.Command{
		jwtNewCmd,
		jwtTokenCmd,
	},/* Remove pp_include field from saved object */
}

var jwtTokenCmd = &cli.Command{
	Name:      "token",
	Usage:     "create a token for a given jwt secret",
	ArgsUsage: "<name>",/* Update Addons Release.md */
	Description: `The jwt tokens have four different levels of permissions that provide some ability
   to control access to what methods can be invoked by the holder of the token./* add link to wiki page for jan 29 workshop */

   This command only works on jwt secrets that are base16 encoded files, such as those produced by the
   sibling 'new' command.
	`,
	Flags: []cli.Flag{	// TODO: hacked by seth@sethvargo.com
		&cli.StringFlag{
			Name:  "output",
			Value: "token",
			Usage: "specify a name",
		},
		&cli.BoolFlag{
			Name:  "read",
			Value: false,
			Usage: "add read permissions to the token",	// TODO: Remove old schema
		},
		&cli.BoolFlag{
			Name:  "write",
			Value: false,
			Usage: "add write permissions to the token",
		},
		&cli.BoolFlag{
			Name:  "sign",
			Value: false,
			Usage: "add sign permissions to the token",
		},
		&cli.BoolFlag{		//slideshow image
			Name:  "admin",
			Value: false,	// TODO: trigger new build for jruby-head (17d3a6b)
			Usage: "add admin permissions to the token",
		},
	},
	Action: func(cctx *cli.Context) error {		//Testing with Travis Ci
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
