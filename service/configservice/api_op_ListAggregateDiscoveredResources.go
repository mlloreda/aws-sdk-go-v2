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

// Accepts a resource type and returns a list of resource identifiers that are
// aggregated for a specific resource type across accounts and regions. A resource
// identifier includes the resource type, ID, (if available) the custom resource
// name, source account, and source region. You can narrow the results to include
// only resources that have specific resource IDs, or a resource name, or source
// account ID, or source region. For example, if the input consists of accountID
// 12345678910 and the region is us-east-1 for resource type AWS::EC2::Instance
// then the API returns all the EC2 instance identifiers of accountID 12345678910
// and region us-east-1.
func (c *Client) ListAggregateDiscoveredResources(ctx context.Context, params *ListAggregateDiscoveredResourcesInput, optFns ...func(*Options)) (*ListAggregateDiscoveredResourcesOutput, error) {
	if params == nil {
		params = &ListAggregateDiscoveredResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAggregateDiscoveredResources", params, optFns, addOperationListAggregateDiscoveredResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAggregateDiscoveredResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAggregateDiscoveredResourcesInput struct {

	// The name of the configuration aggregator.
	//
	// This member is required.
	ConfigurationAggregatorName *string

	// The type of resources that you want AWS Config to list in the response.
	//
	// This member is required.
	ResourceType types.ResourceType

	// Filters the results based on the ResourceFilters object.
	Filters *types.ResourceFilters

	// The maximum number of resource identifiers returned on each page. The default is
	// 100. You cannot specify a number greater than 100. If you specify 0, AWS Config
	// uses the default.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
}

type ListAggregateDiscoveredResourcesOutput struct {

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// Returns a list of ResourceIdentifiers objects.
	ResourceIdentifiers []types.AggregateResourceIdentifier

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListAggregateDiscoveredResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAggregateDiscoveredResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAggregateDiscoveredResources{}, middleware.After)
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
	if err = addOpListAggregateDiscoveredResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAggregateDiscoveredResources(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListAggregateDiscoveredResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "ListAggregateDiscoveredResources",
	}
}