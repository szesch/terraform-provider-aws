package account_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/account"
	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

func TestAccAccountPrimary_basic(t *testing.T) {
	resourceName := "aws_account_primary_contact.test"
	rName1 := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rName2 := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t); testAccPreCheck(t) },
		ErrorCheck:               acctest.ErrorCheck(t, account.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckPrimaryContactDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPrimaryConfig_basic(rName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "account_id", ""),
					resource.TestCheckResourceAttr(resourceName, "address_line_1", "123 Any Street"),
					resource.TestCheckResourceAttr(resourceName, "city", "Seattle"),
					resource.TestCheckResourceAttr(resourceName, "company_name", "Example Corp, Inc."),
					resource.TestCheckResourceAttr(resourceName, "country_code", "US"),
					resource.TestCheckResourceAttr(resourceName, "district_or_county", "King"),
					resource.TestCheckResourceAttr(resourceName, "full_name", rName1),
					resource.TestCheckResourceAttr(resourceName, "phone_number", "+64211111111"),
					resource.TestCheckResourceAttr(resourceName, "postal_code", "98101"),
					resource.TestCheckResourceAttr(resourceName, "state_or_region", "WA"),
					resource.TestCheckResourceAttr(resourceName, "website_url", "https://www.examplecorp.com"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccPrimaryConfig_basic(rName2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "account_id", ""),
					resource.TestCheckResourceAttr(resourceName, "address_line_1", "123 Any Street"),
					resource.TestCheckResourceAttr(resourceName, "city", "Seattle"),
					resource.TestCheckResourceAttr(resourceName, "company_name", "Example Corp, Inc."),
					resource.TestCheckResourceAttr(resourceName, "country_code", "US"),
					resource.TestCheckResourceAttr(resourceName, "district_or_county", "King"),
					resource.TestCheckResourceAttr(resourceName, "full_name", rName2),
					resource.TestCheckResourceAttr(resourceName, "phone_number", "+64211111111"),
					resource.TestCheckResourceAttr(resourceName, "postal_code", "98101"),
					resource.TestCheckResourceAttr(resourceName, "state_or_region", "WA"),
					resource.TestCheckResourceAttr(resourceName, "website_url", "https://www.examplecorp.com"),
				),
			},
		},
	})
}

func testAccCheckPrimaryContactDestroy(s *terraform.State) error {
	return nil
}

func testAccPrimaryConfig_basic(name string) string {
	return fmt.Sprintf(`
resource "aws_account_primary_contact" "test" {
  address_line_1     = "123 Any Street"
  city               = "Seattle"
  company_name       = "Example Corp, Inc."
  country_code       = "US"
  district_or_county = "King"
  full_name          = %[1]q
  phone_number       = "+64211111111"
  postal_code        = "98101"
  state_or_region    = "WA"
  website_url        = "https://www.examplecorp.com"
}
`, name)
}
