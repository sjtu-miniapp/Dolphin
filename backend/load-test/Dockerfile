FROM python:2.7-wheezy

# Install locust
RUN pip install locustio

ADD *.py /usr/local/bin/
ADD execute.sh /usr/local/bin/execute.sh

WORKDIR /usr/local/bin/

EXPOSE 8089

ENTRYPOINT ["/usr/local/bin/execute.sh"]
