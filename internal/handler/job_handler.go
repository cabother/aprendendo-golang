package handler

import (
	"cabother/aula/internal/dto"
	"cabother/aula/internal/service"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func NewJob(c *gin.Context) {

	receivedBody := dto.CreateJobRequestBody{}

	err := c.BindJSON(&receivedBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	job := dto.CreateJobService{
		Position: receivedBody.Position,
		Salary:   receivedBody.Salary,
		UserID:   receivedBody.UserID,
	}

	err = service.CreateJob(job)
	if err != nil {
		res := gin.H{"message": fmt.Sprintf("erro registrando o trabalho %v", job), "error": err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("trabalho %v registrado", job.Position)})
}
func RemoveJob(c *gin.Context) {
	jobID := c.Param("id")

	err := service.DeleteJob(jobID)
	if err != nil {
		res := gin.H{"message": fmt.Sprintf("erro saindo do trabalho %v", jobID), "error": err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("trabalho de id %v removido", jobID)})
}

func GetAllJobs(c *gin.Context) {

	JobModel, err := service.GetAllJobs()
	if err != nil {
		resp := gin.H{"message": "error getint the jobs", "error": err.Error()}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	list := []dto.GetJobResponse{}

	for _, item := range JobModel {

		list = append(list, dto.GetJobResponse{
			Position: item.Position,
			Salary:   item.Salary,
			UserID:   item.UserId,
		})

	}

	c.JSON(http.StatusOK, list)
}
func GetJobByID(c *gin.Context) {

	id := c.Param("id")

	idNumber, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		res := gin.H{"error": fmt.Sprintf("invalid id %s", id)}
		c.JSON(http.StatusBadRequest, res)

		return
	}

	JobModel, err := service.GetJobByID(idNumber)
	if err != nil {
		res := gin.H{"message": fmt.Sprintf("error geting id %s", id), "error": err.Error()}
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, res)
			return
		}

		c.JSON(http.StatusInternalServerError, res)

		return
	}

	JobResponse := dto.GetJobResponse{
		Position: JobModel.Position,
		Salary:   JobModel.Salary,
		UserID:   JobModel.UserId,
	}

	c.JSON(http.StatusOK, JobResponse)
}
func UpdateJobByID(c *gin.Context) {
	var receivedBody dto.UpdateJobRequestBody

	err := c.BindJSON(&receivedBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	job := dto.UpdateJobRequestBody{
		Position: receivedBody.Position,
		Salary:   receivedBody.Salary,
		UserID:   receivedBody.UserID,
	}

	jobID := c.Param("id")
	idNumber, err := strconv.ParseInt(jobID, 10, 64)
	if err != nil {
		res := gin.H{"error": fmt.Sprintf("invalid id %s", jobID)}
		c.JSON(http.StatusBadRequest, res)

		return
	}

	err = service.UpdateJobByID(idNumber, job)
	if err != nil {
		res := gin.H{"message": fmt.Sprintf("error updating id %s", jobID), "error": err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("trabalho %s atualizado com sucesso", job.Position)})
}
