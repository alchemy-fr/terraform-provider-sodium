#!/bin/bash

set -e

rm -f ./src/terraform-provider-sodium

#docker run -e GOCACHE=/usr/app/.cache --rm -ti -v ./src:/usr/app sodium-gobuilder bash
docker run -e GOCACHE=/usr/app/.cache --rm -v ./src:/usr/app sodium-gobuilder go build -o terraform-provider-sodium

target=~/.terraform.d/plugins/alchemy.fr/alchemy/sodium/1.0.0/linux_amd64
mkdir -p $target
cp ./src/terraform-provider-sodium $target/

rm .terraform.lock.hcl
rm -r .terraform

terraform init
TF_LOG=stderr terraform apply
