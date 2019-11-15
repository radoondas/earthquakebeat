// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ApplySecurityGroupsToClientVpnTargetNetworkRequest
type ApplySecurityGroupsToClientVpnTargetNetworkInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Client VPN endpoint.
	//
	// ClientVpnEndpointId is a required field
	ClientVpnEndpointId *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The IDs of the security groups to apply to the associated target network.
	// Up to 5 security groups can be applied to an associated target network.
	//
	// SecurityGroupIds is a required field
	SecurityGroupIds []string `locationName:"SecurityGroupId" locationNameList:"item" type:"list" required:"true"`

	// The ID of the VPC in which the associated target network is located.
	//
	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ApplySecurityGroupsToClientVpnTargetNetworkInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ApplySecurityGroupsToClientVpnTargetNetworkInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ApplySecurityGroupsToClientVpnTargetNetworkInput"}

	if s.ClientVpnEndpointId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientVpnEndpointId"))
	}

	if s.SecurityGroupIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecurityGroupIds"))
	}

	if s.VpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ApplySecurityGroupsToClientVpnTargetNetworkResult
type ApplySecurityGroupsToClientVpnTargetNetworkOutput struct {
	_ struct{} `type:"structure"`

	// The IDs of the applied security groups.
	SecurityGroupIds []string `locationName:"securityGroupIds" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s ApplySecurityGroupsToClientVpnTargetNetworkOutput) String() string {
	return awsutil.Prettify(s)
}

const opApplySecurityGroupsToClientVpnTargetNetwork = "ApplySecurityGroupsToClientVpnTargetNetwork"

// ApplySecurityGroupsToClientVpnTargetNetworkRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Applies a security group to the association between the target network and
// the Client VPN endpoint. This action replaces the existing security groups
// with the specified security groups.
//
//    // Example sending a request using ApplySecurityGroupsToClientVpnTargetNetworkRequest.
//    req := client.ApplySecurityGroupsToClientVpnTargetNetworkRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ApplySecurityGroupsToClientVpnTargetNetwork
func (c *Client) ApplySecurityGroupsToClientVpnTargetNetworkRequest(input *ApplySecurityGroupsToClientVpnTargetNetworkInput) ApplySecurityGroupsToClientVpnTargetNetworkRequest {
	op := &aws.Operation{
		Name:       opApplySecurityGroupsToClientVpnTargetNetwork,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ApplySecurityGroupsToClientVpnTargetNetworkInput{}
	}

	req := c.newRequest(op, input, &ApplySecurityGroupsToClientVpnTargetNetworkOutput{})
	return ApplySecurityGroupsToClientVpnTargetNetworkRequest{Request: req, Input: input, Copy: c.ApplySecurityGroupsToClientVpnTargetNetworkRequest}
}

// ApplySecurityGroupsToClientVpnTargetNetworkRequest is the request type for the
// ApplySecurityGroupsToClientVpnTargetNetwork API operation.
type ApplySecurityGroupsToClientVpnTargetNetworkRequest struct {
	*aws.Request
	Input *ApplySecurityGroupsToClientVpnTargetNetworkInput
	Copy  func(*ApplySecurityGroupsToClientVpnTargetNetworkInput) ApplySecurityGroupsToClientVpnTargetNetworkRequest
}

// Send marshals and sends the ApplySecurityGroupsToClientVpnTargetNetwork API request.
func (r ApplySecurityGroupsToClientVpnTargetNetworkRequest) Send(ctx context.Context) (*ApplySecurityGroupsToClientVpnTargetNetworkResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ApplySecurityGroupsToClientVpnTargetNetworkResponse{
		ApplySecurityGroupsToClientVpnTargetNetworkOutput: r.Request.Data.(*ApplySecurityGroupsToClientVpnTargetNetworkOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ApplySecurityGroupsToClientVpnTargetNetworkResponse is the response type for the
// ApplySecurityGroupsToClientVpnTargetNetwork API operation.
type ApplySecurityGroupsToClientVpnTargetNetworkResponse struct {
	*ApplySecurityGroupsToClientVpnTargetNetworkOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ApplySecurityGroupsToClientVpnTargetNetwork request.
func (r *ApplySecurityGroupsToClientVpnTargetNetworkResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
