package autoscaling

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface"
	"github.com/pkg/errors"
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

// ListGroups return the list of Auto Scaling Groups
func (c *Client) ListGroups() ([]string, error) {
	resp, err := c.api.DescribeAutoScalingGroups(&autoscaling.DescribeAutoScalingGroupsInput{})
	if err != nil {
		return []string{}, errors.Wrap(err, "failed to fetch Auto Scaling Groups")
	}

	groups := []string{}

	for _, g := range resp.AutoScalingGroups {
		groups = append(groups, aws.StringValue(g.AutoScalingGroupName))
	}

	return groups, nil
}
