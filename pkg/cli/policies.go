package cli

import (
	"context"
	"io"
	"strings"

	"github.com/replicatedhq/gatekeeper/pkg/config"
	"github.com/replicatedhq/gatekeeper/pkg/gatekeeper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Policies(c *config.Config, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "policies",
		Short: "view and manage the policies deployed to the cluster",
		Long: `
View and perform some management operations on the policies installed on the cluster
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			g, err := gatekeeper.Get(c, out)
			if err != nil {
				return err
			}

			return g.RunPolicies(context.Background())
		},
	}

	viper.BindPFlags(cmd.Flags())
	viper.BindPFlags(cmd.PersistentFlags())
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	return cmd
}
