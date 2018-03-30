// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// vendor tool to copy external source code from GOPATH or remote location to the
// local vendor folder. See README.md for usage.
package main

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"fmt"
)

func main() {
	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	// Set the AWS Region that the service clients should use
	cfg.Region = endpoints.UsWest2RegionID
	iamSvc := iam.New(cfg)
	userName := "bbradbury@xevo.com"
	keysListReq := iamSvc.ListSSHPublicKeysRequest(&iam.ListSSHPublicKeysInput{
		Marker:   nil,
		MaxItems: nil,
		UserName: &userName,
	})

	// Send the request, and get the response or error back
	keysListResp, err := keysListReq.Send()
	if err != nil {
		panic("failed to describe table, "+err.Error())
	}

	fmt.Println("Response", keysListResp)
	for _, r := range keysListResp.SSHPublicKeys {
		fmt.Println(r.SSHPublicKeyId)

		keyReq := iamSvc.GetSSHPublicKeyRequest(&iam.GetSSHPublicKeyInput{
			Encoding:       "SSH",
			SSHPublicKeyId: r.SSHPublicKeyId,
			UserName:       &userName,
		})

		keysResp, err := keyReq.Send()
		if err != nil {
			panic("Failed to fetch key, "+err.Error())
		}

		body := keysResp.SSHPublicKey.SSHPublicKeyBody
		fmt.Printf(*body)
	}

}