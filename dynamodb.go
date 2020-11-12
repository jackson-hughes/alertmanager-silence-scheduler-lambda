package main

/*func getSilencesFromInputEvent(svc dynamodbiface.DynamoDBAPI) ([]ScheduledSilence, error) {
	scheduleTableName := "Alertmanager-Scheduled-Silences"

	input := &dynamodb.ScanInput{
		TableName: aws.String(scheduleTableName),
	}

	result, err := svc.Scan(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				log.Error(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				log.Error(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				log.Error(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				log.Error(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				log.Error(aerr.Error())
			}
		}
		return nil, err
	}
	log.Debug("found records in DynamoDB:\n", result)

	records := []ScheduledSilence{}
	if err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &records); err != nil {
		return nil, err
	}

	for i := 0; i < len(records); i++ {
		r := &records[i]
		startTime, err := parseCronSchedule(r.StartScheduleCron, time.Now().UTC())
		if err != nil {
			log.Error(err)
		}
		r.StartsAt = startTime
		endTime, err := parseCronSchedule(r.EndScheduleCron, time.Now().UTC())
		if err != nil {
			log.Error(err)
		}
		r.EndsAt = endTime
	}

	recordsPretty, err := json.MarshalIndent(records, "", "    ")
	if err != nil {
		log.Error(err)
	}

	log.Debug("Records:\n", string(recordsPretty))
	return records, nil
}*/
