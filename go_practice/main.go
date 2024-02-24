package main

import (
	"log"
	"net/http"
)

// type anyTypeYouLike struct {
// }

type firstHandler struct {
}

func (f *firstHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`

	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title> Change images </title>
	</head>
	<style>
		.container {
			text-align: center;
			line-height: 3;
		}

		.image {
			max-width: 300px;
			max-height: 300px;
			transition: opacity 0.9s ease;

		}

		#btn {
			height: 40px;
			background-color: rgb(217, 160, 245);
			font-weight: 800;
		}
	</style>

	<body>

		<div class="container">
			<img id="image" class="image" src="https://pbncanvas.com/wp-content/uploads/2021/11/purple-Cute-dinosaur-paint-by-numbers.jpg" alt="Image 1">
			<br>
			<button id="btn">Change Image</button>
		</div>

		<script>

			document.addEventListener("DOMContentLoaded", function() {
				const button = document.getElementById("btn");
				const image = document.getElementById("image");
				const images = ["https://pbncanvas.com/wp-content/uploads/2021/11/purple-Cute-dinosaur-paint-by-numbers.jpg" , "https://i.pinimg.com/474x/bc/1c/ff/bc1cff991cf899beed8ae4a5569344d2.jpg"]; // Array of image URLs
				let currentIndex = 0;

				button.addEventListener("click", function() {
					currentIndex = (currentIndex + 1) % images.length;
					image.src = images[currentIndex];

    			});

			});
		</script>

	</body>
	</html>`))
}

// func (t *anyTypeYouLike) printAMsg(msg string) {
// 	fmt.Println(msg)
// }

func main() {
	// fmt.Println("Hello world!")
	//http.Handle("/", http.FileServer(http.Dir("templates")))

	//if err := http.ListenAndServe(":8080", nil); err != nil {
	//	log.Fatal("ListenAndServe:", err.Error())
	//}

	// myObj := anyTypeYouLike{}
	// myObj.printAMsg("Hello!")

	http.Handle("/", &firstHandler{})

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
