// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

// Contains the parameters for DeleteVpnConnectionRoute.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteVpnConnectionRouteRequest
type DeleteVpnConnectionRouteInput struct {
	_ struct{} `type:"structure"`

	// The CIDR block associated with the local subnet of the customer network.
	//
	// DestinationCidrBlock is a required field
	DestinationCidrBlock *string `type:"string" required:"true"`

	// The ID of the VPN connection.
	//
	// VpnConnectionId is a required field
	VpnConnectionId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteVpnConnectionRouteInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVpnConnectionRouteInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteVpnConnectionRouteInput"}

	if s.DestinationCidrBlock == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationCidrBlock"))
	}

	if s.VpnConnectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpnConnectionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteVpnConnectionRouteOutput
type DeleteVpnConnectionRouteOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteVpnConnectionRouteOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteVpnConnectionRoute = "DeleteVpnConnectionRoute"

// DeleteVpnConnectionRouteRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified static route associated with a VPN connection between
// an existing virtual private gateway and a VPN customer gateway. The static
// route allows traffic to be routed from the virtual private gateway to the
// VPN customer gateway.
//
//    // Example sending a request using DeleteVpnConnectionRouteRequest.
//    req := client.DeleteVpnConnectionRouteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteVpnConnectionRoute
func (c *Client) DeleteVpnConnectionRouteRequest(input *DeleteVpnConnectionRouteInput) DeleteVpnConnectionRouteRequest {
	op := &aws.Operation{
		Name:       opDeleteVpnConnectionRoute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteVpnConnectionRouteInput{}
	}

	req := c.newRequest(op, input, &DeleteVpnConnectionRouteOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteVpnConnectionRouteRequest{Request: req, Input: input, Copy: c.DeleteVpnConnectionRouteRequest}
}

// DeleteVpnConnectionRouteRequest is the request type for the
// DeleteVpnConnectionRoute API operation.
type DeleteVpnConnectionRouteRequest struct {
	*aws.Request
	Input *DeleteVpnConnectionRouteInput
	Copy  func(*DeleteVpnConnectionRouteInput) DeleteVpnConnectionRouteRequest
}

// Send marshals and sends the DeleteVpnConnectionRoute API request.
func (r DeleteVpnConnectionRouteRequest) Send(ctx context.Context) (*DeleteVpnConnectionRouteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVpnConnectionRouteResponse{
		DeleteVpnConnectionRouteOutput: r.Request.Data.(*DeleteVpnConnectionRouteOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVpnConnectionRouteResponse is the response type for the
// DeleteVpnConnectionRoute API operation.
type DeleteVpnConnectionRouteResponse struct {
	*DeleteVpnConnectionRouteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVpnConnectionRoute request.
func (r *DeleteVpnConnectionRouteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
