package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
)

var schedule map[string]*gocron.Scheduler

func init() {
	schedule = map[string]*gocron.Scheduler{}
}

func CreateJob(c *gin.Context) {
	tag := c.Param("tag")
	if schedule[tag] != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("the job tagged as %s is already exists", tag),
		})
		return
	}
	s := gocron.NewScheduler(time.UTC)
	if _, err := s.Every(2).Seconds().Do(func() {
		fmt.Printf("2 sec: hello! tagged as %s", tag)
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unexpected error",
		})
		return
	}
	s.StartAsync()
	schedule[tag] = s
	c.JSON(http.StatusOK, gin.H{
		"message": "create new job, check console",
	})
}

func KillJob(c *gin.Context) {
	tag := c.Param("tag")
	if schedule[tag] == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("the job tagged as %s was not found", tag),
		})
		return
	}
	schedule[tag].Stop()
	schedule[tag] = nil
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("kill the job tagged as %s, check console", tag),
	})

}

func KillAllJob(c *gin.Context) {
	for _, s := range schedule {
		s.Stop()
	}
	schedule = map[string]*gocron.Scheduler{}
	c.JSON(http.StatusOK, gin.H{
		"message": "kill all job, check console",
	})
}
