include .env
generate_mocks:
	mockery --dir internal/ports --with-expecter=true --output internal/mocks --all

download_mmdb:
	wget -P ${mmdb_files} "https://github.com/P3TERX/GeoLite.mmdb/releases/latest/download/GeoLite2-Country.mmdb"
	wget -P ${mmdb_files} "https://github.com/P3TERX/GeoLite.mmdb/releases/latest/download/GeoLite2-City.mmdb"
	wget -P ${mmdb_files} "https://github.com/P3TERX/GeoLite.mmdb/releases/latest/download/GeoLite2-ASN.mmdb"

run:
	docker-compose up

build:
	docker-compose build
