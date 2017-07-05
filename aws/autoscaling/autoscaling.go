package autoscaling

import (
	"github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface"
)

// Client represents the wrapper of Auto Scaling API client
type Client struct {
	api autoscalingiface.AutoScalingAPI
}

// NewClient creates new Client object
func NewClient(api autoscalingiface.AutoScalingAPI) *Client {
	return &Client{
		api: api,
	}
}
