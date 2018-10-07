/*
Copyright Suzhou Tongji Fintech Research Institute 2017 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package factory

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/internal/github.com/hyperledger/fabric/bccsp"

	"github.com/hyperledger/fabric-sdk-go/internal/github.com/hyperledger/fabric/bccsp/gm"
	"github.com/hyperledger/fabric-sdk-go/internal/github.com/hyperledger/fabric/bccsp/factory/sw"
)

const (
	// GuomiBasedFactoryName is the name of the factory of the GuoMi BCCSP implementation
	GuomiBasedFactoryName = "GM"
)

// GMFactory is the factory of the guomi-based BCCSP.
type GMFactory struct{}

// Name returns the name of this factory
func (f *GMFactory) Name() string {
	return GuomiBasedFactoryName
}

// Get returns an instance of BCCSP using Opts.
func (f *GMFactory) Get(config *sw.SwOpts) (bccsp.BCCSP, error) {
	// Validate arguments
	if config == nil {
		return nil, errors.New("Invalid config. It must not be nil.")
	}

	gmOpts := config

	var ks bccsp.KeyStore
	if gmOpts.Ephemeral == true {
		ks = gm.NewDummyKeyStore()
	} else if gmOpts.FileKeystore != nil {
		fks, err := gm.NewFileBasedKeyStore(nil, gmOpts.FileKeystore.KeyStorePath, false)
		if err != nil {
			return nil, fmt.Errorf("Failed to initialize gm software key store: %s", err)
		}
		ks = fks
	} else {
		// Default to DummyKeystore
		ks = gm.NewDummyKeyStore()
	}
	// SM3，国密哈希函数算法标准
	return gm.New(gmOpts.SecLevel, "SM3", ks)
}
