// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the number of AWS Config rules that are compliant and noncompliant, up
// to a maximum of 25 for each.
func (c *Client) GetComplianceSummaryByConfigRule(ctx context.Context, params *GetComplianceSummaryByConfigRuleInput, optFns ...func(*Options)) (*GetComplianceSummaryByConfigRuleOutput, error) {
	if params == nil {
		params = &GetComplianceSummaryByConfigRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetComplianceSummaryByConfigRule", params, optFns, addOperationGetComplianceSummaryByConfigRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetComplianceSummaryByConfigRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetComplianceSummaryByConfigRuleInput struct {
}

//
type GetComplianceSummaryByConfigRuleOutput struct {

	// The number of AWS Config rules that are compliant and the number that are
	// noncompliant, up to a maximum of 25 for each.
	ComplianceSummary *types.ComplianceSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetComplianceSummaryByConfigRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetComplianceSummaryByConfigRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetComplianceSummaryByConfigRule{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetComplianceSummaryByConfigRule(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetComplianceSummaryByConfigRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "GetComplianceSummaryByConfigRule",
	}
}