package main

import (
	"encoding/json"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Record struct {
	Service           string
	StartScheduleCron string
	EndScheduleCron   string
	Matchers          []string
	StartsAt          time.Time
	EndsAt            time.Time
}

func getScheduledSilences() ([]Record, error) {
	scheduleTableName := "Alertmanager-Scheduled-Silences"
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)
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
	log.Debug("found the following records in DynamoDB:\n", result)

	records := []Record{}
	if err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &records); err != nil {
		return nil, err
	}

	for i := 0; i < len(records); i++ {
		r := &records[i]
		startTime, err := parseCronSchedule(r.StartScheduleCron)
		if err != nil {
			log.Error(err)
		}
		r.StartsAt = startTime
		endTime, err := parseCronSchedule(r.EndScheduleCron)
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
}
