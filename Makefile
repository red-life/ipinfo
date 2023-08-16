generate_mocks:
	mockery --dir internal/ports --with-expecter=true --output internal/mocks --all