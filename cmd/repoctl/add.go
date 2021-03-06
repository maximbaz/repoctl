// Copyright (c) 2016, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package main

import "github.com/spf13/cobra"

var movePackages bool

func init() {
	MainCmd.AddCommand(addCmd)

	addCmd.Flags().BoolVarP(&movePackages, "move", "m", false, "move packages into repository")
}

var addCmd = &cobra.Command{
	Use:   "add <pkgfile...>",
	Short: "copy and add packages to the repository",
	Long: `Add (and copy if necessary) the package file to the repository.

  Similarly to the repo-add script, this command copies the package
  file to the repository (if not already there) and adds the package to
  the database. Exactly this package is added to the database, this
  allows you to downgrade a package in the repository.

  Any other package files in the repository are deleted or backed up,
  depending on whether the backup option is given. If the backup option
  is given, the "obsolete" package files are moved to a backup
  directory of choice.

  Note: since version 0.14, the semantic meaning of this command has
        changed. See the update command for the old behavior.
`,
	Example: `  repoctl add -m ./fairsplit-1.0.pkg.tar.gz`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if movePackages {
			return Repo.Move(nil, args...)
		} else {
			return Repo.Copy(nil, args...)
		}
	},
}
