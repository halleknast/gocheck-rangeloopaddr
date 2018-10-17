package a

func RangeLoopAddrTests() {
	var s []int

	// No loop variables.
	for range s {
	}
	for _ = range s {
	}
	for _, _ = range s {
	}

	for i, j := range s {
		// Legal access to values.
		println(i, j)

		// Illegal access to addresses.
		println(&i, j)  // want "taking address of range variable 'i'"
		println(i, &j)  // want "taking address of range variable 'j'"
		println(&i, &j) // want "taking address of range variable 'i'" "taking address of range variable 'j'"

		// Legal access to shadowing variables.
		i := i
		j := j
		println(&i)
		println(&j)
	}

	{
		var i, j int

		for i = range s {
			println(i)
			println(&i) // want "taking address of range variable 'i'"
		}

		for _, j = range s {
			println(j)
			println(&j) // want "taking address of range variable 'j'"
		}

		for i, j = range s {
			println(i, j)
			println(&i) // want "taking address of range variable 'i'"
			println(&j) // want "taking address of range variable 'j'"
		}
	}

	for i := range s {
		func(i int) {
			// Legal access of shadowing variable.
			println(&i)
		}(i)
	}
	{
		var i int
		ip := &i
		for i = range s {
			println(ip)
		}
	}

	func() *int {
		for i, j := range s {
			// Legal access in return statement.
			for i, _ := range s {
				return &i
			}
			for _, j := range s {
				return &j
			}

			// Illegal access in return statement of nested function.
			func() *int {
				return &i // want "taking address of range variable 'i'"
			}()

			func() *int {
				return &j // want "taking address of range variable 'j'"
			}()
			func() *int {
				for range s {
					return &j // want "taking address of range variable 'j'"
				}
				return nil
			}()

			func() *int {
				for i := range s {
					// Legal access when in return statement.
					return &i
				}
				return nil
			}()

			// Legal access in return statement.
			return &i
		}
		return nil
	}()

	func() func() *int {
		for _, j := range s {
			return func() *int {
				// Legal access in return statement in defining loop.
				return &j
			}
		}
		return nil
	}()
}