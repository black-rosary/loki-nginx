# Monitoring

Do not use in production. Created as a proof of concept

Contains:

1. Prometheus - for metrics 
2. Grafana - for graphs
3. Loki - for logs
4. AlertManager - for alerts
5. Promtail - for logs delivery
5. Cadviser - scraps information about host
6. Nodeexporter - scraps information about containers
7. Lebotic - The simplest Telegram bot for alerting (It is switched off by default)
8. Prober - Contains nginx and two Go daemons which can generate logs for testing. In prober folder you can find docker-compose.yml which allow to start log generator  
