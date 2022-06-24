// Code generated by go-swagger; DO NOT EDIT.

//
//   Copyright (c) 2022 Intel Corporation.
//
//   SPDX-License-Identifier: Apache-2.0
//
//
//

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Files files
//
// swagger:model files
type Files struct {

	// files
	Files []*FilesItems0 `json:"files"`
}

// Validate validates this files
func (m *Files) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Files) validateFiles(formats strfmt.Registry) error {
	if swag.IsZero(m.Files) { // not required
		return nil
	}

	for i := 0; i < len(m.Files); i++ {
		if swag.IsZero(m.Files[i]) { // not required
			continue
		}

		if m.Files[i] != nil {
			if err := m.Files[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this files based on the context it is used
func (m *Files) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Files) contextValidateFiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Files); i++ {

		if m.Files[i] != nil {
			if err := m.Files[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Files) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Files) UnmarshalBinary(b []byte) error {
	var res Files
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FilesItems0 files items0
//
// swagger:model FilesItems0
type FilesItems0 struct {

	// filename
	Filename string `json:"filename,omitempty"`

	// hash
	Hash string `json:"hash,omitempty"`

	// hashtype
	// Enum: [md5 sha256]
	Hashtype string `json:"hashtype,omitempty"`

	// mirrorurl
	// Pattern: (?:(?:https?|http|ftp|file|oci)://|www.|ftp.)(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[-A-Z0-9+&@#/%=~_|$?!:,.])*(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[A-Z0-9+&@#/%=~_|$])
	Mirrorurl string `json:"mirrorurl,omitempty"`

	// url
	// Pattern: (?:(?:https?|http|ftp|file|oci)://|www.|ftp.)(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[-A-Z0-9+&@#/%=~_|$?!:,.])*(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[A-Z0-9+&@#/%=~_|$])
	URL string `json:"url,omitempty"`

	// urlreplacement
	Urlreplacement *FilesItems0Urlreplacement `json:"urlreplacement,omitempty"`
}

// Validate validates this files items0
func (m *FilesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHashtype(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMirrorurl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUrlreplacement(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var filesItems0TypeHashtypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["md5","sha256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		filesItems0TypeHashtypePropEnum = append(filesItems0TypeHashtypePropEnum, v)
	}
}

const (

	// FilesItems0HashtypeMd5 captures enum value "md5"
	FilesItems0HashtypeMd5 string = "md5"

	// FilesItems0HashtypeSha256 captures enum value "sha256"
	FilesItems0HashtypeSha256 string = "sha256"
)

// prop value enum
func (m *FilesItems0) validateHashtypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, filesItems0TypeHashtypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FilesItems0) validateHashtype(formats strfmt.Registry) error {
	if swag.IsZero(m.Hashtype) { // not required
		return nil
	}

	// value enum
	if err := m.validateHashtypeEnum("hashtype", "body", m.Hashtype); err != nil {
		return err
	}

	return nil
}

func (m *FilesItems0) validateMirrorurl(formats strfmt.Registry) error {
	if swag.IsZero(m.Mirrorurl) { // not required
		return nil
	}

	if err := validate.Pattern("mirrorurl", "body", m.Mirrorurl, `(?:(?:https?|http|ftp|file|oci)://|www.|ftp.)(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[-A-Z0-9+&@#/%=~_|$?!:,.])*(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[A-Z0-9+&@#/%=~_|$])`); err != nil {
		return err
	}

	return nil
}

func (m *FilesItems0) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.Pattern("url", "body", m.URL, `(?:(?:https?|http|ftp|file|oci)://|www.|ftp.)(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[-A-Z0-9+&@#/%=~_|$?!:,.])*(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[A-Z0-9+&@#/%=~_|$])`); err != nil {
		return err
	}

	return nil
}

func (m *FilesItems0) validateUrlreplacement(formats strfmt.Registry) error {
	if swag.IsZero(m.Urlreplacement) { // not required
		return nil
	}

	if m.Urlreplacement != nil {
		if err := m.Urlreplacement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("urlreplacement")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("urlreplacement")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this files items0 based on the context it is used
func (m *FilesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUrlreplacement(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FilesItems0) contextValidateUrlreplacement(ctx context.Context, formats strfmt.Registry) error {

	if m.Urlreplacement != nil {
		if err := m.Urlreplacement.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("urlreplacement")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("urlreplacement")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FilesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilesItems0) UnmarshalBinary(b []byte) error {
	var res FilesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FilesItems0Urlreplacement files items0 urlreplacement
//
// swagger:model FilesItems0Urlreplacement
type FilesItems0Urlreplacement struct {

	// new
	New string `json:"new,omitempty"`

	// origin
	Origin string `json:"origin,omitempty"`
}

// Validate validates this files items0 urlreplacement
func (m *FilesItems0Urlreplacement) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this files items0 urlreplacement based on context it is used
func (m *FilesItems0Urlreplacement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FilesItems0Urlreplacement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilesItems0Urlreplacement) UnmarshalBinary(b []byte) error {
	var res FilesItems0Urlreplacement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}