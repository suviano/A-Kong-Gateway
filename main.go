package main

func main() {
	// http://www.michaelbach.de/ot/sze_silhouette/index.html
	requestMaker(KongAPI{
		Name:        "silhouette",
		UpstreamURL: "http://www.michaelbach.de/ot/sze_silhouette/index.html",
		Hosts:       []string{"www.michaelbach.de"},
	})
	// http://www.exploratorium.edu/files/exhibits/fading_dot/fading_dot.html
	requestMaker(KongAPI{
		Name:        "fading dot",
		UpstreamURL: "http://www.exploratorium.edu/files/exhibits/fading_dot/fading_dot.html",
		Methods:     []string{"GET", "OPTIONS"},
	})
	// http://www.visnos.com/demos/fractal
	requestMaker(KongAPI{
		Name:        "fractal",
		UpstreamURL: "http://www.visnos.com/demos/fractal",
		Uris:        []string{"/fractal"},
	})
	// http://imaninja.com
	requestMaker(KongAPI{
		Name:        "iamninja",
		UpstreamURL: "http://imaninja.com",
		Hosts:       []string{"imaninja.com"},
		Methods:     []string{"GET", "OPTIONS"},
	})
}
