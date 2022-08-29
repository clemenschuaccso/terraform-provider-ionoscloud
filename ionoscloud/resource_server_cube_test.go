//go:build compute || all || server

package ionoscloud

import (
	"context"
	"fmt"
	ionoscloud "github.com/ionos-cloud/sdk-go/v6"
	"github.com/ionos-cloud/terraform-provider-ionoscloud/v6/utils"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const bootCdromImageIdCube = "83f21679-3321-11eb-a681-1e659523cb7b"

func TestAccCubeServerBasic(t *testing.T) {
	var server ionoscloud.Server

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCubeServerDestroyCheck,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCubeServerConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCubeServerExists(ServerCubeResource+"."+ServerTestResource, &server),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "name", ServerTestResource),
					//resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "cores", "1"),
					//resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "ram", "1024"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "availability_zone", "ZONE_1"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "cpu_family", "INTEL_SKYLAKE"),
					utils.TestImageNotNull(ServerCubeResource, "boot_image"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "image_password", "K3tTj8G14a3EgKyNeeiY"),
					//resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "type", "CUBE"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "volume.0.name", "system"),
					//resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "volume.0.size", "5"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "volume.0.disk_type", "DAS"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "volume.0.bus", "VIRTIO"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "volume.0.availability_zone", "AUTO"),
					resource.TestCheckResourceAttrPair(ServerCubeResource+"."+ServerTestResource, "volume.0.boot_server", ServerCubeResource+"."+ServerTestResource, "id"),
					resource.TestCheckResourceAttrPair(ServerCubeResource+"."+ServerTestResource, "nic.0.lan", LanResource+"."+LanTestResource, "id"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.name", "system"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.dhcp", "true"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall_active", "true"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall_type", "BIDIRECTIONAL"),
					resource.TestCheckResourceAttrPair(ServerCubeResource+"."+ServerTestResource, "nic.0.ips.0", "ionoscloud_ipblock.webserver_ipblock", "ips.0"),
					resource.TestCheckResourceAttrPair(ServerCubeResource+"."+ServerTestResource, "nic.0.ips.1", "ionoscloud_ipblock.webserver_ipblock", "ips.1"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.protocol", "TCP"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.name", "SSH"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.port_range_start", "22"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.port_range_end", "22"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.source_mac", "00:0a:95:9d:68:17"),
					resource.TestCheckResourceAttrPair(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.source_ip", "ionoscloud_ipblock.webserver_ipblock", "ips.2"),
					resource.TestCheckResourceAttrPair(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.target_ip", "ionoscloud_ipblock.webserver_ipblock", "ips.3"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.type", "EGRESS"),
				),
			},
			{
				Config: testAccDataSourceCubeServerMatchId,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "name", ServerCubeResource+"."+ServerTestResource, "name"),
					//resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "cores", ServerCubeResource+"."+ServerTestResource, "cores"),
					//resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "ram", ServerCubeResource+"."+ServerTestResource, "ram"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "availability_zone", ServerCubeResource+"."+ServerTestResource, "availability_zone"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "cpu_family", ServerCubeResource+"."+ServerTestResource, "cpu_family"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "type", ServerCubeResource+"."+ServerTestResource, "type"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "volumes.0.name", ServerCubeResource+"."+ServerTestResource, "volume.0.name"),
					//resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "volumes.0.size", ServerCubeResource+"."+ServerTestResource, "volume.0.size"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "volumes.0.type", ServerCubeResource+"."+ServerTestResource, "volume.0.disk_type"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "volumes.0.bus", ServerCubeResource+"."+ServerTestResource, "volume.0.bus"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "volumes.0.availability_zone", ServerCubeResource+"."+ServerTestResource, "volume.0.availability_zone"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "volumes.0.boot_server", ServerCubeResource+"."+ServerTestResource, "id"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "nics.0.lan", ServerCubeResource+"."+ServerTestResource, "nic.0.lan"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "nics.0.name", ServerCubeResource+"."+ServerTestResource, "nic.0.name"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "nics.0.dhcp", ServerCubeResource+"."+ServerTestResource, "nic.0.dhcp"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "nics.0.firewall_active", ServerCubeResource+"."+ServerTestResource, "nic.0.firewall_active"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "nics.0.firewall_type", ServerCubeResource+"."+ServerTestResource, "nic.0.firewall_type"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "nics.0.ips.0", ServerCubeResource+"."+ServerTestResource, "nic.0.ips.0"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "nics.0.ips.1", ServerCubeResource+"."+ServerTestResource, "nic.0.ips.1"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "nics.0.firewall_rules.0.protocol", ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.protocol"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "nics.0.firewall_rules.0.name", ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.name"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "nics.0.firewall_rules.0.port_range_start", ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.port_range_start"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "nics.0.firewall_rules.0.port_range_end", ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.port_range_end"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "nics.0.firewall_rules.0.source_mac", ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.source_mac"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "nics.0.firewall_rules.0.source_ip", ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.source_ip"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "nics.0.firewall_rules.0.target_ip", ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.target_ip"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceById, "nics.0.firewall_rules.0.type", ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.type"),
				),
			},
			{
				Config: testAccDataSourceCubeServerMatchName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "name", ServerCubeResource+"."+ServerTestResource, "name"),
					//resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "cores", ServerCubeResource+"."+ServerTestResource, "cores"),
					//resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "ram", ServerCubeResource+"."+ServerTestResource, "ram"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "availability_zone", ServerCubeResource+"."+ServerTestResource, "availability_zone"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "cpu_family", ServerCubeResource+"."+ServerTestResource, "cpu_family"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "type", ServerCubeResource+"."+ServerTestResource, "type"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "volumes.0.name", ServerCubeResource+"."+ServerTestResource, "volume.0.name"),
					//resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "volumes.0.size", ServerCubeResource+"."+ServerTestResource, "volume.0.size"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "volumes.0.type", ServerCubeResource+"."+ServerTestResource, "volume.0.disk_type"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "volumes.0.bus", ServerCubeResource+"."+ServerTestResource, "volume.0.bus"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "volumes.0.boot_server", ServerCubeResource+"."+ServerTestResource, "id"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "volumes.0.availability_zone", ServerCubeResource+"."+ServerTestResource, "volume.0.availability_zone"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "nics.0.lan", ServerCubeResource+"."+ServerTestResource, "nic.0.lan"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "nics.0.name", ServerCubeResource+"."+ServerTestResource, "nic.0.name"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "nics.0.dhcp", ServerCubeResource+"."+ServerTestResource, "nic.0.dhcp"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "nics.0.firewall_active", ServerCubeResource+"."+ServerTestResource, "nic.0.firewall_active"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "nics.0.firewall_type", ServerCubeResource+"."+ServerTestResource, "nic.0.firewall_type"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "nics.0.ips.0", ServerCubeResource+"."+ServerTestResource, "nic.0.ips.0"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "nics.0.ips.1", ServerCubeResource+"."+ServerTestResource, "nic.0.ips.1"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "nics.0.firewall_rules.0.protocol", ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.protocol"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "nics.0.firewall_rules.0.name", ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.name"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "nics.0.firewall_rules.0.port_range_start", ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.port_range_start"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "nics.0.firewall_rules.0.port_range_end", ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.port_range_end"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "nics.0.firewall_rules.0.source_mac", ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.source_mac"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "nics.0.firewall_rules.0.source_ip", ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.source_ip"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "nics.0.firewall_rules.0.target_ip", ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.target_ip"),
					resource.TestCheckResourceAttrPair(DataSource+"."+ServerCubeResource+"."+ServerDataSourceByName, "nics.0.firewall_rules.0.type", ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.type"),
				),
			},
			{
				Config:      testAccDataSourceCubeServerWrongNameError,
				ExpectError: regexp.MustCompile(`no server found with the specified criteria: name`),
			},
			//{
			//	Config: testAccCheckCubeServerConfigUpdate,
			//	Check: resource.ComposeTestCheckFunc(
			//		testAccCheckServerExists(ServerCubeResource+"."+ServerTestResource, &server),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "name", UpdatedResources),
			//		//resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "cores", "2"),
			//		//resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "ram", "2048"),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "availability_zone", "ZONE_1"),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "cpu_family", "INTEL_SKYLAKE"),
			//		utils.TestImageNotNull(ServerCubeResource, "boot_image"),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "image_password", "K3tTj8G14a3EgKyNeeiYsasad"),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "volume.0.name", UpdatedResources),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "type", "CUBE"),
			//		//resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "volume.0.size", "6"),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "volume.0.disk_type", "SSD Standard"),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "volume.0.bus", "IDE"),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "volume.0.availability_zone", "ZONE_1"),
			//		resource.TestCheckResourceAttrPair(ServerCubeResource+"."+ServerTestResource, "nic.0.lan", LanResource+"."+LanTestResource, "id"),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.name", UpdatedResources),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.dhcp", "false"),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall_active", "false"),
			//		resource.TestCheckResourceAttrPair(ServerCubeResource+"."+ServerTestResource, "nic.0.ips.0", "ionoscloud_ipblock.webserver_ipblock_update", "ips.0"),
			//		resource.TestCheckResourceAttrPair(ServerCubeResource+"."+ServerTestResource, "nic.0.ips.1", "ionoscloud_ipblock.webserver_ipblock_update", "ips.1"),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.protocol", "TCP"),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.name", UpdatedResources),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.port_range_start", "21"),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.port_range_end", "23"),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.source_mac", "00:0a:95:9d:68:18"),
			//		resource.TestCheckResourceAttrPair(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.source_ip", "ionoscloud_ipblock.webserver_ipblock_update", "ips.2"),
			//		resource.TestCheckResourceAttrPair(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.target_ip", "ionoscloud_ipblock.webserver_ipblock_update", "ips.3"),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.type", "INGRESS"),
			//	),
			//},
		},
	})
}

func TestAccCubeServerBootCdromNoImage(t *testing.T) {
	var server ionoscloud.Server

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCubeServerDestroyCheck,
		Steps:             []resource.TestStep{
			//{
			//	Config: testAccCheckCubeServerConfigBootCdromNoImage,
			//	Check: resource.ComposeTestCheckFunc(
			//		testAccCheckServerExists(ServerCubeResource+"."+ServerTestResource, &server),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "name", ServerTestResource),
			//		//resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "cores", "1"),
			//		//resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "ram", "1024"),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "availability_zone", "ZONE_1"),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "cpu_family", "INTEL_SKYLAKE"),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "volume.0.name", ServerTestResource),
			//		//resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "volume.0.size", "5"),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "volume.0.disk_type", "DAS"),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "volume.0.licence_type", "OTHER"),
			//		resource.TestCheckResourceAttrPair(ServerCubeResource+"."+ServerTestResource, "nic.0.lan", LanResource+"."+LanTestResource, "id"),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.dhcp", "true"),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall_active", "true"),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.protocol", "TCP"),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.name", ServerTestResource),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.port_range_start", "22"),
			//		resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.port_range_end", "22"),
			//	),
			//},
		},
	})
}

func TestAccCubeServerResolveImageName(t *testing.T) {
	var server ionoscloud.Server

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCubeServerDestroyCheck,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCubeServerResolveImageName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(ServerCubeResource+"."+ServerTestResource, &server),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "name", ServerTestResource),
					//resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "cores", "1"),
					//resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "ram", "1024"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "availability_zone", "ZONE_1"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "cpu_family", "INTEL_SKYLAKE"),
					utils.TestImageNotNull(ServerCubeResource, "boot_image"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "image_password", "pass123456"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "volume.0.name", ServerTestResource),
					//resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "volume.0.size", "5"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "volume.0.disk_type", "SSD Standard"),
					resource.TestCheckResourceAttrPair(ServerCubeResource+"."+ServerTestResource, "nic.0.lan", LanResource+"."+LanTestResource, "id"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.dhcp", "true"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall_active", "true"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.protocol", "TCP"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.name", ServerTestResource),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.port_range_start", "22"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.port_range_end", "22"),
				),
			},
		},
	})
}

func TestAccCubeServerWithSnapshot(t *testing.T) {
	var server ionoscloud.Server

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCubeServerDestroyCheck,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccCheckCubeServerWithSnapshot),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(ServerCubeResource+"."+ServerTestResource, &server),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "name", ServerTestResource),
					//resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "cores", "1"),
					//resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "ram", "1024"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "availability_zone", "ZONE_1"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "cpu_family", "INTEL_SKYLAKE"),
					utils.TestImageNotNull(ServerCubeResource, "boot_image"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "volume.0.name", ServerTestResource),
					//resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "volume.0.size", "5"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "volume.0.disk_type", "SSD Standard"),
					resource.TestCheckResourceAttrPair(ServerCubeResource+"."+ServerTestResource, "nic.0.lan", LanResource+"."+LanTestResource, "id"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.dhcp", "true"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall_active", "true"),
				),
			},
		},
	})
}

func TestAccServerCubeServer(t *testing.T) {

	var server ionoscloud.Server

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCubeServerDestroyCheck,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCubeServerAndServersDataSource,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(ServerCubeResource+"."+ServerTestResource, &server),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "name", ServerTestResource),
					//resource.TestCheckResourceAttrPair(ServerCubeResource+"."+ServerTestResource, "cores", "data.ionoscloud_template."+ServerTestResource, "cores"),
					//resource.TestCheckResourceAttrPair(ServerCubeResource+"."+ServerTestResource, "ram", "data.ionoscloud_template."+ServerTestResource, "ram"),
					resource.TestCheckResourceAttrPair(ServerCubeResource+"."+ServerTestResource, "template_uuid", "data.ionoscloud_template."+ServerTestResource, "id"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "availability_zone", "ZONE_2"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "type", "CUBE"),
					utils.TestImageNotNull("ionoscloud_server", "boot_image"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "volume.0.name", ServerTestResource),
					//resource.TestCheckResourceAttrPair(ServerCubeResource+"."+ServerTestResource, "volume.0.size", "data.ionoscloud_template."+ServerTestResource, "storage_size"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "volume.0.disk_type", "DAS"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "volume.0.licence_type", "LINUX"),
					resource.TestCheckResourceAttrPair(ServerCubeResource+"."+ServerTestResource, "nic.0.lan", "ionoscloud_lan.webserver_lan", "id"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.name", ServerTestResource),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.dhcp", "true"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall_active", "true"),
					resource.TestCheckResourceAttr(DataSource+"."+ServersDataSource+"."+ServerDataSourceByName, "servers.#", "1"),
				),
			},
		},
	})
}

func TestAccCubeServerWithICMP(t *testing.T) {
	var server ionoscloud.Server

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCubeServerDestroyCheck,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCubeServerNoFirewall,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(ServerCubeResource+"."+ServerTestResource, &server),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "name", ServerTestResource),
					//resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "cores", "1"),
					//resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "ram", "1024"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "availability_zone", "ZONE_1"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "cpu_family", "INTEL_SKYLAKE"),
					utils.TestImageNotNull(ServerCubeResource, "boot_image"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "image_password", "K3tTj8G14a3EgKyNeeiY"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "volume.0.name", "system"),
					//resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "volume.0.size", "5"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "volume.0.disk_type", "HDD"),
					resource.TestCheckResourceAttrPair(ServerCubeResource+"."+ServerTestResource, "nic.0.lan", LanResource+"."+LanTestResource, "id"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.name", "system"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.dhcp", "true"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall_active", "false"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.icmp_type", "10"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.icmp_code", "1"),
				),
			},
			{
				Config: testAccCheckCubeServerICMP,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(ServerCubeResource+"."+ServerTestResource, &server),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.dhcp", "true"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall_active", "true"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.protocol", "ICMP"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.name", ServerTestResource),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.icmp_type", "12"),
					resource.TestCheckResourceAttr(ServerCubeResource+"."+ServerTestResource, "nic.0.firewall.0.icmp_code", "0"),
				),
			},
		},
	})
}

func testAccCheckCubeServerDestroyCheck(s *terraform.State) error {
	client := testAccProvider.Meta().(SdkBundle).CloudApiClient

	ctx, cancel := context.WithTimeout(context.Background(), *resourceDefaultTimeouts.Default)

	if cancel != nil {
		defer cancel()
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != ServerCubeResource {
			continue
		}

		dcId := rs.Primary.Attributes["datacenter_id"]

		_, apiResponse, err := client.ServersApi.DatacentersServersFindById(ctx, dcId, rs.Primary.ID).Execute()
		logApiRequestTime(apiResponse)

		if err != nil {
			if !httpNotFound(apiResponse) {
				return fmt.Errorf("unable to fetch server %s: %s", rs.Primary.ID, err)
			}
		} else {
			return fmt.Errorf("server still exists %s", rs.Primary.ID)

		}
	}

	return nil
}

func testAccCheckCubeServerExists(n string, server *ionoscloud.Server) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(SdkBundle).CloudApiClient

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("testAccCheckServerExists: Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		ctx, cancel := context.WithTimeout(context.Background(), *resourceDefaultTimeouts.Default)

		if cancel != nil {
			defer cancel()
		}

		foundServer, apiResponse, err := client.ServersApi.DatacentersServersFindById(ctx, rs.Primary.Attributes["datacenter_id"], rs.Primary.ID).Execute()
		logApiRequestTime(apiResponse)

		if err != nil {
			return fmt.Errorf("error occured while fetching Server: %s", rs.Primary.ID)
		}
		if *foundServer.Id != rs.Primary.ID {
			return fmt.Errorf("record not found")
		}

		server = &foundServer

		return nil
	}
}

const testAccCheckCubeServerConfigUpdate = `
data "ionoscloud_template" ` + ServerTestResource + ` {
    name = "CUBES XS"
    cores = 1
    ram   = 1024
    storage_size = 30
}

resource ` + DatacenterResource + ` ` + DatacenterTestResource + ` {
	name       = "server-test"
	location = "de/txl"
}

resource "ionoscloud_ipblock" "webserver_ipblock" {
  location = ` + DatacenterResource + `.` + DatacenterTestResource + `.location
  size = 4
  name = "webserver_ipblock"
}

resource "ionoscloud_ipblock" "webserver_ipblock_update" {
  location = ` + DatacenterResource + `.` + DatacenterTestResource + `.location
  size = 4
  name = "webserver_ipblock"
}
resource ` + LanResource + ` ` + LanTestResource + ` {
  datacenter_id = ` + DatacenterResource + `.` + DatacenterTestResource + `.id
  public = true
  name = "public"
}
resource ` + ServerCubeResource + ` ` + ServerTestResource + ` {
  name = "` + UpdatedResources + `"
  datacenter_id = ` + DatacenterResource + `.` + DatacenterTestResource + `.id
  availability_zone = "ZONE_1"
  cpu_family = "INTEL_SKYLAKE"
  image_name ="ubuntu:latest"
  image_password = "K3tTj8G14a3EgKyNeeiYsasad"
  template_uuid     = data.ionoscloud_template.` + ServerTestResource + `.id

  volume {
    name            = "` + ServerTestResource + `"
    licence_type    = "LINUX"
    disk_type = "DAS"
	}
  nic {
    lan = ` + LanResource + `.` + LanTestResource + `.id
    name = "` + UpdatedResources + `"
    dhcp = false
    firewall_active = false
    ips            = [ ionoscloud_ipblock.webserver_ipblock_update.ips[0], ionoscloud_ipblock.webserver_ipblock_update.ips[1] ]
    firewall {
      protocol = "TCP"
      name = "` + UpdatedResources + `"
      port_range_start = 21
      port_range_end = 23
	  source_mac = "00:0a:95:9d:68:18"
	  source_ip = ionoscloud_ipblock.webserver_ipblock_update.ips[2]
	  target_ip = ionoscloud_ipblock.webserver_ipblock_update.ips[3]
	  type = "INGRESS"
    }
  }
}`

const testAccDataSourceCubeServerMatchId = testAccCheckCubeServerConfigBasic + `
data ` + ServerCubeResource + ` ` + ServerDataSourceById + ` {
  datacenter_id = ` + DatacenterResource + `.` + DatacenterTestResource + `.id
  id			= ` + ServerCubeResource + `.` + ServerTestResource + `.id
}
`

const testAccDataSourceCubeServerMatchName = testAccCheckCubeServerConfigBasic + `
data ` + ServerCubeResource + ` ` + ServerDataSourceByName + ` {
  datacenter_id = ` + DatacenterResource + `.` + DatacenterTestResource + `.id
  name			= "` + ServerTestResource + `"
}
`
const testAccDataSourceCubeServerWrongNameError = testAccCheckCubeServerConfigBasic + `
data ` + ServerCubeResource + ` ` + ServerDataSourceByName + ` {
  datacenter_id = ` + DatacenterResource + `.` + DatacenterTestResource + `.id
  name			= "wrong_name"
}
`

const testAccCheckCubeServerConfigBootCdromNoImage = `
data "ionoscloud_template" ` + ServerTestResource + ` {
    name = "CUBES XS"
    cores = 1
    ram   = 1024
    storage_size = 30
}

resource ` + DatacenterResource + ` ` + DatacenterTestResource + ` {
	name       = "server-test"
	location   = "de/txl"
}
resource ` + LanResource + ` ` + LanTestResource + ` {
  datacenter_id = ` + DatacenterResource + `.` + DatacenterTestResource + `.id
  public = true
  name = "public"
}
resource ` + ServerCubeResource + ` ` + ServerTestResource + ` {
  template_uuid     = data.ionoscloud_template.` + ServerTestResource + `.id
  name = "` + ServerTestResource + `"
  datacenter_id = ` + DatacenterResource + `.` + DatacenterTestResource + `.id
  availability_zone = "ZONE_1"
  cpu_family = "INTEL_SKYLAKE"
  boot_cdrom = "` + bootCdromImageIdCube + `" 
  volume {
    name = "` + ServerTestResource + `"
    disk_type = "SSD Standard"
	licence_type = "OTHER"
  }
  nic {
    lan = ` + LanResource + `.` + LanTestResource + `.id
    dhcp = true
    firewall_active = true 
	firewall {
      protocol = "TCP"
      name = "` + ServerTestResource + `"
      port_range_start = 22
      port_range_end = 22
    }
  }
}`

const testAccCheckCubeServerResolveImageName = `
resource ` + DatacenterResource + ` ` + DatacenterTestResource + ` {
  name        = "test_server"
  location    = "de/fra"
  description = "Test datacenter done by TF"
}
resource ` + LanResource + ` ` + LanTestResource + ` {
  datacenter_id = ` + DatacenterResource + `.` + DatacenterTestResource + `.id
  public        = true
}
resource ` + ServerCubeResource + ` ` + ServerTestResource + ` {
  name              = "` + ServerTestResource + `"
  datacenter_id     = ` + DatacenterResource + `.` + DatacenterTestResource + `.id
  availability_zone = "ZONE_1"
  cpu_family        = "INTEL_SKYLAKE" 
  image_name        = "ubuntu:latest"
  image_password    = "pass123456"
  volume {
    name           = "` + ServerTestResource + `"
    disk_type      = "DAS"
  }
  nic {
    lan             = ` + LanResource + `.` + LanTestResource + `.id
    dhcp            = true
    firewall_active = true
    firewall {
      protocol         = "TCP"
      name             = "` + ServerTestResource + `"
      port_range_start = 22
      port_range_end   = 22
    }
  }
}`

const testAccCheckCubeServerWithSnapshot = `
resource ` + DatacenterResource + ` ` + DatacenterTestResource + ` {
	name       = "volume-test"
	location   = "de/fra"
}
resource ` + LanResource + ` ` + LanTestResource + ` {
  datacenter_id = ` + DatacenterResource + `.` + DatacenterTestResource + `.id
  public = true
  name = "public"
}
resource ` + ServerCubeResource + ` "webserver" {
  name = "webserver"
  datacenter_id = ` + DatacenterResource + `.` + DatacenterTestResource + `.id
  availability_zone = "ZONE_1"
  cpu_family = "INTEL_SKYLAKE"
	image_name = "ubuntu:latest"
	image_password = "K3tTj8G14a3EgKyNeeiY"
  volume {
    name = "system"
    disk_type = "SSD Standard"
  }
  nic {
    lan = ` + LanResource + `.` + LanTestResource + `.id
    dhcp = true
    firewall_active = true
  }
}
resource ` + SnapshotResource + ` "test_snapshot" {
  datacenter_id = ` + DatacenterResource + `.` + DatacenterTestResource + `.id
  volume_id = ` + ServerCubeResource + `.webserver.boot_volume
  name = "terraform_snapshot"
}
resource ` + ServerCubeResource + ` ` + ServerTestResource + ` {
  depends_on = [` + SnapshotResource + `.test_snapshot]
  name = "` + ServerTestResource + `"
  datacenter_id = ` + DatacenterResource + `.` + DatacenterTestResource + `.id
  availability_zone = "ZONE_1"
  cpu_family = "INTEL_SKYLAKE"
  image_name = "terraform_snapshot"
  volume {
    name = "` + ServerTestResource + `"
    disk_type = "SSD Standard"
  }
  nic {
    lan = ` + LanResource + `.` + LanTestResource + `.id
    dhcp = true
    firewall_active = true
  }
}
`

const testAccCheckCubeServerAndServersDataSource = `
data "ionoscloud_template" ` + ServerTestResource + ` {
    name = "CUBES XS"
    cores = 1
    ram   = 1024
    storage_size = 30
}

resource ` + DatacenterResource + " " + DatacenterTestResource + `{
	name       = "volume-test"
	location   = "de/txl"
}

resource "ionoscloud_lan" "webserver_lan" {
  datacenter_id = ` + DatacenterResource + `.` + DatacenterTestResource + `.id
  public = true
  name = "public"
}

resource "ionoscloud_server" ` + ServerTestResource + ` {
  name              = "` + ServerTestResource + `"
  availability_zone = "ZONE_2"
  image_name        = "ubuntu:latest"
  type              = "CUBE"
  template_uuid     = data.ionoscloud_template.` + ServerTestResource + `.id
  image_password = "K3tTj8G14a3EgKyNeeiY"  
  datacenter_id = ` + DatacenterResource + `.` + DatacenterTestResource + `.id
  volume {
    name            = "` + ServerTestResource + `"
    licence_type    = "LINUX" 
    disk_type = "DAS"
	}
  nic {
    lan             = ionoscloud_lan.webserver_lan.id
    name            = "` + ServerTestResource + `"
    dhcp            = true
    firewall_active = true
  }
}
data ` + ServersDataSource + ` ` + ServerDataSourceByName + ` {
 depends_on = [` + ServerCubeResource + `.` + ServerTestResource + `]
  datacenter_id = ` + DatacenterResource + `.` + DatacenterTestResource + `.id
  filter {
   name = "type"
   value = "CUBE" 
  }
}`

const testAccCheckCubeServerNoFirewall = `
resource ` + DatacenterResource + ` ` + DatacenterTestResource + ` {
	name       = "server-test"
	location = "de/fra"
}
resource ` + LanResource + ` ` + LanTestResource + ` {
  datacenter_id = ` + DatacenterResource + `.` + DatacenterTestResource + `.id
  public = true
  name = "public"
}
resource ` + ServerCubeResource + ` ` + ServerTestResource + ` {
  name = "` + ServerTestResource + `"
  datacenter_id = ` + DatacenterResource + `.` + DatacenterTestResource + `.id
  cores = 1
  ram = 1024
  availability_zone = "ZONE_1"
  cpu_family = "INTEL_SKYLAKE"
  image_name ="ubuntu:latest"
  image_password = "K3tTj8G14a3EgKyNeeiY"
  volume {
    name = "system"
    size = 5
	disk_type = "HDD"
}
  nic {
    lan = ` + LanResource + `.` + LanTestResource + `.id
    name = "system"
    dhcp = true
    firewall_active = false
    firewall {
      protocol         = "ICMP"
      name             = "` + ServerTestResource + `"
      icmp_type        = "10"
      icmp_code        = "1"
	  }
  }
}`

const testAccCheckCubeServerICMP = `
resource ` + DatacenterResource + ` ` + DatacenterTestResource + ` {
	name       = "server-test"
	location = "de/fra"
}
resource ` + LanResource + ` ` + LanTestResource + ` {
  datacenter_id = ` + DatacenterResource + `.` + DatacenterTestResource + `.id
  public = true
  name = "public"
}
resource ` + ServerCubeResource + ` ` + ServerTestResource + ` {
  name = "` + ServerTestResource + `"
  datacenter_id = ` + DatacenterResource + `.` + DatacenterTestResource + `.id
  cores = 1
  ram = 1024
  availability_zone = "ZONE_1"
  cpu_family = "INTEL_SKYLAKE"
  image_name ="ubuntu:latest"
  image_password = "K3tTj8G14a3EgKyNeeiY"
  volume {
    name = "system"
    
	licence_type    = "LINUX"
    disk_type = "DAS"
}
  nic {
    lan             = ` + LanResource + `.` + LanTestResource + `.id
    name 			= "system"
    dhcp            = true
    firewall_active = true
    firewall {
      protocol         = "ICMP"
      name             = "` + ServerTestResource + `"
      icmp_type        = "12"
      icmp_code        = "0"
	  }
    }
}`

// INTEL_XEON
