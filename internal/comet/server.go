package comet


type Server struct {
	bucket *Bucket
}

func NewServer() *Server {
	return &Server{
		bucket: NewBucket(1024),
	}
}

func (s *Server) Bucket() *Bucket {
	return s.bucket
}
