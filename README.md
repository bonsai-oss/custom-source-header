# Traefik Plugin: Custom Source Header

Custom Source Header is a middleware plugin for Traefik. It injects the remote address of the client into a selectable header.

## Why does this exist?
When operating a kubernetes environment inside a private network combined with a second layer of Traefik, the remote address of the client is not available to the application due to header overwriting.

## Install the Plugin

To use the plugin please install it according to the following instructions from the Traefic Plugin Catalog: https://plugins.traefik.io/plugins/633a02e3eff06de33b092377/custom-source-header
