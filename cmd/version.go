package cmd

var Version string

func init() {
	// buidVersion gets replaced during build with make
	var buildVersion = "dev-20140512-7337982"
	Version = "0.1"
	Version += buildVersion
}
