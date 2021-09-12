package main

import (
	"adjust/repos"
	"adjust/users_activity"
	"reflect"
	"testing"
)

func Test_checkFiles(t *testing.T) {
	t.Run("check files error", func(t *testing.T) {
		got := checkFiles("path")
		if got.Error() != "file actors.csv required" {
			t.Errorf("got = %v, want %v", got, "file actors.csv required")
		}
	})
}

func Test_usersByPushEventSorting(t *testing.T) {
	users := []users_activity.UsersActivity{
		{
			Id:                "22",
			Name:              "Test1",
			PushEventNumber:   22,
			WatchEventNumber:  22,
			PullEventNumber:   22,
			CreateEventNumber: 22,
			IssuesEventNumber: 22,
			DeleteEventNumber: 22,
		},
		{
			Id:                "33",
			Name:              "Test2",
			PushEventNumber:   33,
			WatchEventNumber:  33,
			PullEventNumber:   33,
			CreateEventNumber: 33,
			IssuesEventNumber: 33,
			DeleteEventNumber: 33,
		},
	}
	wantUsers := []users_activity.UsersActivity{
		{
			Id:                "33",
			Name:              "Test2",
			PushEventNumber:   33,
			WatchEventNumber:  33,
			PullEventNumber:   33,
			CreateEventNumber: 33,
			IssuesEventNumber: 33,
			DeleteEventNumber: 33,
		},
		{
			Id:                "22",
			Name:              "Test1",
			PushEventNumber:   22,
			WatchEventNumber:  22,
			PullEventNumber:   22,
			CreateEventNumber: 22,
			IssuesEventNumber: 22,
			DeleteEventNumber: 22,
		},
	}
	t.Run("", func(t *testing.T) {
		testUsers := users_activity.SortUsersByPushEvents(users)
		if !reflect.DeepEqual(wantUsers, testUsers) {
			t.Errorf("Test_usersByPushEventSorting want = %v, got %v", wantUsers, testUsers)
		}
	})
}

func Test_usersByPullEventSorting(t *testing.T) {
	users := []users_activity.UsersActivity{
		{
			Id:                "22",
			Name:              "Test1",
			PushEventNumber:   22,
			WatchEventNumber:  22,
			PullEventNumber:   22,
			CreateEventNumber: 22,
			IssuesEventNumber: 22,
			DeleteEventNumber: 22,
		},
		{
			Id:                "33",
			Name:              "Test2",
			PushEventNumber:   33,
			WatchEventNumber:  33,
			PullEventNumber:   33,
			CreateEventNumber: 33,
			IssuesEventNumber: 33,
			DeleteEventNumber: 33,
		},
	}
	wantUsers := []users_activity.UsersActivity{
		{
			Id:                "33",
			Name:              "Test2",
			PushEventNumber:   33,
			WatchEventNumber:  33,
			PullEventNumber:   33,
			CreateEventNumber: 33,
			IssuesEventNumber: 33,
			DeleteEventNumber: 33,
		},
		{
			Id:                "22",
			Name:              "Test1",
			PushEventNumber:   22,
			WatchEventNumber:  22,
			PullEventNumber:   22,
			CreateEventNumber: 22,
			IssuesEventNumber: 22,
			DeleteEventNumber: 22,
		},
	}
	t.Run("", func(t *testing.T) {
		testUsers := users_activity.SortUsersByPullEvent(users)
		if !reflect.DeepEqual(wantUsers, testUsers) {
			t.Errorf("Test_usersByPushEventSorting want = %v, got %v", wantUsers, testUsers)
		}
	})

}

func Test_usersByWatchEventSorting(t *testing.T) {
	users := []users_activity.UsersActivity{
		{
			Id:                "22",
			Name:              "Test1",
			PushEventNumber:   22,
			WatchEventNumber:  22,
			PullEventNumber:   22,
			CreateEventNumber: 22,
			IssuesEventNumber: 22,
			DeleteEventNumber: 22,
		},
		{
			Id:                "33",
			Name:              "Test2",
			PushEventNumber:   33,
			WatchEventNumber:  33,
			PullEventNumber:   33,
			CreateEventNumber: 33,
			IssuesEventNumber: 33,
			DeleteEventNumber: 33,
		},
	}
	wantUsers := []users_activity.UsersActivity{
		{
			Id:                "33",
			Name:              "Test2",
			PushEventNumber:   33,
			WatchEventNumber:  33,
			PullEventNumber:   33,
			CreateEventNumber: 33,
			IssuesEventNumber: 33,
			DeleteEventNumber: 33,
		},
		{
			Id:                "22",
			Name:              "Test1",
			PushEventNumber:   22,
			WatchEventNumber:  22,
			PullEventNumber:   22,
			CreateEventNumber: 22,
			IssuesEventNumber: 22,
			DeleteEventNumber: 22,
		},
	}
	t.Run("", func(t *testing.T) {
		testUsers := users_activity.SortUsersByWatchEvents(users)
		if !reflect.DeepEqual(wantUsers, testUsers) {
			t.Errorf("Test_usersByPushEventSorting want = %v, got %v", wantUsers, testUsers)
		}
	})

}

func Test_usersByCreateEventSorting(t *testing.T) {
	users := []users_activity.UsersActivity{
		{
			Id:                "22",
			Name:              "Test1",
			PushEventNumber:   22,
			WatchEventNumber:  22,
			PullEventNumber:   22,
			CreateEventNumber: 22,
			IssuesEventNumber: 22,
			DeleteEventNumber: 22,
		},
		{
			Id:                "33",
			Name:              "Test2",
			PushEventNumber:   33,
			WatchEventNumber:  33,
			PullEventNumber:   33,
			CreateEventNumber: 33,
			IssuesEventNumber: 33,
			DeleteEventNumber: 33,
		},
	}
	wantUsers := []users_activity.UsersActivity{
		{
			Id:                "33",
			Name:              "Test2",
			PushEventNumber:   33,
			WatchEventNumber:  33,
			PullEventNumber:   33,
			CreateEventNumber: 33,
			IssuesEventNumber: 33,
			DeleteEventNumber: 33,
		},
		{
			Id:                "22",
			Name:              "Test1",
			PushEventNumber:   22,
			WatchEventNumber:  22,
			PullEventNumber:   22,
			CreateEventNumber: 22,
			IssuesEventNumber: 22,
			DeleteEventNumber: 22,
		},
	}
	t.Run("", func(t *testing.T) {
		testUsers := users_activity.SortUsersByCreateEvent(users)
		if !reflect.DeepEqual(wantUsers, testUsers) {
			t.Errorf("Test_usersByPushEventSorting want = %v, got %v", wantUsers, testUsers)
		}
	})

}

func Test_usersByIssuesEventSorting(t *testing.T) {
	users := []users_activity.UsersActivity{
		{
			Id:                "22",
			Name:              "Test1",
			PushEventNumber:   22,
			WatchEventNumber:  22,
			PullEventNumber:   22,
			CreateEventNumber: 22,
			IssuesEventNumber: 22,
			DeleteEventNumber: 22,
		},
		{
			Id:                "33",
			Name:              "Test2",
			PushEventNumber:   33,
			WatchEventNumber:  33,
			PullEventNumber:   33,
			CreateEventNumber: 33,
			IssuesEventNumber: 33,
			DeleteEventNumber: 33,
		},
	}
	wantUsers := []users_activity.UsersActivity{
		{
			Id:                "33",
			Name:              "Test2",
			PushEventNumber:   33,
			WatchEventNumber:  33,
			PullEventNumber:   33,
			CreateEventNumber: 33,
			IssuesEventNumber: 33,
			DeleteEventNumber: 33,
		},
		{
			Id:                "22",
			Name:              "Test1",
			PushEventNumber:   22,
			WatchEventNumber:  22,
			PullEventNumber:   22,
			CreateEventNumber: 22,
			IssuesEventNumber: 22,
			DeleteEventNumber: 22,
		},
	}
	t.Run("", func(t *testing.T) {
		testUsers := users_activity.SortUsersByIssuesEvent(users)
		if !reflect.DeepEqual(wantUsers, testUsers) {
			t.Errorf("Test_usersByPushEventSorting want = %v, got %v", wantUsers, testUsers)
		}
	})

}

func Test_usersByDeleteEventSorting(t *testing.T) {
	users := []users_activity.UsersActivity{
		{
			Id:                "22",
			Name:              "Test1",
			PushEventNumber:   22,
			WatchEventNumber:  22,
			PullEventNumber:   22,
			CreateEventNumber: 22,
			IssuesEventNumber: 22,
			DeleteEventNumber: 22,
		},
		{
			Id:                "33",
			Name:              "Test2",
			PushEventNumber:   33,
			WatchEventNumber:  33,
			PullEventNumber:   33,
			CreateEventNumber: 33,
			IssuesEventNumber: 33,
			DeleteEventNumber: 33,
		},
	}
	wantUsers := []users_activity.UsersActivity{
		{
			Id:                "33",
			Name:              "Test2",
			PushEventNumber:   33,
			WatchEventNumber:  33,
			PullEventNumber:   33,
			CreateEventNumber: 33,
			IssuesEventNumber: 33,
			DeleteEventNumber: 33,
		},
		{
			Id:                "22",
			Name:              "Test1",
			PushEventNumber:   22,
			WatchEventNumber:  22,
			PullEventNumber:   22,
			CreateEventNumber: 22,
			IssuesEventNumber: 22,
			DeleteEventNumber: 22,
		},
	}
	t.Run("", func(t *testing.T) {
		testUsers := users_activity.SortUsersByDeleteEvent(users)
		if !reflect.DeepEqual(wantUsers, testUsers) {
			t.Errorf("Test_usersByPushEventSorting want = %v, got %v", wantUsers, testUsers)
		}
	})

}

func Test_usersByPushAndCreateEventSorting(t *testing.T) {
	users := []users_activity.UsersActivity{
		{
			Id:                "22",
			Name:              "Test1",
			PushEventNumber:   22,
			WatchEventNumber:  22,
			PullEventNumber:   22,
			CreateEventNumber: 22,
			IssuesEventNumber: 22,
			DeleteEventNumber: 22,
		},
		{
			Id:                "33",
			Name:              "Test2",
			PushEventNumber:   33,
			WatchEventNumber:  33,
			PullEventNumber:   33,
			CreateEventNumber: 33,
			IssuesEventNumber: 33,
			DeleteEventNumber: 33,
		},
	}
	wantUsers := []users_activity.UsersActivity{
		{
			Id:                "33",
			Name:              "Test2",
			PushEventNumber:   33,
			WatchEventNumber:  33,
			PullEventNumber:   33,
			CreateEventNumber: 33,
			IssuesEventNumber: 33,
			DeleteEventNumber: 33,
		},
		{
			Id:                "22",
			Name:              "Test1",
			PushEventNumber:   22,
			WatchEventNumber:  22,
			PullEventNumber:   22,
			CreateEventNumber: 22,
			IssuesEventNumber: 22,
			DeleteEventNumber: 22,
		},
	}
	t.Run("", func(t *testing.T) {
		testUsers := users_activity.SortUsersByPushAndCreateEventNumber(users)
		if !reflect.DeepEqual(wantUsers, testUsers) {
			t.Errorf("Test_usersByPushEventSorting want = %v, got %v", wantUsers, testUsers)
		}
	})

}

func Test_reposByPushEventSorting(t *testing.T) {
	r := []repos.Repos{
		{
			Id:                "22",
			Name:              "Test1",
			PushEventNumber:   22,
			WatchEventNumber:  22,
			PullEventNumber:   22,
			CreateEventNumber: 22,
			IssuesEventNumber: 22,
			DeleteEventNumber: 22,
		},
		{
			Id:                "33",
			Name:              "Test2",
			PushEventNumber:   33,
			WatchEventNumber:  33,
			PullEventNumber:   33,
			CreateEventNumber: 33,
			IssuesEventNumber: 33,
			DeleteEventNumber: 33,
		},
	}
	wantR := []repos.Repos{
		{
			Id:                "33",
			Name:              "Test2",
			PushEventNumber:   33,
			WatchEventNumber:  33,
			PullEventNumber:   33,
			CreateEventNumber: 33,
			IssuesEventNumber: 33,
			DeleteEventNumber: 33,
		},
		{
			Id:                "22",
			Name:              "Test1",
			PushEventNumber:   22,
			WatchEventNumber:  22,
			PullEventNumber:   22,
			CreateEventNumber: 22,
			IssuesEventNumber: 22,
			DeleteEventNumber: 22,
		},
	}
	t.Run("", func(t *testing.T) {
		testR := repos.SortByPushEvents(r)
		if !reflect.DeepEqual(wantR, testR) {
			t.Errorf("Test_usersByPushEventSorting want = %v, got %v", wantR, testR)
		}
	})
}

func Test_reposByPullEventSorting(t *testing.T) {
	r := []repos.Repos{
		{
			Id:                "22",
			Name:              "Test1",
			PushEventNumber:   22,
			WatchEventNumber:  22,
			PullEventNumber:   22,
			CreateEventNumber: 22,
			IssuesEventNumber: 22,
			DeleteEventNumber: 22,
		},
		{
			Id:                "33",
			Name:              "Test2",
			PushEventNumber:   33,
			WatchEventNumber:  33,
			PullEventNumber:   33,
			CreateEventNumber: 33,
			IssuesEventNumber: 33,
			DeleteEventNumber: 33,
		},
	}
	wantR := []repos.Repos{
		{
			Id:                "33",
			Name:              "Test2",
			PushEventNumber:   33,
			WatchEventNumber:  33,
			PullEventNumber:   33,
			CreateEventNumber: 33,
			IssuesEventNumber: 33,
			DeleteEventNumber: 33,
		},
		{
			Id:                "22",
			Name:              "Test1",
			PushEventNumber:   22,
			WatchEventNumber:  22,
			PullEventNumber:   22,
			CreateEventNumber: 22,
			IssuesEventNumber: 22,
			DeleteEventNumber: 22,
		},
	}
	t.Run("", func(t *testing.T) {
		testR := repos.SortByPullEvent(r)
		if !reflect.DeepEqual(wantR, testR) {
			t.Errorf("Test_usersByPushEventSorting want = %v, got %v", wantR, testR)
		}
	})
}

func Test_reposByWatchEventSorting(t *testing.T) {
	r := []repos.Repos{
		{
			Id:                "22",
			Name:              "Test1",
			PushEventNumber:   22,
			WatchEventNumber:  22,
			PullEventNumber:   22,
			CreateEventNumber: 22,
			IssuesEventNumber: 22,
			DeleteEventNumber: 22,
		},
		{
			Id:                "33",
			Name:              "Test2",
			PushEventNumber:   33,
			WatchEventNumber:  33,
			PullEventNumber:   33,
			CreateEventNumber: 33,
			IssuesEventNumber: 33,
			DeleteEventNumber: 33,
		},
	}
	wantR := []repos.Repos{
		{
			Id:                "33",
			Name:              "Test2",
			PushEventNumber:   33,
			WatchEventNumber:  33,
			PullEventNumber:   33,
			CreateEventNumber: 33,
			IssuesEventNumber: 33,
			DeleteEventNumber: 33,
		},
		{
			Id:                "22",
			Name:              "Test1",
			PushEventNumber:   22,
			WatchEventNumber:  22,
			PullEventNumber:   22,
			CreateEventNumber: 22,
			IssuesEventNumber: 22,
			DeleteEventNumber: 22,
		},
	}
	t.Run("", func(t *testing.T) {
		testR := repos.SortByWatchEvents(r)
		if !reflect.DeepEqual(wantR, testR) {
			t.Errorf("Test_usersByPushEventSorting want = %v, got %v", wantR, testR)
		}
	})
}

func Test_reposByCreateEventSorting(t *testing.T) {
	r := []repos.Repos{
		{
			Id:                "22",
			Name:              "Test1",
			PushEventNumber:   22,
			WatchEventNumber:  22,
			PullEventNumber:   22,
			CreateEventNumber: 22,
			IssuesEventNumber: 22,
			DeleteEventNumber: 22,
		},
		{
			Id:                "33",
			Name:              "Test2",
			PushEventNumber:   33,
			WatchEventNumber:  33,
			PullEventNumber:   33,
			CreateEventNumber: 33,
			IssuesEventNumber: 33,
			DeleteEventNumber: 33,
		},
	}
	wantR := []repos.Repos{
		{
			Id:                "33",
			Name:              "Test2",
			PushEventNumber:   33,
			WatchEventNumber:  33,
			PullEventNumber:   33,
			CreateEventNumber: 33,
			IssuesEventNumber: 33,
			DeleteEventNumber: 33,
		},
		{
			Id:                "22",
			Name:              "Test1",
			PushEventNumber:   22,
			WatchEventNumber:  22,
			PullEventNumber:   22,
			CreateEventNumber: 22,
			IssuesEventNumber: 22,
			DeleteEventNumber: 22,
		},
	}
	t.Run("", func(t *testing.T) {
		testR := repos.SortByCreateEvent(r)
		if !reflect.DeepEqual(wantR, testR) {
			t.Errorf("Test_usersByPushEventSorting want = %v, got %v", wantR, testR)
		}
	})
}

func Test_reposByIssuesEventSorting(t *testing.T) {
	r := []repos.Repos{
		{
			Id:                "22",
			Name:              "Test1",
			PushEventNumber:   22,
			WatchEventNumber:  22,
			PullEventNumber:   22,
			CreateEventNumber: 22,
			IssuesEventNumber: 22,
			DeleteEventNumber: 22,
		},
		{
			Id:                "33",
			Name:              "Test2",
			PushEventNumber:   33,
			WatchEventNumber:  33,
			PullEventNumber:   33,
			CreateEventNumber: 33,
			IssuesEventNumber: 33,
			DeleteEventNumber: 33,
		},
	}
	wantR := []repos.Repos{
		{
			Id:                "33",
			Name:              "Test2",
			PushEventNumber:   33,
			WatchEventNumber:  33,
			PullEventNumber:   33,
			CreateEventNumber: 33,
			IssuesEventNumber: 33,
			DeleteEventNumber: 33,
		},
		{
			Id:                "22",
			Name:              "Test1",
			PushEventNumber:   22,
			WatchEventNumber:  22,
			PullEventNumber:   22,
			CreateEventNumber: 22,
			IssuesEventNumber: 22,
			DeleteEventNumber: 22,
		},
	}
	t.Run("", func(t *testing.T) {
		testR := repos.SortByIssuesEvent(r)
		if !reflect.DeepEqual(wantR, testR) {
			t.Errorf("Test_usersByPushEventSorting want = %v, got %v", wantR, testR)
		}
	})
}

func Test_reposByDeleteEventSorting(t *testing.T) {
	r := []repos.Repos{
		{
			Id:                "22",
			Name:              "Test1",
			PushEventNumber:   22,
			WatchEventNumber:  22,
			PullEventNumber:   22,
			CreateEventNumber: 22,
			IssuesEventNumber: 22,
			DeleteEventNumber: 22,
		},
		{
			Id:                "33",
			Name:              "Test2",
			PushEventNumber:   33,
			WatchEventNumber:  33,
			PullEventNumber:   33,
			CreateEventNumber: 33,
			IssuesEventNumber: 33,
			DeleteEventNumber: 33,
		},
	}
	wantR := []repos.Repos{
		{
			Id:                "33",
			Name:              "Test2",
			PushEventNumber:   33,
			WatchEventNumber:  33,
			PullEventNumber:   33,
			CreateEventNumber: 33,
			IssuesEventNumber: 33,
			DeleteEventNumber: 33,
		},
		{
			Id:                "22",
			Name:              "Test1",
			PushEventNumber:   22,
			WatchEventNumber:  22,
			PullEventNumber:   22,
			CreateEventNumber: 22,
			IssuesEventNumber: 22,
			DeleteEventNumber: 22,
		},
	}
	t.Run("", func(t *testing.T) {
		testR := repos.SortByDeleteEvent(r)
		if !reflect.DeepEqual(wantR, testR) {
			t.Errorf("Test_usersByPushEventSorting want = %v, got %v", wantR, testR)
		}
	})
}
