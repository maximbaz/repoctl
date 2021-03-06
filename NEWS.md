Repoctl Releases
================

## Version 0.17 (31 January 2018)
This release adds dependency resolution for the `down` command
and fixes a bug that occurs when trying to update a repository that
has more than 250 packages.

  - New: `down` command learned `-r` and `-o` flags that resolve dependencies
    and write a recommended order of compilation for any downloaded packages.
  - Bugfix: issue #28, in which AUR queries for a local database with more
    than 250 packages failed.
  - Update: better error messages when pre/post command actions fail.
  - Update: somewhat improved zsh completion (contributed by KoHcoJlb).
  - New: generated bash completion via the cobra library.

## Version 0.16 (22 November 2016)
This release adds action hooks to the configuration, and shows the configuration
when the `version` command is used.

  - New: `pre_action` and `post_action` string options have been added to the
    configuration. These commands are run in a local shell. They can be used
    to mount a remote filesystem where the database is located and dismount
    it afterwards.
  - Change: `version` command prints the values of the active configuration.
  - Change: `new config` command doesn't try to be smart about database
    extension anymore. It's just confusing.
  - Update: removing unnecessary error messages during repository creation.

## Version 0.15 (2 June 2016)
This release adds regex filtering support to the `list` command. A small
bug in the `status` command has been fixed, as well as with the pacman library.
Nothing major however.

  - New: `list` command learned to filter with regex argument
  - Update: documentation of list command improved.
  - Bugfix: status -m does not read AUR
  - Bugfix: reading repository without database failing

## Version 0.14 (26 April 2016)
This release rewrites repoctl to use the cobra library from spf13. Several
new commands are defined: `down` and `reset`, as well as two commands being
redefined `update` and `add`. Also, the core functionality is put into a
library to keep the main program small.

This release changes the AUR lookup functionality to use AUR4 instead of AUR.
When https://aur4.archlinux.org is the same as https://aur.archlinux.org, we
will revert this change (done).

Sorry this release has taken a while, that was not to annoy you. ;-)

With Go 1.5 allowing vendoring, we stop using gb (sorry).

  - New: repoctl learned command `host` for temporarily serving the repository
    on the network on a specified address and port. Something like darkhttpd may
    be better suited if the repository is to be hosted for longer periods of
    time.
  - New: repoctl learned command `down` with several flags. See the help
    message for more information on this. In short, we can download and extract
    tarballs for updated packages, all packages, and specified packages.
  - New: repoctl learned command `reset`.
  - Change: command `add` has completely different semantics. See the help.
  - Change: command `update` inherited the semantics of the old `add` command
    in addition to its current functionality.
  - Change: short form of `--outdated` flag for list has been changed from
    `-u` to `-o`
  - Update: AUR retrieval has been improved
  - Update: using spf13/cobra as our commandline engine now
  - Bugfix: license information correction (was BSD, is MIT)
  - Bugfix: typographical errors

## Version 0.13 (19 July 2015)
This release fixes a critical bug and updates a few other non-functional
files.

  - Bugfix: was not in correct directory when removing package files
  - Update: Zsh completion understands reset command

## Version 0.12 (17 July 2015)
This marks the first release where gb is used to build the project. That means
that all the dependencies for this project (apart from gb itself) are contained
within the project. Other changes are:

  - New: simple shortcutting filter commands is now possible. Instead of
    `aur.new` you can write `a.n`. At the moment, both parts are required.
  - New: library `shortry` to implement some of shortcutting behavior.
  - New: filter `db.missing` shows packages in local database which do not have
    respective files. These are candidates for deletion.
  - New: filter command can negate filters.
  - New: default configuration is written if there is no configuration. If
    default configuration is not edited, repoctl refuses to run.

There are probably many more changes, but at the moment I can't be bothered to
hunt them all down.

## Version 0.11 (18 December 2014)

  - New: The configuration file learned the field `ignore_aur`, which affects
    status and filter commands.

## Version 0.10 (1 October 2014)
This release changes the versioning scheme to use semantic versioning. Since
we are still changing a lot of program details and functionality, we are in
an unstable state, hence the major version number being 0. We chose 10 as
the minor number just because.

Additionally, there are some changes and updates:

  - New: Repoctl learned the `filter` command, which can take certain criteria.
  - Change: The `status` command takes no arguments anymore.
  - Change: The configuration file is in the [TOML](https://github.com/toml-lang/toml)
    format now, and it is being read (but not created yet).
