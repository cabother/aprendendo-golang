package business

import (
	"cabother/aula/internal/dto"
	"cabother/aula/internal/models"
	"cabother/aula/internal/repository"
	"fmt"
)

func CreateJob(job dto.CreateJobService) error {
	if len(job.Position) < 1 {
		return fmt.Errorf("invalid Position %s", job.Position)
	}

	jobModel := models.JobModel{
		Position: job.Position,
		Salary:   job.Salary,
		UserId:   job.UserID,
	}

	err := repository.CreateJob(jobModel)

	return err
}
func DeleteJob(jobID string) error {
	if len(jobID) < 1 {
		return fmt.Errorf("id de serviço não encontrado %s", jobID)
	}

	err := repository.DeleteJob(jobID)
	if err != nil {
		return fmt.Errorf("erro saindo do seviço %s: %w", jobID, err)
	}

	return nil
}
func GetAllJobs() ([]models.JobModel, error) {
	jobs, err := repository.GetAllJobs()

	return jobs, err
}
func GetJobByID(id int64) (models.JobModel, error) {

	if id < 1 {
		return models.JobModel{}, fmt.Errorf("invalid id %d", id)
	}

	Job, err := repository.GetJobByID(id)
	return Job, err
}
func UpdateJobByID(id int64, job dto.UpdateJobRequestBody) error {

	jobModel := models.JobModel{
		Position: job.Position,
		Salary:   job.Salary,
		UserId:   job.UserID,
	}

	if len(job.Position) < 3 {
		return fmt.Errorf("invalid position %s", job.Position)
	}

	if job.UserID < 1 {
		return fmt.Errorf("invalid id %d", id)
	}

	err := repository.UpdateJobByID(id, jobModel)
	return err
}
