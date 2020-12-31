# AWS S3 Bucket Creator

## Setup

### Clone Repository
```
$ cd ~/var/www
$ git clone git@github.com:muzfr7/aws-s3-bucket-creator.git
```

### Create env file
```
$ cd ~/var/www/aws-s3-bucket-creator
$ mv .env.dist .env
```

### Create aws profile
Open `~/.aws/credentials` file and add following contents into it
```
[default]
aws_access_key_id = <YOUR_ACCESS_KEY_ID>
aws_secret_access_key = <YOUR_SECRET_ACCESS_KEY>
```
> Note: whatever profile name you set here, update it in `.env` file as well..

## Usage
```
$ go run
```
