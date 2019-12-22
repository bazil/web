package fs

import "io"

func sample() error {
	var c interface{}
	// START OMIT
	for {
		req, err := c.ReadRequest()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		go c.serve(fs, req)
	}
	// END OMIT
}
