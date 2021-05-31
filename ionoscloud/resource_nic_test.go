package ionoscloud

import (
	"context"
	"fmt"
	ionoscloud "github.com/ionos-cloud/sdk-go/v6"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccNic_Basic(t *testing.T) {
	var nic ionoscloud.Nic
	volumeName := "volume"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNicDestroyCheck,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccCheckNicConfig_basic, volumeName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNICExists("ionoscloud_nic.database_nic", &nic),
					testAccCheckNicAttributes("ionoscloud_nic.database_nic", volumeName),
					resource.TestCheckResourceAttrSet("ionoscloud_nic.database_nic", "mac"),
					resource.TestCheckResourceAttr("ionoscloud_nic.database_nic", "name", volumeName),
				),
			},
			{
				Config: testAccCheckNicConfig_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNicAttributes("ionoscloud_nic.database_nic", "updated"),
					resource.TestCheckResourceAttr("ionoscloud_nic.database_nic", "name", "updated"),
				),
			},
		},
	})
}

func testAccCheckNicDestroyCheck(s *terraform.State) error {
	client := testAccProvider.Meta().(*ionoscloud.APIClient)

	ctx, _ := context.WithTimeout(context.Background(), *resourceDefaultTimeouts.Delete)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ionoscloud_nic" {
			continue
		}

		dcId := rs.Primary.Attributes["datacenter_id"]
		serverId := rs.Primary.Attributes["server_id"]
<<<<<<< HEAD
		_, apiResponse, err := client.NetworkInterfacesApi.DatacentersServersNicsFindById(ctx, dcId, serverId, rs.Primary.ID).Execute()

		if apiError, ok := err.(ionoscloud.GenericOpenAPIError); ok {
			if apiResponse.Response.StatusCode != 404 {
				return fmt.Errorf("NIC still exists %s %s", rs.Primary.ID, apiError)
=======
		_, apiResponse, _ := client.NicApi.DatacentersServersNicsFindById(ctx, dcId, serverId, rs.Primary.ID).Execute()

		if apiResponse == nil || apiResponse.Response.StatusCode != 404 {
			var payload = "<nil>"
			var statusCode = 0
			if apiResponse != nil {
				payload = string(apiResponse.Payload)
				statusCode = apiResponse.StatusCode
>>>>>>> master
			}
			return fmt.Errorf("NIC still exists %s: %d %s", rs.Primary.ID, statusCode, payload)
		}
	}

	return nil
}

func testAccCheckNicAttributes(n string, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("testAccCheckNicAttributes: Not found: %s", n)
		}
		if rs.Primary.Attributes["name"] != name {
			return fmt.Errorf("bad name: %s", rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func testAccCheckNICExists(n string, nic *ionoscloud.Nic) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*ionoscloud.APIClient)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("testAccCheckVolumeExists: Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		ctx, _ := context.WithTimeout(context.Background(), *resourceDefaultTimeouts.Default)
		dcId := rs.Primary.Attributes["datacenter_id"]
		serverId := rs.Primary.Attributes["server_id"]
		foundNic, _, err := client.NetworkInterfacesApi.DatacentersServersNicsFindById(ctx, dcId, serverId, rs.Primary.ID).Execute()

		if err != nil {
			return fmt.Errorf("error occured while fetching Volume: %s", rs.Primary.ID)
		}
		if *foundNic.Id != rs.Primary.ID {
			return fmt.Errorf("Record not found")
		}

		nic = &foundNic

		return nil
	}
}

const testAccCheckNicConfig_basic = `
resource "ionoscloud_datacenter" "foobar" {
	name       = "nic-test"
	location = "us/las"
}

resource "ionoscloud_server" "webserver" {
  name = "webserver"
  datacenter_id = "${ionoscloud_datacenter.foobar.id}"
  cores = 1
  ram = 1024
  availability_zone = "ZONE_1"
  cpu_family = "AMD_OPTERON"
    image = "81e054dd-a347-11eb-b70c-7ade62b52cc0"
	image_password = "K3tTj8G14a3EgKyNeeiY"
  volume {
    name = "system"
    size = 14
    disk_type = "SSD"

}
  nic {
    lan = "1"
    dhcp = true
    firewall_active = true
  }
}

resource "ionoscloud_nic" "database_nic" {
  datacenter_id = "${ionoscloud_datacenter.foobar.id}"
  server_id = "${ionoscloud_server.webserver.id}"
  lan = 2
  dhcp = false
  firewall_active = true
  firewall_type = "INGRESS"
  name = "%s"
}`

const testAccCheckNicConfig_update = `
resource "ionoscloud_datacenter" "foobar" {
	name       = "nic-test"
	location = "us/las"
}

resource "ionoscloud_server" "webserver" {
  name = "webserver"
  datacenter_id = "${ionoscloud_datacenter.foobar.id}"
  cores = 1
  ram = 1024
  availability_zone = "ZONE_1"
  cpu_family = "AMD_OPTERON"
	image = "81e054dd-a347-11eb-b70c-7ade62b52cc0"
	image_password = "K3tTj8G14a3EgKyNeeiY"
  volume {
    name = "system"
    size = 14
    disk_type = "SSD"
}
  nic {
    lan = "1"
    dhcp = true
    firewall_active = true
  }
}

resource "ionoscloud_nic" "database_nic" {
  datacenter_id = "${ionoscloud_datacenter.foobar.id}"
  server_id = "${ionoscloud_server.webserver.id}"
  lan = 2
  dhcp = false
  firewall_active = true
  firewall_type = "INGRESS"
  name = "updated"
}
`
