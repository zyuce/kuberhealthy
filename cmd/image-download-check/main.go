package main

func main() {

	/*

			c := config.New()

			s := plugin.NewStore()

			ctx := context.Background()

			d, err := daemon.NewDaemon(ctx, c, s)
			if err != nil {
				log.Fatalln("error creating new docker daemon", err)
			}

			image, err := d.ImageService().LookupImage("nginx")
			if err != nil {
				log.Fatalln("error looking up image", err)
			}

			b := new(bytes.Buffer)

			d.ImageService().PullImage(ctx, image.Config.Image, "", nil, nil, nil, b)


				// TODO: make this run from a docker server instead of a client
				// make a local docker client
				endpoint := "unix:///var/run/docker.sock"
				dockerClient, err := docker.NewClient(endpoint)
				if err != nil {
					log.Errorln("failed to created new docker client", err)
					os.Exit(1)
				}

				// create a bytes buffer for our docker pull options output stream
				b := new(bytes.Buffer)

				// create channel that will read bytes from our buffer
				// c := make(chan bytes.Buffer)

				// set options/auth for docker image pull
				opts := docker.PullImageOptions{
					Repository:   "nginx",
					Tag:          "1.21",
					OutputStream: b,
				}
				auth := docker.AuthConfiguration{}

				// perform docker pull on target image
				pullImage(dockerClient, opts, auth)

				fmt.Println(b.Len())
	}

	// timeTrack tracks the time it takes for a function to complete
	func timeTrack(start time.Time, name string) time.Duration {
		elapsed := time.Since(start)
		log.Printf("%s took %s", name, elapsed)
		return elapsed
	}

	func pullImage(dockerClient *docker.Client, opts docker.PullImageOptions, auth docker.AuthConfiguration) {
		// start timeTrack timer
		defer timeTrack(time.Now(), "pullImage")

		// pull image
		err := dockerClient.PullImage(opts, auth)
		if err != nil {
			log.Errorln("failed to pull image ", err)
			os.Exit(1)
		}

	*/
}
