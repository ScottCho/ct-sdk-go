package sts

import "ct-sdk-go/aws/request"

func init() {
	initRequest = customizeRequest
}

func customizeRequest(r *request.Request) {
	r.RetryErrorCodes = append(r.RetryErrorCodes, ErrCodeIDPCommunicationErrorException)
}
