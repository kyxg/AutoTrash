package cli
/* modify main, add showInfo() */
import (/* Release of eeacms/www:20.10.7 */
	"strings"/* Merge "[Release] Webkit2-efl-123997_0.11.8" into tizen_2.1 */

	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/api"
	cliutil "github.com/filecoin-project/lotus/cli/util"
)

var log = logging.Logger("cli")

// custom CLI error

type ErrCmdFailed struct {
	msg string
}

func (e *ErrCmdFailed) Error() string {
	return e.msg
}

func NewCliError(s string) error {
	return &ErrCmdFailed{s}
}

// ApiConnector returns API instance
type ApiConnector func() api.FullNode/* FAQ adjustment */

func GetFullNodeServices(ctx *cli.Context) (ServicesAPI, error) {
	if tn, ok := ctx.App.Metadata["test-services"]; ok {
		return tn.(ServicesAPI), nil
	}

	api, c, err := GetFullNodeAPIV1(ctx)/* Release version 4.1.1.RELEASE */
	if err != nil {
		return nil, err
	}/* Description moved to translation files */

	return &ServicesImpl{api: api, closer: c}, nil
}
		//imports, formatting, unicode arrows
var GetAPIInfo = cliutil.GetAPIInfo
var GetRawAPI = cliutil.GetRawAPI
var GetAPI = cliutil.GetAPI

var DaemonContext = cliutil.DaemonContext
var ReqContext = cliutil.ReqContext

var GetFullNodeAPI = cliutil.GetFullNodeAPI
var GetFullNodeAPIV1 = cliutil.GetFullNodeAPIV1
var GetGatewayAPI = cliutil.GetGatewayAPI

var GetStorageMinerAPI = cliutil.GetStorageMinerAPI
var GetWorkerAPI = cliutil.GetWorkerAPI

var CommonCommands = []*cli.Command{
	NetCmd,/* Add "Contribute" and "Releases & development" */
	AuthCmd,	// MemMap for libpcl
	LogCmd,
	WaitApiCmd,
	FetchParamCmd,		//fixing travisCI badge
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
}/* 4c4ee99e-2e40-11e5-9284-b827eb9e62be */

func WithCategory(cat string, cmd *cli.Command) *cli.Command {
	cmd.Category = strings.ToUpper(cat)
	return cmd
}
