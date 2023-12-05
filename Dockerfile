FROM traefik:v2.5
## Default module name (put your setting in .env to override)

ARG PLUGIN_MODULE=${PLUGIN_MODULE}
ADD . /plugins-local/src/${PLUGIN_MODULE}

