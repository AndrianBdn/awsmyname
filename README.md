# awsmyname 

Tool written in golang that fetches AWS instance Name tag 


## Usage 

Just run it. On success it will emit 

```
export NICKNAME=aws-Name-tag-value
```

and exit code 0 (success). Otherwise it will print something else and set non-0 (most likely 1) exit code.


## IAM Role 

This tool expects instance to have IAM Role that allows ec2:DescribeTags action. See policy below:

```json 
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "Stmt1467041390000",
            "Effect": "Allow",
            "Action": [
                "ec2:DescribeTags"
            ],
            "Resource": [
                "*"
            ]
        }
    ]
}
```
 

## Build 

Can be build using Docker, just run build.sh on your host and get your binary from 'binary' folder. 