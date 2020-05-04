#!/bin/bash
cd ..
go run service/group/service/service.go -cfg=cfg.yaml
go run service/group/api/api.go service/group/api/rest.go -cfg=cfg.yaml

