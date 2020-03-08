// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateFargateProfileInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientRequestToken *string `locationName:"clientRequestToken" type:"string" idempotencyToken:"true"`

	// The name of the Amazon EKS cluster to apply the Fargate profile to.
	//
	// ClusterName is a required field
	ClusterName *string `location:"uri" locationName:"name" type:"string" required:"true"`

	// The name of the Fargate profile.
	//
	// FargateProfileName is a required field
	FargateProfileName *string `locationName:"fargateProfileName" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the pod execution role to use for pods
	// that match the selectors in the Fargate profile. The pod execution role allows
	// Fargate infrastructure to register with your cluster as a node, and it provides
	// read access to Amazon ECR image repositories. For more information, see Pod
	// Execution Role (https://docs.aws.amazon.com/eks/latest/userguide/pod-execution-role.html)
	// in the Amazon EKS User Guide.
	//
	// PodExecutionRoleArn is a required field
	PodExecutionRoleArn *string `locationName:"podExecutionRoleArn" type:"string" required:"true"`

	// The selectors to match for pods to use this Fargate profile. Each selector
	// must have an associated namespace. Optionally, you can also specify labels
	// for a namespace. You may specify up to five selectors in a Fargate profile.
	Selectors []FargateProfileSelector `locationName:"selectors" type:"list"`

	// The IDs of subnets to launch your pods into. At this time, pods running on
	// Fargate are not assigned public IP addresses, so only private subnets (with
	// no direct route to an Internet Gateway) are accepted for this parameter.
	Subnets []string `locationName:"subnets" type:"list"`

	// The metadata to apply to the Fargate profile to assist with categorization
	// and organization. Each tag consists of a key and an optional value, both
	// of which you define. Fargate profile tags do not propagate to any other resources
	// associated with the Fargate profile, such as the pods that are scheduled
	// with it.
	Tags map[string]string `locationName:"tags" min:"1" type:"map"`
}

// String returns the string representation
func (s CreateFargateProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateFargateProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateFargateProfileInput"}

	if s.ClusterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterName"))
	}

	if s.FargateProfileName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FargateProfileName"))
	}

	if s.PodExecutionRoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("PodExecutionRoleArn"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateFargateProfileInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	var ClientRequestToken string
	if s.ClientRequestToken != nil {
		ClientRequestToken = *s.ClientRequestToken
	} else {
		ClientRequestToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientRequestToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientRequestToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FargateProfileName != nil {
		v := *s.FargateProfileName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "fargateProfileName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PodExecutionRoleArn != nil {
		v := *s.PodExecutionRoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "podExecutionRoleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Selectors != nil {
		v := s.Selectors

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "selectors", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Subnets != nil {
		v := s.Subnets

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "subnets", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.ClusterName != nil {
		v := *s.ClusterName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateFargateProfileOutput struct {
	_ struct{} `type:"structure"`

	// The full description of your new Fargate profile.
	FargateProfile *FargateProfile `locationName:"fargateProfile" type:"structure"`
}

// String returns the string representation
func (s CreateFargateProfileOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateFargateProfileOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FargateProfile != nil {
		v := s.FargateProfile

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "fargateProfile", v, metadata)
	}
	return nil
}

const opCreateFargateProfile = "CreateFargateProfile"

// CreateFargateProfileRequest returns a request value for making API operation for
// Amazon Elastic Kubernetes Service.
//
// Creates an AWS Fargate profile for your Amazon EKS cluster. You must have
// at least one Fargate profile in a cluster to be able to run pods on Fargate.
//
// The Fargate profile allows an administrator to declare which pods run on
// Fargate and specify which pods run on which Fargate profile. This declaration
// is done through the profile’s selectors. Each profile can have up to five
// selectors that contain a namespace and labels. A namespace is required for
// every selector. The label field consists of multiple optional key-value pairs.
// Pods that match the selectors are scheduled on Fargate. If a to-be-scheduled
// pod matches any of the selectors in the Fargate profile, then that pod is
// run on Fargate.
//
// When you create a Fargate profile, you must specify a pod execution role
// to use with the pods that are scheduled with the profile. This role is added
// to the cluster's Kubernetes Role Based Access Control (https://kubernetes.io/docs/admin/authorization/rbac/)
// (RBAC) for authorization so that the kubelet that is running on the Fargate
// infrastructure can register with your Amazon EKS cluster so that it can appear
// in your cluster as a node. The pod execution role also provides IAM permissions
// to the Fargate infrastructure to allow read access to Amazon ECR image repositories.
// For more information, see Pod Execution Role (https://docs.aws.amazon.com/eks/latest/userguide/pod-execution-role.html)
// in the Amazon EKS User Guide.
//
// Fargate profiles are immutable. However, you can create a new updated profile
// to replace an existing profile and then delete the original after the updated
// profile has finished creating.
//
// If any Fargate profiles in a cluster are in the DELETING status, you must
// wait for that Fargate profile to finish deleting before you can create any
// other profiles in that cluster.
//
// For more information, see AWS Fargate Profile (https://docs.aws.amazon.com/eks/latest/userguide/fargate-profile.html)
// in the Amazon EKS User Guide.
//
//    // Example sending a request using CreateFargateProfileRequest.
//    req := client.CreateFargateProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eks-2017-11-01/CreateFargateProfile
func (c *Client) CreateFargateProfileRequest(input *CreateFargateProfileInput) CreateFargateProfileRequest {
	op := &aws.Operation{
		Name:       opCreateFargateProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/clusters/{name}/fargate-profiles",
	}

	if input == nil {
		input = &CreateFargateProfileInput{}
	}

	req := c.newRequest(op, input, &CreateFargateProfileOutput{})
	return CreateFargateProfileRequest{Request: req, Input: input, Copy: c.CreateFargateProfileRequest}
}

// CreateFargateProfileRequest is the request type for the
// CreateFargateProfile API operation.
type CreateFargateProfileRequest struct {
	*aws.Request
	Input *CreateFargateProfileInput
	Copy  func(*CreateFargateProfileInput) CreateFargateProfileRequest
}

// Send marshals and sends the CreateFargateProfile API request.
func (r CreateFargateProfileRequest) Send(ctx context.Context) (*CreateFargateProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateFargateProfileResponse{
		CreateFargateProfileOutput: r.Request.Data.(*CreateFargateProfileOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateFargateProfileResponse is the response type for the
// CreateFargateProfile API operation.
type CreateFargateProfileResponse struct {
	*CreateFargateProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateFargateProfile request.
func (r *CreateFargateProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}