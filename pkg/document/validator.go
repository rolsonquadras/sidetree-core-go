/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package document

import (
	"errors"
	"fmt"
)

const (
	// ops defines key usage as operations key
	ops = "ops"
	// auth defines key usage as authentication key
	auth = "auth"
	// assertion defines key usage as assertion key
	assertion = "assertion"
	// assertion defines key usage as assertion key
	agreement = "agreement"
	// general defines key usage as general key
	general = "general"

	jwsVerificationKey2020            = "JwsVerificationKey2020"
	ecdsaSecp256k1VerificationKey2019 = "EcdsaSecp256k1VerificationKey2019"
	x25519KeyAgreementKey2019         = "X25519KeyAgreementKey2019"

	// Ed25519VerificationKey2018 requires special handling (convert to base58)
	Ed25519VerificationKey2018 = "Ed25519VerificationKey2018"

	maxJwkProperties       = 4
	maxPublicKeyProperties = 4
)

var allowedOps = map[string]string{
	ops:       ops,
	auth:      auth,
	general:   general,
	assertion: assertion,
	agreement: agreement,
}

type existenceMap map[string]string

var allowedKeyTypesOps = existenceMap{
	jwsVerificationKey2020:            jwsVerificationKey2020,
	ecdsaSecp256k1VerificationKey2019: ecdsaSecp256k1VerificationKey2019,
}

var allowedKeyTypesGeneral = existenceMap{
	jwsVerificationKey2020:            jwsVerificationKey2020,
	ecdsaSecp256k1VerificationKey2019: ecdsaSecp256k1VerificationKey2019,
	Ed25519VerificationKey2018:        Ed25519VerificationKey2018,
	x25519KeyAgreementKey2019:         x25519KeyAgreementKey2019,
}

var allowedKeyTypesVerification = existenceMap{
	jwsVerificationKey2020:            jwsVerificationKey2020,
	ecdsaSecp256k1VerificationKey2019: ecdsaSecp256k1VerificationKey2019,
	Ed25519VerificationKey2018:        Ed25519VerificationKey2018,
}

var allowedKeyTypesAgreement = existenceMap{
	// TODO: Verify appropriate agreement key types for JWS and Secp256k1
	jwsVerificationKey2020:            jwsVerificationKey2020,
	ecdsaSecp256k1VerificationKey2019: ecdsaSecp256k1VerificationKey2019,
	x25519KeyAgreementKey2019:         x25519KeyAgreementKey2019,
}

var allowedKeyTypes = map[string]existenceMap{
	ops:       allowedKeyTypesOps,
	general:   allowedKeyTypesGeneral,
	auth:      allowedKeyTypesVerification,
	assertion: allowedKeyTypesVerification,
	agreement: allowedKeyTypesAgreement,
}

// ValidatePublicKeys validates public keys
func ValidatePublicKeys(pubKeys []PublicKey) error {
	ids := make(map[string]string)

	// the expected fields are id, usage, type and jwk
	for _, pubKey := range pubKeys {
		kid := pubKey.ID()
		if kid == "" {
			return errors.New("public key id is missing")
		}

		if len(pubKey) != maxPublicKeyProperties {
			return errors.New("invalid number of public key properties")
		}

		if _, ok := ids[kid]; ok {
			return fmt.Errorf("duplicate public key id: %s", kid)
		}
		ids[kid] = kid

		if err := validateKeyUsage(pubKey); err != nil {
			return err
		}

		if IsOperationsKey(pubKey.Usage()) {
			if err := ValidateOperationsKey(pubKey); err != nil {
				return err
			}
		}

		if !validateKeyTypeUsage(pubKey) {
			return fmt.Errorf("invalid key type: %s", pubKey.Type())
		}

		if err := ValidateJWK(pubKey.JWK()); err != nil {
			return err
		}
	}

	return nil
}

// validateKeyTypeUsage validates if the public key type is valid for a certain usage
func validateKeyTypeUsage(pubKey PublicKey) bool {
	for _, usage := range pubKey.Usage() {
		allowed, ok := allowedKeyTypes[usage]
		if !ok {
			return false
		}

		_, ok = allowed[pubKey.Type()]
		if !ok {
			return false
		}
	}

	return true
}

// ValidateOperationsKey validates operation key
func ValidateOperationsKey(pubKey PublicKey) error {
	if !IsOperationsKey(pubKey.Usage()) {
		return fmt.Errorf("key '%s' is not an operations key", pubKey.ID())
	}

	return ValidateJWK(pubKey.JWK())
}

// ValidateJWK validates JWK
func ValidateJWK(jwk JWK) error {
	if jwk == nil {
		return errors.New("key has to be in JWK format")
	}

	if len(jwk) != maxJwkProperties {
		return errors.New("invalid number of JWK properties")
	}

	if jwk.Crv() == "" {
		return errors.New("JWK crv is missing")
	}

	if jwk.Kty() == "" {
		return errors.New("JWK kty is missing")
	}

	if jwk.X() == "" {
		return errors.New("JWK x is missing")
	}

	return nil
}

// IsOperationsKey returns true if key is an operations key
func IsOperationsKey(usages []string) bool {
	return isUsageKey(usages, ops)
}

// IsGeneralKey returns true if key is a general key
func IsGeneralKey(usages []string) bool {
	return isUsageKey(usages, general)
}

// IsAuthenticationKey returns true if key is an authentication key
func IsAuthenticationKey(usages []string) bool {
	return isUsageKey(usages, auth)
}

// IsAssertionKey returns true if key is an assertion key
func IsAssertionKey(usages []string) bool {
	return isUsageKey(usages, assertion)
}

// IsAgreementKey returns true if key is an agreement key
func IsAgreementKey(usages []string) bool {
	return isUsageKey(usages, agreement)
}

func isUsageKey(usages []string, mode string) bool {
	for _, usage := range usages {
		if usage == mode {
			return true
		}
	}

	return false
}

// The object MUST include a usage property, and its value MUST be an array that includes one or more of the following:
// - ops: the key is allowed to generate DID operations for the DID.
// - general: the key is to be included in the publicKeys section of the resolved DID Document.
// - auth: the key is to be included in the authentication section of the resolved DID Document
func validateKeyUsage(pubKey PublicKey) error {
	if len(pubKey.Usage()) == 0 {
		return fmt.Errorf("key '%s' is missing usage", pubKey.ID())
	}

	if len(pubKey.Usage()) > len(allowedOps) {
		return fmt.Errorf("public key usage exceeds maximum length: %d", len(allowedOps))
	}

	for _, usage := range pubKey.Usage() {
		if _, ok := allowedOps[usage]; !ok {
			return fmt.Errorf("invalid usage: %s", usage)
		}
	}

	return nil
}