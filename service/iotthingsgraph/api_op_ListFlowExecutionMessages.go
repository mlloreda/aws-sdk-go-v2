// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of objects that contain information about events in a flow
// execution.
func (c *Client) ListFlowExecutionMessages(ctx context.Context, params *ListFlowExecutionMessagesInput, optFns ...func(*Options)) (*ListFlowExecutionMessagesOutput, error) {
	if params == nil {
		params = &ListFlowExecutionMessagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFlowExecutionMessages", params, optFns, addOperationListFlowExecutionMessagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFlowExecutionMessagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFlowExecutionMessagesInput struct {

	// The ID of the flow execution.
	//
	// This member is required.
	FlowExecutionId *string

	// The maximum number of results to return in the response.
	MaxResults *int32

	// The string that specifies the next page of results. Use this when you're
	// paginating results.
	NextToken *string
}

type ListFlowExecutionMessagesOutput struct {

	// A list of objects that contain information about events in the specified flow
	// execution.
	Messages []types.FlowExecutionMessage

	// The string to specify as nextToken when you request the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListFlowExecutionMessagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListFlowExecutionMessages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFlowExecutionMessages{}, middleware.After)
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
	if err = addOpListFlowExecutionMessagesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFlowExecutionMessages(options.Region), middleware.Before); err != nil {
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

// ListFlowExecutionMessagesAPIClient is a client that implements the
// ListFlowExecutionMessages operation.
type ListFlowExecutionMessagesAPIClient interface {
	ListFlowExecutionMessages(context.Context, *ListFlowExecutionMessagesInput, ...func(*Options)) (*ListFlowExecutionMessagesOutput, error)
}

var _ ListFlowExecutionMessagesAPIClient = (*Client)(nil)

// ListFlowExecutionMessagesPaginatorOptions is the paginator options for
// ListFlowExecutionMessages
type ListFlowExecutionMessagesPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFlowExecutionMessagesPaginator is a paginator for ListFlowExecutionMessages
type ListFlowExecutionMessagesPaginator struct {
	options   ListFlowExecutionMessagesPaginatorOptions
	client    ListFlowExecutionMessagesAPIClient
	params    *ListFlowExecutionMessagesInput
	nextToken *string
	firstPage bool
}

// NewListFlowExecutionMessagesPaginator returns a new
// ListFlowExecutionMessagesPaginator
func NewListFlowExecutionMessagesPaginator(client ListFlowExecutionMessagesAPIClient, params *ListFlowExecutionMessagesInput, optFns ...func(*ListFlowExecutionMessagesPaginatorOptions)) *ListFlowExecutionMessagesPaginator {
	options := ListFlowExecutionMessagesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListFlowExecutionMessagesInput{}
	}

	return &ListFlowExecutionMessagesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFlowExecutionMessagesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListFlowExecutionMessages page.
func (p *ListFlowExecutionMessagesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFlowExecutionMessagesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListFlowExecutionMessages(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListFlowExecutionMessages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "ListFlowExecutionMessages",
	}
}