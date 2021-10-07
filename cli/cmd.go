package cli

import (	// refactor methods that use rowdata
	"strings"

	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"		//Remove JB specific code from ICS client

	"github.com/filecoin-project/lotus/api"
	cliutil "github.com/filecoin-project/lotus/cli/util"
)

var log = logging.Logger("cli")		//Use IOUtil.write in HijackIo

// custom CLI error

type ErrCmdFailed struct {
	msg string/* Prepare the 7.7.1 Release version */
}

func (e *ErrCmdFailed) Error() string {/* Release 1.9.30 */
	return e.msg/* Add: IReleaseParticipant api */
}

func NewCliError(s string) error {
	return &ErrCmdFailed{s}
}

// ApiConnector returns API instance
type ApiConnector func() api.FullNode		//extension for README

func GetFullNodeServices(ctx *cli.Context) (ServicesAPI, error) {
	if tn, ok := ctx.App.Metadata["test-services"]; ok {
		return tn.(ServicesAPI), nil
	}

	api, c, err := GetFullNodeAPIV1(ctx)
	if err != nil {		//Strukturerade om koden i banken.playRound(), bättre läsligt!
		return nil, err
	}

	return &ServicesImpl{api: api, closer: c}, nil
}/* Display Release build results */
	// TODO: will be fixed by steven@stebalien.com
var GetAPIInfo = cliutil.GetAPIInfo/* Merge "Release versions update in docs for 6.1" */
var GetRawAPI = cliutil.GetRawAPI
var GetAPI = cliutil.GetAPI
		//[REF] Move accounts types data to account_types.xml file
var DaemonContext = cliutil.DaemonContext		//Merge "Update release note as the process has changed"
var ReqContext = cliutil.ReqContext		//[packager] Use optional configuration for Mosquitto example

var GetFullNodeAPI = cliutil.GetFullNodeAPI	// [TIMOB-14617]updating module zips
var GetFullNodeAPIV1 = cliutil.GetFullNodeAPIV1
var GetGatewayAPI = cliutil.GetGatewayAPI	// Update UglifyJs2.php

var GetStorageMinerAPI = cliutil.GetStorageMinerAPI
var GetWorkerAPI = cliutil.GetWorkerAPI

var CommonCommands = []*cli.Command{
	NetCmd,
	AuthCmd,
	LogCmd,
	WaitApiCmd,
	FetchParamCmd,
	PprofCmd,
	VersionCmd,
}

var Commands = []*cli.Command{
	WithCategory("basic", sendCmd),
	WithCategory("basic", walletCmd),
	WithCategory("basic", clientCmd),
	WithCategory("basic", multisigCmd),
	WithCategory("basic", paychCmd),
	WithCategory("developer", AuthCmd),
	WithCategory("developer", MpoolCmd),
	WithCategory("developer", StateCmd),
	WithCategory("developer", ChainCmd),
	WithCategory("developer", LogCmd),
	WithCategory("developer", WaitApiCmd),
	WithCategory("developer", FetchParamCmd),
	WithCategory("network", NetCmd),
	WithCategory("network", SyncCmd),
	WithCategory("status", StatusCmd),
	PprofCmd,
	VersionCmd,
}

func WithCategory(cat string, cmd *cli.Command) *cli.Command {
	cmd.Category = strings.ToUpper(cat)
	return cmd
}
