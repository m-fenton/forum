forum

This project consists of creating a web forum that allows registered users to :

Create posts with images and gifs, optionally assigning categories.
Like/dislike posts and comments.
Access your posts and liked posts via your username.
Implemented Google and GitHub authentication.

Unregistered users have view only access.


Usage: How to run

1. Open the terminal and navigate to the root directory.
2. Run the following command: "go run ."
3. Wait for the program to download the required packages. This might take a few minutes based on your internet speed.
4. Once the message "HTTPS(!)Server Starting on port 7000..." appears in the terminal, the packages have downloaded.
5. Open your browser and navigate to https://localhost:7000/ to access the forum.


Wiping the Database
For testing purposes, a method to wipe the database, including accounts and most pictures (except the logo), is provided. 
Run the following command: "go run . new" 
A new database will be created during the initialization process.


Created by 

Martin Fenton
Rupert Cheetham
Nikoi Kwasie