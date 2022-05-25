
go mod init name_mod
go tidy
npm install -g nodemon
nodemon --exec go run main.go --signal SIGTERM
godotenv.Load()
port := os.Getenv("TESTE_T")