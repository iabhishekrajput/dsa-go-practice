package stacks

// AsteroidCollision simulates asteroid collisions.
// Positive = moving right, negative = moving left.
// Returns the state of asteroids after all collisions.
func AsteroidCollision(asteroids []int) []int {
	// Your solution here

	stack := []int{}

	for _, asteroid := range asteroids {
		if len(stack) == 0 || asteroid > 0 {
			stack = append(stack, asteroid)
			continue
		}

		asteroidAbs := -1 * asteroid
		survived := true

		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if top < 0 {
				break
			}
			if asteroidAbs < top {
				survived = false
				break
			} else if asteroidAbs == top {
				stack = stack[:len(stack)-1]
				survived = false
				break
			}
			stack = stack[:len(stack)-1]
		}

		if survived {
			stack = append(stack, asteroid)
		}
	}

	return stack
}
