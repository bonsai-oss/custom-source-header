# custom-source-header

... is a middleware plugin for traefik. It injects the remote address of the client into a selectable header.

## Why does this exist?
When operating a kubernetes environment inside a private network combined with a second layer of traefik,
the remote address of the client is not available to the application due to header overwriting.