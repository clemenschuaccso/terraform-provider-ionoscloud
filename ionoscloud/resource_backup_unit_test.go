package ionoscloud

import (
	"context"
	"fmt"
	ionoscloud "github.com/ionos-cloud/sdk-go/v6"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccBackupUnitBasic(t *testing.T) {
	var backupUnit ionoscloud.BackupUnit

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckBackupUnitDestroyCheck,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckBackupUnitConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBackupUnitExists(BackupUnitResource+"."+BackupUnitTestResource, &backupUnit),
					resource.TestCheckResourceAttr(BackupUnitResource+"."+BackupUnitTestResource, "name", BackupUnitTestResource),
					resource.TestCheckResourceAttr(BackupUnitResource+"."+BackupUnitTestResource, "email", "example@ionoscloud.com"),
					resource.TestCheckResourceAttr(BackupUnitResource+"."+BackupUnitTestResource, "password", "DemoPassword123$"),
				),
			},
			{
				Config: testAccCheckBackupUnitConfigUpdatePassword,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBackupUnitExists(BackupUnitResource+"."+BackupUnitTestResource, &backupUnit),
					resource.TestCheckResourceAttr(BackupUnitResource+"."+BackupUnitTestResource, "name", BackupUnitTestResource),
					resource.TestCheckResourceAttr(BackupUnitResource+"."+BackupUnitTestResource, "email", "example@ionoscloud.com"),
					resource.TestCheckResourceAttr(BackupUnitResource+"."+BackupUnitTestResource, "password", "DemoPassword1234$Updated"),
				),
			},
			{
				Config: testAccCheckBackupUnitConfigUpdateEmail,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBackupUnitExists(BackupUnitResource+"."+BackupUnitTestResource, &backupUnit),
					resource.TestCheckResourceAttr(BackupUnitResource+"."+BackupUnitTestResource, "name", BackupUnitTestResource),
					resource.TestCheckResourceAttr(BackupUnitResource+"."+BackupUnitTestResource, "email", "example-updated@ionoscloud.com"),
					resource.TestCheckResourceAttr(BackupUnitResource+"."+BackupUnitTestResource, "password", "DemoPassword1234$Updated"),
				),
			},
		},
	})
}

func testAccCheckBackupUnitDestroyCheck(s *terraform.State) error {
	client := testAccProvider.Meta().(SdkBundle).CloudApiClient

	ctx, cancel := context.WithTimeout(context.Background(), *resourceDefaultTimeouts.Default)

	if cancel != nil {
		defer cancel()
	}

	for _, rs := range s.RootModule().Resources {

		if rs.Type != BackupUnitResource {
			continue
		}

		_, apiResponse, err := client.BackupUnitsApi.BackupunitsFindById(ctx, rs.Primary.ID).Execute()
		logApiRequestTime(apiResponse)

		if err != nil {
			if apiResponse == nil || apiResponse.Response != nil && apiResponse.StatusCode != 404 {
				return fmt.Errorf("an error occurred while checking for the destruction of backup unit %s: %s",
					rs.Primary.ID, err)
			}
		} else {
			return fmt.Errorf("backup unit %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckBackupUnitExists(n string, backupUnit *ionoscloud.BackupUnit) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(SdkBundle).CloudApiClient

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		ctx, cancel := context.WithTimeout(context.Background(), *resourceDefaultTimeouts.Default)

		if cancel != nil {
			defer cancel()
		}

		foundBackupUnit, apiResponse, err := client.BackupUnitsApi.BackupunitsFindById(ctx, rs.Primary.ID).Execute()
		logApiRequestTime(apiResponse)

		if err != nil {
			return fmt.Errorf("error occured while fetching backup unit: %s", rs.Primary.ID)
		}
		if *foundBackupUnit.Id != rs.Primary.ID {
			return fmt.Errorf("record not found")
		}
		backupUnit = &foundBackupUnit

		return nil
	}
}

const testAccCheckBackupUnitConfigBasic = `
resource ` + BackupUnitResource + ` ` + BackupUnitTestResource + ` {
	name        = "` + BackupUnitTestResource + `"
	password    = "DemoPassword123$"
	email       = "example@ionoscloud.com"
}
`

const testAccCheckBackupUnitConfigUpdatePassword = `
resource ` + BackupUnitResource + ` ` + BackupUnitTestResource + ` {
	name        = "` + BackupUnitTestResource + `"
	password    = "DemoPassword1234$Updated"
	email       = "example@ionoscloud.com"
}
`

const testAccCheckBackupUnitConfigUpdateEmail = `
resource ` + BackupUnitResource + ` ` + BackupUnitTestResource + ` {
	name        = "` + BackupUnitTestResource + `"
	password    = "DemoPassword1234$Updated"
	email       = "example-updated@ionoscloud.com"
}
`
