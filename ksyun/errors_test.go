package ksyun

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

func TestRetryError(t *testing.T) {
	err := awserr.New("INTERNAL_FAILURE", "test", nil)

	requestFailure := awserr.NewRequestFailure(err, 500, "2efb29f4-ad46-4b22-996b-7499d4b0871e")

	err2 := awserr.New("InvalidActionOrVesion", "test", nil)
	requestFailure2 := awserr.NewRequestFailure(err2, 400, "7ba0ad63-a4a2-4db5-8e7f-00cf862f9464")
	requestFailureList := []awserr.RequestFailure{requestFailure, requestFailure2}
	for _, req := range requestFailureList {
		resourceRetry := retryError(req)
		t.Log("retryable: ", resourceRetry.Retryable)
	}

}
