package main

// DOES NOT COMPILE
// used only to illustrates a pattern
func tooManyChecks() {
	_, err = fd.Write(p0[a:b])
	if err != nil {
		return err
	}
	_, err = fd.Write(p1[c:d])
	if err != nil {
		return err
	}
	_, err = fd.Write(p2[e:f])
	if err != nil {
		return err
	}
}

func closurePattern() {
	var err error
	write := func(buf []byte) {
		if err != nil {
			return
		}
		_, err = w.Write(buf)
	}
	write(p0[a:b])
	write(p1[c:d])
	write(p2[e:f])
	if err != nil {
		return err
	}
}
