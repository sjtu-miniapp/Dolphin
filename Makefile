.ONESHELL:
.PHONY: help build docker deploy transfer api
SHELL=/bin/bash
ENV_FILE=.env
DOCKER_SQL=backend/database/sql
SRV=backend/srv
GOOUT=backend
GOSCP=backend/scripts
API=backend/api
help:
build:
	@source ${GOSCP}/source.sh
	@SRVLIST=$$(find ${SRV}/* -maxdepth 0 -type d | sed  "s/^.*\///")
	@cd ${SRV}/
	@for srv in $${SRVLIST}; do \
  		cd $${srv}; \
#  		if [[ $${srv} == group ]]; then  \
#  		cd .. ;\
#  		continue ; \
#		fi ; \
  		go build -a -o ../../build/$${srv} cmd/server.go; \
  		cd .. ;\
  	done
docker:
	@echo $DOCKER_PASSWORD | sudo docker login -u $DOCKER_USERNAME --password-stdin
	# build all the images
	@sudo docker build -t nihplod/mysql ${DOCKER_SQL}
	@sudo docker push nihplod/mysql
transfer:
	@#sshpass -p ${SERVER1_PASSWORD} scp -o StrictHostKeyChecking=no -P ${SERVER1_SSH} -r * ${SERVER1_USER}@${SERVER1_IP}:~/dolphin
	@#echo "transfer to server1"
	@sshpass -p ${SERVER2_PASSWORD} scp -o StrictHostKeyChecking=no -P ${SERVER2_SSH} -r * ${SERVER2_USER}@${SERVER2_IP}:~/dolphin
	@echo "transfer to server2"
	@sshpass -p ${SERVER3_PASSWORD} scp -o StrictHostKeyChecking=no -P ${SERVER3_SSH} -r * ${SERVER3_USER}@${SERVER3_IP}:~/dolphin
	@echo "transfer to server3"
deploy:	build transfer
#	@servercnt=$$(sed -n /^SERVER/p ${ENV_FILE}| sed -n '$$p'| cut -b 7)
	@sshpass -p ${SERVER1_PASSWORD} ssh -o StrictHostKeyChecking=no -P ${SERVER1_SSH} ${SERVER1_USER}@${SERVER1_IP} 'cd dolphin/scripts && ./setup.sh server'

api:
	@version=$$(cat ${API}/VERSION)
#make api up=1
	@if [[ ! -z "${up}" ]]; then \
   		version='v'$$(expr $$(cut -b 2- ${API}/VERSION) + 1); \
   		echo $${version} > "${API}/VERSION";\
   		mkdir -p "${API}/$${version}";\
   	  fi;
	@SRVLIST=$$(find ${SRV}/* -maxdepth 0 -type d | sed  "s/^.*\///")
	@cd backend/
	@for srv in $${SRVLIST}; do \
		protoc --proto_path=api/$${version} --proto_path=thirdparty --go_out=plugins=grpc:srv/$${srv}/pb $${srv}.proto; \
		protoc --proto_path=api/$${version} --proto_path=thirdparty --grpc-gateway_out=logtostderr=true:srv/$${srv}/pb $${srv}.proto; \
		protoc --proto_path=api/$${version} --proto_path=thirdparty --swagger_out=json_names_for_fields=true:api/$${version} $${srv}.proto; \
   	done
	@#go run swagger/main.go swagger > api/api.swagger.json
clean:
	@rm -rf $(GOOUT)/build