build:
	@go build -o terraform-provider-person
	@mv terraform-provider-person ~/.terraform.d/plugins/kiril.com/personprovider/person/1.0.0/linux_amd64 