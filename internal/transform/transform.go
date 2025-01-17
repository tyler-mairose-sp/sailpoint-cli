package transform

import (
	"context"
	"os"

	"github.com/olekukonko/tablewriter"
	sailpoint "github.com/sailpoint-oss/golang-sdk"
	sailpointsdk "github.com/sailpoint-oss/golang-sdk/v3"
	transmodel "github.com/sailpoint-oss/sailpoint-cli/cmd/transform/model"
	"github.com/sailpoint-oss/sailpoint-cli/internal/config"
	"github.com/sailpoint-oss/sailpoint-cli/internal/sdk"
)

func GetTransforms() ([]sailpointsdk.Transform, error) {
	var transforms []sailpointsdk.Transform

	apiClient, err := config.InitAPIClient()
	if err != nil {
		return transforms, err
	}

	transforms, resp, err := sailpoint.PaginateWithDefaults[sailpointsdk.Transform](apiClient.V3.TransformsApi.ListTransforms(context.TODO()))
	if err != nil {
		return transforms, sdk.HandleSDKError(resp, err)
	}

	return transforms, nil
}

func ListTransforms() error {

	transforms, err := GetTransforms()
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(transmodel.TransformColumns)
	for _, v := range transforms {
		table.Append([]string{*v.Id, v.Name})
	}
	table.Render()

	return nil
}
