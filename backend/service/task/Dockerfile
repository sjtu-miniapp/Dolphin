FROM scratch
COPY build/task-srv ./srv
COPY cfg.yaml .
CMD ["./srv", "-cfg=cfg.yaml"]