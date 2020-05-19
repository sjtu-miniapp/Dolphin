FROM scratch
COPY build/auth-srv ./srv
COPY cfg.yaml .
CMD ["./srv", "-cfg=cfg.yaml"]