package ionoscloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceShareMatchFields(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckShareConfigBasic,
			},
			{
				Config: testAccDataSourceShareConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(shareResourceFullName, "id"),
					resource.TestCheckResourceAttrPair(shareResourceFullName, "id", DataSource+"."+ShareResource+"."+sourceShareName, "id"),
					resource.TestCheckResourceAttrPair(shareResourceFullName, "edit_privilege",
						DataSource+"."+ShareResource+"."+sourceShareName, "edit_privilege"),
					resource.TestCheckResourceAttrPair(shareResourceFullName, "share_privilege",
						DataSource+"."+ShareResource+"."+sourceShareName, "share_privilege"),
					resource.TestCheckResourceAttr(DataSource+"."+ShareResource+"."+sourceShareName, "edit_privilege", "true"),
				),
			},
		},
	})
}

var testAccDataSourceShareConfigBasic = testAccCheckShareConfigBasic + `
data ` + ShareResource + " " + sourceShareName + `{
  group_id    = "${ionoscloud_group.group.id}"
  resource_id = "${ionoscloud_datacenter.foobar.id}"
  id		  = ` + shareResourceFullName + `.id
}
`
