echo "Running API test script"
echo ""

# Create a user. Pray that nobody is already called 'SuperEpicUsername45' (obviously I wouldn't make that assumption in the real world).
echo "Creating a new user called SuperEpicUsername45"
curl -d "firstname=Johnny&lastname=Smalls&username=SuperEpicUsername45&authcode=a764bcjd" -X POST 127.0.0.1:8080/createuser
echo ""
echo ""

#List users.
curl -d "authcode=a764bcjd" -X POST 127.0.0.1:8080/listusers
echo ""
echo ""

# Rename user.
echo "Renaming that user"
curl -d "oldusername=SuperEpicUsername45&newusername=SuperEpicUsername99&authcode=a764bcjd" -X POST 127.0.0.1:8080/updatename
echo ""
echo ""

#List users.
curl -d "authcode=a764bcjd" -X POST 127.0.0.1:8080/listusers
echo ""
echo ""

# Set darkmode to false.
echo "Setting their darkmode to false"
curl -d "username=SuperEpicUsername99&on=0&authcode=a764bcjd" -X POST 127.0.0.1:8080/setdarkmode
echo ""
echo ""

#Search for this user.
echo "Searching for this user with search term 'Super'"
curl -d "term=Super&authcode=a764bcjd" -X POST 127.0.0.1:8080/searchusers
echo ""
echo ""

#Delete user.
echo "Deleting this user"
curl -d "username=SuperEpicUsername99&authcode=a764bcjd" -X POST 127.0.0.1:8080/deleteuser
echo ""
echo ""

#List users.
curl -d "authcode=a764bcjd" -X POST 127.0.0.1:8080/listusers