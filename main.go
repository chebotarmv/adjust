package main

import (
	"adjust/repos"
	"adjust/users_activity"
	"flag"
	"log"
	"os"
)

func main() {
	path := flag.String("path", "", "Path to data directory. required - example [/home/example/]")
	us := flag.Bool("users", false, "Users statistics. optional")
	rep := flag.Bool("repos", false, "Repositories statistics. optional")
	top := flag.Int("top", 10, "Amount of users or repositories statistics to show. optional - max 100")
	sortBy := flag.String("sort", "push", "Sort by event {push|pull|watch|create|issue|delete| pull-push(only for users)}. optional")
	flag.Parse()
	if *path == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *top > 100 {
		*top = 100
	}

	start(*path, *us, *rep, *top, *sortBy)
}

func start(path string, us bool, rep bool, top int, sortBy string) {
	err := checkFiles(path)
	if err != nil {
		log.Fatal(err)
	}
	ev, err := getEvents(path)
	if err != nil {
		log.Fatal(err)
	}

	if us {
		users, err := users_activity.GetUsersActivity(ev)
		if err != nil {
			log.Fatal(err)
		}
		switch sortBy {
		case "push":
			users_activity.SortUsersByPushEvents(users)
		case "pull":
			users_activity.SortUsersByPullEvent(users)
		case "watch":
			users_activity.SortUsersByWatchEvents(users)
		case "create":
			users_activity.SortUsersByCreateEvent(users)
		case "issue":
			users_activity.SortUsersByIssuesEvent(users)
		case "delete":
			users_activity.SortUsersByDeleteEvent(users)
		case "pull-push":
			users_activity.SortUsersByPushAndCreateEventNumber(users)
		}
		users_activity.PrettyPrintUsers(users[:top])
	}

	if rep {
		repo, err := repos.GetReposInfo(ev)
		if err != nil {
			log.Fatal(err)
		}
		switch sortBy {
		case "push":
			repos.SortByPushEvents(repo)
		case "pull":
			repos.SortByPullEvent(repo)
		case "watch":
			repos.SortByWatchEvents(repo)
		case "create":
			repos.SortByCreateEvent(repo)
		case "issue":
			repos.SortByIssuesEvent(repo)
		case "delete":
			repos.SortByDeleteEvent(repo)
		}
		repos.PrettyPrintRepos(repo[:top])
	}

	os.Exit(1)

}
