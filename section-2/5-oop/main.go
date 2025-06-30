package main

import "learn-go/section-2/5-oop/usecase"

// Contoh DI
func main() {
	senderUsecase := usecase.NewSmtpSenderUsecase()
	userUsecase := usecase.NewUserUsecase(senderUsecase)

	userUsecase.RegistrasiUser("ibp@gmail.com")
}
