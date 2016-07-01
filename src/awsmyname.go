package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/keegancsmith/shell"
	"os"
)

func main() {
	updatePathPtr := flag.String("sh", "", "a writable path to the sh file that will contain export NICKNAME")
	flag.Parse()

	sess := session.New()
	meta := ec2metadata.New(sess)

	instanceMeta, err := meta.GetInstanceIdentityDocument()

	if err != nil {
		fmt.Println("Unable to get AWS instance identity, error: ", err)
		os.Exit(1)
	}

	svc := ec2.New(sess, &aws.Config{Region: aws.String(instanceMeta.Region)})

	params := &ec2.DescribeTagsInput{
		DryRun: aws.Bool(false),
		Filters: []*ec2.Filter{
			{
				Name: aws.String("resource-id"),
				Values: []*string{
					aws.String(instanceMeta.InstanceID), // Required
					// More values...
				},
			},
			{
				Name: aws.String("key"),
				Values: []*string{
					aws.String("Name"), // Required
					// More values...
				},
			},
		},
		MaxResults: aws.Int64(256),
	}
	resp, err := svc.DescribeTags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		os.Exit(1)
		return
	}

	for i := 0; i < len(resp.Tags); i++ {
		key := aws.StringValue(resp.Tags[i].Key)
		value := aws.StringValue(resp.Tags[i].Value)

		if key == "Name" {
			exportLine := fmt.Sprintf("export NICKNAME=%s\n", shell.ReadableEscapeArg(value))

			if len(*updatePathPtr) > 0 {
				f, err := os.Create(*updatePathPtr)
				if err != nil {
					fmt.Println(err.Error())
					os.Exit(1)
				}
				defer f.Close()
				f.WriteString(exportLine)
			}

			fmt.Print(exportLine)
			os.Exit(0)
		}

	}

	os.Exit(1)
}
