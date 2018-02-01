package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/pkg/errors"
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
}

func init() {
	rootCmd.AddCommand(backupCmd)
	backupCmd.Flags().StringVarP(&destDir, "dest", "d", "", "full path to backup directory")
	backupCmd.Flags().StringVarP(&srcDir, "src", "s", "", "full path to GITBUCKET_HOME")
	backupCmd.PreRun = validateArgs
	backupCmd.RunE = doBackup
}

func validateArgs(cmd *cobra.Command, args []string) {
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
		util.ErrorExit("could not find git installed\n")
	}
}

func doBackup(cmd *cobra.Command, args []string) error {
	if err := util.CreateDir(destDir); err != nil {
		msg := fmt.Sprintf("%#v", err)
		util.ErrorExit(msg)
	}
	fmt.Printf("Cloning all repositories...\n")
	return nil
}

// createMirrorRepository creates git clone into dest directory
// but skip cloning if dest directory already exists
func createMirrorRepository(src string, dest string) error {
	if util.IsExist(dest) {
		fmt.Printf("  %s is already exists, skip to git clone\n", dest)
	} else {
		fmt.Printf("  cloning %s into %s\n", src, dest)
		if err := util.GitCloneWithMirrorOpt(src, dest); err != nil {
			return errors.Wrap(err, "could not create a repository into a new directory\n")
		}
	}
	return nil
}
