build:
	@go build -o terraform-provider-person
	@mkdir -p ~/.terraform.d/plugins/kiril.com/personprovider/person/1.0.0/linux_amd64 
	@mv terraform-provider-person ~/.terraform.d/plugins/kiril.com/personprovider/person/1.0.0/linux_amd64 
