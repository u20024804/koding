package social_message

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPISocialMessageListRepliesReader is a Reader for the PostRemoteAPISocialMessageListReplies structure.
type PostRemoteAPISocialMessageListRepliesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPISocialMessageListRepliesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPISocialMessageListRepliesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPISocialMessageListRepliesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPISocialMessageListRepliesOK creates a PostRemoteAPISocialMessageListRepliesOK with default headers values
func NewPostRemoteAPISocialMessageListRepliesOK() *PostRemoteAPISocialMessageListRepliesOK {
	return &PostRemoteAPISocialMessageListRepliesOK{}
}

/*PostRemoteAPISocialMessageListRepliesOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPISocialMessageListRepliesOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPISocialMessageListRepliesOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialMessage.listReplies][%d] postRemoteApiSocialMessageListRepliesOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPISocialMessageListRepliesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPISocialMessageListRepliesUnauthorized creates a PostRemoteAPISocialMessageListRepliesUnauthorized with default headers values
func NewPostRemoteAPISocialMessageListRepliesUnauthorized() *PostRemoteAPISocialMessageListRepliesUnauthorized {
	return &PostRemoteAPISocialMessageListRepliesUnauthorized{}
}

/*PostRemoteAPISocialMessageListRepliesUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPISocialMessageListRepliesUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPISocialMessageListRepliesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialMessage.listReplies][%d] postRemoteApiSocialMessageListRepliesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPISocialMessageListRepliesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
