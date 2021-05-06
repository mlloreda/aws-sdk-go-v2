// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentity

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use GetPrincipalTagAttributeMap to list all mappings between PrincipalTags and
// user attributes.
func (c *Client) GetPrincipalTagAttributeMap(ctx context.Context, params *GetPrincipalTagAttributeMapInput, optFns ...func(*Options)) (*GetPrincipalTagAttributeMapOutput, error) {
	if params == nil {
		params = &GetPrincipalTagAttributeMapInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPrincipalTagAttributeMap", params, optFns, addOperationGetPrincipalTagAttributeMapMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPrincipalTagAttributeMapOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPrincipalTagAttributeMapInput struct {

	// You can use this operation to get the ID of the Identity Pool you setup
	// attribute mappings for.
	//
	// This member is required.
	IdentityPoolId *string

	// You can use this operation to get the provider name.
	//
	// This member is required.
	IdentityProviderName *string
}

type GetPrincipalTagAttributeMapOutput struct {

	// You can use this operation to get the ID of the Identity Pool you setup
	// attribute mappings for.
	IdentityPoolId *string

	// You can use this operation to get the provider name.
	IdentityProviderName *string

	// You can use this operation to add principal tags. The PrincipalTagsoperation
	// enables you to reference user attributes in your IAM permissions policy.
	PrincipalTags map[string]string

	// You can use this operation to list
	UseDefaults *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetPrincipalTagAttributeMapMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetPrincipalTagAttributeMap{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPrincipalTagAttributeMap{}, middleware.After)
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
	if err = addOpGetPrincipalTagAttributeMapValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPrincipalTagAttributeMap(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetPrincipalTagAttributeMap(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-identity",
		OperationName: "GetPrincipalTagAttributeMap",
	}
}