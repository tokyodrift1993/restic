package migrate

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// MachinesClient is the move your workloads to Azure.
type MachinesClient struct {
	BaseClient
}

// NewMachinesClient creates an instance of the MachinesClient client.
func NewMachinesClient(subscriptionID string, acceptLanguage string) MachinesClient {
	return NewMachinesClientWithBaseURI(DefaultBaseURI, subscriptionID, acceptLanguage)
}

// NewMachinesClientWithBaseURI creates an instance of the MachinesClient client.
func NewMachinesClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) MachinesClient {
	return MachinesClient{NewWithBaseURI(baseURI, subscriptionID, acceptLanguage)}
}

// Get get the machine with the specified name. Returns a json object of type 'machine' defined in Models section.
//
// resourceGroupName is name of the Azure Resource Group that project is part of. projectName is name of the Azure
// Migrate project. machineName is unique name of a machine in private datacenter.
func (client MachinesClient) Get(ctx context.Context, resourceGroupName string, projectName string, machineName string) (result Machine, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, projectName, machineName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.MachinesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.MachinesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.MachinesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client MachinesClient) GetPreparer(ctx context.Context, resourceGroupName string, projectName string, machineName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"machineName":       autorest.Encode("path", machineName),
		"projectName":       autorest.Encode("path", projectName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/machines/{machineName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(client.AcceptLanguage) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Accept-Language", autorest.String(client.AcceptLanguage)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client MachinesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client MachinesClient) GetResponder(resp *http.Response) (result Machine, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusUnauthorized, http.StatusNotFound, http.StatusInternalServerError, http.StatusServiceUnavailable),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByProject get data of all the machines available in the project. Returns a json array of objects of type
// 'machine' defined in Models section.
//
// resourceGroupName is name of the Azure Resource Group that project is part of. projectName is name of the Azure
// Migrate project.
func (client MachinesClient) ListByProject(ctx context.Context, resourceGroupName string, projectName string) (result MachineResultList, err error) {
	req, err := client.ListByProjectPreparer(ctx, resourceGroupName, projectName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.MachinesClient", "ListByProject", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByProjectSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.MachinesClient", "ListByProject", resp, "Failure sending request")
		return
	}

	result, err = client.ListByProjectResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.MachinesClient", "ListByProject", resp, "Failure responding to request")
	}

	return
}

// ListByProjectPreparer prepares the ListByProject request.
func (client MachinesClient) ListByProjectPreparer(ctx context.Context, resourceGroupName string, projectName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"projectName":       autorest.Encode("path", projectName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/machines", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(client.AcceptLanguage) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Accept-Language", autorest.String(client.AcceptLanguage)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByProjectSender sends the ListByProject request. The method will close the
// http.Response Body if it receives an error.
func (client MachinesClient) ListByProjectSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByProjectResponder handles the response to the ListByProject request. The method always
// closes the http.Response Body.
func (client MachinesClient) ListByProjectResponder(resp *http.Response) (result MachineResultList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusUnauthorized, http.StatusNotFound, http.StatusInternalServerError, http.StatusServiceUnavailable),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}