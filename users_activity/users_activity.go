package users_activity

import (
	"adjust/helpers"
	"errors"
	"fmt"
	"sort"
)

type UsersActivity struct {
	Id                string
	Name              string
	PushEventNumber   int
	WatchEventNumber  int
	PullEventNumber   int
	CreateEventNumber int
	IssuesEventNumber int
	DeleteEventNumber int
}

func GetUsersActivity(events [][]string) ([]UsersActivity, error) {
	usersInfo, err := helpers.ReadCsv("./data/actors.csv")
	if err != nil {
		return nil, errors.New("can't get info from actors.csv file")
	}

	users := make(map[string]string, len(usersInfo)-1)
	for _, line := range usersInfo[1:] {
		users[line[0]] = line[1]
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
			PushEventNumberMap[event[2]] += 1
		case string(helpers.WatchEvent):
			WatchEventNumberMap[event[2]] += 1
		case string(helpers.PullRequestEvent):
			PullEventNumberMap[event[2]] += 1
		case string(helpers.CreateEvent):
			CreateEventNumberMap[event[2]] += 1
		case string(helpers.IssuesEvent):
			IssuesEventNumberMap[event[2]] += 1
		case string(helpers.DeleteEvent):
			DeleteEventNumberMap[event[2]] += 1
		default:

		}
	}

	response := make([]UsersActivity, 0, len(users))
	for userID, userName := range users {
		response = append(response, UsersActivity{
			Id:                userID,
			Name:              userName,
			PushEventNumber:   PushEventNumberMap[userID],
			WatchEventNumber:  WatchEventNumberMap[userID],
			PullEventNumber:   PullEventNumberMap[userID],
			CreateEventNumber: CreateEventNumberMap[userID],
			IssuesEventNumber: IssuesEventNumberMap[userID],
			DeleteEventNumber: DeleteEventNumberMap[userID],
		})

	}

	return response, nil
}

type UsersByPushEvents []UsersActivity

func (e UsersByPushEvents) Len() int           { return len(e) }
func (e UsersByPushEvents) Less(i, j int) bool { return e[i].PushEventNumber > e[j].PushEventNumber }
func (e UsersByPushEvents) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

type UsersByWatchEvents []UsersActivity

func (e UsersByWatchEvents) Len() int           { return len(e) }
func (e UsersByWatchEvents) Less(i, j int) bool { return e[i].WatchEventNumber > e[j].WatchEventNumber }
func (e UsersByWatchEvents) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

type UsersByPullEventNumber []UsersActivity

func (e UsersByPullEventNumber) Len() int { return len(e) }
func (e UsersByPullEventNumber) Less(i, j int) bool {
	return e[i].PullEventNumber > e[j].PullEventNumber
}
func (e UsersByPullEventNumber) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

type UsersByCreateEventNumber []UsersActivity

func (e UsersByCreateEventNumber) Len() int { return len(e) }
func (e UsersByCreateEventNumber) Less(i, j int) bool {
	return e[i].CreateEventNumber > e[j].CreateEventNumber
}
func (e UsersByCreateEventNumber) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

type UsersByIssuesEventNumber []UsersActivity

func (e UsersByIssuesEventNumber) Len() int { return len(e) }
func (e UsersByIssuesEventNumber) Less(i, j int) bool {
	return e[i].IssuesEventNumber > e[j].IssuesEventNumber
}
func (e UsersByIssuesEventNumber) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

type UsersByDeleteEventNumber []UsersActivity

func (e UsersByDeleteEventNumber) Len() int { return len(e) }
func (e UsersByDeleteEventNumber) Less(i, j int) bool {
	return e[i].DeleteEventNumber > e[j].DeleteEventNumber
}
func (e UsersByDeleteEventNumber) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

type UsersByPushAndCreateEventNumber []UsersActivity

func (e UsersByPushAndCreateEventNumber) Len() int { return len(e) }
func (e UsersByPushAndCreateEventNumber) Less(i, j int) bool {
	if e[i].PushEventNumber > e[j].PushEventNumber {
		return true
	}
	if e[i].CreateEventNumber < e[j].CreateEventNumber {
		return false
	}
	return e[i].PushEventNumber < e[j].PushEventNumber
}
func (e UsersByPushAndCreateEventNumber) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

func SortUsersByPushEvents(users []UsersActivity) []UsersActivity {
	sort.Sort(UsersByPushEvents(users))
	return users
}

func SortUsersByWatchEvents(users []UsersActivity) []UsersActivity {
	sort.Sort(UsersByWatchEvents(users))
	return users
}

func SortUsersByPullEvent(users []UsersActivity) []UsersActivity {
	sort.Sort(UsersByPullEventNumber(users))
	return users
}

func SortUsersByCreateEvent(users []UsersActivity) []UsersActivity {
	sort.Sort(UsersByCreateEventNumber(users))
	return users
}

func SortUsersByIssuesEvent(users []UsersActivity) []UsersActivity {
	sort.Sort(UsersByIssuesEventNumber(users))
	return users
}

func SortUsersByDeleteEvent(users []UsersActivity) []UsersActivity {
	sort.Sort(UsersByDeleteEventNumber(users))
	return users
}

func SortUsersByPushAndCreateEventNumber(users []UsersActivity) []UsersActivity {
	sort.Sort(UsersByPushAndCreateEventNumber(users))
	return users
}

func PrettyPrintUsers(users []UsersActivity) {
	fmt.Println(fmt.Sprintf("%s %s %s %s %s %s %s ", "Name", "PushEventNumber", "PullEventNumber", "WatchEventNumber", "CreateEventNumber", "DeleteEventNumber", "IssuesEventNumber"))
	for _, user := range users {
		fmt.Println(fmt.Sprintf("%s %d %d %d %d %d %d ", user.Name, user.PushEventNumber, user.PullEventNumber, user.WatchEventNumber, user.CreateEventNumber, user.DeleteEventNumber, user.IssuesEventNumber))
	}
}
