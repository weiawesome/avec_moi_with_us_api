package utils

func GetChromaClientUrl() string {
	return "http://" + EnvChromaHost() + ":" + EnvChromaPort()
}
func GetChromaNum() string {
	return EnvChromaNum()
}
