package services

func GetServiceInfo() (name, version string, err error) {
	name = "Monorepo App"
	version = "0.1.0"
	return name, version, err
}
