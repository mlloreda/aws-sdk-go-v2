// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes a source identifier from an existing Amazon DocumentDB event
// notification subscription.
func (c *Client) RemoveSourceIdentifierFromSubscription(ctx context.Context, params *RemoveSourceIdentifierFromSubscriptionInput, optFns ...func(*Options)) (*RemoveSourceIdentifierFromSubscriptionOutput, error) {
	if params == nil {
		params = &RemoveSourceIdentifierFromSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveSourceIdentifierFromSubscription", params, optFns, c.addOperationRemoveSourceIdentifierFromSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveSourceIdentifierFromSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to RemoveSourceIdentifierFromSubscription.
type RemoveSourceIdentifierFromSubscriptionInput struct {

	// The source identifier to be removed from the subscription, such as the instance
	// identifier for an instance, or the name of a security group.
	//
	// This member is required.
	SourceIdentifier *string

	// The name of the Amazon DocumentDB event notification subscription that you want
	// to remove a source identifier from.
	//
	// This member is required.
	SubscriptionName *string

	noSmithyDocumentSerde
}

type RemoveSourceIdentifierFromSubscriptionOutput struct {

	// Detailed information about an event to which you have subscribed.
	EventSubscription *types.EventSubscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRemoveSourceIdentifierFromSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRemoveSourceIdentifierFromSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRemoveSourceIdentifierFromSubscription{}, middleware.After)
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpRemoveSourceIdentifierFromSubscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveSourceIdentifierFromSubscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRemoveSourceIdentifierFromSubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "RemoveSourceIdentifierFromSubscription",
	}
}