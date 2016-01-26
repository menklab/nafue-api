# Nafue Security Services
# Menklab LLC

## Building/Running
1. Clone this repository.
2. Open the root directory in your favorite CLI.
* Note: If you're not running on Linux, you'll most likley be running the Docker Toolbox Terminal
3. Run `docker build -t [name] .` to build the image
* Note: The name can be anything you want; it'll be used as a quick reference for other Docker commands.
4. Run `docker run --publish [host-port]:[container-port] --name [container-name] --env-file [file-name]  --rm [image-name]`
* We've got a few things going on here:
* The --publish flag forwards a port on the host to the docker container, useful for the majority of us that are runnning Docker in a VM.
* Like images, the container name can be anything you want. It's just a way to more easily reference it for other Docker commands.
* --env-file is used to define a file for reading Environment Variables. Files with extension *.env are ignored by git, so use that for defining security credentials and not worry about commiting them to source control. It can be used multiple times if desired.
* --rm deletes the container after it exits. No need to worry about extra processes taking up system resources when you're done!
* The image name is the image you want to run in the container (which should be the one you built earlier).
5. If you're on Linux and can run the container natively, the API can be accessed through `localhost:[host-port]`. Otherwise you'll have to use the VM's public IP (which you can find by running `docker-machine ip default`.