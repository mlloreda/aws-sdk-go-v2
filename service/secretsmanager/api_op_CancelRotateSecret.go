// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables automatic scheduled rotation and cancels the rotation of a secret if
// currently in progress. To re-enable scheduled rotation, call RotateSecret with
// AutomaticallyRotateAfterDays set to a value greater than 0. This immediately
// rotates your secret and then enables the automatic schedule. If you cancel a
// rotation while in progress, it can leave the VersionStage labels in an
// unexpected state. Depending on the step of the rotation in progress, you might
// need to remove the staging label AWSPENDING from the partially created version,
// specified by the VersionId response value. You should also evaluate the
// partially rotated new version to see if it should be deleted, which you can do
// by removing all staging labels from the new version VersionStage field. To
// successfully start a rotation, the staging label AWSPENDING must be in one of
// the following states:
//
// * Not attached to any version at all
//
// * Attached to the
// same version as the staging label AWSCURRENT
//
// If the staging label AWSPENDING
// attached to a different version than the version with AWSCURRENT then the
// attempt to rotate fails. Minimum permissions To run this command, you must have
// the following permissions:
//
// * secretsmanager:CancelRotateSecret
//
// Related
// operations
//
// * To configure rotation for a secret or to manually trigger a
// rotation, use RotateSecret.
//
// * To get the rotation configuration details for a
// secret, use DescribeSecret.
//
// * To list all of the currently available secrets,
// use ListSecrets.
//
// * To list all of the versions currently associated with a
// secret, use ListSecretVersionIds.
func (c *Client) CancelRotateSecret(ctx context.Context, params *CancelRotateSecretInput, optFns ...func(*Options)) (*CancelRotateSecretOutput, error) {
	if params == nil {
		params = &CancelRotateSecretInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelRotateSecret", params, optFns, addOperationCancelRotateSecretMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelRotateSecretOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelRotateSecretInput struct {

	// Specifies the secret to cancel a rotation request. You can specify either the
	// Amazon Resource Name (ARN) or the friendly name of the secret. If you specify an
	// ARN, we generally recommend that you specify a complete ARN. You can specify a
	// partial ARN too—for example, if you don’t include the final hyphen and six
	// random characters that Secrets Manager adds at the end of the ARN when you
	// created the secret. A partial ARN match can work as long as it uniquely matches
	// only one secret. However, if your secret has a name that ends in a hyphen
	// followed by six characters (before Secrets Manager adds the hyphen and six
	// characters to the ARN) and you try to use that as a partial ARN, then those
	// characters cause Secrets Manager to assume that you’re specifying a complete
	// ARN. This confusion can cause unexpected results. To avoid this situation, we
	// recommend that you don’t create secret names ending with a hyphen followed by
	// six characters. If you specify an incomplete ARN without the random suffix, and
	// instead provide the 'friendly name', you must not include the random suffix. If
	// you do include the random suffix added by Secrets Manager, you receive either a
	// ResourceNotFoundException or an AccessDeniedException error, depending on your
	// permissions.
	//
	// This member is required.
	SecretId *string
}

type CancelRotateSecretOutput struct {

	// The ARN of the secret for which rotation was canceled.
	ARN *string

	// The friendly name of the secret for which rotation was canceled.
	Name *string

	// The unique identifier of the version of the secret created during the rotation.
	// This version might not be complete, and should be evaluated for possible
	// deletion. At the very least, you should remove the VersionStage value AWSPENDING
	// to enable this version to be deleted. Failing to clean up a cancelled rotation
	// can block you from successfully starting future rotations.
	VersionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCancelRotateSecretMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCancelRotateSecret{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCancelRotateSecret{}, middleware.After)
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
	if err = addOpCancelRotateSecretValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelRotateSecret(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelRotateSecret(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "secretsmanager",
		OperationName: "CancelRotateSecret",
	}
}