package test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

func TestTerraformVirtualNetwork(t *testing.T) {
	t.Parallel()

	resourceGroupNameStrings := []string{"test01", "mns", "rg", "01"}

	vnetNameStrings := []string{"test01", "mns", "vnet", "01"}
	vnetName := strings.Join(vnetNameStrings, "-")
	vnetLocation := "westeurope"
	vnetAddressSpace := []string{"10.0.0.0/16"}

	// Root folder where terraform files should be (relative to the test folder)
	rootFolder := ".."
	// Relative path to terraform module being tested from the root folder
	terraformFolderRelativeToRoot := "calling-code"
	// Copy the terraform folder to a temp folder
	tempTestFolder := test_structure.CopyTerraformFolderToTemp(t, rootFolder, terraformFolderRelativeToRoot)

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: tempTestFolder,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"vnet_name_strings":  vnetNameStrings,
			"name_strings":       resourceGroupNameStrings,
			"name_separator":     "-",
			"vnet_address_space": vnetAddressSpace,
		},

		BackendConfig: map[string]interface{}{
			"bucket": "byt-terraform-state-bucket",
			"key":    "testing/test01.tfstate",
			"region": "eu-west-2",
		},
		// Variables to pass to our Terraform code using -var-file options
		// VarFiles: []string{"varfile.tfvars"},

		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	actualVnetName := terraform.Output(t, terraformOptions, "virtual_network_name")
	actualVnetLocation := terraform.Output(t, terraformOptions, "virtual_network_location")
	actualVnetAddressSpace := terraform.OutputList(t, terraformOptions, "virtual_network_address_space")

	// Verify we're getting back the outputs we expect
	assert.Equal(t, vnetName, actualVnetName)
	assert.Equal(t, vnetLocation, actualVnetLocation)
	assert.Equal(t, vnetAddressSpace, actualVnetAddressSpace)
}
func TestTerraformVirtualNetworkTags(t *testing.T) {
	t.Parallel()

	resourceGroupNameStrings := []string{"test02", "mns", "rg", "01"}

	vnetNameStrings := []string{"test02", "mns", "vnet", "01"}
	vnetName := strings.Join(vnetNameStrings, "-")
	vnetLocation := "westeurope"
	vnetAddressSpace := []string{"10.0.0.0/16"}
	vnetDNSServers := []string{"10.0.1.1", "10.0.2.1"}
	vnetTags := map[string]string{"environment": "development", "location": "westeurope"}

	// Root folder where terraform files should be (relative to the test folder)
	rootFolder := ".."
	// Relative path to terraform module being tested from the root folder
	terraformFolderRelativeToRoot := "calling-code"
	// Copy the terraform folder to a temp folder
	tempTestFolder := test_structure.CopyTerraformFolderToTemp(t, rootFolder, terraformFolderRelativeToRoot)

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: tempTestFolder,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"vnet_name_strings":  vnetNameStrings,
			"name_strings":       resourceGroupNameStrings,
			"name_separator":     "-",
			"vnet_address_space": vnetAddressSpace,
			"vnet_dns_servers":   vnetDNSServers,
			"tags":               vnetTags,
		},

		BackendConfig: map[string]interface{}{
			"bucket": "byt-terraform-state-bucket",
			"key":    "testing/test02.tfstate",
			"region": "eu-west-2",
		},
		// Variables to pass to our Terraform code using -var-file options
		// VarFiles: []string{"varfile.tfvars"},

		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	actualVnetName := terraform.Output(t, terraformOptions, "virtual_network_name")
	actualVnetLocation := terraform.Output(t, terraformOptions, "virtual_network_location")
	actualVnetAddressSpace := terraform.OutputList(t, terraformOptions, "virtual_network_address_space")
	actualVnetDNSServers := terraform.OutputList(t, terraformOptions, "virtual_network_dns_servers")
	actualVnetTags := terraform.OutputMap(t, terraformOptions, "virtual_network_tags")

	// Verify we're getting back the outputs we expect
	assert.Equal(t, vnetName, actualVnetName)
	assert.Equal(t, vnetLocation, actualVnetLocation)
	assert.Equal(t, vnetAddressSpace, actualVnetAddressSpace)
	assert.Equal(t, vnetDNSServers, actualVnetDNSServers)
	assert.Equal(t, vnetTags, actualVnetTags)
}
