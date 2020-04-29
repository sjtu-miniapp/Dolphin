.ONESHELL:
.PHONY: help build docker deploy transfer
SHELL=/bin/bash
ENV_FILE=.env
DOCKER_SQL=backend/database/sql
SRV=backend/srv
GOOUT=backend
GOSCP=backend/scripts
help:
build:
	@source ${GOSCP}/source.sh
	@SRVLIST=$$(find ${SRV}/* -maxdepth 0 -type d | sed  "s/^.*\///")
	@cd ${SRV}/
	@for srv in $${SRVLIST}; do \
  		cd $${srv}; \
  		if [[ $${srv} == group ]]; then  \
  		cd .. ;\
  		continue ; \
		fi ; \
#  		go build -a -o $(GOOUT)/build/$${srv} ${SRV}/$${srv}/cmd/server.go; \
  		go build -a -o ../../build/$${srv} cmd/server.go; \
  		cd .. ;\
  	done
docker:
	@echo $DOCKER_PASSWORD | sudo docker login -u $DOCKER_USERNAME --password-stdin
	# build all the images
	@sudo docker build -t nihplod/mysql ${DOCKER_SQL}
	@sudo docker push nihplod/mysql
transfer:
	@sshpass -p ${SERVER1_PASSWORD} scp -o StrictHostKeyChecking=no -P ${SERVER1_SSH} -r * ${SERVER1_USER}@${SERVER1_IP}:~/dolphin
	@echo "transfer to server1"
	@sshpass -p ${SERVER2_PASSWORD} scp -o StrictHostKeyChecking=no -P ${SERVER2_SSH} -r * ${SERVER2_USER}@${SERVER2_IP}:~/dolphin
	@echo "transfer to server2"
	@sshpass -p ${SERVER3_PASSWORD} scp -o StrictHostKeyChecking=no -P ${SERVER3_SSH} -r * ${SERVER3_USER}@${SERVER3_IP}:~/dolphin
	@echo "transfer to server3"
deploy:
#	@servercnt=$$(sed -n /^SERVER/p ${ENV_FILE}| sed -n '$$p'| cut -b 7)
	@sshpass -p ${SERVER1_PASSWORD} ssh -o StrictHostKeyChecking=no -P ${SERVER1_SSH} ${SERVER1_USER}@${SERVER1_IP} 'cd dolphin/scripts && ./setup.sh sql'

clean:
	@rm -rf $(GOOUT)/build