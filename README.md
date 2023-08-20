# IP Info
A web application for getting information of an IP address based on MaxMind Database.\

## Techologies and architecture
- It's written in golang
- using Gin framework for http routing and serving
- Implemented Hexagonal Architecture and Dependency Injection
- Unit tested all functionalities
- Integration tested the http adapter.
- Mocked MaxMind repository and Business Logic

## Deployment
- Clone the repository
- Create an .env file
```bash
mv .env.sample .env
```
- Change environment variables as desired
- Download MaxMind Database files
```bash
make download_mmdb
```
- Build and run using docker-compose
```bash
make build && make run
```

## Run tests
First of all generate mocks using `make generate_mocks`. Then you must put ".mmdb" files into internal/adapters/maxmind/tmp.
Now you can run tests using command below:\
`make test`

## Endpoints
### /
<b>METHOD:</b> <i>GET</i>\
<b>QUERY STRINGS:</b>
- continent: bool (true|false, default: true)
- country: bool (true|false, default: true)
- city: bool (true|false, default: true)
- asn: bool (true|false, default: true)\

<b>DESCRIPTION:</b> Returns the full information of the <i>client</i> IP\
<b>USAGE:</b>
```bash
curl "http://localhost:5000/" | jq
```
### /:ip
<b>METHOD:</b> <i>GET</i>\
<b>QUERY STRINGS:</b> Same as above\
<b>DESCRIPTION:</b> Returns the full information of the <i>given</i> IP in path parameter\
<b>USAGE:</b>
```bash
curl "http://localhost:5000/8.8.8.8" | jq
```
### /short
<b>METHOD:</b> <i>GET</i>\
<b>QUERY STRINGS:</b> NONE
<b>DESCRIPTION:</b> Returns the short information of the <i>client</i> IP in path parameter\
<b>USAGE:</b>
```bash
curl "http://localhost:5000/short" | jq
```

### /:ip/short
<b>METHOD:</b> <i>GET</i>\
<b>QUERY STRINGS:</b> NONE\
<b>DESCRIPTION:</b> Returns the short information of the <i>given</i> IP in path parameter\
<b>USAGE:</b>
```bash
curl "http://localhost:5000/8.8.8.8/short" | jq
```

## Alias in Terminal
If you want to get information of an ip just using a command, instead of writing `curl https://localhost:5000`
you can just set an alias in your shell configuration file:\
`alias ipinfo="curl -sk http://localhost:5000/$1/short | jq"`\
Just put this in your shell configuration file like `~/.bashrc` or `~/.zshrc` and run `source ~/.bashrc(or zshrc)`


## Future Plans
- [ ] Documenting using Swagger
- [ ] Cleaner tests and more test cases
