# goSellix
Golang binding for the official e-commerce Sellix API.
This project is still on development and my code is far from being perfect, so if you want to help feel free to do it.

# How to download
    go get -u github.com/Taptiive/goSellix

# How to use
_First, you need to create a Sellix Client with your API key that you can find [here](https://dashboard.sellix.io/settings/security)_
```
sellixClient := goSellix.NewClient(YOUR_API_KEY)
```

# License
This software is Licensed under GPL version 3. A copy of the license can be found in the file LICENSE
