package main

/*
import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type mockDynamoDBClient struct {
	dynamodbiface.DynamoDBAPI
}

func (m *mockDynamoDBClient) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	respContent := []byte(`{
	  Count: 2,
	  Items: [{
	      service: {
	        S: "mobilev1"
	      },
	      endScheduleCron: {
	        S: "0 4 * * 0"
	      },
	      matchers: {
	        L: [{
	            M: {
	              name: {
	                S: "account"
	              },
	              value: {
	                S: "production-services"
	              }
	            }
	          },{
	            M: {
	              name: {
	                S: "alertname"
	              },
	              value: {
	                S: "BoomiApiAlerts"
	              }
	            }
	          }]
	      },
	      startScheduleCron: {
	        S: "0 3 * * 0"
	      }
	    },{
	      startScheduleCron: {
	        S: "0 2 * * 0"
	      }],
	  ScannedCount: 1
	}`)

	testRecordStartTime, _ := time.Parse(time.RFC3339, "2020-08-09T03:00:00Z")
	testRecordEndTime, _ := time.Parse(time.RFC3339, "2020-08-09T04:00:00Z")
	// items := []ScheduledSilence{}
	r := ScheduledSilence{
		Service:           "test",
		StartScheduleCron: "0 3 * * 0",
		EndScheduleCron:   "0 4 * * 0",
		Matchers: []Matcher{
			{
				IsRegex: false,
				Name:    "environment",
				Value:   "test",
			},
		},
		StartsAt: testRecordStartTime,
		EndsAt:   testRecordEndTime,
	}
	// items = append(items, r)

	av, err := dynamodbattribute.Marshal(r)
	if err != nil {
		fmt.Println(err)
	}

	return &dynamodb.ScanOutput{
		ConsumedCapacity: nil,
		Count:            nil,
		Items:            []map[string]*dynamodb.AttributeValue{},
		LastEvaluatedKey: nil,
		ScannedCount:     nil,
	}, nil
 }*/

/*func TestGetScheduledSilences(t *testing.T) {
	testRecordStartTime, _ := time.Parse(time.RFC3339, "2020-08-09T03:00:00Z")
	testRecordEndTime, _ := time.Parse(time.RFC3339, "2020-08-09T04:00:00Z")

	want := []ScheduledSilence{
		{
			Service:           "test",
			StartScheduleCron: "0 3 * * 0",
			EndScheduleCron:   "0 4 * * 0",
			Matchers: []Matcher{
				{
					IsRegex: false,
					Name:    "environment",
					Value:   "test",
				},
				{
					IsRegex: false,
					Name:    "alertname",
					Value:   "HostDown",
				},
			},
			StartsAt: testRecordStartTime,
			EndsAt:   testRecordEndTime,
		},
	}

	mockSvc := &mockDynamoDBClient{}
	got, _ := getSilencesFromInputEvent(mockSvc)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v but got %v", got, want)
	}
}*/
