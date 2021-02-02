Copyright Anthony Sharp 2021

-------- PREREQUISITES --------

You'll need a few things in order to run this. At the very least, you'll need MySQL, curl, Go and probably a Linux machine.

You'll also need a MySQL account on your computer with sufficient access. To do this, you will probably need to do the following:

sudo mysql
CREATE USER 'admin'@'localhost' IDENTIFIED BY 'secretpassword';
GRANT ALL PRIVILEGES ON JDIAPIDB . * TO 'admin'@'localhost';
FLUSH PRIVILEGES;

NOTE: the account name and password can be changed in database.go by simply changing the variables at the top of the file. If you do that, presumably, all you'll need
      to run are the last two lines above.

-------- USAGE --------

To start the server, run the executable file (do it from the command line or you probably won't see any output). The database is automatically created if it doesn't already exist.

To run the curl tests, start the server and then run the .sh file.
To run the Go unit tests, cd into the src directory and run go test (you do not need to start the server; it will start itself).

-------- TIME SPENT --------

I timed myself making this and it took me 4 hours and 44 minutes in total. This doesn't include the time it took to install all the prerequisites or write this readme.

-------- API ENDPOINTS --------

The server listens on port 8080 and accepts POST requests only (for security reasons) on the following paths:

/createuser     requires fields: firstname, lastname, username, authcode    returns: 0/1 (1 = success, 0 = failure)
/updatename     requires fields: oldusername, newusername, authcode         returns: 0/1
/setdarkmode    requires fields: username, on, authcode                     returns: 0/1
/deleteuser     requires fields: username, authcode                         returns: 0/1
/listusers      requires fields: authcode                                   returns: JSON (or nothing if nothing found)
/searchusers    requires fields: term, authcode                             returns: JSON (or nothing if nothing found)

I preferred this design over doing something like /deleteuser/32 or DELETE users/32. I feel that this way is cleaner and is safer since it exposes less of the user's data through the URL
(which anyone could be sniffing).

I thought it was overkill to format the "0/1" that the first four endpoints return as JSON, which is why they don't return JSON.

In the real world I would not have indented the JSON since that consumes extra bandwidth and bandwidth costs money.

-------- IMPROVEMENTS TO MAKE --------

There are plenty of corners I cut for the sake of simplicity. Here are some of the things that could be improved:

- /listusers and /searchusers should return something even if nothing was found.

- Implement a way to close the server whilst it is running without having to force-close it (helps prevent database corruptions etc).

- Use of transactions (although, to be fair, none of the current API endpoints would really benefit from transactions, since most [in fact, I think all] of them access the 
    database in a single step at a time, so there is not really any opportunity for corruption, other than perhaps the odd warning [e.g. if /updatename were to be called on a
    user right after /deleteuser had been called on them ... actually, that probably wouldn't even generate a warning, it just wouldn't do anything]).

- A better test function for searching users. Real users might have the phrase "test" in their names, which sort of makes it an imperfect test (although it's good enough for
    99% of use cases).

- Perhaps better file management. Maybe. In a real-life project with more code, I would have made more of an effort to organise the code. I also likely would have added more
    comments and maybe even a header describing what each file does at the top of each file.

- Use a more elegant method to check whether the database already exists or not.
