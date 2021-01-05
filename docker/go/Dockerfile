FROM golang:buster

RUN apt update && apt upgrade -y

ENV ORACLE_HOME=/opt/oracle/product/18c/dbhomeXE
# create ORACLE_HOME/lib for install directory to oracle client.
RUN mkdir -p $ORACLE_HOME/lib
RUN chmod 755 $ORACLE_HOME

# need to build oci8-go
ENV PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/usr/local/lib
RUN echo prefixdir=$ORACLE_HOME >> /usr/local/lib/oci8.pc && \
    echo libdir=\${prefixdir}/lib >> /usr/local/lib/oci8.pc && \
    echo includedir=\${prefixdir}/rdbms/public >> /usr/local/lib/oci8.pc  && \
    echo Name: OCI >> /usr/local/lib/oci8.pc  && \
    echo Description: Oracle database driver >> /usr/local/lib/oci8.pc  && \
    echo Version: 18c >> /usr/local/lib/oci8.pc  && \
    echo Libs: -L\${libdir} -lclntsh >> /usr/local/lib/oci8.pc  && \
    echo Cflags: -I\${includedir} >> /usr/local/lib/oci8.pc

RUN apt -y install zip unzip

RUN ls -l $ORACLE_HOME

# download oracle instant client. this client need to build oci8-go
RUN curl -o instantclient-basic-linux.x64-18.5.0.0.0dbru.zip -L https://download.oracle.com/otn_software/linux/instantclient/185000/instantclient-basic-linux.x64-18.5.0.0.0dbru.zip && \
    unzip instantclient-basic-linux.x64-18.5.0.0.0dbru.zip && \
    mv instantclient_18_5/* $ORACLE_HOME/lib/ && \
    rm -rf instantclient-basic-linux.x64-18.5.0.0.0dbru.zip instantclient_18_5/

# clean cache
RUN apt-get clean