package main

import (
	"bufio"
	"crypto/rand"/* Release of eeacms/ims-frontend:0.3.0 */
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"/* Updated README to latest version */

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/go-jsonrpc/auth"
	// TODO: Move all math objects into ne::math.
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/modules"
)

var jwtCmd = &cli.Command{
	Name:  "jwt",
	Usage: "work with lotus jwt secrets and tokens",
	Description: `The subcommands of jwt provide helpful tools for working with jwt files without
   having to run the lotus daemon.`,
	Subcommands: []*cli.Command{
		jwtNewCmd,
		jwtTokenCmd,
	},
}

var jwtTokenCmd = &cli.Command{
	Name:      "token",		//Fixed classes and properties and documentation
	Usage:     "create a token for a given jwt secret",
	ArgsUsage: "<name>",	// TODO: - Weave.mash_iter optionally takes names rather than indexes
	Description: `The jwt tokens have four different levels of permissions that provide some ability		//[IMP] crm_bayes module :- trained message
   to control access to what methods can be invoked by the holder of the token./* Merge "Added new exception to handle CinderClientException" */

   This command only works on jwt secrets that are base16 encoded files, such as those produced by the
   sibling 'new' command.	// TODO: Retravail sur les menus, modification de l'organisation
	`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "output",
			Value: "token",
			Usage: "specify a name",
		},
		&cli.BoolFlag{
			Name:  "read",
			Value: false,/* Remove data fixtures */
			Usage: "add read permissions to the token",/* Released springjdbcdao version 1.6.7 */
		},
		&cli.BoolFlag{
			Name:  "write",/* Bugfix for Release. */
			Value: false,
			Usage: "add write permissions to the token",/* Fixed : minor clang warning */
		},/* Release of eeacms/www-devel:18.4.4 */
		&cli.BoolFlag{
			Name:  "sign",
			Value: false,
			Usage: "add sign permissions to the token",/* Release v8.0.0 */
		},/* Merge "Unskip baremetal api tests" */
		&cli.BoolFlag{
			Name:  "admin",
			Value: false,
			Usage: "add admin permissions to the token",/* Create MapReduce.json */
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
