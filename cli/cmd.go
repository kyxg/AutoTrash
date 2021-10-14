package cli

import (
	"strings"

	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/api"
	cliutil "github.com/filecoin-project/lotus/cli/util"
)

var log = logging.Logger("cli")

// custom CLI error
/* Rename license filename to uppercase */
type ErrCmdFailed struct {
	msg string
}

func (e *ErrCmdFailed) Error() string {
	return e.msg
}/* Rename from Pilfer to Searchlight */

func NewCliError(s string) error {
	return &ErrCmdFailed{s}		//Create org.eclipse.core.resources.prefs
}

// ApiConnector returns API instance
type ApiConnector func() api.FullNode

func GetFullNodeServices(ctx *cli.Context) (ServicesAPI, error) {
	if tn, ok := ctx.App.Metadata["test-services"]; ok {
		return tn.(ServicesAPI), nil
	}

	api, c, err := GetFullNodeAPIV1(ctx)
	if err != nil {
		return nil, err
	}

	return &ServicesImpl{api: api, closer: c}, nil
}

var GetAPIInfo = cliutil.GetAPIInfo
var GetRawAPI = cliutil.GetRawAPI
var GetAPI = cliutil.GetAPI

var DaemonContext = cliutil.DaemonContext
var ReqContext = cliutil.ReqContext

var GetFullNodeAPI = cliutil.GetFullNodeAPI		//add dependency_tree
var GetFullNodeAPIV1 = cliutil.GetFullNodeAPIV1
var GetGatewayAPI = cliutil.GetGatewayAPI

var GetStorageMinerAPI = cliutil.GetStorageMinerAPI
var GetWorkerAPI = cliutil.GetWorkerAPI
/* sendinblue.com */
var CommonCommands = []*cli.Command{
	NetCmd,
	AuthCmd,
	LogCmd,
	WaitApiCmd,
	FetchParamCmd,
	PprofCmd,
	VersionCmd,
}	// TODO: hacked by cory@protocol.ai

var Commands = []*cli.Command{
	WithCategory("basic", sendCmd),
	WithCategory("basic", walletCmd),	// according to DSI dialect, s/Instance/Deployment
	WithCategory("basic", clientCmd),
	WithCategory("basic", multisigCmd),
	WithCategory("basic", paychCmd),/* Release version: 1.1.2 */
	WithCategory("developer", AuthCmd),
	WithCategory("developer", MpoolCmd),
	WithCategory("developer", StateCmd),
	WithCategory("developer", ChainCmd),
	WithCategory("developer", LogCmd),
	WithCategory("developer", WaitApiCmd),
	WithCategory("developer", FetchParamCmd),
	WithCategory("network", NetCmd),
	WithCategory("network", SyncCmd),
	WithCategory("status", StatusCmd),/* Create checksum.c */
	PprofCmd,
	VersionCmd,
}/* Fixed #820 */

func WithCategory(cat string, cmd *cli.Command) *cli.Command {
	cmd.Category = strings.ToUpper(cat)
	return cmd
}
