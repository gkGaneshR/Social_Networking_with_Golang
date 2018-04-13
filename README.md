# Social_Networking_with_Golang
The objective is to develop a Golang server so that users can upload image, add comments and view image like a social network web application.
It used BoltDB at the backend to store user details, image details and comments entered by users image-wise. The images uploaded of different size are resized to a common size and a secure cookie is also maintained for every session login.
Dependencies:
"github.com/boltdb/bolt"
"github.com/eahydra/gouuid"
"github.com/nfnt/resize"
"github.com/gorilla/mux"
