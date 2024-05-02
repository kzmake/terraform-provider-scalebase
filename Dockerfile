FROM hashicorp/terraform:latest

WORKDIR /root/work
COPY examples/minimal/** /root/work

RUN terraform init

ENTRYPOINT [ "/bin/sh" ]
