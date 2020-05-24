FROM scratch
COPY build/auth-srv ./srv
COPY cfg.yaml .
CMD ["./srv", "-cfg=cfg.yaml"]
# FIXME
#FROM alpine
#COPY build/auth-srv ./srv
#COPY cfg.yaml .
#RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
#ADD nginx/1_isalb.xyz_bundle.crt /usr/local/share/ca-certificates/mycert.crt
#RUN chmod 777 /usr/local/share/ca-certificates/mycert.crt && update-ca-certificates
#RUN ["chmod", "777", "./srv"]
#RUN ["chmod", "777", "./cfg.yaml"]
#CMD ["./srv", "-cfg=cfg.yaml"]
