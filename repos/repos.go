package repos

import (
	"adjust/helpers"
	"errors"
	"fmt"
	"sort"
)

type Repos struct {
	Id                string
	Name              string
	PushEventNumber   int
	WatchEventNumber  int
	PullEventNumber   int
	CreateEventNumber int
	IssuesEventNumber int
	DeleteEventNumber int
}

type ByPushEvents []Repos

func (e ByPushEvents) Len() int           { return len(e) }
func (e ByPushEvents) Less(i, j int) bool { return e[i].PushEventNumber > e[j].PushEventNumber }
func (e ByPushEvents) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

type ByWatchEvents []Repos

func (e ByWatchEvents) Len() int           { return len(e) }
func (e ByWatchEvents) Less(i, j int) bool { return e[i].WatchEventNumber > e[j].WatchEventNumber }
func (e ByWatchEvents) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

type ByPullEventNumber []Repos

func (e ByPullEventNumber) Len() int           { return len(e) }
func (e ByPullEventNumber) Less(i, j int) bool { return e[i].PullEventNumber > e[j].PullEventNumber }
func (e ByPullEventNumber) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

type ByCreateEventNumber []Repos

func (e ByCreateEventNumber) Len() int { return len(e) }
func (e ByCreateEventNumber) Less(i, j int) bool {
	return e[i].CreateEventNumber > e[j].CreateEventNumber
}
func (e ByCreateEventNumber) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

type ByIssuesEventNumber []Repos

func (e ByIssuesEventNumber) Len() int { return len(e) }
func (e ByIssuesEventNumber) Less(i, j int) bool {
	return e[i].IssuesEventNumber > e[j].IssuesEventNumber
}
func (e ByIssuesEventNumber) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

type ByDeleteEventNumber []Repos

func (e ByDeleteEventNumber) Len() int { return len(e) }
func (e ByDeleteEventNumber) Less(i, j int) bool {
	return e[i].DeleteEventNumber > e[j].DeleteEventNumber
}
func (e ByDeleteEventNumber) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

func SortByPushEvents(repos []Repos) []Repos {
	sort.Sort(ByPushEvents(repos))
	return repos
}

func SortByWatchEvents(repos []Repos) []Repos {
	sort.Sort(ByWatchEvents(repos))
	return repos
}

func SortByPullEvent(repos []Repos) []Repos {
	sort.Sort(ByPullEventNumber(repos))
	return repos
}

func SortByCreateEvent(repos []Repos) []Repos {
	sort.Sort(ByCreateEventNumber(repos))
	return repos
}

func SortByIssuesEvent(repos []Repos) []Repos {
	sort.Sort(ByIssuesEventNumber(repos))
	return repos
}

func SortByDeleteEvent(repos []Repos) []Repos {
	sort.Sort(ByDeleteEventNumber(repos))
	return repos
}

func GetReposInfo(events [][]string) ([]Repos, error) {
	reposInfo, err := helpers.ReadCsv("./data/repos.csv")
	if err != nil {
		return nil, errors.New("can't get info from repos.csv file")
	}

	repos := make(map[string]string, len(reposInfo)-1)
	for _, line := range reposInfo[1:] {
		repos[line[0]] = line[1]
	}

	PushEventNumberMap := make(map[string]int)
	WatchEventNumberMap := make(map[string]int)
	PullEventNumberMap := make(map[string]int)
	CreateEventNumberMap := make(map[string]int)
	IssuesEventNumberMap := make(map[string]int)
	DeleteEventNumberMap := make(map[string]int)

	for _, event := range events[1:] {
		switch event[1] {
		case string(helpers.PushEvent):
			PushEventNumberMap[event[3]] += 1
		case string(helpers.WatchEvent):
			WatchEventNumberMap[event[3]] += 1
		case string(helpers.PullRequestEvent):
			PullEventNumberMap[event[3]] += 1
		case string(helpers.CreateEvent):
			CreateEventNumberMap[event[3]] += 1
		case string(helpers.IssuesEvent):
			IssuesEventNumberMap[event[3]] += 1
		case string(helpers.DeleteEvent):
			DeleteEventNumberMap[event[3]] += 1
		default:

		}
	}

	response := make([]Repos, 0, len(repos))
	for repoID, repoName := range repos {
		response = append(response, Repos{
			Id:                repoID,
			Name:              repoName,
			PushEventNumber:   PushEventNumberMap[repoID],
			WatchEventNumber:  WatchEventNumberMap[repoID],
			PullEventNumber:   PullEventNumberMap[repoID],
			CreateEventNumber: CreateEventNumberMap[repoID],
			IssuesEventNumber: IssuesEventNumberMap[repoID],
			DeleteEventNumber: DeleteEventNumberMap[repoID],
		})

	}

	return response, nil

}

func PrettyPrintRepos(repos []Repos) {
	fmt.Println(fmt.Sprintf("%s %s %s %s %s %s %s ", "Name", "PushEventNumber", "PullEventNumber", "WatchEventNumber", "CreateEventNumber", "DeleteEventNumber", "IssuesEventNumber"))
	for _, repo := range repos {
		fmt.Println(fmt.Sprintf("%s %d %d %d %d %d %d ", repo.Name, repo.PushEventNumber, repo.PullEventNumber, repo.WatchEventNumber, repo.CreateEventNumber, repo.DeleteEventNumber, repo.IssuesEventNumber))
	}
}
