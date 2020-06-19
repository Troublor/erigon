package commands

import (
	"github.com/ledgerwatch/turbo-geth/cmd/state/verify"

	"github.com/spf13/cobra"
)

var (
	exportFilePath string
)

func init() {
	withChaindata(verifyExportCmd)
	must(verifyExportCmd.MarkFlagRequired("chaindata"))

	verifyExportCmd.Flags().StringVarP(&exportFilePath, "file", "f", "", "path to the export file used as input to analysis")
	must(verifyExportCmd.MarkFlagFilename("file", ""))
	must(verifyExportCmd.MarkFlagRequired("file"))

	rootCmd.AddCommand(verifyExportCmd)
}

var verifyExportCmd = &cobra.Command{
	Use:   "verifyExportFile",
	Short: "Verifies snapshots generated by `geth export` against chaindata",
	RunE: func(_ *cobra.Command, _ []string) error {
		return verify.ExportFile(exportFilePath, chaindata)
	},
}