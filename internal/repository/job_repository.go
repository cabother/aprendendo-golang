package repository

import (
	"cabother/aula/internal/database"
	"cabother/aula/internal/models"
	"fmt"
)

func CreateJob(job models.JobModel) error {
	banco, err := database.ConnectDB()
	if err != nil {
		return err
	}

	execution := `insert into jobs(position, salary, user_id) values($1, $2, $3)`

	result, err := banco.Exec(execution, job.Position, job.Salary, job.UserId)
	if err != nil {
		return err
	}

	rowsAffected, errRowsAffected := result.RowsAffected()
	if errRowsAffected != nil {
		return errRowsAffected
	}

	if rowsAffected == 0 {
		return fmt.Errorf("job not registered")
	}

	return nil
}

func DeleteJob(jobID string) error {
	banco, err := database.ConnectDB()
	if err != nil {
		return err
	}

	execution := `delete from jobs where id = $1`

	result, err := banco.Exec(execution, jobID)
	if err != nil {
		return err
	}

	rowsAffected, errRowsAffected := result.RowsAffected()
	if errRowsAffected != nil {
		return errRowsAffected
	}

	if rowsAffected == 0 {
		return fmt.Errorf("job not found")
	}

	return nil
}

func GetAllJobs() ([]models.JobModel, error) {
	execution := `Select id, position, salary, user_id from jobs`
	banco, _ := database.ConnectDB()
	result, err := banco.Query(execution)
	if err != nil {
		return []models.JobModel{}, err
	}

	var jobs []models.JobModel

	for result.Next() {
		var job models.JobModel
		err = result.Scan(&job.ID, &job.Position, &job.Salary, &job.UserId)
		if err != nil {
			return []models.JobModel{}, err
		}

		jobs = append(jobs, job)
	}

	if len(jobs) == 0 {
		return []models.JobModel{}, fmt.Errorf("not found")
	}

	return jobs, nil
}

func GetJobByID(id int64) (models.JobModel, error) {
	execution := `Select id, position, salary from jobs where id = $1`
	banco, _ := database.ConnectDB()
	result, err := banco.Query(execution, id)
	if err != nil {
		return models.JobModel{}, err
	}

	var job models.JobModel
	for result.Next() {
		err = result.Scan(&job.ID, &job.Position, &job.Salary)
		if err != nil {
			return models.JobModel{}, err
		}
	}

	if job.ID == 0 {
		return models.JobModel{}, fmt.Errorf("id %d not found", id)
	}

	return job, nil
}

func UpdateJobByID(id int64, job models.JobModel) error {
	banco, err := database.ConnectDB()
	if err != nil {
		return err
	}

	query := `update jobs set position = $1, salary = $2, user_id = $3 where id = $4`

	result, err := banco.Exec(query, job.Position, job.Salary, job.UserId, id)
	if err != nil {
		return err
	}

	rowsAffected, errRowsAffected := result.RowsAffected()
	if errRowsAffected != nil {
		return errRowsAffected
	}

	if rowsAffected == 0 {
		return fmt.Errorf("job not updated")
	}

	return nil
}
