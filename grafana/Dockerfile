ARG GRAFANA_VERSION
FROM grafana/grafana:${GRAFANA_VERSION}

ARG NGINX_GRAFANA_HOST
USER root
COPY grafana.ini /etc/grafana/grafana.ini
RUN sed -i "s/%NGINX_GRAFANA_HOST%/${NGINX_GRAFANA_HOST}/" /etc/grafana/grafana.ini
# do projects as default dashboard
RUN curl 'http://127.0.0.1:3000/api/org/preferences' -X PUT  'http://127.0.0.1:3000' -H 'Accept-Encoding: gzip, deflate, br' -H 'X-Grafana-Org-Id: 1' -H 'Accept-Language: en-IN,en;q=0.9,en-GB;q=0.8,en-US;q=0.7,ml;q=0.6,mt;q=0.5' -H 'Content-Type: application/json;charset=UTF-8' -H 'Accept: application/json, text/plain, */*' -H 'Cache-Control: no-cache' -H 'Connection: keep-alive' --data-binary '{"homeDashboardId":6,"theme":"","timezone":""}' --compressed