FROM registry.cn-hangzhou.aliyuncs.com/acejilam/centos:8
# RUN yum install epel-release vim wget htop -y
RUN rm -rf /etc/yum.repos.d/* && curl -o /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-8.repo
RUN yum clean all && yum -y install nmap-ncat net-tools
WORKDIR /
COPY ./bin/* /usr/local/bin/
