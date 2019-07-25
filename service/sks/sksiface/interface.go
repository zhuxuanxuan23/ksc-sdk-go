// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package sksiface provides an interface to enable mocking the sks service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package sksiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/ksc/ksc-sdk-go/service/sks"
)

// SksAPI provides an interface to enable mocking the
// sks.Sks service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // sks.
//    func myFunc(svc sksiface.SksAPI) bool {
//        // Make svc.CreateKey request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := sks.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockSksClient struct {
//        sksiface.SksAPI
//    }
//    func (m *mockSksClient) CreateKey(input *map[string]interface{}) (*map[string]interface{}, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockSksClient{}
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
type SksAPI interface {
	CreateKey(*map[string]interface{}) (*map[string]interface{}, error)
	CreateKeyWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateKeyRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteKey(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteKeyWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteKeyRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeKeys(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeKeysWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeKeysRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ImportKey(*map[string]interface{}) (*map[string]interface{}, error)
	ImportKeyWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ImportKeyRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyKey(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyKeyWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyKeyRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})
}

var _ SksAPI = (*sks.Sks)(nil)