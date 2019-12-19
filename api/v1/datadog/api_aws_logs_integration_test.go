package datadog_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

func generateUniqueAwsLambdaAccounts() (datadog.AwsAccount, datadog.AwsAccountAndLambdaRequest, datadog.AwsLogsServicesRequest) {
	accountID := fmt.Sprintf("go_%09d", time.Now().UnixNano()%1000000000)
	var uniqueAwsAccount = datadog.AwsAccount{
		AccountId:                     &accountID,
		RoleName:                      datadog.PtrString("DatadogAWSIntegrationRole"),
		AccountSpecificNamespaceRules: &map[string]bool{"opsworks": true},
		FilterTags:                    &[]string{"testTag", "test:Tag2"},
		HostTags:                      &[]string{"filter:one", "filtertwo"},
	}

	var testLambdaArn = datadog.AwsAccountAndLambdaRequest{
		AccountId: *uniqueAwsAccount.AccountId,
		LambdaArn: "arn:aws:lambda:us-east-1:123456789101:function:GoClientTest",
	}

	var savedServices = datadog.AwsLogsServicesRequest{
		AccountId: *uniqueAwsAccount.AccountId,
		Services:  []string{"s3", "elb", "elbv2", "cloudfront", "redshift", "lambda"},
	}

	return uniqueAwsAccount, testLambdaArn, savedServices
}

// Test AddAWSLambdaARN and EnableServices endpoints
func TestAddAndSaveAWSLogs(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	testawsacc, testLambdaAcc, testServices := generateUniqueAwsLambdaAccounts()
	defer uninstallAWSIntegration(testawsacc)

	// Assert AWS Integration Created with proper fields
	_, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.CreateAWSAccount(TESTAUTH).AwsAccount(testawsacc).Execute()
	if err != nil {
		t.Fatalf("Error creating AWS Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	_, httpresp, err = TESTAPICLIENT.AWSLogsIntegrationApi.AddAWSLambdaARN(TESTAUTH).AwsAccountAndLambdaInput(testLambdaAcc).Execute()
	if err != nil {
		t.Errorf("Error adding lamda ARN: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpResp.StatusCode, 200)

	_, httpResp, err = TESTAPICLIENT.AWSLogsIntegrationApi.EnableAWSLogServices(TESTAUTH).AwsLogsServicesInput(testServices).Execute()
	if err != nil {
		t.Errorf("Error enabling log services: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
}

func TestListAWSLogsServices(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	listServicesOutput, httpResp, err := TESTAPICLIENT.AWSLogsIntegrationApi.AWSLogsServicesList(TESTAUTH).Execute()
	if err != nil {
		t.Errorf("Error listing log services: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpResp.StatusCode, 200)

	// There are currently 6 supported AWS Logs services as noted in the docs
	// https://docs.datadoghq.com/api/?lang=bash#get-list-of-aws-log-ready-services
	assert.Check(t, len(listServicesOutput) >= 6)
}

func TestListAndDeleteAWSLogs(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	testAWSAcc, testLambdaAcc, testServices := generateUniqueAwsLambdaAccounts()
	defer uninstallAWSIntegration(testAWSAcc)

	// Create the AWS integration.
	_, httpResp, err := TESTAPICLIENT.AWSIntegrationApi.CreateAWSAccount(TESTAUTH).AwsAccount(testAWSAcc).Execute()
	if err != nil {
		t.Fatalf("Error creating AWS Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpResp.StatusCode, 200)

	// Add Lambda to Account
	addOutput, httpresp, err := TESTAPICLIENT.AWSLogsIntegrationApi.AddAWSLambdaARN(TESTAUTH).AwsAccountAndLambdaInput(testLambdaAcc).Execute()
	if err != nil || httpresp.StatusCode != 200 {
		t.Errorf("Error Adding Lambda %v: Response %s: %v", addOutput, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpResp.StatusCode, 200)

	// Enable services for Lambda
	_, httpResp, err = TESTAPICLIENT.AWSLogsIntegrationApi.EnableAWSLogServices(TESTAUTH).AwsLogsServicesInput(testServices).Execute()
	if err != nil {
		t.Fatalf("Error enabling log services: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpResp.StatusCode, 200)

	// List AWS Logs integrations before deleting
	listOutput1, _, err := TESTAPICLIENT.AWSLogsIntegrationApi.AWSLogsList(TESTAUTH).Execute()
	if err != nil {
		t.Errorf("Error listing log services: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpResp.StatusCode, 200)
	// Iterate over output and list Lambdas
	var accountExists = false
	for _, Account := range listOutput1 {
		if Account.GetAccountId() == *testAWSAcc.AccountId {
			if Account.GetLambdas()[0].GetArn() == testLambdaAcc.LambdaArn {
				accountExists = true
			}
		}
	}
	// Test that variable is true as expected
	assert.Equal(t, accountExists, true)

	// Delete newly added Lambda
	deleteOutput, httpresp, err := TESTAPICLIENT.AWSLogsIntegrationApi.DeleteAWSLambdaARN(TESTAUTH, testLambdaAcc)
	if err != nil || httpresp.StatusCode != 200 {
		t.Errorf("Error Deleting Lambda %v: Status: %v: %v", deleteOutput, httpresp.StatusCode, err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	// List AWS logs integrations after deleting
	listOutput2, httpResp, err := TESTAPICLIENT.AWSLogsIntegrationApi.AWSLogsList(TESTAUTH).Execute()
	if err != nil {
		t.Errorf("Error listing log services: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpResp.StatusCode, 200)

	var listOfARNs2 []datadog.AwsLogsListResponseLambdas
	var accountExistsAfterDelete = false
	for _, Account := range listOutput2 {
		if Account.GetAccountId() == *testAWSAcc.AccountId {
			listOfARNs2 = Account.GetLambdas()
		}
	}
	for _, lambda := range listOfARNs2 {
		if lambda.GetArn() == testLambdaAcc.LambdaArn {
			accountExistsAfterDelete = true
		}
	}
	// Check that ARN no longer exists after delete
	assert.Assert(t, accountExistsAfterDelete != true)
}

func TestCheckLambdaAsync(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testAWSAcc, testLambdaAcc, _ := generateUniqueAwsLambdaAccounts()
	defer uninstallAWSIntegration(testAWSAcc)

	// Assert AWS Integration Created with proper fields
	_, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.CreateAWSAccount(TESTAUTH, testAWSAcc)
	if err != nil || httpresp.StatusCode != 200 {
		t.Fatalf("Error creating AWS Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}

	status, httpresp, err := TESTAPICLIENT.AWSLogsIntegrationApi.AWSLogsCheckLambdaAsync(TESTAUTH, testLambdaAcc)
	if err != nil || httpresp.StatusCode != 200 {
		t.Fatalf("Error checking the AWS Lambda: %v", err)
	}

	assert.Equal(t, len(status.GetErrors()), 0)
	assert.Equal(t, status.GetStatus(), "created")

	// Give the async call time to finish
	time.Sleep(10 * time.Second)

	status, httpresp, err = TESTAPICLIENT.AWSLogsIntegrationApi.AWSLogsCheckLambdaAsync(TESTAUTH, testLambdaAcc)
	if err != nil || httpresp.StatusCode != 200 {
		t.Fatalf("Error checking the AWS Lambda: %v", err)
	}

	assert.Assert(t, status.GetErrors()[0].GetCode() != "")
	assert.Assert(t, status.GetErrors()[0].GetMessage() != "")
	assert.Equal(t, status.GetStatus(), "error")
}

func TestCheckServicesAsync(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	testAWSAcc, _, testServices := generateUniqueAwsLambdaAccounts()
	defer uninstallAWSIntegration(testAWSAcc)

	// Assert AWS Integration Created with proper fields
	_, httpresp, err := TESTAPICLIENT.AWSIntegrationApi.CreateAWSAccount(TESTAUTH, testAWSAcc)
	if err != nil || httpresp.StatusCode != 200 {
		t.Fatalf("Error creating AWS Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}

	status, httpresp, err := TESTAPICLIENT.AWSLogsIntegrationApi.AWSLogsCheckServicesAsync(TESTAUTH, testServices)
	if err != nil || httpresp.StatusCode != 200 {
		t.Fatalf("Error checking the AWS Logs Services: %v", err)
	}
	assert.Assert(t, status.GetErrors()[0].GetCode() != "")
	assert.Assert(t, status.GetErrors()[0].GetMessage() != "")
}