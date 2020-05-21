FROM python:3.6-slim
  
RUN apt-get update && apt-get install -y \
        build-essential \
        libxml2-dev \
        libxslt1-dev \
        libffi-dev\
        libpq-dev \
        libssl-dev \
        zlib1g-dev \
        git \
        libsasl2-dev \
        python-dev \
        libldap2-dev \
        libssl-dev \
        vim \
        && rm -rf /var/lib/apt/lists/*

RUN pip install --no-warn-script-location \
# gunicorn is used for launching netbox
      gunicorn \
      greenlet \
      eventlet \
# napalm is used for gathering information from network devices
      napalm \
# ruamel is used in startup_scripts
      'ruamel.yaml>=0.15,<0.16' \
# django_auth_ldap is required for ldap
      django_auth_ldap \
# django-storages was introduced in 2.7 and is optional
      django-storages


WORKDIR /opt/

RUN ls -a
RUN git clone -b master https://github.com/netbox-community/netbox.git .
WORKDIR /opt/netbox/

COPY ./requirements.txt /
COPY ./configuration.py /opt/netbox/netbox/


RUN pip install --no-warn-script-location -r /requirements.txt


