package server

type Server interface {
	NewRouter()
	Run() error
}
