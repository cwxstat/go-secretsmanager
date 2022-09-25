//Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
//SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"

	"github.com/mchirico/go-secretsmanager/common"
	"github.com/mchirico/go-secretsmanager/config"

	"github.com/google/uuid"
)

func main() {

	var secretArn string
	var secretName string
	var value string

	cfg, err := config.Config()

	if err != nil {
		panic("Couldn't load config!")
	}

	secretName = uuid.NewString()
	value = "s00pers33kr1t"

	conn := secretsmanager.NewFromConfig(cfg)

	if secretArn, err = common.CreateSecret(context.TODO(), conn, secretName, value, "desc"); err != nil {
		panic("Couldn't create secret!: " + err.Error())
	}
	fmt.Printf("Created the arn %v\n", secretArn)

	if value, err = common.GetSecret(context.TODO(), conn, secretArn); err != nil {
		panic("Couldn't get secret value!")
	}
	fmt.Printf("it has the value \"%v\"\n", value)

	if err = common.UpdateSecret(context.TODO(), conn, secretArn, "correct horse battery staple"); err != nil {
		panic("Couldn't update secret!")
	}
	fmt.Println("The secret has been updated.")

	if value, err = common.GetSecret(context.TODO(), conn, secretArn); err != nil {
		panic("Couldn't get secret value!")
	}
	fmt.Printf("it has the value \"%v\"\n", value)

	var secretIds []string

	if secretIds, err = common.ListSecrets(context.TODO(), conn); err != nil {
		panic("Couldn't list secrets!")
	}

	fmt.Printf("There are %v secrets -- here's their IDs: \n", len(secretIds))
	for _, id := range secretIds {
		fmt.Println(id)
	}

	if err = common.DeleteSecret(context.TODO(), conn, secretArn); err != nil {
		panic("Couldn't delete secret!")
	}
	fmt.Printf("Deleted the secret with arn %v\n", secretArn)
}
