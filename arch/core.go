package arch

type Core interface {
	Run(sys Sys, step int) int
}
