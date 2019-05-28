// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package groundstation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/UpdateMissionProfileRequest
type UpdateMissionProfileInput struct {
	_ struct{} `type:"structure"`

	// Amount of time after a contact ends that you’d like to receive a CloudWatch
	// event indicating the pass has finished.
	ContactPostPassDurationSeconds *int64 `locationName:"contactPostPassDurationSeconds" min:"1" type:"integer"`

	// Amount of time after a contact ends that you’d like to receive a CloudWatch
	// event indicating the pass has finished.
	ContactPrePassDurationSeconds *int64 `locationName:"contactPrePassDurationSeconds" min:"1" type:"integer"`

	// A list of lists of ARNs. Each list of ARNs is an edge, with a from Config
	// and a to Config.
	DataflowEdges [][]string `locationName:"dataflowEdges" type:"list"`

	// Smallest amount of time in seconds that you’d like to see for an available
	// contact. AWS Ground Station will not present you with contacts shorter than
	// this duration.
	MinimumViableContactDurationSeconds *int64 `locationName:"minimumViableContactDurationSeconds" min:"1" type:"integer"`

	// ID of a mission profile.
	//
	// MissionProfileId is a required field
	MissionProfileId *string `location:"uri" locationName:"missionProfileId" type:"string" required:"true"`

	// Name of a mission profile.
	Name *string `locationName:"name" min:"1" type:"string"`

	// ARN of a tracking Config.
	TrackingConfigArn *string `locationName:"trackingConfigArn" type:"string"`
}

// String returns the string representation
func (s UpdateMissionProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateMissionProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateMissionProfileInput"}
	if s.ContactPostPassDurationSeconds != nil && *s.ContactPostPassDurationSeconds < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("ContactPostPassDurationSeconds", 1))
	}
	if s.ContactPrePassDurationSeconds != nil && *s.ContactPrePassDurationSeconds < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("ContactPrePassDurationSeconds", 1))
	}
	if s.MinimumViableContactDurationSeconds != nil && *s.MinimumViableContactDurationSeconds < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MinimumViableContactDurationSeconds", 1))
	}

	if s.MissionProfileId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MissionProfileId"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateMissionProfileInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ContactPostPassDurationSeconds != nil {
		v := *s.ContactPostPassDurationSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "contactPostPassDurationSeconds", protocol.Int64Value(v), metadata)
	}
	if s.ContactPrePassDurationSeconds != nil {
		v := *s.ContactPrePassDurationSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "contactPrePassDurationSeconds", protocol.Int64Value(v), metadata)
	}
	if len(s.DataflowEdges) > 0 {
		v := s.DataflowEdges

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "dataflowEdges", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls1 := ls0.List()
			ls1.Start()
			for _, v2 := range v1 {
				ls1.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v2)})
			}
			ls1.End()
		}
		ls0.End()

	}
	if s.MinimumViableContactDurationSeconds != nil {
		v := *s.MinimumViableContactDurationSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "minimumViableContactDurationSeconds", protocol.Int64Value(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TrackingConfigArn != nil {
		v := *s.TrackingConfigArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "trackingConfigArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MissionProfileId != nil {
		v := *s.MissionProfileId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "missionProfileId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/MissionProfileIdResponse
type UpdateMissionProfileOutput struct {
	_ struct{} `type:"structure"`

	// ID of a mission profile.
	MissionProfileId *string `locationName:"missionProfileId" type:"string"`
}

// String returns the string representation
func (s UpdateMissionProfileOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateMissionProfileOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.MissionProfileId != nil {
		v := *s.MissionProfileId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "missionProfileId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateMissionProfile = "UpdateMissionProfile"

// UpdateMissionProfileRequest returns a request value for making API operation for
// AWS Ground Station.
//
// Updates a mission profile.
//
// Updating a mission profile will not update the execution parameters for existing
// future contacts.
//
//    // Example sending a request using UpdateMissionProfileRequest.
//    req := client.UpdateMissionProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/UpdateMissionProfile
func (c *Client) UpdateMissionProfileRequest(input *UpdateMissionProfileInput) UpdateMissionProfileRequest {
	op := &aws.Operation{
		Name:       opUpdateMissionProfile,
		HTTPMethod: "PUT",
		HTTPPath:   "/missionprofile/{missionProfileId}",
	}

	if input == nil {
		input = &UpdateMissionProfileInput{}
	}

	req := c.newRequest(op, input, &UpdateMissionProfileOutput{})
	return UpdateMissionProfileRequest{Request: req, Input: input, Copy: c.UpdateMissionProfileRequest}
}

// UpdateMissionProfileRequest is the request type for the
// UpdateMissionProfile API operation.
type UpdateMissionProfileRequest struct {
	*aws.Request
	Input *UpdateMissionProfileInput
	Copy  func(*UpdateMissionProfileInput) UpdateMissionProfileRequest
}

// Send marshals and sends the UpdateMissionProfile API request.
func (r UpdateMissionProfileRequest) Send(ctx context.Context) (*UpdateMissionProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateMissionProfileResponse{
		UpdateMissionProfileOutput: r.Request.Data.(*UpdateMissionProfileOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateMissionProfileResponse is the response type for the
// UpdateMissionProfile API operation.
type UpdateMissionProfileResponse struct {
	*UpdateMissionProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateMissionProfile request.
func (r *UpdateMissionProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
