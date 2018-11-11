
/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package msp

import (
	"path"
	"strings"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/core"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/keyvaluestore"
	"github.com/pkg/errors"
)

// NewFileCertStore ...
func NewFileCertStore(cryptoConfigMSPPath string) (core.KVStore, error) {
	opts := &keyvaluestore.FileKeyValueStoreOptions{
		Path: cryptoConfigMSPPath,
		KeySerializer: func(key interface{}) (string, error) {
			ck, ok := key.(*msp.IdentityIdentifier)
			if !ok {
				return "", errors.New("converting key to CertKey failed")
			}
			if ck == nil || ck.MSPID == "" || ck.ID == "" {
				return "", errors.New("invalid key")
			}
			// TODO: refactor to case insensitive or remove eventually.
			// 文件夹去除Mspid名称
			r := strings.NewReplacer("{userName}", ck.ID)
			certDir := path.Join(r.Replace(cryptoConfigMSPPath), "signcerts")
			// 文件名精简为cert.pem
			return path.Join(certDir, "cert.pem"), nil
		},
	}
	return keyvaluestore.New(opts)
}