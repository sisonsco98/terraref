# Turbo Sad Simulator

## Introduction

Turbo Sad Simulator (TSS) is a GoLang application created for CS 460 & CS 461 -Software Engineering I and II. The purpose of TSS is to transform a given Virtual Private Cloud infrastructure defined in a terraform file (.tfstate) into a reference architecture diagram defined in draw.io. 



## Installation 

While not actually required to install, you will need to download [Go](https://go.dev/dl/) - ideally v1.17 or higher. 

### Github CLI

If you don't have the GitHub CLI installed you can follow the quickstart guide [here](https://docs.github.com/en/github-cli/github-cli/quickstart). 

After nagivating to the repository where you want to install it you can do so with the following command. 

`gh repo clone sisonsco98/terraref`

### Github

You can also download it as a .zip file from the [repository](https://github.com/sisonsco98/terraref) and save it to a location of your choosing later. 


## Usage 

From the main directory you can run the following command 

`go run main.go -in (fileLocation) -out (fileLocation)`


### Flags

-in (fileLocation) specifies the path where the input .tfstate file is located.


-out (fileLocation) specifies where you want the output file to be saved.  


## Roadmap 

- [ ] Add support for Azure & AWS
- [ ] Automated Build Pipeline
- [ ] Expand range of supported terraform resources (VMs)

Please see the [current known issues](https://github.com/sisonsco98/terraref/issues) for a list of fixes that still need to be implemented. 

## Contact

Scott Sison - @sisonsco98 - sisonsco@hawaii.edu

Kaley Fujii - @kaleyf - kaleyf@hawaii.edu

Daniel Hoshino - @Reeseslover123 - danielmh@hawaii.edu

## License 

This project is open source under the MIT Open Source License. You can read the details in the [LICENSE.txt](https://github.com/sisonsco98/terraref/blob/ReadmeBranch/LICENSE.txt) file. 



