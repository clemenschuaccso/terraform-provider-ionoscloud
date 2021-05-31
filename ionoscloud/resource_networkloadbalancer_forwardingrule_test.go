package ionoscloud

import (
	"context"
	"fmt"
	ionoscloud "github.com/ionos-cloud/sdk-go/v6"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccNetworkLoadBalancerForwardingRule_Basic(t *testing.T) {
	var networkLoadBalancerForwardingRule ionoscloud.NetworkLoadBalancerForwardingRule
	networkLoadBalancerForwardingRuleName := "networkLoadBalancerForwardingRule"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkLoadBalancerForwardingRuleDestroyCheck,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccCheckNetworkLoadBalancerForwardingRuleConfig_basic, networkLoadBalancerForwardingRuleName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkLoadBalancerForwardingRuleExists("ionoscloud_networkloadbalancer_forwardingrule.forwarding_rule", &networkLoadBalancerForwardingRule),
					resource.TestCheckResourceAttr("ionoscloud_networkloadbalancer_forwardingrule.forwarding_rule", "name", networkLoadBalancerForwardingRuleName),
				),
			},
			{
				Config: testAccCheckNetworkLoadBalancerForwardingRuleConfig_update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ionoscloud_networkloadbalancer_forwardingrule.forwarding_rule", "name", "updated"),
				),
			},
		},
	})
}

func testAccCheckNetworkLoadBalancerForwardingRuleDestroyCheck(s *terraform.State) error {
	client := testAccProvider.Meta().(SdkBundle).Client
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ionoscloud_datacenter" {
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), *resourceDefaultTimeouts.Delete)

		if cancel != nil {
			defer cancel()
		}

		_, apiResponse, err := client.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesDelete(ctx, rs.Primary.Attributes["datacenter_id"], rs.Primary.Attributes["networkloadbalancer_id"], rs.Primary.ID).Execute()

		if _, ok := err.(ionoscloud.GenericOpenAPIError); ok {
			if apiResponse.Response.StatusCode != 404 {
				return fmt.Errorf("Network loadbalancer forwarding rule still exists %s %s", rs.Primary.ID, string(apiResponse.Payload))
			}
		} else {
			return fmt.Errorf("unable to fetch network loadbalancer forwarding rule %s %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckNetworkLoadBalancerForwardingRuleExists(n string, networkLoadBalancerForwardingRule *ionoscloud.NetworkLoadBalancerForwardingRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(SdkBundle).Client
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("testAccCheckNetworkLoadBalancerForwardingRuleExists: Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		ctx, cancel := context.WithTimeout(context.Background(), *resourceDefaultTimeouts.Delete)

		if cancel != nil {
			defer cancel()
		}

		foundNetworkLoadBalancerForwardingRule, _, err := client.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesFindByForwardingRuleId(ctx, rs.Primary.Attributes["datacenter_id"], rs.Primary.Attributes["networkloadbalancer_id"], rs.Primary.ID).Execute()

		if err != nil {
			return fmt.Errorf("error occured while fetching NetworkLoadBalancerForwardingRule: %s", rs.Primary.ID)
		}
		if *foundNetworkLoadBalancerForwardingRule.Id != rs.Primary.ID {
			return fmt.Errorf("Record not found")
		}

		networkLoadBalancerForwardingRule = &foundNetworkLoadBalancerForwardingRule

		return nil
	}
}

const testAccCheckNetworkLoadBalancerForwardingRuleConfig_basic = `
resource "ionoscloud_datacenter" "nlb_fr_datacenter" {
  name              = "test_nlb_fr"
  location          = "gb/lhr"
  description       = "datacenter for hosting "
}

resource "ionoscloud_lan" "nlb_fr_lan_1" {
  datacenter_id = ionoscloud_datacenter.nlb_fr_datacenter.id
  public        = false
  name          = "test_nlb_fr_lan_1"
}

resource "ionoscloud_lan" "nlb_fr_lan_2" {
  datacenter_id = ionoscloud_datacenter.nlb_fr_datacenter.id
  public        = false
  name          = "test_nlb_fr_lan_1"
}


resource "ionoscloud_networkloadbalancer" "test_nbl_fr" {
  datacenter_id = ionoscloud_datacenter.nlb_fr_datacenter.id
  name          = "test_nlb_fr"
  listener_lan  = ionoscloud_lan.nlb_fr_lan_1.id
  target_lan    = ionoscloud_lan.nlb_fr_lan_2.id
  ips           = ["10.12.118.224"]
  lb_private_ips = ["10.13.72.225/24"]
}

resource "ionoscloud_networkloadbalancer_forwardingrule" "forwarding_rule" {
 datacenter_id = ionoscloud_datacenter.nlb_fr_datacenter.id
 networkloadbalancer_id = ionoscloud_networkloadbalancer.test_nbl_fr.id
 name = "%s"
 algorithm = "SOURCE_IP"
 protocol = "TCP"
 listener_ip = "10.12.118.224"
 listener_port = "8081"
 targets {
   ip = "22.231.2.2"
   port = "8080"
   weight = "123"
   health_check {
     check = true
     check_interval = 1000
   }
 }
}
`

const testAccCheckNetworkLoadBalancerForwardingRuleConfig_update = `
resource "ionoscloud_datacenter" "nlb_fr_datacenter" {
  name              = "test_nlb_fr"
  location          = "gb/lhr"
  description       = "datacenter for hosting "
}

resource "ionoscloud_lan" "nlb_fr_lan_1" {
  datacenter_id = ionoscloud_datacenter.nlb_fr_datacenter.id
  public        = false
  name          = "test_nlb_fr_lan_1"
}

resource "ionoscloud_lan" "nlb_fr_lan_2" {
  datacenter_id = ionoscloud_datacenter.nlb_fr_datacenter.id
  public        = false
  name          = "test_nlb_fr_lan_1"
}


resource "ionoscloud_networkloadbalancer" "test_nbl_fr" {
  datacenter_id = ionoscloud_datacenter.nlb_fr_datacenter.id
  name          = "test_nlb_fr"
  listener_lan  = ionoscloud_lan.nlb_fr_lan_1.id
  target_lan    = ionoscloud_lan.nlb_fr_lan_2.id
  ips           = ["10.12.118.224"]
  lb_private_ips = ["10.13.72.225/24"]
}

resource "ionoscloud_networkloadbalancer_forwardingrule" "forwarding_rule" {
 datacenter_id = ionoscloud_datacenter.nlb_fr_datacenter.id
 networkloadbalancer_id = ionoscloud_networkloadbalancer.test_nbl_fr.id
 name = "updated"
 algorithm = "SOURCE_IP"
 protocol = "TCP"
 listener_ip = "10.12.118.224"
 listener_port = "8081"
 targets {
   ip = "22.231.2.2"
   port = "8080"
   weight = "123"
   health_check {
     check = true
     check_interval = 1000
   }
 }
}`
