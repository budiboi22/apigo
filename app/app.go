package app

const BASE_PATH  = "/Users/arifbudiman/go/src/github.com/budiboi22/apigo"

func Setup() (err error) {
	err = setupConfig(BASE_PATH)
	err = setupRoute(err)

	return  err
}
