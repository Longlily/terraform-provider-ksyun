// Copyright 2022 NotOne Lv <aiphalv0010@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package ksyun

import (
	"fmt"
	"time"

	"github.com/KscSDK/ksc-sdk-go/service/kec"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

type DataGuardSrv struct {
	client *KsyunClient
}

func (d *DataGuardSrv) describeDataGuardGroup(input map[string]interface{}) (data []interface{}, err error) {
	var (
		resp *map[string]interface{}
	)

	resp, err = d.GetConn().DescribeDataGuardGroup(&input)

	if err != nil {
		return nil, err
	}
	results, err := getSdkValue("DataGuardsSet", *resp)
	if err != nil || results == nil {
		return nil, fmt.Errorf("the current available zone not exsits any data guard group")
	}
	data = results.([]interface{})

	return data, err
}

// createDataGuardGroup will create data guard group and it returns this data guard group id
func (d *DataGuardSrv) createDataGuardGroup(input map[string]interface{}) (string, error) {
	var (
		resp *map[string]interface{}
		err  error
	)
	resp, err = d.GetConn().CreateDataGuardGroup(&input)

	results, err := getSdkValue("DataGuardId", *resp)
	if err != nil || results == nil {
		return "", err
	}
	guardId := results.(string)
	return guardId, err
}

func (d *DataGuardSrv) deleteDataGuardGroup(input map[string]interface{}) error {

	return resource.Retry(3*time.Minute, func() *resource.RetryError {
		if _, ok := input["DBParameterGroupId"]; !ok && input["DBParameterGroupId"].(string) != "" {
			_, err := d.GetConn().DeleteDataGuardGroups(&input)
			// logger.Debug("test %s %s %s", "DeleteDBParameterGroup", inUseError(deleteErr), deleteErr)
			if err == nil || notFoundErrorNew(err) || inUseError(err) {
				return nil
			} else {
				return resource.RetryableError(err)
			}
		}
		return nil // resource.RetryableError(nil)
	})
}

func (d *DataGuardSrv) modifyModifyDataGuardGroups(input map[string]interface{}) (map[string]interface{}, error) {
	var (
		resp *map[string]interface{}
		err  error
	)
	resp, err = d.GetConn().ModifyDataGuardGroups(&input)
	return *resp, err
}

func (d *DataGuardSrv) GetConn() *kec.Kec {
	return d.client.kecconn
}
