package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ApplicationTransferParam represents the ApplicationTransferParam schema from the OpenAPI specification
type ApplicationTransferParam struct {
	Key string `json:"key,omitempty"` // The type of the transfer parameter, such as `PRIVACY_LEVEL`.
	Value []string `json:"value,omitempty"` // The value of the transfer parameter, such as `PRIVATE` or `SHARED`.
}

// ApplicationsListResponse represents the ApplicationsListResponse schema from the OpenAPI specification
type ApplicationsListResponse struct {
	Applications []Application `json:"applications,omitempty"` // The list of applications that support data transfer and are also installed for the customer.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Kind string `json:"kind,omitempty"` // Identifies the resource as a collection of Applications.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token to specify the next page in the list.
}

// DataTransfer represents the DataTransfer schema from the OpenAPI specification
type DataTransfer struct {
	Applicationdatatransfers []ApplicationDataTransfer `json:"applicationDataTransfers,omitempty"` // The list of per-application data transfer resources. It contains details of the applications associated with this transfer resource, and also specifies the applications for which data transfer has to be done at the time of the transfer resource creation.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Id string `json:"id,omitempty"` // Read-only. The transfer's ID.
	Kind string `json:"kind,omitempty"` // Identifies the resource as a DataTransfer request.
	Newowneruserid string `json:"newOwnerUserId,omitempty"` // ID of the user to whom the data is being transferred.
	Oldowneruserid string `json:"oldOwnerUserId,omitempty"` // ID of the user whose data is being transferred.
	Overalltransferstatuscode string `json:"overallTransferStatusCode,omitempty"` // Read-only. Overall transfer status.
	Requesttime string `json:"requestTime,omitempty"` // Read-only. The time at which the data transfer was requested.
}

// DataTransfersListResponse represents the DataTransfersListResponse schema from the OpenAPI specification
type DataTransfersListResponse struct {
	Datatransfers []DataTransfer `json:"dataTransfers,omitempty"` // List of data transfer requests.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Kind string `json:"kind,omitempty"` // Identifies the resource as a collection of data transfer requests.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token to specify the next page in the list.
}

// Application represents the Application schema from the OpenAPI specification
type Application struct {
	Id string `json:"id,omitempty"` // The application's ID. Retrievable by using the [`applications.list()`](/admin-sdk/data-transfer/reference/rest/v1/applications/list) method.
	Kind string `json:"kind,omitempty"` // Identifies the resource as a DataTransfer Application Resource.
	Name string `json:"name,omitempty"` // The application's name.
	Transferparams []ApplicationTransferParam `json:"transferParams,omitempty"` // The list of all possible transfer parameters for this application. These parameters select which categories of the user's data to transfer.
	Etag string `json:"etag,omitempty"` // Etag of the resource.
}

// ApplicationDataTransfer represents the ApplicationDataTransfer schema from the OpenAPI specification
type ApplicationDataTransfer struct {
	Applicationtransferstatus string `json:"applicationTransferStatus,omitempty"` // Read-only. Current status of transfer for this application.
	Applicationid string `json:"applicationId,omitempty"` // The application's ID.
	Applicationtransferparams []ApplicationTransferParam `json:"applicationTransferParams,omitempty"` // The transfer parameters for the application. These parameters are used to select the data which will get transferred in context of this application. For more information about the specific values available for each application, see the [Transfer parameters](/admin-sdk/data-transfer/v1/parameters) reference.
}
