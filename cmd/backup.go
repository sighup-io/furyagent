package cmd

import (
	"log"

	"github.com/sighupio/furyagent/pkg/component"
	"github.com/spf13/cobra"
)

// backupCmd represents the `furyctl backup` command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Executes backups",
	Long:  ``,
}

// etcdBackupCmd represents the `furyctl backup etcd` command
var etcdBackupCmd = &cobra.Command{
	Use:   "etcd",
	Short: "Backups etcd node",
	Long:  `Backups etcd node`,
	Run: func(cmd *cobra.Command, args []string) {
		// Does what is suppose to do
		var etcd component.ClusterComponent = component.Etcd{data}
		err := etcd.Backup()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
	backupCmd.AddCommand(etcdBackupCmd)
}
