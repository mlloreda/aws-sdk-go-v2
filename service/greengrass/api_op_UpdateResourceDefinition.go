// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/UpdateResourceDefinitionRequest
type UpdateResourceDefinitionInput struct {
	_ struct{} `type:"structure"`

	Name *string `type:"string"`

	// ResourceDefinitionId is a required field
	ResourceDefinitionId *string `location:"uri" locationName:"ResourceDefinitionId" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateResourceDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateResourceDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateResourceDefinitionInput"}

	if s.ResourceDefinitionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceDefinitionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateResourceDefinitionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceDefinitionId != nil {
		v := *s.ResourceDefinitionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ResourceDefinitionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/UpdateResourceDefinitionResponse
type UpdateResourceDefinitionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateResourceDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateResourceDefinitionOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateResourceDefinition = "UpdateResourceDefinition"

// UpdateResourceDefinitionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Updates a resource definition.
//
//    // Example sending a request using UpdateResourceDefinitionRequest.
//    req := client.UpdateResourceDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/UpdateResourceDefinition
func (c *Client) UpdateResourceDefinitionRequest(input *UpdateResourceDefinitionInput) UpdateResourceDefinitionRequest {
	op := &aws.Operation{
		Name:       opUpdateResourceDefinition,
		HTTPMethod: "PUT",
		HTTPPath:   "/greengrass/definition/resources/{ResourceDefinitionId}",
	}

	if input == nil {
		input = &UpdateResourceDefinitionInput{}
	}

	req := c.newRequest(op, input, &UpdateResourceDefinitionOutput{})
	return UpdateResourceDefinitionRequest{Request: req, Input: input, Copy: c.UpdateResourceDefinitionRequest}
}

// UpdateResourceDefinitionRequest is the request type for the
// UpdateResourceDefinition API operation.
type UpdateResourceDefinitionRequest struct {
	*aws.Request
	Input *UpdateResourceDefinitionInput
	Copy  func(*UpdateResourceDefinitionInput) UpdateResourceDefinitionRequest
}

// Send marshals and sends the UpdateResourceDefinition API request.
func (r UpdateResourceDefinitionRequest) Send(ctx context.Context) (*UpdateResourceDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateResourceDefinitionResponse{
		UpdateResourceDefinitionOutput: r.Request.Data.(*UpdateResourceDefinitionOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateResourceDefinitionResponse is the response type for the
// UpdateResourceDefinition API operation.
type UpdateResourceDefinitionResponse struct {
	*UpdateResourceDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateResourceDefinition request.
func (r *UpdateResourceDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}