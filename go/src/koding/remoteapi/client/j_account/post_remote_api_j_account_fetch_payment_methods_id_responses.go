package j_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJAccountFetchPaymentMethodsIDReader is a Reader for the PostRemoteAPIJAccountFetchPaymentMethodsID structure.
type PostRemoteAPIJAccountFetchPaymentMethodsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJAccountFetchPaymentMethodsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJAccountFetchPaymentMethodsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJAccountFetchPaymentMethodsIDOK creates a PostRemoteAPIJAccountFetchPaymentMethodsIDOK with default headers values
func NewPostRemoteAPIJAccountFetchPaymentMethodsIDOK() *PostRemoteAPIJAccountFetchPaymentMethodsIDOK {
	return &PostRemoteAPIJAccountFetchPaymentMethodsIDOK{}
}

/*PostRemoteAPIJAccountFetchPaymentMethodsIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJAccountFetchPaymentMethodsIDOK struct {
	Payload PostRemoteAPIJAccountFetchPaymentMethodsIDOKBody
}

func (o *PostRemoteAPIJAccountFetchPaymentMethodsIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JAccount.fetchPaymentMethods/{id}][%d] postRemoteApiJAccountFetchPaymentMethodsIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJAccountFetchPaymentMethodsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJAccountFetchPaymentMethodsIDOKBody post remote API j account fetch payment methods ID o k body
swagger:model PostRemoteAPIJAccountFetchPaymentMethodsIDOKBody
*/
type PostRemoteAPIJAccountFetchPaymentMethodsIDOKBody struct {
	models.JAccount

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRemoteAPIJAccountFetchPaymentMethodsIDOKBody) UnmarshalJSON(raw []byte) error {

	var postRemoteAPIJAccountFetchPaymentMethodsIDOKBodyAO0 models.JAccount
	if err := swag.ReadJSON(raw, &postRemoteAPIJAccountFetchPaymentMethodsIDOKBodyAO0); err != nil {
		return err
	}
	o.JAccount = postRemoteAPIJAccountFetchPaymentMethodsIDOKBodyAO0

	var postRemoteAPIJAccountFetchPaymentMethodsIDOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &postRemoteAPIJAccountFetchPaymentMethodsIDOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = postRemoteAPIJAccountFetchPaymentMethodsIDOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRemoteAPIJAccountFetchPaymentMethodsIDOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	postRemoteAPIJAccountFetchPaymentMethodsIDOKBodyAO0, err := swag.WriteJSON(o.JAccount)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJAccountFetchPaymentMethodsIDOKBodyAO0)

	postRemoteAPIJAccountFetchPaymentMethodsIDOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJAccountFetchPaymentMethodsIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post remote API j account fetch payment methods ID o k body
func (o *PostRemoteAPIJAccountFetchPaymentMethodsIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JAccount.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.DefaultResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
