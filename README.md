# ShortURL
`ShortURL`- the mini-website to shorten long URL. The website was developed on [golang v1.21.0](https://go.dev/dl/) with [fiber framework](https://docs.gofiber.io).

The service requires [docker](https://www.docker.com) for easy run this project.

## Install and Run
1) Clone this repository: `git clone https://github.com/Gamafoot/ShortURL.git`.
2) Change directory: `cd ShortURL`.
3) Run the app in docker:  `docker run --name ShortURL -p 8000:8000 -d short_url` (you can change working port for this service use var enviroment: PORT)

## Done
Now you can use the web service in your browser: `http://localhost:8000`.
