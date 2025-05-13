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

[Air](https://github.com/air-verse/air) has been used to make development smoother with live-reload.
[Make](https://www.gnu.org/software/make/) has been used to set up some targets to not have to type commands very often.

## Makefile

The make file contains a few targets to simplify development. Things like generating CSS, templ-templates and starting the webserver will be found here.

```
make templ-gen
```
Generates Templ templates from the .templ files.

```
make templ-watch
```
Watches the .templ files for changes and generates new ones on save.

```
make taildwind-gen
```
Generates CSS from the .templ files

```
make tailwind-watch
```
Watches the .templ files for changes and generates new CSS on save.

```
make dev
```
Builds the application, generates .templ.go files and runs Air for live-reloading.

## Run locally

First of all you will need a few thing in an .env file to make things work:

```
CLIENT_ID=
CLIENT_SECRET=
MONGO_DB_CONN_STRING=
```
The client Id and secret requires creating a client on Blizzard entertainments developer-portal: https://develop.battle.net/
Once done it should give you the credentials to paste into the .env file.

The MONGO_DB_CONN_STRING is a connection string to a mongodb database. I've used a cloud version while building this project: https://www.mongodb.com/atlas
But a local one or one run through a docker container should work as well, in theory.

To run the application locally you'll need to have Go 1.24 or above installed on your system.

Then just run:
```
go run cmd/main.go
```
from the root of the project should start the webserver up, if you just want to demo the application this should suffice.
Go should install any missing dependencies when running/building.

# If you want to use Make or change some code:

You will need to install Templ to change stuff in the .templ files and be able to build new ones, check their documentation for more help but just running:
```
go install github.com/a-h/templ/cmd/templ@latest
```
should be enough to start.


You can download the [TailwindCSS CLI](https://tailwindcss.com/blog/standalone-cli) if you want to make CSS/Styling changes. 
If you do, make sure to download the one for your OS, rename it to just tailwindcss and put it in the root of the project. Also on Mac and Linux you might need to make it executable, check the link for more info.

Air is only required if you want to use live-reloading, you can install it on your machine using: 
```
go install github.com/air-verse/air@latest
```
check their documentation for more info. The repo contains a config file for Air so hopefully it should work.

If you want to use Make you can install it using [chocolatey](https://chocolatey.org/install) on windows by running (with Chocolatey installed): 
```
choco install make
```


On Mac you can use Homebrew to install Make: 
```
brew install make
```

On Linux Make should be available through most package managers:
```
sudo apt install make
```
on Ubuntu, for instance.

Once installed you should be able to run: 
```
make dev
```
from the root of the project which should build the app and start Air, given that Air is installed on the system.



## Inspiration and Help
[Raider.io](https://raider.io/) and other similar sites have stood as inspiration for general layout and what type of content to include.

The general layout, structure and setup of the Go code has been heavily inspired/helped by https://github.com/sigrdrifa/gotth-example and https://github.com/TomDoesTech/GOTTH and their acompanying youtube videos. So a massive thanks to both of them.

## Note

This is not production ready code, and should not be treated as such.


