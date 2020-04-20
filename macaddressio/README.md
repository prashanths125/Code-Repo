# Find Company by MAC Address

This command line utility finds the company name associated with a MAC address and writes it to the console output.

*Important*: You will need an API key obtained from https://macaddress.io to run

## Simply run the utility - Run from an existing [Dockerhub image](https://hub.docker.com/repository/docker/prashanths125/macaddressio)

`docker run -eMACADDRESSIO_API_KEY=<your api key> prashanths125/macaddressio:v3 44:38:39:ff:ef:57`

## Get module
`go get github.com/prashanths125/macaddressio`

## Run 
`export MACADDRESSIO_API_KEY=<Your api key>`

`go run github.com/prashanths125/macaddressio <Mac Address>`

e.g.

`go run github.com/prashanths125/macaddressio 44:38:39:ff:ef:57`

Sample output

`Cumulus Networks, Inc`

## Test
`go test github.com/prashanths125/macaddressio`

## Build your own docker image

`docker build -t <your_tag_name>:<version> .`

## Run the docker image

`docker run -eMACADDRESSIO_API_KEY=<your api key> <your_tag_name>:<version> <MAC_ADDRESS_TO_SEARCH>`








