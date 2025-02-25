package main

import (
	"github.com/urfave/cli/v2"
)

func commands() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "list-followers",
			Aliases: []string{"followers"},
			Usage:   "List your followers",
			Action:  executeCmd(cmdListFollowers),
			Flags:   []cli.Flag{addListFlag()},
		},
		{
			Name:    "list-followings",
			Aliases: []string{"followings"},
			Usage:   "List your followings",
			Action:  executeCmd(cmdListFollowings),
			Flags:   []cli.Flag{addListFlag()},
		},
		{
			Name:    "clean-followings",
			Aliases: []string{"clean", "unfollow-unmutual", "remove-unmutual", "rm-unmutual"},
			Usage:   "Un follow not mutual followings, except of whitelisted",
			Action:  executeCmd(cmdCleanFollowings),
		},
		{
			Name:    "remove-followers",
			Aliases: []string{"rm", "remove"},
			Usage:   "Remove a list of followers, by username.",
			Action:  executeCmd(cmdRemoveFollowers),
			Flags:   []cli.Flag{addUsersFlag()},
		},
		{
			Name:    "unfollow-users",
			Aliases: []string{"unfollow", "remove-followings"},
			Usage:   "Unfollow a list of followings, by username.",
			Action:  executeCmd(cmdUnfollowUsers),
			Flags:   []cli.Flag{addUsersFlag()},
		},
		{
			Name:    "follow-users",
			Aliases: []string{"follow", "add-followings"},
			Usage:   "Follow a list of followings, by username.",
			Action:  executeCmd(cmdFollowUsers),
			Flags:   []cli.Flag{addUsersFlag()},
		},
		{
			Name:    "list-unmutual",
			Aliases: []string{"unmutual"},
			Usage:   "List all not mutual followings",
			Action:  executeCmd(cmdListNotMutual),
			Flags:   []cli.Flag{addListFlag()},
		},
		{
			Name:    "list-useless",
			Aliases: []string{"useless, bots"},
			Usage:   "List all statistic-useless accounts (bots, business accounts or mass-followers) (alpha)",
			Action:  executeCmd(cmdListUseless),
			Flags:   []cli.Flag{addListFlag()},
		},
		{
			Name:    "list-diff",
			Aliases: []string{"diff"},
			Usage:   "List diff for account (lost and new followers and followings)",
			Action:  executeCmd(cmdListDiff),
			Flags:   []cli.Flag{addListFlag()},
		},
		{
			Name:    "diff-history",
			Aliases: []string{"history"},
			Usage:   "List diff account history (lost and new followers and followings)",
			Action:  executeCmd(cmdListHistoryDiff),
		},
		{
			Name:    "upload",
			Aliases: []string{"u"},
			Usage:   "Upload media to profile",
			Action:  executeCmd(cmdUploadMedia),
			Flags:   uploadMediaFlags(),
		},
	}
}
