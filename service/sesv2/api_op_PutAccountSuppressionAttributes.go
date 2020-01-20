// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sesv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to change your account's suppression preferences.
type PutAccountSuppressionAttributesInput struct {
	_ struct{} `type:"structure"`

	// A list that contains the reasons that email addresses will be automatically
	// added to the suppression list for your account. This list can contain any
	// or all of the following:
	//
	//    * COMPLAINT – Amazon SES adds an email address to the suppression list
	//    for your account when a message sent to that address results in a complaint.
	//
	//    * BOUNCE – Amazon SES adds an email address to the suppression list
	//    for your account when a message sent to that address results in a hard
	//    bounce.
	SuppressedReasons []SuppressionListReason `type:"list"`
}

// String returns the string representation
func (s PutAccountSuppressionAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutAccountSuppressionAttributesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.SuppressedReasons != nil {
		v := s.SuppressedReasons

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "SuppressedReasons", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
type PutAccountSuppressionAttributesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutAccountSuppressionAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutAccountSuppressionAttributesOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutAccountSuppressionAttributes = "PutAccountSuppressionAttributes"

// PutAccountSuppressionAttributesRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Change the settings for the account-level suppression list.
//
//    // Example sending a request using PutAccountSuppressionAttributesRequest.
//    req := client.PutAccountSuppressionAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sesv2-2019-09-27/PutAccountSuppressionAttributes
func (c *Client) PutAccountSuppressionAttributesRequest(input *PutAccountSuppressionAttributesInput) PutAccountSuppressionAttributesRequest {
	op := &aws.Operation{
		Name:       opPutAccountSuppressionAttributes,
		HTTPMethod: "PUT",
		HTTPPath:   "/v2/email/account/suppression",
	}

	if input == nil {
		input = &PutAccountSuppressionAttributesInput{}
	}

	req := c.newRequest(op, input, &PutAccountSuppressionAttributesOutput{})
	return PutAccountSuppressionAttributesRequest{Request: req, Input: input, Copy: c.PutAccountSuppressionAttributesRequest}
}

// PutAccountSuppressionAttributesRequest is the request type for the
// PutAccountSuppressionAttributes API operation.
type PutAccountSuppressionAttributesRequest struct {
	*aws.Request
	Input *PutAccountSuppressionAttributesInput
	Copy  func(*PutAccountSuppressionAttributesInput) PutAccountSuppressionAttributesRequest
}

// Send marshals and sends the PutAccountSuppressionAttributes API request.
func (r PutAccountSuppressionAttributesRequest) Send(ctx context.Context) (*PutAccountSuppressionAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutAccountSuppressionAttributesResponse{
		PutAccountSuppressionAttributesOutput: r.Request.Data.(*PutAccountSuppressionAttributesOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutAccountSuppressionAttributesResponse is the response type for the
// PutAccountSuppressionAttributes API operation.
type PutAccountSuppressionAttributesResponse struct {
	*PutAccountSuppressionAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutAccountSuppressionAttributes request.
func (r *PutAccountSuppressionAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}