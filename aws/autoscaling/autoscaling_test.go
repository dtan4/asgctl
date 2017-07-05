package autoscaling

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/dtan4/asgctl/aws/mock"
	"github.com/golang/mock/gomock"
)

func TestListGroups(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	api := mock.NewMockAutoScalingAPI(ctrl)

	api.EXPECT().DescribeAutoScalingGroups(&autoscaling.DescribeAutoScalingGroupsInput{}).Return(&autoscaling.DescribeAutoScalingGroupsOutput{
		AutoScalingGroups: []*autoscaling.Group{
			&autoscaling.Group{
				AutoScalingGroupARN:  aws.String("arn:aws:autoscaling:ap-northeast-1:012345678901:autoScalingGroup:cd25aef2-a44c-4e27-bb9b-123456abcdef:autoScalingGroupName/asg-01"),
				AutoScalingGroupName: aws.String("asg-01"),
			},
			&autoscaling.Group{
				AutoScalingGroupARN:  aws.String("arn:aws:autoscaling:ap-northeast-1:096233016669:autoScalingGroup:d4204d07-3694-494f-b665-abcdef123456:autoScalingGroupName/asg-02"),
				AutoScalingGroupName: aws.String("asg-02"),
			},
		},
	}, nil)
	client := &Client{
		api: api,
	}

	got, err := client.ListGroups()
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	expected := []string{
		"asg-01",
		"asg-02",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("groups mismatch, expected: %#v, got: %#v", expected, got)
	}
}
