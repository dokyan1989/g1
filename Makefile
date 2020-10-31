PROTOC_VERSION = 3.15.6
PROTOC_LINUX_ZIP = protoc-$(PROTOC_VERSION)-linux-x86_64.zip
MYSQL_URL = mysql://root:123456@tcp(127.0.0.1:3307)/service1?charset=utf8mb4&parseTime=true
SEED_PATH = app/service1/migrations/seed

install_pb_tools:
	curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/$(PROTOC_LINUX_ZIP)
	unzip -o $(PROTOC_LINUX_ZIP) -d /usr/local bin/protoc
	unzip -o $(PROTOC_LINUX_ZIP) -d /usr/local 'include/*'
	rm -f $(PROTOC_LINUX_ZIP)

gen_pb:
	rm $(path)/pb/*.go && \
	protoc --proto_path=$(path)/pb --go_out=. --go-grpc_out=. --grpc-gateway_out=. $(path)/pb/*.proto

clean_pb:
	rm $(path)/pb/*.go

create_mysql_container:
	sudo docker run --name g1mysql -e MYSQL_ROOT_PASSWORD=123456 -d -p 3307:3306 mysql

start_mysql_container:
	sudo docker container start g1mysql

migrate_file:
	migrate create -ext sql -dir $(path)/migrations $(name)

migrate_up:
	migrate -source file://$(path)/migrations -database '$(MYSQL_URL)' up

migrate_down:
	migrate -source file://$(path)/migrations -database '$(MYSQL_URL)' down $(num)

migrate_force:
	migrate -source file://$(path)/migrations -database '$(MYSQL_URL)' force 20210325104212

dump_data:
	sudo docker exec g1mysql mysqldump -uroot -p123456 service1 > service1.sql

restore_data:
	sudo docker exec -i g1mysql sh -c 'exec mysql -uroot -p123456 service1' < $(SEED_PATH)/load_employees.dump && \
	sudo docker exec -i g1mysql sh -c 'exec mysql -uroot -p123456 service1' < $(SEED_PATH)/load_departments.dump && \
	sudo docker exec -i g1mysql sh -c 'exec mysql -uroot -p123456 service1' < $(SEED_PATH)/load_dept_emp.dump && \
	sudo docker exec -i g1mysql sh -c 'exec mysql -uroot -p123456 service1' < $(SEED_PATH)/load_dept_manager.dump && \
	sudo docker exec -i g1mysql sh -c 'exec mysql -uroot -p123456 service1' < $(SEED_PATH)/load_titles.dump && \
	sudo docker exec -i g1mysql sh -c 'exec mysql -uroot -p123456 service1' < $(SEED_PATH)/load_salaries1.dump && \
	sudo docker exec -i g1mysql sh -c 'exec mysql -uroot -p123456 service1' < $(SEED_PATH)/load_salaries2.dump && \
	sudo docker exec -i g1mysql sh -c 'exec mysql -uroot -p123456 service1' < $(SEED_PATH)/load_salaries3.dump

server:
	go run app/$(name)/cmd/server/main.go

client: 
	go run app/$(name)/cmd/client/main.go

.PHONY:
	install_pb_tools gen_pb clean_pb create_mysql_container start_mysql_container	\
	migration_file migrate_up migrate_down dump_data server client