---
name: example
description: Work with the example togo plugin (endpoints, UI, service).
---

# example plugin

This plugin exposes `GET /api/example/ping` and injects `web/app/example/page.tsx`.
Backend logic lives in `internal/example`. Extend the service, add routes in
`plugin.go`, and ship UI under `web/`.
