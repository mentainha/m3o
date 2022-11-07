package job

import (
	"fmt"
	"sync"

	"go.m3o.com/cron"
	"m3o.dev/apps/home/server/api"
)

var (
	mtx  sync.Mutex
	jobs = map[string]*cron.Job{
		"daily": {
			Id:       "checkin",
			Name:     "Daily Check-in",
			Interval: "0 9 * * *",
			Callback: api.URL(nil) + "/api/checkin",
		},
		"notify": {
			Id:       "notify",
			Name:     "Send Notifications",
			Interval: "*/5 * * * *",
			Callback: api.URL(nil) + "/api/notify",
		},
	}
)

// Register jobs
func Register() {
	mtx.Lock()
	defer mtx.Unlock()

	// lookup jobs
	rsp, err := api.Client.Cron.Jobs(&cron.JobsRequest{})
	if err != nil {
		fmt.Printf("Error retrieving existing jobs: %v\n", err)
		return
	}

	exists := make(map[string]*cron.Job)

	for _, job := range rsp.Jobs {
		exists[job.Id] = &job
	}

	// schedule all the jobs
	for _, job := range jobs {
		// lookup the job
		if _, ok := exists[job.Id]; ok {
			fmt.Printf("Job %s %s already scheduled\n", job.Id, job.Name)
			continue
		}

		fmt.Println("Scheduling", job.Id, job.Name)

		if _, err := api.Client.Cron.Schedule(&cron.ScheduleRequest{
			Id:       job.Id,
			Name:     job.Name,
			Interval: job.Interval,
			Callback: job.Callback,
		}); err != nil {
			fmt.Printf("Erroring scheduling %s: %v\n", job.Id, err)
		}
	}
}
