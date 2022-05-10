# Terraref

## Introduction

Terraref is a GoLang application created for CS 460 & CS 461 -Software Engineering I and II. The purpose of terraref is to transform a given Virtual Private Cloud infrastructure defined in a terraform file (.tfstate) into a reference architecture diagram defined in draw.io. 



## Installation 

While not actually required to install, you will need to download [Go](https://go.dev/dl/) - This project was built on Go 1.17.2, so we suggest downloading the same version. 

### Github CLI

If you don't have the GitHub CLI installed you can follow the quickstart guide [here](https://docs.github.com/en/github-cli/github-cli/quickstart). 

After nagivating to the repository where you want to install it you can do so with the following command. 

`gh repo clone sisonsco98/terraref`

### Github

You can also download it as a .zip file from the [repository](https://github.com/sisonsco98/terraref) and save it to a location of your choosing later. 


## Usage 

Brief description of the main files

main.go - Pilot code. Also where you're able to modify the default flag values, useful when testing. 

parser.go - Takes in the .tfstate file and maps elements to GoLang structs. Also stores dependencies between elements. 
          Significant output from this is struct T of draw.io elements. 
          
mapper.go - Creates a XML tree and attaches elements with coordinates based on an underlying grid structure. Also draws arrows and zones to highlight relationships               between elements. 

validator.go - Removes invalid elements and redraws arrows as necessary to present a logical draw.io diagram. 






From the main directory you can run the following command with the flags -in and -out. If not specified, they will default to a pre-set location (see below). 

`go run main.go -in (fileLocation) -out (fileLocation)`


### Flags

-in (fileLocation) specifies the path where the input .tfstate file is located. Defaults to 'inputs/shared_vpc.tfstate', an example file. 


-out (fileLocation) specifies where you want the output file to be saved. Defaults to '"outputs/out.drawio"'. 

### Libraries 

It might seem a bit silly that we have the GCP folder buried under a series of subfolders with only one item each but this in an attempt to increase the extensibility - by segmenting the resources so throughly it greatly increases lookup speed if you ever need to do a deep dive in the code. It was also intended to seperate the provider specific functions from each other so if someone wanted to add support for AWS, they could do so without having to look at the GCP code.  


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



