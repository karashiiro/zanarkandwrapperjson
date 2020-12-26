package sapphire

import "os"

func exists(f string) bool {
	_, err := os.Stat(f)
	return !os.IsNotExist(err)
}
