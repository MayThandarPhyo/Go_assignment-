
document.addEventListener("DOMContentLoaded", function() {
    const button = document.getElementById("btn");
    const image = document.getElementById("image");
    const images = ["https://pbncanvas.com/wp-content/uploads/2021/11/purple-Cute-dinosaur-paint-by-numbers.jpg" , "https://i.pinimg.com/474x/bc/1c/ff/bc1cff991cf899beed8ae4a5569344d2.jpg"]; // Array of image URLs
    let currentIndex = 0;

    button.addEventListener("click", function() {
        currentIndex = (currentIndex + 1) % images.length; 
        image.src = images[currentIndex]; 
        image.alt = `Image ${currentIndex + 1}`; 
    });
    
});

