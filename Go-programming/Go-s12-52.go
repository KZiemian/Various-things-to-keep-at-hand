jobID, err := store.PollNextJob()

if err != nil {
	return nil, fmt.Errorf("polling for next job: %w", err)
}

owner, err := store.FindOwnerByJobID(jobID)

if err != nil {
	return nil, fmt.Errorf("fetching job owner for job %s: %w", jobID,
		err)
}

j := jobs.New(jobID, owner)
res, errr := j.Start()

if err != nil {
	return nil, fmt.Errorf("starting job %s: %w", jobID, err)
}
