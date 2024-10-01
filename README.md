# README

Hello there!

This is the repository where your SDSE TA's will upload any code seen during exercise livecoding sessions!

Do you have questions regarding the exercises, lectures or the course in general?

Feel free to ask them using the Q&A forum on the SDSE course page.


# Upcoming structure

When more livecoding sessions occur, we will restructure this repo's code into folders per livecoding session and make a table for overview.

E.g:

| Week no. | Code contents |
| --- | --- |
| 4 | linkedlist |
| 5 | temperatureConverter |

# Nix support
**If you are not using nix please ignore these files in the root folder of the repo:**  
/  
├── .envrc  
├── .devenv  
├── devenv.nix  
└── devenv.yaml  
  
If you are using nix you can start a devenv shell with go using the following command:
```bash
$ devenv shell
```
You can even allow the devenv to be started automatically when you enter the directory by running:
```bash
$ direnv allow
```

