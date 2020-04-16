// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDownloadClusterKubeconfigParams creates a new DownloadClusterKubeconfigParams object
// with the default values initialized.
func NewDownloadClusterKubeconfigParams() *DownloadClusterKubeconfigParams {
	var ()
	return &DownloadClusterKubeconfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadClusterKubeconfigParamsWithTimeout creates a new DownloadClusterKubeconfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDownloadClusterKubeconfigParamsWithTimeout(timeout time.Duration) *DownloadClusterKubeconfigParams {
	var ()
	return &DownloadClusterKubeconfigParams{

		timeout: timeout,
	}
}

// NewDownloadClusterKubeconfigParamsWithContext creates a new DownloadClusterKubeconfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewDownloadClusterKubeconfigParamsWithContext(ctx context.Context) *DownloadClusterKubeconfigParams {
	var ()
	return &DownloadClusterKubeconfigParams{

		Context: ctx,
	}
}

// NewDownloadClusterKubeconfigParamsWithHTTPClient creates a new DownloadClusterKubeconfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDownloadClusterKubeconfigParamsWithHTTPClient(client *http.Client) *DownloadClusterKubeconfigParams {
	var ()
	return &DownloadClusterKubeconfigParams{
		HTTPClient: client,
	}
}

/*DownloadClusterKubeconfigParams contains all the parameters to send to the API endpoint
for the download cluster kubeconfig operation typically these are written to a http.Request
*/
type DownloadClusterKubeconfigParams struct {

	/*ClusterID
	  The ID of the cluster whose kubeconfig to download

	*/
	ClusterID strfmt.UUID
	/*FileName
	  The kubeconfig file name

	*/
	FileName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the download cluster kubeconfig params
func (o *DownloadClusterKubeconfigParams) WithTimeout(timeout time.Duration) *DownloadClusterKubeconfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download cluster kubeconfig params
func (o *DownloadClusterKubeconfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download cluster kubeconfig params
func (o *DownloadClusterKubeconfigParams) WithContext(ctx context.Context) *DownloadClusterKubeconfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download cluster kubeconfig params
func (o *DownloadClusterKubeconfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download cluster kubeconfig params
func (o *DownloadClusterKubeconfigParams) WithHTTPClient(client *http.Client) *DownloadClusterKubeconfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download cluster kubeconfig params
func (o *DownloadClusterKubeconfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the download cluster kubeconfig params
func (o *DownloadClusterKubeconfigParams) WithClusterID(clusterID strfmt.UUID) *DownloadClusterKubeconfigParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the download cluster kubeconfig params
func (o *DownloadClusterKubeconfigParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WithFileName adds the fileName to the download cluster kubeconfig params
func (o *DownloadClusterKubeconfigParams) WithFileName(fileName string) *DownloadClusterKubeconfigParams {
	o.SetFileName(fileName)
	return o
}

// SetFileName adds the fileName to the download cluster kubeconfig params
func (o *DownloadClusterKubeconfigParams) SetFileName(fileName string) {
	o.FileName = fileName
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadClusterKubeconfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterId
	if err := r.SetPathParam("clusterId", o.ClusterID.String()); err != nil {
		return err
	}

	// query param fileName
	qrFileName := o.FileName
	qFileName := qrFileName
	if qFileName != "" {
		if err := r.SetQueryParam("fileName", qFileName); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
