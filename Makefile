
phony: list

list:
	@echo "Commands:"
	@echo "init-db - initialize local databases"

init-db:
	docker-compose exec postgres bash -c "createdb -U usr auth; createdb -U usr betting; createdb -U usr kyc; createdb -U usr promotions; createdb -U usr referral;"