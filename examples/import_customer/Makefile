SHELL = /bin/bash

.PHONY: all
all: import apply

.PHONY: import
import:
	echo "required xsv: brew install xsv"
	@xsv select optional_id customers.csv \
	  | xsv table \
	  | tail -n +2 \
	  | xargs -I__ID__ echo "terraform import 'scalebase_customer.example[\"__ID__\"]' __ID__"; terraform import 'scalebase_customer.example["__ID__"]' __ID__ > /dev/null || true

.PHONY: apply
apply:
	terraform apply -auto-approve

.PHONY: plan
plan:
	terraform plan

.PHONY: destroy
destroy:
	terraform destroy -auto-approve
