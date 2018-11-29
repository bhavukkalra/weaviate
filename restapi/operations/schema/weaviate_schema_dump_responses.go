/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */
// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// WeaviateSchemaDumpOKCode is the HTTP code returned for type WeaviateSchemaDumpOK
const WeaviateSchemaDumpOKCode int = 200

/*WeaviateSchemaDumpOK Successfully dumped the database schema.

swagger:response weaviateSchemaDumpOK
*/
type WeaviateSchemaDumpOK struct {

	/*
	  In: Body
	*/
	Payload *WeaviateSchemaDumpOKBody `json:"body,omitempty"`
}

// NewWeaviateSchemaDumpOK creates WeaviateSchemaDumpOK with default headers values
func NewWeaviateSchemaDumpOK() *WeaviateSchemaDumpOK {

	return &WeaviateSchemaDumpOK{}
}

// WithPayload adds the payload to the weaviate schema dump o k response
func (o *WeaviateSchemaDumpOK) WithPayload(payload *WeaviateSchemaDumpOKBody) *WeaviateSchemaDumpOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate schema dump o k response
func (o *WeaviateSchemaDumpOK) SetPayload(payload *WeaviateSchemaDumpOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateSchemaDumpOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateSchemaDumpUnauthorizedCode is the HTTP code returned for type WeaviateSchemaDumpUnauthorized
const WeaviateSchemaDumpUnauthorizedCode int = 401

/*WeaviateSchemaDumpUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateSchemaDumpUnauthorized
*/
type WeaviateSchemaDumpUnauthorized struct {
}

// NewWeaviateSchemaDumpUnauthorized creates WeaviateSchemaDumpUnauthorized with default headers values
func NewWeaviateSchemaDumpUnauthorized() *WeaviateSchemaDumpUnauthorized {

	return &WeaviateSchemaDumpUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateSchemaDumpUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}