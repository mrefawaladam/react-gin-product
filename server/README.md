├── cmd
│   └── main.go
├── config
│   └── config.go
├── pkg
│   ├── entity
│   │   ├── product.go
│   │   ├── productBrand.go
│   │   └── user.go
│   ├── handler
│   │   ├── productHandler.go
│   │   ├── productBrandHandler.go
│   │   └── userHandler.go
│   ├── repository
│   │   ├── gorm
│   │   │   ├── productRepository.go
│   │   │   ├── productBrandRepository.go
│   │   │   └── userRepository.go
│   ├── usecase
│   │   ├── productUsecase.go
│   │   ├── productBrandUsecase.go
│   │   └── userUsecase.go
│   └── service
│       ├── productService.go
│       ├── productBrandService.go
│       └── userService.go
└── routes
    ├── productRoutes.go
    ├── productBrandRoutes.go
    └── userRoutes.go
