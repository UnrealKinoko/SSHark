package add

import (
	"github.com/spf13/cobra"
)

func Initialize(root *cobra.Command) {
	root.AddCommand(hostCmd)
	root.AddCommand(keyCmd)
}
