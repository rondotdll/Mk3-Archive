package main

func main() {
	vault := &Vault{}

	vault.Init(TEMPFILEDIR + RandStringBytes(32) + ".lvtn")

	vault.CreateTable("SYS_INFO", map[string]string{
		"hostname":  "TEXT",
		"platform":  "TEXT",
		"CPU_Arch":  "TEXT",
		"RAM":       "TEXT",
		"disk_size": "INT",
	}).Store()
}
