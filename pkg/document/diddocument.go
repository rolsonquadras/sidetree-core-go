/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package document

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

const (

	// ContextProperty defines key for context property
	ContextProperty = "@context"

	// ServiceProperty defines key for service property
	ServiceProperty = "service"

	// PublicKeyProperty defines key for public key property
	PublicKeyProperty = "publicKey"

	// AuthenticationProperty defines key for authentication property
	AuthenticationProperty = "authentication"

	// ControllerProperty defines key for controller
	ControllerProperty = "controller"

	jsonldType         = "type"
	jsonldServicePoint = "serviceEndpoint"

	// various public key encodings
	jsonldPublicKeyBase64 = "publicKeyBase64"
	jsonldPublicKeyBase58 = "publicKeyBase58"
	jsonldPublicKeyHex    = "publicKeyHex"
	jsonldPublicKeyPem    = "publicKeyPem"
	jsonldPublicKeyJwk    = "publicKeyJwk"

	// key usage
	jsonldPublicKeyUsage = "usage"
)

// DIDDocument Defines DID Document data structure used by Sidetree for basic type safety checks.
type DIDDocument map[string]interface{}

// ID is identifier for DID subject (what DID Document is about)
func (doc DIDDocument) ID() string {
	return stringEntry(doc[IDProperty])
}

// Context is the context of did document
func (doc DIDDocument) Context() []string {
	return stringArray(doc[ContextProperty])
}

// PublicKeys are used for digital signatures, encryption and other cryptographic operations
func (doc DIDDocument) PublicKeys() []PublicKey {
	entry, ok := doc[PublicKeyProperty]
	if !ok {
		return nil
	}

	typedEntry, ok := entry.([]interface{})
	if !ok {
		return nil
	}

	var result []PublicKey
	for _, e := range typedEntry {
		emap, ok := e.(map[string]interface{})
		if !ok {
			continue
		}
		result = append(result, NewPublicKey(emap))
	}
	return result
}

// Services is an array of service endpoints
func (doc DIDDocument) Services() []Service {
	entry, ok := doc[ServiceProperty]
	if !ok {
		return nil
	}

	typedEntry, ok := entry.([]interface{})
	if !ok {
		return nil
	}

	var result []Service
	for _, e := range typedEntry {
		emap, ok := e.(map[string]interface{})
		if !ok {
			continue
		}
		result = append(result, NewService(emap))
	}
	return result
}

// JSONLdObject returns map that represents JSON LD Object
func (doc DIDDocument) JSONLdObject() map[string]interface{} {
	return doc
}

// Authentication return authentication array (mixture of strings and objects)
func (doc DIDDocument) Authentication() []interface{} {
	return interfaceArray(doc[AuthenticationProperty])
}

// DIDDocumentFromReader creates an instance of DIDDocument by reading a JSON document from Reader
func DIDDocumentFromReader(r io.Reader) (DIDDocument, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return DidDocumentFromBytes(data)
}

// DidDocumentFromBytes creates an instance of DIDDocument by reading a JSON document from bytes
func DidDocumentFromBytes(data []byte) (DIDDocument, error) {
	doc := make(DIDDocument)
	err := json.Unmarshal(data, &doc)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// DidDocumentFromJSONLDObject creates an instance of DIDDocument from json ld object
func DidDocumentFromJSONLDObject(jsonldObject map[string]interface{}) DIDDocument {
	return jsonldObject
}

// interfaceArray
func interfaceArray(entry interface{}) []interface{} {
	if entry == nil {
		return nil
	}

	entries, ok := entry.([]interface{})
	if !ok {
		return nil
	}

	return entries
}
