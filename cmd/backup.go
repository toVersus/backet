package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/toversus/backet/util"
)

var (
	destDir string
	srcDir  string
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "backup git repositories",
	Long:  `backup is used for obtaining backup of git repositories.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if srcDir != "" {
			if !filepath.IsAbs(srcDir) {
				util.ErrorExit("specify the absolute path for 'src' argument\n")
			}
			if !util.IsExist(srcDir) {
				util.ErrorExit("cannot find the specified GITBUCKET_HOME directory\n")
			}
		}
		if destDir != "" {
			if !filepath.IsAbs(destDir) {
				util.ErrorExit("specify the absolute path for 'dest' argument\n")
			}
		} else {
			util.ErrorExit("specify 'dest' argument\n")
		}
		if !util.IsGitInstalled() {
			util.ErrorExit("could not find git tool\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
	backupCmd.Flags().StringVarP(&destDir, "dest", "d", "", "full path to backup directory")
	backupCmd.Flags().StringVarP(&srcDir, "src", "s", "", "full path to GITBUCKET_HOME")
	backupCmd.RunE = doBackup
}

func doBackup(cmd *cobra.Command, args []string) error {
	return nil
}
