# Requirements
Gmail account
An address

# This is a script/program for sending emails to colleges for free merchandise
The main binary is the actual application, if you want to look at the source code, it is in the main.go file
I scrapped a website for all of the emails, and just like in the YouTube video, I just added admissions into the emails
I wanted to keep things unique, so I sort of randomized some sentences, so it wouldn't be too obvious

# changing the phrases
I tried to add some custom phrases but it wasn't worth the time, I added a couple, but if you want to add more 
you are going to have to update the main.go file, and run the program with go run main.go [arguments]
if you want to add custom phrases looking into main.go file, yo

# usage

```
./main [google username] [google password] [first last name] [address]

./main "gmailusername" "gmail_password" "jack ma" "123 New York Street, New York, New York 21251" 

```

You need to provide your gmail username, password, "first + last name", as well as address for actually shipping the items!
