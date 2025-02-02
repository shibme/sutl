package commands

import "github.com/spf13/cobra"

func XipherCommand() *cobra.Command {
	if xipherCmd != nil {
		return xipherCmd
	}
	xipherCmd = &cobra.Command{
		Use:   appNameLowerCase,
		Short: "Xipher is a curated collection of cryptographic primitives put together to perform password-based asymmetric encryption. It is written in Go and can be used as a library or a CLI tool.",
		Run: func(cmd *cobra.Command, args []string) {
			jsonFormat, _ := cmd.Flags().GetBool(jsonFlag.name)
			version, _ := cmd.Flags().GetBool(versionFlag.name)
			if version {
				showVersionInfo(jsonFormat)
			} else {
				cmd.Help()
			}
		},
	}
	xipherCmd.PersistentFlags().BoolP(jsonFlag.fields())
	xipherCmd.Flags().BoolP(versionFlag.fields())
	xipherCmd.AddCommand(versionCommand())
	xipherCmd.AddCommand(keygenCommand())
	xipherCmd.AddCommand(encryptCommand())
	xipherCmd.AddCommand(decryptCommand())
	return xipherCmd
}
