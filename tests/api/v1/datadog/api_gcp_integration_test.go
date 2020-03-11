/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

func generateUniqueGCPAccount() (datadog.GCPAccount, datadog.GCPAccount) {
	projectID := fmt.Sprintf("go_%d", TESTCLOCK.Now().Unix())
	var testGCPAccount = datadog.GCPAccount{
		Type:                    datadog.PtrString("service_account"),
		ProjectId:               datadog.PtrString(projectID),
		PrivateKeyId:            datadog.PtrString("fake_private_key_id"),
		PrivateKey:              datadog.PtrString("fake_key"),
		ClientEmail:             datadog.PtrString("api-test@fake-sandbox.iam.gserviceaccount.com"),
		ClientId:                datadog.PtrString("123456712345671234567"),
		AuthUri:                 datadog.PtrString("fake_uri"),
		TokenUri:                datadog.PtrString("fake_uri"),
		AuthProviderX509CertUrl: datadog.PtrString("fake_url"),
		ClientX509CertUrl:       datadog.PtrString("fake_url"),
		HostFilters:             datadog.PtrString("fake:tag,example:test"),
		Automute:                datadog.PtrBool(false),
	}

	var testUpdateGCPAccount = datadog.GCPAccount{
		ProjectId:   datadog.PtrString(projectID),
		ClientEmail: testGCPAccount.ClientEmail,
		HostFilters: datadog.PtrString("fake:update,example:update"),
		Automute:    datadog.PtrBool(true),
	}

	return testGCPAccount, testUpdateGCPAccount
}

func TestGCPCreate(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)
	testGCPAcct, _ := generateUniqueGCPAccount()
	defer uninstallGCPIntegration(testGCPAcct)

	_, httpresp, err := TESTAPICLIENT.GCPIntegrationApi.CreateGCPIntegration(TESTAUTH).Body(testGCPAcct).Execute()
	if err != nil {
		t.Fatalf("Error creating GCP integration: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
}

func TestGCPListandDelete(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	testGCPAcct, _ := generateUniqueGCPAccount()
	defer uninstallGCPIntegration(testGCPAcct)

	// Setup GCP Account to List
	TESTAPICLIENT.GCPIntegrationApi.CreateGCPIntegration(TESTAUTH).Body(testGCPAcct).Execute()

	gcpListOutput, httpresp, err := TESTAPICLIENT.GCPIntegrationApi.ListGCPIntegration(TESTAUTH).Execute()
	if err != nil {
		t.Fatalf("Error listing GCP Accounts: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	var x datadog.GCPAccount
	for _, Account := range gcpListOutput {
		if Account.GetProjectId() == *testGCPAcct.ProjectId {
			x = Account
		}
	}
	assert.Equal(t, x.GetClientEmail(), *testGCPAcct.ClientEmail)
	assert.Equal(t, x.GetHostFilters(), *testGCPAcct.HostFilters)

	// Assert returned list is greater than or equal to 1
	assert.Assert(t, len(gcpListOutput) >= 1)

	// Test account deletion as well
	_, httpresp, err = TESTAPICLIENT.GCPIntegrationApi.DeleteGCPIntegration(TESTAUTH).Body(testGCPAcct).Execute()
	if err != nil {
		t.Fatalf("Error uninstalling GCP Account: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
}

func TestUpdateGCPAccount(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)
	testGCPAcct, testGCPUpdateAcct := generateUniqueGCPAccount()
	defer uninstallGCPIntegration(testGCPAcct)

	// Setup GCP Account to Update
	_, httpresp, err := TESTAPICLIENT.GCPIntegrationApi.CreateGCPIntegration(TESTAUTH).Body(testGCPAcct).Execute()
	if err != nil {
		t.Fatalf("Error creating GCP integration: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	_, httpresp, err = TESTAPICLIENT.GCPIntegrationApi.UpdateGCPIntegration(TESTAUTH).Body(testGCPUpdateAcct).Execute()
	if err != nil {
		t.Fatalf("Error updating GCP integration: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	// List account to ensure update worked.
	gcpListOutput, _, err := TESTAPICLIENT.GCPIntegrationApi.ListGCPIntegration(TESTAUTH).Execute()
	if err != nil {
		t.Fatalf("Error listing GCP accounts: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	var x datadog.GCPAccount
	for _, Account := range gcpListOutput {
		if Account.GetClientEmail() == *testGCPAcct.ClientEmail {
			x = Account
		}
	}
	assert.Equal(t, x.GetAutomute(), true)
	assert.Equal(t, x.GetHostFilters(), *testGCPUpdateAcct.HostFilters)
}

func uninstallGCPIntegration(account datadog.GCPAccount) {
	_, httpresp, err := TESTAPICLIENT.GCPIntegrationApi.DeleteGCPIntegration(TESTAUTH).Body(account).Execute()
	if httpresp.StatusCode != 200 || err != nil {
		log.Printf("Error uninstalling GCP Account: %v, Another test may have already removed this account.", account)
	}
}