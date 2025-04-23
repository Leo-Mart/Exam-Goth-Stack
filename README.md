# WoW-Tracker (name pending) fullstack-app written for Exam

## General
WoW-Tracker is a pretty simple webapp for keeping track of some information and stats about your wow-character. 

Generally the aim is not to build something terribly advanced or to fill a niche in the WoW-resources space, but to just build something that I enjoy and learn while doing it, using a stack that is mostly new to me. While also serving as part of the examination of my education to become a fullstack developer.

## Stack
WoW-Tracker is built in the GoTTH stack. This stack uses:
- [Go](https://go.dev/) for the backend/webserver.
- [Templ](https://templ.guide/) which is a package for go that lets me build components and then serve/render them.
- [Tailwindcss](https://tailwindcss.com/) For styling, also using the [DaisyUI](https://daisyui.com/) plugin for some pre-built Tailwind components.
- [HTMX](https://htmx.org/) to handle requests to the backend.

## Makefile

The make file contains a few targets to simplify development. Things like generating CSS, templ-templates and starting the webserver will be found here.

## Inspiration and Help
[Raider.io](https://raider.io/) and other similar sites have stood as inspiration for general layout and what type of content to include.

The general layout, structure and setup of the Go code has been heavily inspired/helped by https://github.com/sigrdrifa/gotth-example and https://github.com/TomDoesTech/GOTTH and their acompanying youtube videos. So a massive thanks to both of them.


