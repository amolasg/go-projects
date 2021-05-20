package global

const (
	dburi       = "mongodb+srv://mongo:mongo@cluster0.9xq9h.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	dbname      = "blog-app"
	performance = 100
)

var (
	jwtScret = []byte("blogscret")
)
