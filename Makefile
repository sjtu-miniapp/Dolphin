.ONESHELL:
.PHONY: help docker deploy transfer
ENV_FILE=.env
DOCKER_SQL=backend/database/sql/
help:

docker:
	@echo $DOCKER_PASSWORD | sudo docker login -u $DOCKER_USERNAME --password-stdin
	# build all the images
	@sudo docker build -t nihplod/mysql ${DOCKER_SQL}
	@sudo docker push nihplod/mysql
transfer:
	@sshpass -p ${SERVER1_PASSWORD} scp -P ${SERVER1_SSH} -r * ${SERVER1_USER}@${SERVER1_IP}:~/dolphin
	@sshpass -p ${SERVER2_PASSWORD} scp -P ${SERVER2_SSH} -r * ${SERVER2_USER}@${SERVER2_IP}:~/dolphin
	@sshpass -p ${SERVER3_PASSWORD} scp -P ${SERVER3_SSH} -r * ${SERVER3_USER}@${SERVER3_IP}:~/dolphin
deploy:
#	@servercnt=$$(sed -n /^SERVER/p ${ENV_FILE}| sed -n '$$p'| cut -b 7)
	@sshpass -p ${SERVER1_PASSWORD} ssh -P ${SERVER1_SSH} ${SERVER1_USER}@${SERVER1_IP} 'cd dolphin/scripts && ./setup.sh sql'