package test

import "github.com/afex/hystrix-go/hystrix"

func main() {
	hystrix.ConfigureCommand("my_command", hystrix.CommandConfig{
		Timeout:               10,
		MaxConcurrentRequests: 100,
		ErrorPercentThreshold: 25,
	})

	err := hystrix.Do("my_command", func() error {
		// talk to other services
		return nil
	}, nil)
	print(err)

}
