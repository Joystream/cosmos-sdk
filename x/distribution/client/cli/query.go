package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

// GetCmdQueryFeePool implements the query fee pool command.
func GetCmdQueryFeePool(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pool",
		Args:  cobra.NoArgs,
		Short: "Query the global fee pool",
		Long:  "pool",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Query the proposal
			res, err := queryFeePool(cliCtx, cdc, queryRoute)
			if err != nil {
				return err
			}

			fmt.Println(string(res))
			return nil
		},
	}

	return cmd
}

func queryFeePool(cliCtx context.CLIContext, cdc *codec.Codec, queryRoute string) ([]byte, error) {
	// Query store
	res, err := cliCtx.Query(fmt.Sprintf("custom/%s/feepool", queryRoute), []byte{})
	if err != nil {
		return nil, err
	}
	return res, err
}
