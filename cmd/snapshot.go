package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var snapshotCmd = &cobra.Command{
	Use:   "snapshot [data-type]",
	Short: "Get a snapshot of data at a specific point in time.",
	Long:  `Retrieves and displays a snapshot of data, such as the pair universe.`, 
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dataType := args[0]

		snapshots := SnapshotManager.GetSnapshots(dataType)

		if len(snapshots) == 0 {
			fmt.Printf("No snapshots found for data type: %s\n", dataType)
			return
		}

		fmt.Printf("-- Snapshots for %s --\n", dataType)
		for i, snapshot := range snapshots {
			fmt.Printf("Snapshot %d at %s:\n", i+1, snapshot.Timestamp.Format("2006-01-02 15:04:05"))
			fmt.Printf("%+v\n", snapshot.Data)
		}
	},
}

func init() {
	rootCmd.AddCommand(snapshotCmd)
}
