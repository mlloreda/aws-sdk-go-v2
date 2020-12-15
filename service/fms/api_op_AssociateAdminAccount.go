// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the AWS Firewall Manager administrator account. AWS Firewall Manager must
// be associated with the master account of your AWS organization or associated
// with a member account that has the appropriate permissions. If the account ID
// that you submit is not an AWS Organizations master account, AWS Firewall Manager
// will set the appropriate permissions for the given member account. The account
// that you associate with AWS Firewall Manager is called the AWS Firewall Manager
// administrator account.
func (c *Client) AssociateAdminAccount(ctx context.Context, params *AssociateAdminAccountInput, optFns ...func(*Options)) (*AssociateAdminAccountOutput, error) {
	if params == nil {
		params = &AssociateAdminAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateAdminAccount", params, optFns, addOperationAssociateAdminAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateAdminAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateAdminAccountInput struct {

	// The AWS account ID to associate with AWS Firewall Manager as the AWS Firewall
	// Manager administrator account. This can be an AWS Organizations master account
	// or a member account. For more information about AWS Organizations and master
	// accounts, see Managing the AWS Accounts in Your Organization
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts.html).
	//
	// This member is required.
	AdminAccount *string
}

type AssociateAdminAccountOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAssociateAdminAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateAdminAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateAdminAccount{}, middleware.After)
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
	if err = addOpAssociateAdminAccountValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateAdminAccount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateAdminAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "AssociateAdminAccount",
	}
}