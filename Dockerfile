FROM ubuntu:14.04
MAINTAINER tobe tobeg3oogle@gmail.com
RUN apt-get update

# Install Go
RUN apt-get install golang

# Install build tools
RUN apt-get install -y git
RUN apt-get install -y maven

# Install HBase
RUN git clone git://git.apache.org/hbase.git /opt/hbase
WORKDIR /opt/hbase
# RUN git checkout 0.94.11 # use other version
RUN mvn clean package -DskipTests

# Port
EXPOSE 6789

CMD ["./bin/start-hbase.sh"]