// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package stsiface provides an interface to enable mocking the AWS Security Token Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package stsiface

import (
	"github.com/ScottCho/ct-sdk-go/aws"
	"github.com/ScottCho/ct-sdk-go/aws/request"
	"github.com/ScottCho/ct-sdk-go/service/sts"
)

// STSAPI provides an interface to enable mocking the
// sts.STS service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Security Token Service.
//    func myFunc(svc stsiface.STSAPI) bool {
//        // Make svc.AssumeRole request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := sts.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockSTSClient struct {
//        stsiface.STSAPI
//    }
//    func (m *mockSTSClient) AssumeRole(input *sts.AssumeRoleInput) (*sts.AssumeRoleOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockSTSClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type STSAPI interface {
	AssumeRole(*sts.AssumeRoleInput) (*sts.AssumeRoleOutput, error)
	AssumeRoleWithContext(aws.Context, *sts.AssumeRoleInput, ...request.Option) (*sts.AssumeRoleOutput, error)
	AssumeRoleRequest(*sts.AssumeRoleInput) (*request.Request, *sts.AssumeRoleOutput)

	AssumeRoleWithSAML(*sts.AssumeRoleWithSAMLInput) (*sts.AssumeRoleWithSAMLOutput, error)
	AssumeRoleWithSAMLWithContext(aws.Context, *sts.AssumeRoleWithSAMLInput, ...request.Option) (*sts.AssumeRoleWithSAMLOutput, error)
	AssumeRoleWithSAMLRequest(*sts.AssumeRoleWithSAMLInput) (*request.Request, *sts.AssumeRoleWithSAMLOutput)

	AssumeRoleWithWebIdentity(*sts.AssumeRoleWithWebIdentityInput) (*sts.AssumeRoleWithWebIdentityOutput, error)
	AssumeRoleWithWebIdentityWithContext(aws.Context, *sts.AssumeRoleWithWebIdentityInput, ...request.Option) (*sts.AssumeRoleWithWebIdentityOutput, error)
	AssumeRoleWithWebIdentityRequest(*sts.AssumeRoleWithWebIdentityInput) (*request.Request, *sts.AssumeRoleWithWebIdentityOutput)

	DecodeAuthorizationMessage(*sts.DecodeAuthorizationMessageInput) (*sts.DecodeAuthorizationMessageOutput, error)
	DecodeAuthorizationMessageWithContext(aws.Context, *sts.DecodeAuthorizationMessageInput, ...request.Option) (*sts.DecodeAuthorizationMessageOutput, error)
	DecodeAuthorizationMessageRequest(*sts.DecodeAuthorizationMessageInput) (*request.Request, *sts.DecodeAuthorizationMessageOutput)

	GetAccessKeyInfo(*sts.GetAccessKeyInfoInput) (*sts.GetAccessKeyInfoOutput, error)
	GetAccessKeyInfoWithContext(aws.Context, *sts.GetAccessKeyInfoInput, ...request.Option) (*sts.GetAccessKeyInfoOutput, error)
	GetAccessKeyInfoRequest(*sts.GetAccessKeyInfoInput) (*request.Request, *sts.GetAccessKeyInfoOutput)

	GetCallerIdentity(*sts.GetCallerIdentityInput) (*sts.GetCallerIdentityOutput, error)
	GetCallerIdentityWithContext(aws.Context, *sts.GetCallerIdentityInput, ...request.Option) (*sts.GetCallerIdentityOutput, error)
	GetCallerIdentityRequest(*sts.GetCallerIdentityInput) (*request.Request, *sts.GetCallerIdentityOutput)

	GetFederationToken(*sts.GetFederationTokenInput) (*sts.GetFederationTokenOutput, error)
	GetFederationTokenWithContext(aws.Context, *sts.GetFederationTokenInput, ...request.Option) (*sts.GetFederationTokenOutput, error)
	GetFederationTokenRequest(*sts.GetFederationTokenInput) (*request.Request, *sts.GetFederationTokenOutput)

	GetSessionToken(*sts.GetSessionTokenInput) (*sts.GetSessionTokenOutput, error)
	GetSessionTokenWithContext(aws.Context, *sts.GetSessionTokenInput, ...request.Option) (*sts.GetSessionTokenOutput, error)
	GetSessionTokenRequest(*sts.GetSessionTokenInput) (*request.Request, *sts.GetSessionTokenOutput)
}

var _ STSAPI = (*sts.STS)(nil)
