module "secret-dimension-uploader"

go 1.14

require (
	github.com/google/uuid v1.1.2
    github.com/KercyLAN/secret-dimension-core v0.0.0-20201012065146-28ede75e5363
)

replace github.com/KercyLAN/secret-dimension-core => ../secret-dimension-core
