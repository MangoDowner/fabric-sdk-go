package gm

import (
	"testing"
	"reflect"
	"github.com/stretchr/testify/assert"
	mocks2 "github.com/hyperledger/fabric/bccsp/mocks"
	"errors"
	"github.com/hyperledger/fabric/bccsp/sw/mocks"
	"fmt"
)

func TestHash(t *testing.T){


	t.Parallel()

	expectetMsg := []byte{1, 2, 3, 4}
	expectedOpts := &mocks2.HashOpts{}
	expectetValue := []byte{1, 2, 3, 4, 5}
	expectedErr := errors.New("Expected Error")

	hashers := make(map[reflect.Type]Hasher)
	hashers[reflect.TypeOf(&mocks2.HashOpts{})] = &mocks.Hasher{
		MsgArg:  expectetMsg,
		OptsArg: expectedOpts,
		Value:   expectetValue,
		Err:     nil,
	}
	csp := impl{hashers: hashers}
	value, err := csp.Hash(expectetMsg, expectedOpts)
	fmt.Println("hash value",value);
	assert.Equal(t, expectetValue, value)
	assert.Nil(t, err)

	hashers = make(map[reflect.Type]Hasher)
	hashers[reflect.TypeOf(&mocks2.HashOpts{})] = &mocks.Hasher{
		MsgArg:  expectetMsg,
		OptsArg: expectedOpts,
		Value:   nil,
		Err:     expectedErr,
	}
	csp = impl{hashers: hashers}
	value, err = csp.Hash(expectetMsg, expectedOpts)
	assert.Nil(t, value)
	assert.Contains(t, err.Error(), expectedErr.Error())
	fmt.Println(value)
	fmt.Println(value)




}
