package controllers

import (
	"net/http"

	"github.com/dhruvpatel-code/JobTrackerAPI/initializers"
	"github.com/dhruvpatel-code/JobTrackerAPI/models"
	"github.com/gin-gonic/gin"
)

func GetJobs(c *gin.Context) {
	var jobs []models.Job // Use a slice of Job objects
	if result := initializers.DB.Find(&jobs); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No jobs found"})
		return
	}
	c.JSON(http.StatusOK, jobs) // Return the slice of Job objects
}

// Add a new job
func AddJob(c *gin.Context) {
	var job models.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	initializers.DB.Create(&job)
	c.JSON(http.StatusCreated, job)
}
