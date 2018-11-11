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

// NewFileKeyStore ...
func NewFileKeyStore(cryptoConfigMSPPath string) (core.KVStore, error) {
	opts := &keyvaluestore.FileKeyValueStoreOptions{
		Path: cryptoConfigMSPPath,
		KeySerializer: func(key interface{}) (string, error) {
			pkk, ok := key.(*msp.PrivKeyKey)
			if !ok {
				return "", errors.New("converting key to PrivKeyKey failed")
			}
			if pkk == nil || pkk.MSPID == "" || pkk.ID == "" || pkk.SKI == nil {
				return "", errors.New("invalid key")
			}

			// TODO: refactor to case insensitive or remove eventually.
			r := strings.NewReplacer("{userName}", pkk.ID, "{username}", pkk.ID)
			keyDir := path.Join(r.Replace(cryptoConfigMSPPath), "keystore")
			// 文件名统一为_sk
			return path.Join(keyDir, "_sk"), nil
		},
	}
	return keyvaluestore.New(opts)
}
