
Small environment for testing parsing nginx logs by promtail.  

### Install

docker-compose up 

### Usage

Contains:

1. Prometheus - for metrics 
2. Grafana - for graphs
3. Loki - for logs
4. Promtail - for logs delivery
5. Prober - nginx and small daemon who requests itself for generating logs.

Prober starts generating logs immediately. You can change its options by modifying conf.yaml.
Conf.yaml defines all possible entry points and what http status, response size, response time they will have

Type http://localhosy:3000 to enter into grafana interface.
Credentials are admin:admin. 
It has two predefined dashboards: Projects and Monitoring services

Dashboard Projects shows metrics from nginx

Dashboard Monitoring services shows metrics about monitoring services such as loki or prometheus


 

 
