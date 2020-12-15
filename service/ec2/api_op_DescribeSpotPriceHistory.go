// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes the Spot price history. For more information, see Spot Instance
// pricing history
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-spot-instances-history.html)
// in the Amazon EC2 User Guide for Linux Instances. When you specify a start and
// end time, this operation returns the prices of the instance types within the
// time range that you specified and the time when the price changed. The price is
// valid within the time period that you specified; the response merely indicates
// the last time that the price changed.
func (c *Client) DescribeSpotPriceHistory(ctx context.Context, params *DescribeSpotPriceHistoryInput, optFns ...func(*Options)) (*DescribeSpotPriceHistoryOutput, error) {
	if params == nil {
		params = &DescribeSpotPriceHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSpotPriceHistory", params, optFns, addOperationDescribeSpotPriceHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSpotPriceHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeSpotPriceHistory.
type DescribeSpotPriceHistoryInput struct {

	// Filters the results by the specified Availability Zone.
	AvailabilityZone *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// The date and time, up to the current date, from which to stop retrieving the
	// price history data, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ).
	EndTime *time.Time

	// One or more filters.
	//
	// * availability-zone - The Availability Zone for which
	// prices should be returned.
	//
	// * instance-type - The type of instance (for example,
	// m3.medium).
	//
	// * product-description - The product description for the Spot price
	// (Linux/UNIX | Red Hat Enterprise Linux | SUSE Linux | Windows | Linux/UNIX
	// (Amazon VPC) | Red Hat Enterprise Linux (Amazon VPC) | SUSE Linux (Amazon VPC) |
	// Windows (Amazon VPC)).
	//
	// * spot-price - The Spot price. The value must match
	// exactly (or use wildcards; greater than or less than comparison is not
	// supported).
	//
	// * timestamp - The time stamp of the Spot price history, in UTC
	// format (for example, YYYY-MM-DDTHH:MM:SSZ). You can use wildcards (* and ?).
	// Greater than or less than comparison is not supported.
	Filters []types.Filter

	// Filters the results by the specified instance types.
	InstanceTypes []types.InstanceType

	// The maximum number of results to return in a single call. Specify a value
	// between 1 and 1000. The default value is 1000. To retrieve the remaining
	// results, make another call with the returned NextToken value.
	MaxResults int32

	// The token for the next set of results.
	NextToken *string

	// Filters the results by the specified basic product descriptions.
	ProductDescriptions []string

	// The date and time, up to the past 90 days, from which to start retrieving the
	// price history data, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ).
	StartTime *time.Time
}

// Contains the output of DescribeSpotPriceHistory.
type DescribeSpotPriceHistoryOutput struct {

	// The token required to retrieve the next set of results. This value is null or an
	// empty string when there are no more results to return.
	NextToken *string

	// The historical Spot prices.
	SpotPriceHistory []types.SpotPrice

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeSpotPriceHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeSpotPriceHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeSpotPriceHistory{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSpotPriceHistory(options.Region), middleware.Before); err != nil {
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

// DescribeSpotPriceHistoryAPIClient is a client that implements the
// DescribeSpotPriceHistory operation.
type DescribeSpotPriceHistoryAPIClient interface {
	DescribeSpotPriceHistory(context.Context, *DescribeSpotPriceHistoryInput, ...func(*Options)) (*DescribeSpotPriceHistoryOutput, error)
}

var _ DescribeSpotPriceHistoryAPIClient = (*Client)(nil)

// DescribeSpotPriceHistoryPaginatorOptions is the paginator options for
// DescribeSpotPriceHistory
type DescribeSpotPriceHistoryPaginatorOptions struct {
	// The maximum number of results to return in a single call. Specify a value
	// between 1 and 1000. The default value is 1000. To retrieve the remaining
	// results, make another call with the returned NextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeSpotPriceHistoryPaginator is a paginator for DescribeSpotPriceHistory
type DescribeSpotPriceHistoryPaginator struct {
	options   DescribeSpotPriceHistoryPaginatorOptions
	client    DescribeSpotPriceHistoryAPIClient
	params    *DescribeSpotPriceHistoryInput
	nextToken *string
	firstPage bool
}

// NewDescribeSpotPriceHistoryPaginator returns a new
// DescribeSpotPriceHistoryPaginator
func NewDescribeSpotPriceHistoryPaginator(client DescribeSpotPriceHistoryAPIClient, params *DescribeSpotPriceHistoryInput, optFns ...func(*DescribeSpotPriceHistoryPaginatorOptions)) *DescribeSpotPriceHistoryPaginator {
	options := DescribeSpotPriceHistoryPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeSpotPriceHistoryInput{}
	}

	return &DescribeSpotPriceHistoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeSpotPriceHistoryPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeSpotPriceHistory page.
func (p *DescribeSpotPriceHistoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeSpotPriceHistoryOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeSpotPriceHistory(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeSpotPriceHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeSpotPriceHistory",
	}
}