package main

func main() {
	vault := new(Vault)
	vault.Init("")

	GetProductKey()
	vault.StoreTable(ToTable(ExtractTokens()))
	return
}
