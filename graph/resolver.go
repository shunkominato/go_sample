package graph

import "src/dataloader"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	CompanyLoader *dataloader.CompanyLoader // 追記
}
